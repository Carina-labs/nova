package keeper

import (
	"context"

	"github.com/Carina-labs/novachain/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// RegisterZone implements the Msg/RegisterZone interface
func (k msgServer) RegisterZone(goCtx context.Context, zone *types.MsgRegisterZone) (*types.MsgRegisterZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.SetRegesterZone(ctx, *zone)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, zone.ConnectionId, zone.OwnerAddress); err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{}, nil
}

func (k msgServer) ICADelegate(goCtx context.Context, msg *types.MsgICADelegate) (*types.MsgICADelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone_info, ok := k.GetRegisteredZone(ctx, msg.ZoneName)

	if !ok {
		panic("zone name not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress, Amount: msg.Amount})
	err := k.SendICATx(ctx, zone_info.OwnerAddress, zone_info.ConnectionId, msgs)

	if err != nil {
		panic("ICADelegate transaction failed to send")
	}

	return nil, nil
}

func (k msgServer) ICAUndelegate(goCtx context.Context, msg *types.MsgICAUndelegate) (*types.MsgICAUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone_info, ok := k.GetRegisteredZone(ctx, msg.ZoneName)

	if !ok {
		panic("zone name not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress, Amount: msg.Amount})
	err := k.SendICATx(ctx, zone_info.OwnerAddress, zone_info.ConnectionId, msgs)

	if err != nil {
		panic("ICAUnDelegate transaction failed to send")
	}

	return nil, nil
}
