package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
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
	err := k.Keeper.RegisterZone(ctx, zone)
	if err != nil {
    return nil, err
  }
  
	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, zone.IcaInfo.ConnectionId, zone.IcaInfo.PortId); err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{}, nil
}

func (k msgServer) DeleteRegisteredZone(goCtx context.Context, zone *types.MsgDeleteRegisteredZone) (*types.MsgDeleteRegisteredZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := k.GetRegisteredZone(ctx, zone.ZoneName)

	if !ok {
		return &types.MsgDeleteRegisteredZoneResponse{}, errors.New("zone name is not found")
	}

	if zoneInfo.IcaAccount.OwnerAddress != zone.SenderAddress {
		return &types.MsgDeleteRegisteredZoneResponse{}, errors.New("")
	}

	k.DeleteRegisterZone(ctx, zone.ZoneName)

	return &types.MsgDeleteRegisteredZoneResponse{}, nil
}

func (k msgServer) ChangeRegisteredZoneInfo(goCtx context.Context, zone *types.MsgChangeRegisteredZoneInfo) (*types.MsgChangeRegisteredZoneInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo := &types.RegisteredZone{
		ZoneId: zone.ZoneName,
		IcaConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: zone.IcaInfo.ConnectionId,
			PortId:       zone.IcaInfo.PortId,
		},
		TransferConnectionInfo: &types.TransferConnectionInfo{
			ConnectionId: zone.TransferInfo.ConnectionId,
			PortId:       zone.TransferInfo.PortId,
			ChannelId:    zone.TransferInfo.ChannelId,
		},
		IcaAccount: &types.IcaAccount{
			OwnerAddress: zone.IcaAccount.OwnerAddress,
			HostAddress:  zone.IcaAccount.HostAddress,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		StDenom:          "st" + zone.BaseDenom,
		SnDenom:          "sn" + zone.BaseDenom,
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)
	return &types.MsgChangeRegisteredZoneInfoResponse{}, nil
}

func (k msgServer) IcaDelegate(goCtx context.Context, msg *types.MsgIcaDelegate) (*types.MsgIcaDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneName)
	if !ok {
		return &types.MsgIcaDelegateResponse{}, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zoneInfo.IcaAccount.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return &types.MsgIcaDelegateResponse{}, errors.New("IcaDelegate transaction failed to send")
	}

	return &types.MsgIcaDelegateResponse{}, nil
}

func (k msgServer) IcaUndelegate(goCtx context.Context, msg *types.MsgIcaUndelegate) (*types.MsgIcaUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneName)

	if !ok {
		return &types.MsgIcaUndelegateResponse{}, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zoneInfo.IcaAccount.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return &types.MsgIcaUndelegateResponse{}, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgIcaUndelegateResponse{}, nil
}

func (k msgServer) IcaAutoStaking(goCtx context.Context, msg *types.MsgIcaAutoStaking) (*types.MsgIcaAutoStakingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneName)
	if !ok {
		return &types.MsgIcaAutoStakingResponse{}, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &distributiontype.MsgWithdrawDelegatorReward{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zoneInfo.ValidatorAddress})
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.SenderAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})

	err := k.SendIcaTx(ctx, msg.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return &types.MsgIcaAutoStakingResponse{}, errors.New("IcaAutoStaking transaction failed to send")
	}

	return &types.MsgIcaAutoStakingResponse{}, nil
}

func (k msgServer) IcaWithdraw(goCtx context.Context, msg *types.MsgIcaWithdraw) (*types.MsgIcaWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneName)
	if !ok {
		return &types.MsgIcaWithdrawResponse{}, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	//transfer msg
	//sourceport, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    zoneInfo.TransferConnectionInfo.PortId,
		SourceChannel: zoneInfo.TransferConnectionInfo.ChannelId,
		Token:         msg.Amount,
		Sender:        msg.SenderAddress,
		Receiver:      msg.ReceiverAddress,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	})

	err := k.SendIcaTx(ctx, msg.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return &types.MsgIcaWithdrawResponse{}, errors.New("IcaWithdraw transaction failed to send")
	}

	return &types.MsgIcaWithdrawResponse{}, nil
}

func (k msgServer) IcaRegisterHostAccount(goCtx context.Context, msg *types.MsgRegisterHostAccount) (*types.MsgRegisterHostAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo := k.GetRegisteredZoneForPortId(ctx, "icacontroller-"+msg.AccountInfo.OwnerAddress)
	if zoneInfo.ZoneId == "" {
		return &types.MsgRegisterHostAccountResponse{}, errors.New("zone is not found")
	}

	zoneInfo.IcaAccount.HostAddress = msg.AccountInfo.HostAddress

	k.Keeper.RegisterZone(ctx, zoneInfo)

	return &types.MsgRegisterHostAccountResponse{}, nil
}
