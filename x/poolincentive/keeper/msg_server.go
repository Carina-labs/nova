package keeper

import (
	"context"
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
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

// CreateCandidatePool handles MsgCreateCandidatePool message, it creates a new candidate poolincentive
// with poolincentive id and contract address of poolincentive.
func (m msgServer) CreateCandidatePool(goCtx context.Context, msg *types.MsgCreateCandidatePool) (*types.MsgCreateCandidatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool := &types.CandidatePool{
		PoolId:              msg.PoolId,
		PoolContractAddress: msg.PoolContractAddress,
	}

	if err := m.keeper.CreateCandidatePool(ctx, pool); err != nil {
		return nil, err
	}

	return &types.MsgCreateCandidatePoolResponse{}, nil
}

// CreateIncentivePool handles MsgCreateIncentivePool message, it creates a new incentive poolincentive
// with poolincentive id and contract address of poolincentive.
// This message only can be handled by operator.
func (m msgServer) CreateIncentivePool(goCtx context.Context, msg *types.MsgCreateIncentivePool) (*types.MsgCreateIncentivePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ok := m.keeper.isValidOperator(ctx, msg.GetSigners()[0].String())
	if !ok {
		return nil, fmt.Errorf("invalid controller address: %s", msg.Operator)
	}
	pool := &types.IncentivePool{
		PoolId:              msg.PoolId,
		PoolContractAddress: msg.PoolContractAddress,
		Weight:              0,
	}

	if err := m.keeper.CreateIncentivePool(ctx, pool); err != nil {
		return nil, err
	}

	return &types.MsgCreateIncentivePoolResponse{}, nil
}

// SetPoolWeight handles MsgSetPoolWeight message,
// it set a weight of poolincentive, and it is used for calculating portion of distribution for newly minted nova token.
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

func (m msgServer) SetMultiplePoolWeight(goCtx context.Context, msg *types.MsgSetMultiplePoolWeight) (*types.MsgSetMultiplePoolWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ok := m.keeper.isValidOperator(ctx, msg.Operator)
	if !ok {
		return nil, fmt.Errorf("invalid controller address: %s", msg.Operator)
	}

	for _, pool := range msg.NewPoolData {
		ok := m.keeper.IsIncentivePool(ctx, pool.PoolId)
		if !ok {
			ctx.Logger().Info(fmt.Sprintf("poolincentive id [%s] is not incentive poolincentive-incentive", pool.PoolId))
			continue
		}

		if err := m.keeper.SetPoolWeight(ctx, pool.PoolId, pool.NewWeight); err != nil {
			return nil, err
		}
	}

	return &types.MsgSetMultiplePoolWeightResponse{}, nil
}
