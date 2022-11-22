package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type QueryServer struct {
	types.QueryServer
	keeper *Keeper
}

func NewQueryServer(keeper *Keeper) *QueryServer {
	return &QueryServer{
		keeper: keeper,
	}
}

func (q QueryServer) NewQuerySingleCandidatePool(goCtx context.Context, request *types.QuerySingleCandidatePoolRequest) *types.QuerySingleCandidatePoolResponse {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := q.keeper.FindCandidatePoolById(ctx, request.PoolId)
	if err != nil {
		return nil
	}

	return &types.QuerySingleCandidatePoolResponse{
		PoolId:      pool.PoolId,
		PoolAddress: pool.PoolContractAddress,
	}
}

func (q QueryServer) NewQueryAllCandidatePool(goCtx context.Context, request *types.QueryAllCandidatePoolRequest) *types.QueryAllCandidatePoolResponse {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pools := q.keeper.GetAllCandidatePool(ctx)

	var ret []types.CandidatePool
	for _, p := range pools {
		ret = append(ret, *p)
	}

	return &types.QueryAllCandidatePoolResponse{
		CandidatePools: ret,
	}
}

func (q QueryServer) NewQuerySingleIncentivePool(goCtx context.Context, request *types.QuerySingleIncentivePoolRequest) *types.QuerySingleIncentivePoolResponse {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := q.keeper.FindIncentivePoolById(ctx, request.PoolId)
	if err != nil {
		return nil
	}

	return &types.QuerySingleIncentivePoolResponse{
		PoolId:      pool.PoolId,
		PoolAddress: pool.PoolContractAddress,
	}
}

func (q QueryServer) NewQueryAllIncentivePool(goCtx context.Context, request *types.QueryAllIncentivePoolRequest) *types.QueryAllIncentivePoolResponse {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pools := q.keeper.GetAllIncentivePool(ctx)

	var ret []types.IncentivePool
	for _, p := range pools {
		ret = append(ret, *p)
	}

	return &types.QueryAllIncentivePoolResponse{
		IncentivePools: ret,
	}
}
