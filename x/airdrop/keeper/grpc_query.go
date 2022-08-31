package keeper

import (
	context "context"
	"github.com/Carina-labs/nova/x/airdrop/types"
)

var _ types.QueryServer = &Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q Querier) TotalAssetForAirdrop(ctx context.Context, request *types.QueryTotalAssetForAirdropRequest) (*types.QueryTotalAssetForAirdropResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q Querier) QuestState(ctx context.Context, request *types.QueryQuestStateRequest) (*types.QueryQuestStateResponse, error) {
	//TODO implement me
	panic("implement me")
}
