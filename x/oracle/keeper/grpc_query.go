package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		Operator:        res.OperatorAddress,
		Decimal:         0,
		LastBlockHeight: res.LastBlockHeight,
		AppHash:         res.AppHash,
		ChainId:         res.ChainId,
	}, nil
}

func (q Querier) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params := q.Keeper.GetParams(sdk.UnwrapSDKContext(ctx))

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

func (q Querier) OracleVersion(goCtx context.Context, request *types.QueryOracleVersionRequest) (*types.QueryOracleVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.Keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	version, height := q.Keeper.GetOracleVersion(ctx, request.ZoneId)

	return &types.QueryOracleVersionResponse{
		Version: version,
		Height:  height,
	}, nil
}
