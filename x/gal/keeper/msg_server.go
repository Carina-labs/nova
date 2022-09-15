package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
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
		Depositor: depositorAcc.String(),
		Amount:    &deposit.Amount,
		State:     types.DepositRequest,
	}

	record, found := m.keeper.GetUserDepositRecord(ctx, zoneInfo.ZoneId, claimAcc)

	if !found {
		m.keeper.SetDepositRecord(ctx, &types.DepositRecord{
			ZoneId:  deposit.ZoneId,
			Claimer: deposit.Claimer,
			Records: []*types.DepositRecordContent{newRecord},
		})
	} else {
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

	if !m.keeper.icaControlKeeper.IsValidDaoModifier(ctx, delegate.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, delegate.ControllerAddress)
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, delegate.ZoneId)

	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}
	ok = m.keeper.ChangeDepositState(ctx, zoneInfo.ZoneId, types.DepositSuccess, types.DelegateRequest)
	if !ok {
		return nil, types.ErrCanNotChangeState
	}

	ibcDenom := m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	delegateAmt := m.keeper.GetTotalDepositAmtForZoneId(ctx, delegate.ZoneId, ibcDenom, types.DelegateRequest)
	delegateAmt.Denom = zoneInfo.BaseDenom

	var msgs []sdk.Msg
	msgs = append(msgs, &stakingtype.MsgDelegate{DelegatorAddress: zoneInfo.IcaAccount.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: delegateAmt})

	err := m.keeper.icaControlKeeper.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, types.ErrDelegateFail
	}

	if err := ctx.EventManager().EmitTypedEvent(
		types.NewEventDelegate(
			zoneInfo.IcaAccount.HostAddress,
			zoneInfo.ValidatorAddress,
			&delegateAmt,
			zoneInfo.TransferInfo.ChannelId,
			zoneInfo.TransferInfo.PortId)); err != nil {
		return nil, err
	}

	return &types.MsgDelegateResponse{}, nil
}

// PendingUndelegate is used when user requests undelegate their staked asset.
// 1. User sends their st-token to module account.
// 2. And GAL coins step 1 to the store.
func (m msgServer) PendingUndelegate(goCtx context.Context, undelegate *types.MsgPendingUndelegate) (*types.MsgPendingUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// change undelegate State
	zoneInfo, found := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, undelegate.ZoneId)
	if !found {
		return nil, types.ErrNotFoundZoneInfo
	}
	if zoneInfo.SnDenom != undelegate.Amount.Denom {
		return nil, types.ErrInvalidDenom
	}

	//send stAsset to GAL moduleAccount
	delegatorAcc, err := sdk.AccAddressFromBech32(undelegate.Delegator)
	if err != nil {
		return nil, err
	}

	oracleVersion := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	newRecord := types.UndelegateRecordContent{
		Withdrawer:    undelegate.Withdrawer,
		SnAssetAmount: &undelegate.Amount,
		State:         types.UndelegateRequestByUser,
		OracleVersion: oracleVersion,
	}

	snAssetAmt := sdk.NewCoin(undelegate.Amount.Denom, undelegate.Amount.Amount)

	undelegateRecord, found := m.keeper.GetUndelegateRecord(ctx, undelegate.ZoneId, undelegate.Delegator)
	if !found {
		undelegateRecord = &types.UndelegateRecord{
			ZoneId:    undelegate.ZoneId,
			Delegator: undelegate.Delegator,
			Records:   []*types.UndelegateRecordContent{&newRecord},
		}
	} else {
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

	if !m.keeper.icaControlKeeper.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone is not found")
	}

	oracleVersion := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)
	burnAssets, undelegateAssets := m.keeper.GetUndelegateAmount(ctx, zoneInfo.SnDenom, zoneInfo, oracleVersion)
	if burnAssets.IsZero() || undelegateAssets.IsZero() {
		return nil, errors.New("no coins to undelegate")
	}

	m.keeper.ChangeUndelegateState(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)

	undelegateAmt := sdk.NewCoin(zoneInfo.BaseDenom, undelegateAssets)

	var msgs []sdk.Msg
	msgs = append(msgs, &stakingtype.MsgUndelegate{
		DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
		ValidatorAddress: zoneInfo.ValidatorAddress,
		Amount:           undelegateAmt})
	err := m.keeper.icaControlKeeper.SendTx(ctx, zoneInfo.IcaConnectionInfo.PortId, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	if err := m.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName,
		sdk.Coins{sdk.Coin{Denom: burnAssets.Denom, Amount: burnAssets.Amount}}); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(
		types.NewEventUndelegate(zoneInfo.ZoneId, &burnAssets, &undelegateAmt)); err != nil {
		return nil, err
	}

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

	controllerAddr, err := sdk.AccAddressFromBech32(zoneInfo.IcaAccount.ControllerAddress)
	if err != nil {
		return nil, err
	}

	// sum of all withdraw records for user
	withdrawAmt := m.keeper.GetWithdrawAmountForUser(ctx, zoneInfo.ZoneId, ibcDenom, withdraw.Withdrawer)
	if withdrawAmt.IsZero() {
		return nil, types.ErrNoWithdrawRecord
	}

	withdrawerAddr, err := sdk.AccAddressFromBech32(withdraw.Withdrawer)
	if err != nil {
		return nil, err
	}

	if err := m.keeper.ClaimWithdrawAsset(ctx, controllerAddr, withdrawerAddr, withdrawAmt); err != nil {
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

	if !m.keeper.icaControlKeeper.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	withdrawAmount := m.keeper.GetTotalWithdrawAmountForZoneId(ctx, msg.ZoneId, zoneInfo.BaseDenom, msg.ChainTime)
	if withdrawAmount.IsZero() {
		return nil, sdkerrors.Wrapf(types.ErrCanNotWithdrawAsset, "total widraw amount: %s", withdrawAmount)
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
		return nil, errors.New("PendingWithdraw transaction failed to send")
	}

	if err = ctx.EventManager().EmitTypedEvent(types.NewEventIcaWithdraw(
		zoneInfo.IcaAccount.HostAddress,
		zoneInfo.IcaAccount.ControllerAddress,
		&withdrawAmount,
		zoneInfo.IcaConnectionInfo.ConnectionId,
		msg.IcaTransferChannelId,
		msg.IcaTransferPortId)); err != nil {
		return nil, err
	}

	return &types.MsgIcaWithdrawResponse{}, nil
}

func (m msgServer) ClaimSnAsset(goCtx context.Context, claimMsg *types.MsgClaimSnAsset) (*types.MsgClaimSnAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	claimerAddr, err := sdk.AccAddressFromBech32(claimMsg.Claimer)
	if err != nil {
		return nil, err
	}

	records, found := m.keeper.GetUserDepositRecord(ctx, claimMsg.ZoneId, claimerAddr)
	if !found {
		return nil, types.ErrNoDepositRecord
	}

	zoneInfo, ok := m.keeper.icaControlKeeper.GetRegisteredZone(ctx, records.ZoneId)
	if !ok {
		return nil, fmt.Errorf("cannot find zone id : %s", records.ZoneId)
	}

	ibcDenom := m.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	totalClaimAsset := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  ibcDenom,
	}

	oracleVersion := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)
	for _, record := range records.Records {
		if record.OracleVersion >= oracleVersion {
			return nil, fmt.Errorf("oracle is not updated. current oracle version: %d", oracleVersion)
		}
		if record.State == types.DelegateSuccess {
			totalClaimAsset = totalClaimAsset.Add(*record.Amount)
		}
	}

	claimSnAsset, err := m.keeper.ClaimShareToken(ctx, &zoneInfo, totalClaimAsset)
	if err != nil {
		return nil, sdkerrors.Wrapf(err,
			"account: %s", claimMsg.Claimer)
	}

	err = m.keeper.MintTo(ctx, claimerAddr, claimSnAsset)
	if err != nil {
		return nil, sdkerrors.Wrapf(err,
			"account: %s", claimMsg.Claimer)
	}

	err = m.keeper.DeleteRecordedDepositItem(ctx, zoneInfo.ZoneId, claimerAddr, types.DelegateSuccess)
	if err != nil {
		return nil, sdkerrors.Wrapf(err,
			"account: %s", claimMsg.Claimer)
	}

	// mark user performed claim action
	//m.keeper.airdropKeeper.PostClaimedSnAsset(ctx, claimerAddr)
	if err = ctx.EventManager().EmitTypedEvent(
		types.NewEventClaimSnToken(claimMsg.Claimer, &claimSnAsset, oracleVersion)); err != nil {
		return nil, err
	}

	return &types.MsgClaimSnAssetResponse{
		Claimer: claimMsg.Claimer,
		Minted:  claimSnAsset,
	}, nil
}
