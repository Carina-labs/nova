package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
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

func (q QueryServer) Params(goCtx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: res}, nil
}

func (q QueryServer) SingleCandidatePool(goCtx context.Context, request *types.QuerySingleCandidatePoolRequest) (*types.QuerySingleCandidatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, err := q.keeper.FindCandidatePoolById(ctx, request.PoolId)
	if err != nil {
		return &types.QuerySingleCandidatePoolResponse{}, err
	}

	return &types.QuerySingleCandidatePoolResponse{
		PoolId:      pool.PoolId,
		PoolAddress: pool.PoolContractAddress,
	}, nil
}

func (q QueryServer) AllCandidatePool(goCtx context.Context, request *types.QueryAllCandidatePoolRequest) (*types.QueryAllCandidatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pools := q.keeper.GetAllCandidatePool(ctx)

	return &types.QueryAllCandidatePoolResponse{
		CandidatePools: pools,
	}, nil
}

func (q QueryServer) SingleIncentivePool(goCtx context.Context, request *types.QuerySingleIncentivePoolRequest) (*types.QuerySingleIncentivePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := q.keeper.FindIncentivePoolById(ctx, request.PoolId)
	if err != nil {
		return &types.QuerySingleIncentivePoolResponse{}, err
	}

	return &types.QuerySingleIncentivePoolResponse{
		PoolId:      pool.PoolId,
		PoolAddress: pool.PoolContractAddress,
		Weight:      strconv.FormatUint(pool.Weight, 10),
	}, nil
}

func (q QueryServer) AllIncentivePool(goCtx context.Context, request *types.QueryAllIncentivePoolRequest) (*types.QueryAllIncentivePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pools := q.keeper.GetAllIncentivePool(ctx)

	return &types.QueryAllIncentivePoolResponse{
		IncentivePools: pools,
	}, nil
}

func (q QueryServer) TotalWeight(goCtx context.Context, request *types.QueryTotalWeightRequest) (*types.QueryTotalWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := q.keeper.GetTotalWeight(ctx)

	return &types.QueryTotalWeightResponse{
		TotalWeight: res,
	}, nil
}
