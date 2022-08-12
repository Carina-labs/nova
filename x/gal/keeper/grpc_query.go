package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

func (q QueryServer) ClaimableAmount(goCtx context.Context, request *types.ClaimableAmountRequest) (*types.ClaimableAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zone, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, "zone not found")
	}

	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}

	assets, err := q.keeper.TotalClaimableAssets(ctx, zone, request.TransferPortId, request.TransferChannelId, addr)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "failed to get total claimable assets: %s", err.Error())
	}

	return &types.ClaimableAmountResponse{
		Amount: *assets,
	}, nil
}

func (q QueryServer) PendingWithdrawals(goCtx context.Context, request *types.PendingWithdrawalsRequest) (*types.PendingWithdrawalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) ActiveWithdrawals(goCtx context.Context, request *types.ActiveWithdrawalsRequest) (*types.ActiveWithdrawalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) Share(goCtx context.Context, request *types.QueryMyShareRequest) (*types.QueryMyShareResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) DepositHistory(goCtx context.Context, request *types.QueryDepositHistoryRequest) (*types.QueryDepositHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) UndelegateHistory(goCtx context.Context, request *types.QueryUndelegateHistoryRequest) (*types.QueryUndelegateHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) WithdrawHistory(goCtx context.Context, request *types.QueryWithdrawHistoryRequest) (*types.QueryWithdrawHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
