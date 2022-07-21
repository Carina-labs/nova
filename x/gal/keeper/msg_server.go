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

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the ibcstaking Msg service interface
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

// Deposit handles deposit action.
// 1. User submits deposit tx.
// 2. User's asset is transferred to the module(gal) account.
// 3. After IBC transfer, GAL coins deposit info.
func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, deposit.ZoneId)
	if !ok {
		return nil, fmt.Errorf("can't find valid IBC zone, input zoneId: %s", deposit.ZoneId)
	}
	depositorAcc, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return nil, err
	}

	oracleVersion, err := m.keeper.oracleKeeper.GetOracleVersion(ctx, zoneInfo.BaseDenom)
	if err != nil {
		return nil, err
	}

	newRecord := &types.DepositRecordContent{
		ZoneId:        zoneInfo.ZoneId,
		Amount:        &deposit.Amount,
		IsTransferred: false,
		BlockHeight:   oracleVersion,
	}

	record, err := m.keeper.GetRecordedDepositAmt(ctx, depositorAcc)
	if err == types.ErrNoDepositRecord {
		m.keeper.SetDepositAmt(ctx, &types.DepositRecord{
			Address: deposit.Depositor,
			Records: []*types.DepositRecordContent{newRecord},
		})
	} else {
		record.Records = append(record.Records, newRecord)
		m.keeper.SetDepositAmt(ctx, record)
	}

	err = m.keeper.TransferToTargetZone(ctx,
		deposit.TransferPortId,
		deposit.TransferChannelId,
		deposit.Depositor,
		zoneInfo.IcaAccount.HostAddress,
		deposit.Amount)

	if err != nil {
		return nil, err
	}

	return &types.MsgDepositResponse{
		Receiver:        deposit.Depositor,
		DepositedAmount: deposit.Amount,
	}, nil
}

// UndelegateRecord is used when user requests undelegate their staked asset.
// 1. User sends their st-token to module account.
// 2. And GAL coins step 1 to the store.
func (m msgServer) UndelegateRecord(goCtx context.Context, undelegate *types.MsgUndelegateRecord) (*types.MsgUndelegateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// change undelegate State
	_, found := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, undelegate.ZoneId)
	if !found {
		return nil, errors.New("zone not found")
	}

	//send stAsset to GAL moduleAccount
	depositorAcc, err := sdk.AccAddressFromBech32(undelegate.Depositor)
	if err != nil {
		return nil, err
	}

	err = m.keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, depositorAcc, types.ModuleName, sdk.Coins{undelegate.Amount})
	if err != nil {
		return nil, err
	}

	undelegateInfo, found := m.keeper.GetUndelegateRecord(ctx, undelegate.ZoneId+undelegate.Depositor)
	if found {
		undelegate.Amount = undelegate.Amount.Add(*undelegateInfo.Amount)
	}

	amt := sdk.NewCoin(undelegate.Amount.Denom, undelegate.Amount.Amount)
	m.keeper.SetUndelegateRecord(ctx, types.UndelegateRecord{
		ZoneId:    undelegate.ZoneId,
		Delegator: undelegate.Depositor,
		Amount:    &amt,
		State:     int64(UNDELEGATE_REQUEST_USER),
	})

	return &types.MsgUndelegateRecordResponse{
		ZoneId: undelegate.ZoneId,
		User:   undelegate.Depositor,
		Amount: amt,
	}, nil
}

// Undelegate used when protocol requests undelegate to the host chain.
// 1. Protocol refers the store that contains user's undelegate request history.
// 2. Using it, controller chain requests undelegate staked asset using ICA.
// 3. And burn share token Module account have.
func (m msgServer) GalUndelegate(goCtx context.Context, msg *types.MsgGalUndelegate) (*types.MsgGalUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.keeper.ibcstakingKeeper.IsValidDaoModifier(ctx, msg.ControllerAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	zoneInfo, ok := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone is not found")
	}

	m.keeper.ChangeUndelegateState(ctx, zoneInfo.ZoneId, UNDELEGATE_REQUEST_ICA)
	totalStAsset := m.keeper.GetUndelegateAmount(ctx, zoneInfo.SnDenom, zoneInfo.ZoneId, UNDELEGATE_REQUEST_ICA)
	totalStAsset.Denom = zoneInfo.SnDenom

	if totalStAsset.Amount.Equal(sdk.NewInt(0)) {
		// TODO: should handle if no coins to undelegate
		return nil, errors.New("no coins to undelegate")
	}

	if err := m.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName,
		sdk.Coins{sdk.Coin{Denom: totalStAsset.Denom, Amount: totalStAsset.Amount}}); err != nil {
		return nil, err
	}

	m.keeper.SetWithdrawRecords(ctx, msg.ZoneId, UNDELEGATE_REQUEST_ICA)

	var msgs []sdk.Msg
	undelegateAmount, err := m.keeper.GetWithdrawAmt(ctx, totalStAsset)
	if err != nil {
		return nil, err
	}
	msgs = append(msgs, &stakingtype.MsgUndelegate{
		DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
		ValidatorAddress: zoneInfo.ValidatorAddress,
		Amount:           undelegateAmount})
	err = m.keeper.ibcstakingKeeper.SendIcaTx(ctx, zoneInfo.IcaAccount.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgGalUndelegateResponse{
		ZoneId:            zoneInfo.ZoneId,
		UndelegatedAmount: undelegateAmount,
	}, nil
}

// Withdraw write user's withdraw requests to the "Withdraw" store.
// It will be used after undelegate, distribute native coin to the user.
func (m msgServer) Withdraw(goCtx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, withdraw.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", withdraw.ZoneId)
	}

	withdrawRecord, err := m.keeper.GetWithdrawRecord(ctx, withdraw.ZoneId, withdraw.Withdrawer)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "account: %s", withdrawRecord.Withdrawer)
	}

	if withdrawRecord.State != int64(TRANSFER_SUCCESS) {
		return nil, sdkerrors.Wrapf(types.ErrCanNotWithdrawAsset, "unbonding time has not passed %s", withdrawRecord.CompletionTime)
	}

	ibcDenom := m.keeper.ibcstakingKeeper.GetIBCHashDenom(ctx, withdraw.TransferPortId, withdraw.TransferChannelId, zoneInfo.BaseDenom)
	withdrawAmount := sdk.NewInt64Coin(ibcDenom, withdrawRecord.Amount.Amount.Int64())
	controllerAddr, err := sdk.AccAddressFromBech32(zoneInfo.IcaAccount.DaomodifierAddress)
	if err != nil {
		return nil, err
	}

	recipientAcc, err := sdk.AccAddressFromBech32(withdraw.Recipient)
	if err != nil {
		return nil, err
	}

	if err := m.keeper.ClaimWithdrawAsset(ctx, controllerAddr, recipientAcc, withdrawAmount); err != nil {
		return nil, err
	}

	m.keeper.DeleteWithdrawRecord(ctx, *withdrawRecord)

	return &types.MsgWithdrawResponse{
		Withdrawer:     withdraw.Withdrawer,
		WithdrawAmount: withdrawAmount,
	}, nil
}

func (m msgServer) PendingWithdraw(goCtx context.Context, msg *types.MsgPendingWithdraw) (*types.MsgPendingWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.keeper.ibcstakingKeeper.IsValidDaoModifier(ctx, msg.DaomodifierAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	zoneInfo, ok := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone name is not found")
	}

	withdrawAmount := m.keeper.GetTotalWithdrawAmountForZoneId(ctx, msg.ZoneId, msg.ChainTime)
	if withdrawAmount.IsZero() {
		return nil, sdkerrors.Wrapf(types.ErrCanNotWithdrawAsset, "total widraw amount: %s", withdrawAmount)
	}

	var msgs []sdk.Msg

	//transfer msg
	//sourceport, Source channel, Token, Sender, receiver, TimeoutHeight, TimeoutTimestamp
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:    msg.TransferPortId,
		SourceChannel: msg.TransferChannelId,
		Token:         withdrawAmount,
		Sender:        zoneInfo.IcaAccount.HostAddress,
		Receiver:      msg.DaomodifierAddress,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	})

	err := m.keeper.ibcstakingKeeper.SendIcaTx(ctx, msg.DaomodifierAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
	if err != nil {
		return nil, errors.New("PendingWithdraw transaction failed to send")
	}

	return &types.MsgPendingWithdrawResponse{}, nil
}

func (m msgServer) ClaimSnAsset(goCtx context.Context, claimMsg *types.MsgClaimSnAsset) (*types.MsgClaimSnAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	claimerAddr, err := sdk.AccAddressFromBech32(claimMsg.Claimer)
	if err != nil {
		return nil, err
	}
	// TODO : add key
	record, err := m.keeper.GetRecordedDepositAmt(ctx, claimerAddr)
	if err != nil {
		return nil, err
	}

	for _, record := range record.Records {
		zone, ok := m.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, record.ZoneId)
		if !ok {
			return nil, fmt.Errorf("cannot find zone id : %s", record.ZoneId)
		}

		oracleVersion, err := m.keeper.oracleKeeper.GetOracleVersion(ctx, zone.BaseDenom)
		if err != nil {
			return nil, err
		}
		if record.BlockHeight >= oracleVersion {
			return nil, fmt.Errorf("oracle is not updated. current oracle version: %d", oracleVersion)
		}

		if record.IsTransferred && record.Amount.Equal(record.Amount) {
			minted, err := m.keeper.ClaimAndMintShareToken(ctx, claimerAddr, *record.Amount)
			if err != nil {
				return nil, err
			}

			// TODO: Delete deposit record
			return &types.MsgClaimSnAssetResponse{
				Claimer: claimMsg.Claimer,
				Minted:  minted,
			}, nil
		}
	}

	return nil, sdkerrors.Wrapf(types.ErrNoDepositRecord,
		"account: %s", claimMsg.Claimer)
}
