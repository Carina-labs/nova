package keeper

import (
	"context"
	"fmt"
	"github.com/Carina-labs/nova/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

func (m msgServer) CreatePool(goCtx context.Context, pool *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.keeper.CreatePool(ctx, pool.PoolId, pool.PoolAddress)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePoolResponse{}, nil
}

func (m msgServer) SetPoolWeight(goCtx context.Context, msg *types.MsgSetPoolWeight) (*types.MsgSetPoolWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ok := m.keeper.isValidOperator(msg)
	if !ok {
		return nil, fmt.Errorf("invalid controller address: %s", msg.Operator)
	}
	err := m.keeper.SetPoolWeight(ctx, msg.PoolId, msg.NewWeight)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetPoolWeightResponse{}, nil
}
