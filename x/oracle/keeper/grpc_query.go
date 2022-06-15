package keeper

import (
	"context"

	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) State(ctx context.Context, request *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	res, err := q.Keeper.GetChainState(sdk.UnwrapSDKContext(ctx), request.ChainDenom)
	if err != nil {
		return nil, err
	}

	return &types.QueryStateResponse{
		Coin:            res.Coin,
		Decimal:         res.Decimal,
		LastBlockHeight: res.LastBlockHeight,
	}, nil
}

func (q Querier) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params := q.Keeper.GetParams(sdk.UnwrapSDKContext(ctx))

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
