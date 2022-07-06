package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	zoneInfo := &types.RegisteredZone{
		ZoneId: zone.ZoneId,
		IcaConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: zone.IcaInfo.ConnectionId,
			PortId:       zone.IcaInfo.PortId,
		},
		IcaAccount: &types.IcaAccount{
			DaomodifierAddress: zone.IcaAccount.DaomodifierAddress,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		SnDenom:          "sn" + zone.BaseDenom,
	}

	if !k.IsValidDaoModifier(ctx, zone.IcaAccount.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.DaomodifierAddress)
	}

	_, ok := k.Keeper.GetRegisteredZone(ctx, zoneInfo.ZoneId)
	if ok {
		return nil, errors.New(zoneInfo.ZoneId + "already registered")
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, zone.IcaInfo.ConnectionId, zone.IcaInfo.PortId); err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{}, nil
}

// DeleteRegisteredZone implements the Msg/MsgDeleteRegisteredZone interface
func (k msgServer) DeleteRegisteredZone(goCtx context.Context, zone *types.MsgDeleteRegisteredZone) (*types.MsgDeleteRegisteredZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, zone.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, zone.ZoneId)

	if !ok {
		return nil, errors.New("zone name is not found")
	}

	if zoneInfo.IcaAccount.DaomodifierAddress != zone.DaomodifierAddress {
		return nil, errors.New("sender is not valid daomodifier address")
	}

	k.Keeper.DeleteRegisteredZone(ctx, zone.ZoneId)
	return &types.MsgDeleteRegisteredZoneResponse{}, nil
}

// ChangeRegisteredZoneInfo implements the Msg/MsgChangeRegisteredZoneInfo interface
func (k msgServer) ChangeRegisteredZoneInfo(goCtx context.Context, zone *types.MsgChangeRegisteredZoneInfo) (*types.MsgChangeRegisteredZoneInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, zone.IcaAccount.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.DaomodifierAddress)
	}

	zoneInfo := &types.RegisteredZone{
		ZoneId: zone.ZoneId,
		IcaConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: zone.IcaInfo.ConnectionId,
			PortId:       zone.IcaInfo.PortId,
		},
		IcaAccount: &types.IcaAccount{
			DaomodifierAddress: zone.IcaAccount.DaomodifierAddress,
			HostAddress:        zone.IcaAccount.HostAddress,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		SnDenom:          "sn" + zone.BaseDenom,
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)
	return &types.MsgChangeRegisteredZoneInfoResponse{}, nil
}

// IcaDelegate implements the Msg/MsgIcaDelegate interface
func (k msgServer) IcaDelegate(goCtx context.Context, msg *types.MsgIcaDelegate) (*types.MsgIcaDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zoneInfo.IcaAccount.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaDelegate transaction failed to send")
	}

	return &types.MsgIcaDelegateResponse{}, nil
}

// IcaUndelegate implements the Msg/MsgIcaUndelegate interface
func (k msgServer) IcaUndelegate(goCtx context.Context, msg *types.MsgIcaUndelegate) (*types.MsgIcaUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendIcaTx(ctx, zoneInfo.IcaAccount.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgIcaUndelegateResponse{}, nil
}

// IcaAutoStaking implements the Msg/MsgIcaAutoStaking interface
func (k msgServer) IcaAutoStaking(goCtx context.Context, msg *types.MsgIcaAutoStaking) (*types.MsgIcaAutoStakingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &distributiontype.MsgWithdrawDelegatorReward{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress})
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})

	err := k.SendIcaTx(ctx, msg.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaAutoStaking transaction failed to send")
	}

	return &types.MsgIcaAutoStakingResponse{}, nil
}

// IcaWithdraw implements the Msg/MsgIcaWithdrawResponse interface
func (k msgServer) IcaWithdraw(goCtx context.Context, msg *types.MsgIcaWithdraw) (*types.MsgIcaWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone name is not found")
	}

	var msgs []sdk.Msg

	//transfer msg
	//sourceport, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    msg.TransferPortId,
		SourceChannel: msg.TransferChannelId,
		Token:         msg.Amount,
		Sender:        msg.HostAddress,
		Receiver:      msg.ReceiverAddress,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	})

	err := k.SendIcaTx(ctx, msg.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaWithdraw transaction failed to send")
	}

	return &types.MsgIcaWithdrawResponse{}, nil
}

// IcaRegisterHostAccount implements the Msg/MsgRegisterHostAccount interface
func (k msgServer) IcaRegisterHostAccount(goCtx context.Context, msg *types.MsgRegisterHostAccount) (*types.MsgRegisterHostAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo := k.GetRegisteredZoneForPortId(ctx, msg.AccountInfo.DaomodifierAddress)
	if zoneInfo == nil {
		return &types.MsgRegisterHostAccountResponse{}, errors.New("zone is not found")
	}

	zoneInfo.IcaAccount.HostAddress = msg.AccountInfo.HostAddress

	k.Keeper.RegisterZone(ctx, zoneInfo)

	return &types.MsgRegisterHostAccountResponse{}, nil
}
