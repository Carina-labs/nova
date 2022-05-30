package keeper

import (
	"context"

	"github.com/Carina-labs/novachain/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
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

	ZoneInfo := &types.RegisteredZone{
		ZoneName: zone.ZoneName,
		ConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: zone.ConnectionId,
			OwnerAddress: zone.OwnerAddress,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		StDenom:          "st" + zone.BaseDenom,
		SnDenom:          "sn" + zone.BaseDenom,
		AuthzAddress:     zone.AuthzAddress,
	}

	k.SetRegesterZone(ctx, *ZoneInfo)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, zone.ConnectionId, zone.OwnerAddress); err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{}, nil
}

func (k msgServer) IcaDelegate(goCtx context.Context, msg *types.MsgIcaDelegate) (*types.MsgIcaDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone_info, ok := k.GetRegisteredZone(ctx, msg.ZoneName)

	if !ok {
		panic("zone name not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zone_info.ConnectionInfo.OwnerAddress, zone_info.ConnectionInfo.ConnectionId, msgs)

	if err != nil {
		panic("IcaDelegate transaction failed to send")
	}

	return nil, nil
}

func (k msgServer) IcaUndelegate(goCtx context.Context, msg *types.MsgIcaUndelegate) (*types.MsgIcaUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone_info, ok := k.GetRegisteredZone(ctx, msg.ZoneName)

	if !ok {
		panic("zone name not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zone_info.ConnectionInfo.OwnerAddress, zone_info.ConnectionInfo.ConnectionId, msgs)

	if err != nil {
		panic("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgIcaUndelegateResponse{}, nil
}

func (k msgServer) IcaAutoStaking(goCtx context.Context, msg *types.MsgIcaAutoStaking) (*types.MsgIcaAutoStakingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone_info, ok := k.GetRegisteredZone(ctx, msg.ZoneName)
	if !ok {
		panic("zone name not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &distributiontype.MsgWithdrawDelegatorReward{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress})
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zone_info.ValidatorAddress, Amount: msg.Amount})

	err := k.SendIcaTx(ctx, msg.OwnerAddress, zone_info.ConnectionInfo.ConnectionId, msgs)
	if err != nil {
		panic("IcaAutoCompound transaction failed to send")
	}

	return &types.MsgIcaAutoStakingResponse{}, nil
}
