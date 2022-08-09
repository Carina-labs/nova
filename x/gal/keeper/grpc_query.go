package keeper

import (
	"context"

	"github.com/Carina-labs/nova/x/gal/types"
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

func (g QueryServer) Params(c context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := g.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (g QueryServer) Share(ctx context.Context, request *types.QueryCacheDepositAmountRequest) (*types.QueryCachedDepositAmountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g QueryServer) DepositHistory(ctx context.Context, request *types.QueryDepositHistoryRequest) (*types.QueryDepositHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g QueryServer) UndelegateHistory(ctx context.Context, request *types.QueryUndelegateHistoryRequest) (*types.QueryUndelegateHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g QueryServer) WithdrawHistory(ctx context.Context, request *types.QueryWithdrawHistoryRequest) (*types.QueryWithdrawHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
