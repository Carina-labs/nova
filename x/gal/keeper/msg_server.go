package keeper

import (
	"context"
	"github.com/Carina-labs/novachain/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func (m msgServer) Deposit(goCtx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.keeper.DepositCoin(ctx, deposit.Depositor,
		deposit.Receiver, "transfer", "channel-0", deposit.Amount); err != nil {
		return nil, err
	}
	return &types.MsgDepositResponse{}, nil
}

func (m msgServer) Withdraw(goCtx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.keeper.WithdrawNovaToken(ctx, withdraw.Withdrawer, withdraw.Amount); err != nil {
		return nil, err
	}
	return &types.MsgWithdrawResponse{}, nil
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{keeper: keeper}
}
