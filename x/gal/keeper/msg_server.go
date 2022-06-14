package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

// Deposit handles deposit action.
// 1. User submits deposit tx.
// 2. GAL module write it to "record" store.
// 3. User's asset is transferred to the module(gal) account.
func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.RecordDepositAmt(ctx, *deposit)
	if err != nil {
		return nil, err
	}

	// IBC transfer
	zoneInfo, ok := m.keeper.interTxKeeper.GetRegisteredZone(ctx, deposit.ZoneId)
	if !ok {
		return nil, err
	}

	err = m.keeper.TransferToTargetZone(ctx,
		zoneInfo.TransferConnectionInfo.PortId,
		zoneInfo.TransferConnectionInfo.ChannelId,
		deposit.Depositor,
		zoneInfo.IcaConnectionInfo.OwnerAddress,
		deposit.Amount[0])

	return &types.MsgDepositResponse{}, nil
}

//denom : stAsset
func (m msgServer) UndelegateRecord(goCtx context.Context, undelegate *types.MsgUndelegateRecord) (*types.MsgUndelegateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// snAtom -> [GAL] -> wAtom
	// change undelegate State
	zoneInfo, found := m.keeper.interTxKeeper.GetRegisteredZone(ctx, undelegate.ZoneId)
	if !found {
		return nil, errors.New("zone not found")
	}

	//send stAsset to GAL moduleAccount
	m.keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(undelegate.Depositor), types.ModuleName, sdk.Coins{undelegate.Amount})

	undelegateInfo, found := m.keeper.GetUndelegateRecord(ctx, undelegate.ZoneId+undelegate.Depositor)
	if found {
		undelegate.Amount = undelegate.Amount.Add(*undelegateInfo.Amount)
	}

	amt := &sdk.Coin{
		Denom:  zoneInfo.BaseDenom,
		Amount: undelegate.Amount.Amount,
	}

	record := &types.UndelegateRecord{
		ZoneId:    undelegate.ZoneId,
		Delegator: undelegate.Depositor,
		Amount:    amt,
	}

	m.keeper.SetUndelegateRecord(ctx, *record)

	return &types.MsgUndelegateRecordResponse{}, nil
}

func (m msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.interTxKeeper.GetRegisteredZone(ctx, msg.ZoneId)
	if !ok {
		return nil, errors.New("zone is not found")
	}

	m.keeper.ChangeUndelegateState(ctx, zoneInfo.ZoneName, UNDELEGATE_REQUEST_ICA)

	totalStAsset := m.keeper.GetUndelegateAmount(ctx, zoneInfo.BaseDenom, zoneInfo.ZoneName, UNDELEGATE_REQUEST_ICA)
	totalStAsset.Denom = zoneInfo.StDenom

	// burn stAsset, wAsset withdraw record에 저장 : 요청 할때 stAsset burn, wAsset 계산
	// if err := m.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName,
	// 	sdk.Coins{sdk.Coin{Denom: totalStAsset.Denom, Amount: totalStAsset.Amount}}); err != nil {
	// 	return nil, err
	// }

	// wAsset계산 + withdraw record 생성
	m.keeper.SetWithdrawRecords(ctx, msg.ZoneId, UNDELEGATE_REQUEST_ICA)

	var msgs []sdk.Msg

	totalRequestAmt := sdk.Coin{
		Amount: totalStAsset.Amount,
		Denom:  zoneInfo.BaseDenom,
	}

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: totalRequestAmt})
	err := m.keeper.interTxKeeper.SendIcaTx(ctx, zoneInfo.IcaConnectionInfo.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgUndelegateResponse{}, nil
}

// 사용자가 withdraw 요청
func (m msgServer) WithdrawRecord(goCtx context.Context, withdraw *types.MsgWithdrawRecord) (*types.MsgWithdrawRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//withdraw record 조회
	withdrawRecord, found := m.keeper.GetWithdrawRecord(ctx, withdraw.ZoneId+withdraw.Withdrawer)
	if !found {
		return nil, errors.New("withdraw record is not found")
	}

	withdrawState := &types.WithdrawRecord{
		ZoneId:         withdrawRecord.ZoneId,
		Withdrawer:     withdrawRecord.Withdrawer,
		Recipient:      withdrawRecord.Recipient,
		Amount:         withdrawRecord.Amount,
		State:          WITHDRAW_REQUEST_USER,
		CompletionTime: withdrawRecord.CompletionTime,
	}

	// state 변경
	m.keeper.SetWithdrawRecord(ctx, *withdrawState)

	//zone 정보 조회
	zoneInfo, ok := m.keeper.interTxKeeper.GetRegisteredZone(ctx, withdraw.ZoneId)
	if !ok {
		return nil, errors.New("zone is not found")
	}

	// module account의 상태 조회
	ok = m.keeper.IsAbleToWithdraw(ctx, *withdrawRecord.Amount)
	if !ok {
		return nil, fmt.Errorf("user cannot withdraw funds : insufficient fund")
	}

	if !ok {
		//ICA transfer 요청
		//transfer msg
		var msgs []sdk.Msg

		msgs = append(msgs, &ibctransfertypes.MsgTransfer{
			SourcePort:    zoneInfo.TransferConnectionInfo.PortId,
			SourceChannel: zoneInfo.TransferConnectionInfo.ChannelId,
			Token:         *withdrawRecord.Amount,
			Sender:        zoneInfo.IcaConnectionInfo.OwnerAddress,
			Receiver:      string(m.keeper.accountKeeper.GetModuleAddress(types.ModuleName)),
			TimeoutHeight: ibcclienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		})

		err := m.keeper.interTxKeeper.SendIcaTx(ctx, zoneInfo.IcaConnectionInfo.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
		if err != nil {
			return &types.MsgWithdrawRecordResponse{}, errors.New("IcaWithdraw transaction failed to send")
		}
	}

	// moduleAccountToAccount
	m.keeper.ClaimWithdrawAsset(ctx, withdraw.Recipient, withdraw.Amount)

	// withdrawRecord 삭제
	m.keeper.DeleteWithdrawRecord(ctx, *withdrawState)
	record := &types.WithdrawRecord{
		ZoneId:         withdraw.ZoneId,
		Withdrawer:     withdraw.Withdrawer,
		Amount:         &withdraw.Amount,
		CompletionTime: withdraw.Time,
	}

	m.keeper.SetWithdrawRecord(ctx, *record)

	return &types.MsgWithdrawRecordResponse{}, nil
}
