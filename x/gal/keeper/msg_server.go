package keeper

import (
	"context"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	depositorAddr, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return nil, err
	}

	if err := m.keeper.DepositCoin(ctx, depositorAddr,
		deposit.Receiver, "transfer", deposit.Channel, deposit.Amount); err != nil {
		return nil, err
	}
	
	return &types.MsgDepositResponse{}, nil
}

func (m msgServer) Withdraw(goCtx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	withdrawerAddr, err := sdk.AccAddressFromBech32(withdraw.Withdrawer)
	if err != nil {
		return nil, err
	}

	if err := m.keeper.WithdrawCoin(ctx, withdrawerAddr, withdraw.Amount); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
