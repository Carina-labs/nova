package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the icacontrol Msg service interface
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

// Deposit handles deposit action.
// 1. User submits deposit tx.
// 2. User's asset is transferred to the module(gal) account.
// 3. After IBC transfer, GAL coins deposit info.
func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, deposit.ZoneId)
	if !ok {
		return nil, fmt.Errorf("can't find valid IBC zone, input zoneId: %s", deposit.ZoneId)
	}

	depositorAcc, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return nil, err
	}

	claimAcc, err := sdk.AccAddressFromBech32(deposit.Claimer)
	if err != nil {
		return nil, err
	}

	// check IBC denom
	if deposit.Amount.Denom != m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom) {
		return nil, types.ErrInvalidDenom
	}

	newRecord := &types.DepositRecordContent{
		Claimer: claimAcc.String(),
		Amount:  &deposit.Amount,
		State:   types.DepositRequest,
	}

	record, found := m.keeper.GetUserDepositRecord(ctx, zoneInfo.ZoneId, depositorAcc)

	if !found {
		m.keeper.SetDepositRecord(ctx, &types.DepositRecord{
			ZoneId:    deposit.ZoneId,
			Depositor: deposit.Claimer,
			Records:   []*types.DepositRecordContent{newRecord},
		})
	} else {
		if len(record.Records) >= int(zoneInfo.DepositMaxEntries) &&
			m.keeper.HasMaxDepositEntries(*record, zoneInfo.DepositMaxEntries) {
			return nil, types.ErrMaxDepositEntries
		}

		record.Records = append(record.Records, newRecord)
		m.keeper.SetDepositRecord(ctx, record)
	}

	err = m.keeper.TransferToTargetZone(ctx, &IBCTransferOption{
		SourcePort:    zoneInfo.TransferInfo.PortId,
		SourceChannel: zoneInfo.TransferInfo.ChannelId,
		Token:         deposit.Amount,
		Sender:        deposit.Depositor,
		Receiver:      zoneInfo.IcaAccount.HostAddress,
	})

	if err != nil {
		return nil, err
	}

	if err = ctx.EventManager().EmitTypedEvent(
		types.NewEventDeposit(deposit.Depositor, deposit.Claimer, &deposit.Amount)); err != nil {
		return nil, err
	}

	return &types.MsgDepositResponse{
		Depositor:       deposit.Depositor,
		Receiver:        deposit.Claimer,
		DepositedAmount: deposit.Amount,
	}, nil
}

func (m msgServer) Delegate(goCtx context.Context, delegate *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, delegate.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	if !m.keeper.icaControlKeeper.IsValidControllerAddr(ctx, delegate.ZoneId, delegate.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, delegate.ControllerAddress)
	}

	// check unreceived ack
	ackSeq, _ := m.keeper.channelKeeper.GetNextSequenceAck(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	packetSeq, _ := m.keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	if ackSeq != packetSeq {
		ctx.Logger().Error("Delegate", "packetSequence", packetSeq, "ackSequence", ackSeq)
		return nil, types.ErrInvalidAck
	}
	ctx.Logger().Info("Delegate", "packetSequence", packetSeq, "ackSequence", ackSeq)

	// version state check
	if !m.keeper.IsValidDelegateVersion(ctx, delegate.ZoneId, delegate.Version) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, strconv.FormatUint(delegate.Version, 10))
	}

	ibcDenom := m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)

	versionInfo := m.keeper.GetDelegateVersion(ctx, zoneInfo.ZoneId)
	version := versionInfo.Record[delegate.Version]

	if version.State == types.IcaPending {
		depoistAmt := m.keeper.GetTotalDepositAmtForZoneId(ctx, zoneInfo.ZoneId, ibcDenom, types.DepositSuccess)
		if depoistAmt.IsZero() {
			return nil, types.ErrNoDepositRecord
		}

		m.keeper.SetDelegateRecords(ctx, zoneInfo.ZoneId)
		m.keeper.DeleteDepositRecords(ctx, zoneInfo.ZoneId, types.DepositSuccess)
	}

	delegateAmt := m.keeper.GetTotalDelegateAmtForZoneId(ctx, delegate.ZoneId, ibcDenom, versionInfo.CurrentVersion, types.DelegateRequest)
	delegateAmt.Denom = zoneInfo.BaseDenom

	if delegateAmt.IsZero() {
		return nil, types.ErrInsufficientFunds
	}

	var msgs []sdk.Msg
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: delegateAmt})

	err := m.keeper.icaControlKeeper.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		versionInfo.Record[delegate.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetDelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, sdkerrors.Wrapf(err, "IcaDelegateFail")
	}

	if err = ctx.EventManager().EmitTypedEvent(
		types.NewEventDelegate(
			zoneInfo.IcaAccount.HostAddress,
			zoneInfo.ValidatorAddress,
			&delegateAmt,
			zoneInfo.TransferInfo.ChannelId,
			zoneInfo.TransferInfo.PortId)); err != nil {
		versionInfo.Record[delegate.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetDelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, err
	}

	versionInfo.Record[delegate.Version] = &types.IBCTrace{
		Version: versionInfo.CurrentVersion,
		State:   types.IcaRequest,
	}
	m.keeper.SetDelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
	return &types.MsgDelegateResponse{}, nil
}

// PendingUndelegate is used when user requests undelegate their staked asset.
// 1. User sends their st-token to module account.
// 2. And GAL coins step 1 to the store.
func (m msgServer) PendingUndelegate(goCtx context.Context, undelegate *types.MsgPendingUndelegate) (*types.MsgPendingUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, found := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, undelegate.ZoneId)
	if !found {
		return nil, types.ErrNotFoundZoneInfo
	}
	if zoneInfo.SnDenom != undelegate.Amount.Denom {
		return nil, types.ErrInvalidDenom
	}

	// check snAsset decimal
	if err := m.keeper.CheckDecimal(undelegate.Amount, zoneInfo.Decimal); err != nil {
		return nil, err
	}

	// check undelegate amount
	wAsset, err := m.keeper.GetWithdrawAmt(ctx, undelegate.Amount)
	if err != nil {
		return nil, err
	}

	if wAsset.IsZero() || wAsset.IsNil() {
		return nil, types.ErrConvertWAssetIsZero
	}

	delegatorAcc, err := sdk.AccAddressFromBech32(undelegate.Delegator)
	if err != nil {
		return nil, err
	}

	oracleVersion, _ := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	newRecord := types.UndelegateRecordContent{
		Withdrawer:    undelegate.Withdrawer,
		SnAssetAmount: &undelegate.Amount,
		State:         types.UndelegateRequestByUser,
		OracleVersion: oracleVersion,
	}

	snAssetAmt := sdk.NewCoin(undelegate.Amount.Denom, undelegate.Amount.Amount)

	//send stAsset to GAL moduleAccount
	undelegateRecord, found := m.keeper.GetUndelegateRecord(ctx, undelegate.ZoneId, undelegate.Delegator)
	if !found {
		undelegateRecord = &types.UndelegateRecord{
			ZoneId:    undelegate.ZoneId,
			Delegator: undelegate.Delegator,
			Records:   []*types.UndelegateRecordContent{&newRecord},
		}
	} else {
		if len(undelegateRecord.Records) >= int(zoneInfo.UndelegateMaxEntries) &&
			m.keeper.HasMaxUndelegateEntries(*undelegateRecord, zoneInfo.UndelegateMaxEntries) {
			return nil, types.ErrMaxUndelegateEntries
		}

		undelegateRecord.Records = append(undelegateRecord.Records, &newRecord)
	}

	m.keeper.SetUndelegateRecord(ctx, undelegateRecord)

	err = m.keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, delegatorAcc, types.ModuleName, sdk.Coins{undelegate.Amount})
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(
		types.NewEventPendingUndelegate(undelegate.ZoneId,
			undelegate.Delegator,
			undelegate.Withdrawer,
			&snAssetAmt,
			&undelegate.Amount)); err != nil {
		return nil, err
	}

	return &types.MsgPendingUndelegateResponse{
		ZoneId:     undelegate.ZoneId,
		Delegator:  undelegate.Delegator,
		Withdrawer: undelegate.Withdrawer,
		BurnAsset:  snAssetAmt,
	}, nil
}

// Undelegate used when protocol requests undelegate to the host chain.
// 1. Protocol refers the store that contains user's undelegate request history.
// 2. Using it, controller chain requests undelegate staked asset using ICA.
// 3. And burn share token Module account have.
func (m msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone is not found")
	}

	if !m.keeper.icaControlKeeper.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	if !m.keeper.IsValidUndelegateVersion(ctx, msg.ZoneId, msg.Version) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, strconv.FormatUint(msg.Version, 10))
	}

	// unreceived ack 확인
	ackSeq, _ := m.keeper.channelKeeper.GetNextSequenceAck(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	packetSeq, _ := m.keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	if ackSeq != packetSeq {
		ctx.Logger().Error("Undelegate", "packetSequence", packetSeq, "ackSequence", ackSeq)
		return nil, types.ErrInvalidAck
	}
	ctx.Logger().Info("Undelegate", "packetSequence", packetSeq, "ackSequence", ackSeq)

	var burnAssets sdk.Coin
	var undelegateAssets sdk.Int

	versionInfo := m.keeper.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	version := versionInfo.Record[msg.Version]
	oracleVersion, _ := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	if version.State == types.IcaPending {
		burnAssets, undelegateAssets = m.keeper.GetUndelegateAmount(ctx, zoneInfo.SnDenom, zoneInfo, oracleVersion)
		if burnAssets.IsZero() || undelegateAssets.IsZero() {
			return nil, errors.New("no coins to undelegate")
		}

		m.keeper.ChangeUndelegateState(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)

		if err := m.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName,
			sdk.Coins{sdk.Coin{Denom: burnAssets.Denom, Amount: burnAssets.Amount}}); err != nil {
			return nil, err
		}
	}

	if version.State == types.IcaFail {
		burnAssets, undelegateAssets = m.keeper.GetReUndelegateAmount(ctx, zoneInfo.SnDenom, zoneInfo, oracleVersion)
		if burnAssets.IsZero() || undelegateAssets.IsZero() {
			return nil, errors.New("no coins to undelegate")
		}
	}

	undelegateAmt := sdk.NewCoin(zoneInfo.BaseDenom, undelegateAssets)

	var msgs []sdk.Msg
	msgs = append(msgs, &stakingtype.MsgUndelegate{
		DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
		ValidatorAddress: zoneInfo.ValidatorAddress,
		Amount:           undelegateAmt})
	err := m.keeper.icaControlKeeper.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		versionInfo.Record[msg.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetUndelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, sdkerrors.Wrapf(err, "IcaUnDelegate transaction failed to send")
	}

	if err := ctx.EventManager().EmitTypedEvent(
		types.NewEventUndelegate(zoneInfo.ZoneId, &burnAssets, &undelegateAmt)); err != nil {
		versionInfo.Record[msg.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetUndelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, err
	}

	versionInfo.Record[msg.Version] = &types.IBCTrace{
		Version: versionInfo.CurrentVersion,
		State:   types.IcaRequest,
	}
	m.keeper.SetUndelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)

	return &types.MsgUndelegateResponse{
		ZoneId:               zoneInfo.ZoneId,
		TotalBurnAsset:       burnAssets,
		TotalUndelegateAsset: undelegateAmt,
	}, nil
}

// Withdraw write user's withdraw requests to the "Withdraw" store.
// It will be used after undelegate, distribute native coin to the user.
func (m msgServer) Withdraw(goCtx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, withdraw.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", withdraw.ZoneId)
	}

	ibcDenom := m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)

	// sum of all withdraw records for user
	withdrawAmt := m.keeper.GetWithdrawAmountForUser(ctx, zoneInfo.ZoneId, ibcDenom, withdraw.Withdrawer)
	if withdrawAmt.IsZero() {
		return nil, types.ErrNoWithdrawRecord
	}

	withdrawerAddr, err := sdk.AccAddressFromBech32(withdraw.Withdrawer)
	if err != nil {
		return nil, err
	}

	if err = m.keeper.ClaimWithdrawAsset(ctx, withdrawerAddr, withdrawAmt); err != nil {
		return nil, err
	}

	withdrawRecord, found := m.keeper.GetWithdrawRecord(ctx, zoneInfo.ZoneId, withdraw.Withdrawer)
	if !found {
		return nil, types.ErrNoWithdrawRecord
	}

	m.keeper.DeleteWithdrawRecord(ctx, withdrawRecord)

	if err := ctx.EventManager().EmitTypedEvent(
		types.NewEventWithdraw(zoneInfo.ZoneId, withdraw.Withdrawer, &withdrawAmt)); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{
		Withdrawer:     withdraw.Withdrawer,
		WithdrawAmount: withdrawAmt.Amount,
	}, nil
}

func (m msgServer) IcaWithdraw(goCtx context.Context, msg *types.MsgIcaWithdraw) (*types.MsgIcaWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.keeper.icaControlKeeper.IsValidControllerAddr(ctx, msg.ZoneId, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	if !m.keeper.IsValidWithdrawVersion(ctx, msg.ZoneId, msg.Version) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, strconv.FormatUint(msg.Version, 10))
	}

	// unreceived ack 확인
	ackSeq, _ := m.keeper.channelKeeper.GetNextSequenceAck(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	packetSeq, _ := m.keeper.channelKeeper.GetNextSequenceSend(ctx, icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ChannelId)
	if ackSeq != packetSeq {
		ctx.Logger().Error("IcaWithdraw", "packetSequence", packetSeq, "ackSequence", ackSeq)
		return nil, types.ErrInvalidAck
	}
	ctx.Logger().Info("IcaWithdraw", "packetSequence", packetSeq, "ackSequence", ackSeq)

	versionInfo := m.keeper.GetWithdrawVersion(ctx, zoneInfo.ZoneId)
	version := versionInfo.Record[msg.Version]

	var withdrawAmount sdk.Coin
	if version.State == types.IcaPending {
		withdrawAmount = m.keeper.GetTotalWithdrawAmountForZoneId(ctx, msg.ZoneId, zoneInfo.BaseDenom, msg.ChainTime)
	}

	if version.State == types.IcaFail {
		withdrawAmount = m.keeper.GetTotalWithdrawAmountForFailCase(ctx, msg.ZoneId, zoneInfo.BaseDenom, msg.ChainTime)
	}

	if withdrawAmount.IsZero() {
		return nil, sdkerrors.Wrapf(types.ErrCanNotWithdrawAsset, "total withdraw amount: %s", withdrawAmount)
	}

	var msgs []sdk.Msg

	//transfer msg
	//Source Port, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    msg.IcaTransferPortId,
		SourceChannel: msg.IcaTransferChannelId,
		Token:         withdrawAmount,
		Sender:        zoneInfo.IcaAccount.HostAddress,
		Receiver:      zoneInfo.IcaAccount.ControllerAddress,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	})

	err := m.keeper.icaControlKeeper.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		versionInfo.Record[msg.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetWithdrawVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, sdkerrors.Wrapf(err, "PendingWithdraw transaction failed to send")
	}

	if err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaWithdraw(
		zoneInfo.IcaAccount.HostAddress,
		zoneInfo.IcaAccount.ControllerAddress,
		&withdrawAmount,
		zoneInfo.IcaConnectionInfo.ConnectionId,
		msg.IcaTransferChannelId,
		msg.IcaTransferPortId)); err != nil {
		versionInfo.Record[msg.Version] = &types.IBCTrace{
			Version: versionInfo.CurrentVersion,
			State:   types.IcaFail,
		}
		m.keeper.SetWithdrawVersion(ctx, zoneInfo.ZoneId, versionInfo)
		return nil, err
	}

	versionInfo.Record[msg.Version] = &types.IBCTrace{
		Version: versionInfo.CurrentVersion,
		State:   types.IcaRequest,
	}
	m.keeper.SetWithdrawVersion(ctx, zoneInfo.ZoneId, versionInfo)

	return &types.MsgIcaWithdrawResponse{}, nil
}

func (m msgServer) ClaimSnAsset(goCtx context.Context, claimMsg *types.MsgClaimSnAsset) (*types.MsgClaimSnAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	claimerAddr, err := sdk.AccAddressFromBech32(claimMsg.Claimer)
	if err != nil {
		return nil, err
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, claimMsg.ZoneId)
	if !ok {
		return nil, fmt.Errorf("cannot find zone id : %s", claimMsg.ZoneId)
	}

	records, found := m.keeper.GetUserDelegateRecord(ctx, claimMsg.ZoneId, claimerAddr)
	if !found {
		return nil, types.ErrNoDelegateRecord
	}

	if records.Records == nil {
		return nil, types.ErrNoDelegateRecord
	}

	ibcDenom := m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	totalClaimAsset := sdk.NewCoin(ibcDenom, sdk.NewInt(0))
	ctx.Logger().Info("ClaimSnAsset", "totalClaimAsset", totalClaimAsset)

	oracleVersion, _ := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)
	for _, record := range records.Records {
		if record.OracleVersion >= oracleVersion {
			return nil, fmt.Errorf("oracle is not updated. current oracle version: %d", oracleVersion)
		}
		if record.State == types.DelegateSuccess {
			totalClaimAsset = totalClaimAsset.Add(*record.Amount)
		}
	}

	claimSnAsset, err := m.keeper.ClaimShareToken(ctx, &zoneInfo, totalClaimAsset)
	ctx.Logger().Info("ClaimSnAsset", "claimSnAsset", claimSnAsset)
	if err != nil {
		ctx.Logger().Error("ClaimSnAsset", "ClaimShareToken", err)
		return nil, sdkerrors.Wrapf(err,
			"account: %s", claimMsg.Claimer)
	}

	err = m.keeper.MintTo(ctx, claimerAddr, *claimSnAsset)
	if err != nil {
		ctx.Logger().Error("ClaimSnAsset", "MintTo", err)
		return nil, sdkerrors.Wrapf(err,
			"account: %s", claimMsg.Claimer)
	}

	delegateRecord, found := m.keeper.GetUserDelegateRecord(ctx, zoneInfo.ZoneId, claimerAddr)
	if !found {
		ctx.Logger().Error("ClaimSnAsset", "delegateRecord", found)
		return nil, types.ErrNoDelegateRecord
	}

	m.keeper.DeleteDelegateRecord(ctx, delegateRecord)

	// mark user performed claim action
	m.keeper.airdropKeeper.PostClaimedSnAsset(ctx, claimerAddr)
	if err = ctx.EventManager().EmitTypedEvent(
		types.NewEventClaimSnToken(claimMsg.Claimer, claimSnAsset, oracleVersion)); err != nil {
		return nil, err
	}

	return &types.MsgClaimSnAssetResponse{
		Claimer: claimMsg.Claimer,
		Minted:  *claimSnAsset,
	}, nil
}
