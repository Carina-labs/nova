package keeper

import (
	"context"
	"errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"time"

	"github.com/Carina-labs/nova/x/icacontrol/types"
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

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the icacontrol Msg service interface
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
			ControllerAddress: zone.IcaAccount.ControllerAddress,
		},
		TransferInfo: &types.TransferConnectionInfo{
			PortId:    zone.TransferInfo.PortId,
			ChannelId: zone.TransferInfo.ChannelId,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		SnDenom:          appendSnPrefix(types.PrefixSnAsset, zone.BaseDenom),
	}

	if !k.IsValidDaoModifier(ctx, zone.IcaAccount.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.ControllerAddress)
	}

	_, ok := k.Keeper.GetRegisteredZone(ctx, zoneInfo.ZoneId)
	if ok {
		return nil, errors.New(zoneInfo.ZoneId + "already registered")
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)

	if err := k.IcaControllerKeeper.RegisterInterchainAccount(ctx, zone.IcaInfo.ConnectionId, zoneInfo.IcaConnectionInfo.PortId); err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{}, nil
}

// DeleteRegisteredZone implements the Msg/MsgDeleteRegisteredZone interface
func (k msgServer) DeleteRegisteredZone(goCtx context.Context, zone *types.MsgDeleteRegisteredZone) (*types.MsgDeleteRegisteredZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, zone.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.ControllerAddress)
	}

	_, ok := k.GetRegisteredZone(ctx, zone.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	k.Keeper.DeleteRegisteredZone(ctx, zone.ZoneId)
	return &types.MsgDeleteRegisteredZoneResponse{}, nil
}

// ChangeRegisteredZoneInfo implements the Msg/MsgChangeRegisteredZoneInfo interface
func (k msgServer) ChangeRegisteredZoneInfo(goCtx context.Context, zone *types.MsgChangeRegisteredZoneInfo) (*types.MsgChangeRegisteredZoneInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, zone.IcaAccount.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.ControllerAddress)
	}

	zoneInfo := &types.RegisteredZone{
		ZoneId: zone.ZoneId,
		IcaConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: zone.IcaInfo.ConnectionId,
			PortId:       zone.IcaInfo.PortId,
		},
		IcaAccount: &types.IcaAccount{
			ControllerAddress: zone.IcaAccount.ControllerAddress,
			HostAddress:       zone.IcaAccount.HostAddress,
		},
		TransferInfo: &types.TransferConnectionInfo{
			PortId:    zone.TransferInfo.PortId,
			ChannelId: zone.TransferInfo.ChannelId,
		},
		ValidatorAddress: zone.ValidatorAddress,
		BaseDenom:        zone.BaseDenom,
		SnDenom:          appendSnPrefix(types.PrefixSnAsset, zone.BaseDenom),
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)
	return &types.MsgChangeRegisteredZoneInfoResponse{}, nil
}

// IcaDelegate implements the Msg/MsgIcaDelegate interface
func (k msgServer) IcaDelegate(goCtx context.Context, msg *types.MsgIcaDelegate) (*types.MsgIcaDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaDelegate transaction failed to send")
	}

	return &types.MsgIcaDelegateResponse{}, nil
}

// IcaUndelegate implements the Msg/MsgIcaUndelegate interface
func (k msgServer) IcaUndelegate(goCtx context.Context, msg *types.MsgIcaUndelegate) (*types.MsgIcaUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

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
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg

	msgs = append(msgs, &distributiontype.MsgWithdrawDelegatorReward{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress})
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})

	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaAutoStaking transaction failed to send")
	}

	return &types.MsgIcaAutoStakingResponse{}, nil
}

// IcaTransfer implements the Msg/MsgIcaTransferResponse interface
func (k msgServer) IcaTransfer(goCtx context.Context, msg *types.MsgIcaTransfer) (*types.MsgIcaTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg

	//transfer msg
	//sourceport, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    msg.IcaTransferPortId,
		SourceChannel: msg.IcaTransferChannelId,
		Token:         msg.Amount,
		Sender:        msg.HostAddress,
		Receiver:      msg.ReceiverAddress,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	})

	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaTransfer transaction failed to send")
	}

	return &types.MsgIcaTransferResponse{}, nil
}

func (k msgServer) IcaAuthzGrant(goCtx context.Context, msg *types.MsgIcaAuthzGrant) (*types.MsgIcaAuthzGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg
	msgs = append(msgs, &authz.MsgGrant{
		Granter: zoneInfo.IcaAccount.HostAddress,
		Grantee: msg.Grantee,
		Grant:   msg.Grant,
	})
	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaAuthzGrant transaction failed to send")
	}

	return &types.MsgIcaAuthzGrantResponse{}, nil
}

func (k msgServer) IcaAuthzRevoke(goCtx context.Context, msg *types.MsgIcaAuthzRevoke) (*types.MsgIcaAuthzRevokeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	var msgs []sdk.Msg
	msgs = append(msgs, &authz.MsgRevoke{
		Granter:    zoneInfo.IcaAccount.HostAddress,
		Grantee:    msg.Grantee,
		MsgTypeUrl: msg.MsgTypeUrl,
	})

	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaAuthzRevoke transaction failed to send")
	}

	return &types.MsgIcaAuthzRevokeResponse{}, nil
}
