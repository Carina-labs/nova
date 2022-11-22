package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/poolincentive/types"
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

func (q QueryServer) NewQuerySingleCandidatePool(ctx context.Context, request *types.QuerySingleCandidatePoolRequest) *types.QuerySingleCandidatePoolResponse {
	return &types.QuerySingleCandidatePoolResponse{}
}

func (q QueryServer) NewQueryAllCandidatePool(ctx context.Context, request *types.QueryAllCandidatePoolRequest) *types.QueryAllCandidatePoolResponse {
	return &types.QueryAllCandidatePoolResponse{}
}

func (q QueryServer) NewQuerySingleIncentivePool(ctx context.Context, request *types.QuerySingleIncentivePoolRequest) *types.QuerySingleIncentivePoolResponse {
	return &types.QuerySingleIncentivePoolResponse{}
}

func (q QueryServer) NewQueryAllIncentivePool(ctx context.Context, request *types.QueryAllIncentivePoolRequest) *types.QueryAllIncentivePoolResponse {
	return &types.QueryAllIncentivePoolResponse{}
}
