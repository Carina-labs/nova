package keeper

import (
	"context"
	"errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"strconv"
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
	*Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the icacontrol Msg service interface
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// RegisterZone implements the Msg/RegisterZone interface
func (k msgServer) RegisterZone(goCtx context.Context, zone *types.MsgRegisterZone) (*types.MsgRegisterZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, zone.ZoneId, zone.IcaAccount.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.ControllerAddress)
	}

	_, ok := k.Keeper.GetRegisteredZone(ctx, zone.ZoneId)
	if ok {
		return nil, errors.New(zone.ZoneId + "already registered")
	}

	zoneId := k.DenomDuplicateCheck(ctx, zone.BaseDenom)

	if zoneId != "" {
		return nil, sdkerrors.Wrap(types.ErrDenomDuplicates, zone.BaseDenom)
	}

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
		ValidatorAddress:     zone.ValidatorAddress,
		BaseDenom:            zone.BaseDenom,
		SnDenom:              appendSnPrefix(types.PrefixSnAsset, zone.BaseDenom),
		Decimal:              zone.Decimal,
		UndelegateMaxEntries: zone.UndelegateMaxEntries,
		DepositMaxEntries:    zone.DepositMaxEntries,
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)

	if err := k.IcaControllerKeeper.RegisterInterchainAccount(ctx, zone.IcaInfo.ConnectionId, zoneInfo.IcaConnectionInfo.PortId); err != nil {
		return nil, err
	}

	//register_zone event
	err := ctx.EventManager().EmitTypedEvent(types.NewEventRegisterZone(zoneInfo))
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterZoneResponse{
		ZoneId:               zoneInfo.ZoneId,
		IcaInfo:              zoneInfo.IcaConnectionInfo,
		TransferInfo:         zoneInfo.TransferInfo,
		ValidatorAddress:     zoneInfo.ValidatorAddress,
		BaseDenom:            zoneInfo.BaseDenom,
		SnDenom:              zoneInfo.BaseDenom,
		Decimal:              zoneInfo.Decimal,
		DepositMaxEntries:    zoneInfo.DepositMaxEntries,
		UndelegateMaxEntries: zoneInfo.UndelegateMaxEntries,
	}, nil
}

// DeleteRegisteredZone implements the Msg/MsgDeleteRegisteredZone interface
func (k msgServer) DeleteRegisteredZone(goCtx context.Context, zone *types.MsgDeleteRegisteredZone) (*types.MsgDeleteRegisteredZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, zone.ZoneId, zone.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.ControllerAddress)
	}

	_, ok := k.GetRegisteredZone(ctx, zone.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	k.Keeper.DeleteRegisteredZone(ctx, zone.ZoneId)

	err := ctx.EventManager().EmitTypedEvent(types.NewEventDeleteZone(zone))
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteRegisteredZoneResponse{}, nil
}

// ChangeRegisteredZoneInfo implements the Msg/MsgChangeRegisteredZoneInfo interface
func (k msgServer) ChangeRegisteredZone(goCtx context.Context, zone *types.MsgChangeRegisteredZone) (*types.MsgChangeRegisteredZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, zone.ZoneId, zone.IcaAccount.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, zone.IcaAccount.ControllerAddress)
	}

	zoneId := k.DenomDuplicateCheck(ctx, zone.BaseDenom)
	if zoneId != "" {
		return nil, sdkerrors.Wrap(types.ErrDenomDuplicates, zone.BaseDenom)
	}

	ok := k.Keeper.IcaControllerKeeper.IsBound(ctx, zone.IcaInfo.PortId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrInvalidPortId, zone.IcaInfo.PortId)
	}

	_, ok = k.Keeper.IcaControllerKeeper.GetOpenActiveChannel(ctx, zone.IcaInfo.ConnectionId, zone.IcaInfo.PortId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrInvalidConnId, zone.IcaInfo.ConnectionId)
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
		ValidatorAddress:     zone.ValidatorAddress,
		BaseDenom:            zone.BaseDenom,
		SnDenom:              appendSnPrefix(types.PrefixSnAsset, zone.BaseDenom),
		Decimal:              zone.Decimal,
		UndelegateMaxEntries: zone.UndelegateMaxEntries,
		DepositMaxEntries:    zone.DepositMaxEntries,
	}

	k.Keeper.RegisterZone(ctx, zoneInfo)

	err := ctx.EventManager().EmitTypedEvent(types.NewEventChangeRegisterZone(zoneInfo))
	if err != nil {
		return nil, err
	}

	return &types.MsgChangeRegisteredZoneResponse{
		ZoneId:               zoneInfo.ZoneId,
		IcaInfo:              zoneInfo.IcaConnectionInfo,
		TransferInfo:         zoneInfo.TransferInfo,
		ValidatorAddress:     zoneInfo.ValidatorAddress,
		BaseDenom:            zoneInfo.BaseDenom,
		SnDenom:              zoneInfo.BaseDenom,
		Decimal:              zoneInfo.Decimal,
		UndelegateMaxEntries: zoneInfo.UndelegateMaxEntries,
		DepositMaxEntries:    zoneInfo.DepositMaxEntries,
	}, nil
}

// IcaDelegate implements the Msg/MsgIcaDelegate interface
func (k msgServer) IcaDelegate(goCtx context.Context, msg *types.MsgIcaDelegate) (*types.MsgIcaDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	// check unreceived ack
	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaDelegate", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaDelegate", "packetSequence", packetSeq, "commitment", commitment)

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaDelegate transaction failed to send")
	}

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaDelegate(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaDelegateResponse{}, nil
}

// IcaUndelegate implements the Msg/MsgIcaUndelegate interface
func (k msgServer) IcaUndelegate(goCtx context.Context, msg *types.MsgIcaUndelegate) (*types.MsgIcaUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaUndelegate", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaUndelegate", "packetSequence", packetSeq, "commitment", commitment)

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})
	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaUndelegate(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaUndelegateResponse{}, nil
}

// IcaAutoStaking implements the Msg/MsgIcaAutoStaking interface
func (k msgServer) IcaAutoStaking(goCtx context.Context, msg *types.MsgIcaAutoStaking) (*types.MsgIcaAutoStakingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	if !k.IsValidAutoStakingVersion(ctx, msg.ZoneId, msg.Version) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, strconv.FormatUint(msg.Version, 10))
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	// check unreceived ack
	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaAutoStaking", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaAutoStaking", "packetSequence", packetSeq, "commitment", commitment)

	var msgs []sdk.Msg

	msgs = append(msgs, &distributiontype.MsgWithdrawDelegatorReward{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress})
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: msg.Amount})

	err := k.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("IcaAutoStaking transaction failed to send")
	}

	// version state 변경
	versionInfo := k.GetAutoStakingVersion(ctx, zoneInfo.ZoneId)
	versionInfo.Record[versionInfo.CurrentVersion] = &types.IBCTrace{
		Version: versionInfo.CurrentVersion,
		State:   types.IcaRequest,
	}

	k.SetAutoStakingVersion(ctx, zoneInfo.ZoneId, versionInfo)

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaAutoStaking(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaAutoStakingResponse{}, nil
}

// IcaTransfer implements the Msg/MsgIcaTransferResponse interface
func (k msgServer) IcaTransfer(goCtx context.Context, msg *types.MsgIcaTransfer) (*types.MsgIcaTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	// check unreceived ack
	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaTransfer", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaTransfer", "packetSequence", packetSeq, "commitment", commitment)

	var msgs []sdk.Msg

	//transfer msg
	//sourceport, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    msg.IcaTransferPortId,
		SourceChannel: msg.IcaTransferChannelId,
		Token:         msg.Amount,
		Sender:        zoneInfo.IcaAccount.HostAddress,
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

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaTransfer(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaTransferResponse{}, nil
}

func (k msgServer) IcaAuthzGrant(goCtx context.Context, msg *types.MsgIcaAuthzGrant) (*types.MsgIcaAuthzGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	// check unreceived ack
	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaAuthzGrant", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaAuthzGrant", "packetSequence", packetSeq, "commitment", commitment)

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

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaAuthzGrant(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaAuthzGrantResponse{}, nil
}

func (k msgServer) IcaAuthzRevoke(goCtx context.Context, msg *types.MsgIcaAuthzRevoke) (*types.MsgIcaAuthzRevokeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := k.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	// check unreceived ack
	packetSeq, _ := k.Keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	commitment := k.Keeper.channelKeeper.GetPacketCommitment(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId,  zoneInfo.IcaConnectionInfo.ChannelId, packetSeq-1)
	if len(commitment) != 0 {
		ctx.Logger().Error("IcaAuthzRevoke", "packetSequence", packetSeq, "commitment", commitment)
	}
	ctx.Logger().Info("IcaAuthzRevoke", "packetSequence", packetSeq, "commitment", commitment)

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

	err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaAuthzRevoke(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgIcaAuthzRevokeResponse{}, nil
}

func (k msgServer) RegisterControllerAddress(goCtx context.Context, msg *types.MsgRegisterControllerAddr) (*types.MsgRegisterControllerAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsValidKeyManager(ctx, msg.FromAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.FromAddress)
	}

	controllerInfo := k.GetControllerAddr(ctx, msg.ZoneId)
	controllerAddrs := controllerInfo.ControllerAddress

	controllerAddrs = append(controllerAddrs, msg.ControllerAddress)
	k.SetControllerAddr(ctx, msg.ZoneId, controllerAddrs)

	err := ctx.EventManager().EmitTypedEvent(types.NewEventRegisterControllerAddress(msg))
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterControllerAddrResponse{}, nil
}
