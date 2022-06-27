package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

// Deposit handles deposit action.
// 1. User submits deposit tx.
// 2. User's asset is transferred to the module(gal) account.
// 3. After IBC transfer, GAL coins deposit info.
func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Deposit(ctx, deposit)

	if err != nil {
		return nil, err
	}

	return &types.MsgDepositResponse{}, nil
}

// UndelegateRecord is used when user requests undelegate their staked asset.
// 1. User sends their st-token to module account.
// 2. And GAL coins step 1 to the store.
func (m msgServer) UndelegateRecord(goCtx context.Context, undelegate *types.MsgUndelegateRecord) (*types.MsgUndelegateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// snAtom -> [GAL] -> wAtom
	// change undelegate State
	_, found := m.keeper.interTxKeeper.GetRegisteredZone(ctx, undelegate.ZoneId)
	if !found {
		return nil, errors.New("zone not found")
	}

	//send stAsset to GAL moduleAccount
	//snBalance := m.keeper.bankKeeper.GetBalance(
	//	ctx, undelegate.Depositor, "snNova")
	//fmt.Printf("sn balance : %s\n", snBalance.String())
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
	})

	return &types.MsgUndelegateRecordResponse{}, nil
}

// Undelegate used when protocol requests undelegate to the host chain.
// 1. Protocol refers the store that contains user's undelegate request history.
// 2. Using it, controller chain requests undelegate staked asset using ICA.
// 3. And burn share token Module account have.
func (m msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zoneInfo, ok := m.keeper.interTxKeeper.GetRegisteredZone(ctx, msg.ZoneId)
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

	totalRequestAmt := sdk.Coin{
		Amount: totalStAsset.Amount,
		Denom:  zoneInfo.BaseDenom,
	}

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: totalRequestAmt})
	err := m.keeper.interTxKeeper.SendIcaTx(ctx, zoneInfo.IcaAccount.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgUndelegateResponse{}, nil
}

// WithdrawRecord write user's withdraw requests to the "WithdrawRecord" store.
// It will be used after undelegating, distribute native coin to the user.
func (m msgServer) WithdrawRecord(goCtx context.Context, withdraw *types.MsgWithdrawRecord) (*types.MsgWithdrawRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// withdraw record 조회
	withdrawRecord, found := m.keeper.GetWithdrawRecord(ctx, withdraw.ZoneId+withdraw.Withdrawer)
	if !found {
		return nil, errors.New("withdraw record is not found")
	}

	withdrawState := &types.WithdrawRecord{
		ZoneId:         withdrawRecord.ZoneId,
		Withdrawer:     withdrawRecord.Withdrawer,
		Recipient:      withdrawRecord.Recipient,
		Amount:         withdrawRecord.Amount,
		State:          int64(WITHDRAW_REQUEST_USER),
		CompletionTime: withdrawRecord.CompletionTime,
	}

	// state 변경
	m.keeper.SetWithdrawRecord(ctx, *withdrawState)

	// zone 정보 조회
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
		ctx.EventManager().EmitTypedEvent(&zoneInfo)
		ctx.EventManager().EmitTypedEvent(&withdrawRecord)

		//ICA transfer 요청(Bot)
		//transfer msg
		// var msgs []sdk.Msg

		// msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		// 	SourcePort:    zoneInfo.TransferConnectionInfo.PortId,
		// 	SourceChannel: zoneInfo.TransferConnectionInfo.ChannelId,
		// 	Token:         *withdrawRecord.Amount,
		// 	Sender:        zoneInfo.IcaConnectionInfo.OwnerAddress,
		// 	Receiver:      string(m.keeper.accountKeeper.GetModuleAddress(types.ModuleName)),
		// 	TimeoutHeight: ibcclienttypes.Height{
		// 		RevisionHeight: 0,
		// 		RevisionNumber: 0,
		// 	},
		// 	TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		// })

		// err := m.keeper.interTxKeeper.SendIcaTx(ctx, zoneInfo.IcaConnectionInfo.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)
		// if err != nil {
		// 	return &types.MsgWithdrawRecordResponse{}, errors.New("IcaWithdraw transaction failed to send")
		// }
	}

	// moduleAccountToAccount
	recipientAcc, err := sdk.AccAddressFromBech32(withdraw.Recipient)
	if err != nil {
		return nil, err
	}
	if err := m.keeper.ClaimWithdrawAsset(ctx, recipientAcc, withdraw.Amount); err != nil {
		return nil, err
	}

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

func (m msgServer) Claim(goCtx context.Context, claimMsg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	claimerAddr, err := sdk.AccAddressFromBech32(claimMsg.Claimer)
	if err != nil {
		return nil, err
	}

	record, err := m.keeper.GetRecordedDepositAmt(ctx, claimerAddr)
	if err != nil {
		return nil, err
	}

	for _, record := range record.Records {
		if record.IsTransferred && record.Amount.Equal(claimMsg.Amount) {
			err = m.keeper.ClaimAndMintShareToken(ctx, claimerAddr, *record.Amount)
			if err != nil {
				return nil, err
			}

			return &types.MsgClaimResponse{}, nil
		}
	}

	return nil, fmt.Errorf("can't find deposit record. address: %s, amount: %s",
		claimMsg.Amount, claimMsg.Amount.String())
}
