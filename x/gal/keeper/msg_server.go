package keeper

import (
	"context"
	"errors"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	depositorAddr, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return nil, err
	}

	receiverAddr, err := sdk.AccAddressFromBech32(deposit.Receiver)
	if err != nil {
		return nil, err
	}

	for _, coin := range deposit.Amount {
		err := m.keeper.DepositCoin(
			ctx, depositorAddr, receiverAddr, "transfer", deposit.Channel, coin)
		if err != nil {
			return nil, err
		}
	}

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
		return nil, errors.New("zone name is not found")
	}
	m.keeper.ChangeUndelegateState(ctx, zoneInfo.ZoneName, UNDELEGATE_REQUEST_ICA)

	amt := m.keeper.GetUndelegateAmount(ctx, zoneInfo.BaseDenom, zoneInfo.ZoneName, UNDELEGATE_REQUEST_ICA)

	// GetShareTokenMintingAmt(amt)
	// burn stAsset
	// if err := m.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName,
	// 	sdk.Coins{sdk.Coin{Denom: undelegate.Amount.Denom, Amount: undelegate.Amount.Amount}}); err != nil {
	// 	return nil, err
	// }

	var msgs []sdk.Msg

	msgs = append(msgs, &stakingtype.MsgUndelegate{DelegatorAddress: msg.HostAddress, ValidatorAddress: zoneInfo.ValidatorAddress, Amount: *amt})
	err := m.keeper.interTxKeeper.SendIcaTx(ctx, zoneInfo.IcaConnectionInfo.OwnerAddress, zoneInfo.IcaConnectionInfo.ConnectionId, msgs)

	if err != nil {
		return nil, errors.New("IcaUnDelegate transaction failed to send")
	}

	return &types.MsgUndelegateResponse{}, nil
}

func (m msgServer) WithdrawRecord(goCtx context.Context, withdraw *types.MsgWithdrawRecord) (*types.MsgWithdrawRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO

	record := &types.WithdrawRecord{
		ZoneId:         withdraw.ZoneId,
		Withdrawer:     withdraw.Withdrawer,
		Amount:         &withdraw.Amount,
		CompletionTime: withdraw.Time,
	}

	m.keeper.SetWithdrawRecord(ctx, *record)

	return &types.MsgWithdrawRecordResponse{}, nil
}
