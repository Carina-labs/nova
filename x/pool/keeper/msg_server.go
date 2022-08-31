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

// CreatePool handles MsgCreatePool message, it creates a new pool
// with pool id and contract address of pool.
func (m msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool := &types.Pool{
		PoolId:              msg.PoolId,
		PoolContractAddress: msg.PoolContractAddress,
		Weight:              0,
	}
	err := m.keeper.CreatePool(ctx, pool)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePoolResponse{}, nil
}

// SetPoolWeight handles MsgSetPoolWeight message,
// it set a weight of pool, and it is used for calculating portion of distribution for newly minted nova token.
func (m msgServer) SetPoolWeight(goCtx context.Context, msg *types.MsgSetPoolWeight) (*types.MsgSetPoolWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ok := m.keeper.isValidOperator(ctx, msg.Operator)
	if !ok {
		return nil, fmt.Errorf("invalid controller address: %s", msg.Operator)
	}
	err := m.keeper.SetPoolWeight(ctx, msg.PoolId, msg.NewWeight)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetPoolWeightResponse{}, nil
}
