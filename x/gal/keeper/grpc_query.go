package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Carina-labs/nova/x/gal/types"
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

func (q QueryServer) Params(c context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q QueryServer) ClaimableAmount(ctx context.Context, request *types.ClaimableAmountRequest) (*types.ClaimableAmountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) PendingWithdrawals(ctx context.Context, request *types.PendingWithdrawalsRequest) (*types.PendingWithdrawalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) ActiveWithdrawals(ctx context.Context, request *types.ActiveWithdrawalsRequest) (*types.ActiveWithdrawalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) Share(ctx context.Context, request *types.QueryMyShareRequest) (*types.QueryMyShareResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) DepositHistory(ctx context.Context, request *types.QueryDepositHistoryRequest) (*types.QueryDepositHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) UndelegateHistory(ctx context.Context, request *types.QueryUndelegateHistoryRequest) (*types.QueryUndelegateHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) WithdrawHistory(ctx context.Context, request *types.QueryWithdrawHistoryRequest) (*types.QueryWithdrawHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
