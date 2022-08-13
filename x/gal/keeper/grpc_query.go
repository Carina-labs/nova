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
	// return sum of all withdraw-able assets with WithdrawStatus_Registerred status
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	ibcDenom := q.keeper.ibcstakingKeeper.GetIBCHashDenom(ctx, request.TransferPortId, request.TransferChannelId, zoneInfo.BaseDenom)
	amount := sdk.NewCoin(ibcDenom, sdk.ZeroInt())

	// if the user has no withdraw-able assets (when transfer success record doesn't exist), return 0
	withdrawRecord, found := q.keeper.GetWithdrawRecord(ctx, request.ZoneId, request.Address)

	// if found is false, withdrawRecord variable is nil
	if !found {
		return &types.PendingWithdrawalsResponse{
			Amount: amount,
		}, nil
	}

	for _, record := range withdrawRecord.Records {
		if record.State == int64(WithdrawStatus_Registerred) {
			amount = amount.Add(*record.Amount)
		}
	}

	return &types.PendingWithdrawalsResponse{
		Amount: amount,
	}, nil
}

func (q QueryServer) ActiveWithdrawals(goCtx context.Context, request *types.ActiveWithdrawalsRequest) (*types.ActiveWithdrawalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		ctx.Logger().Error("zone_id not found", "zone_id", request.ZoneId, "module", types.ModuleName)
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	// sum of all pending withdrawals
	// if the user has no pending withdrawals (when transfer success record doesn't exist), return 0
	withdrawAmt := q.keeper.GetWithdrawAmontForUser(ctx, zoneInfo.ZoneId, zoneInfo.BaseDenom, request.Address)
	ibcDenom := q.keeper.ibcstakingKeeper.GetIBCHashDenom(ctx, request.TransferPortId, request.TransferChannelId, zoneInfo.BaseDenom)
	withdrawAmount := sdk.NewInt64Coin(ibcDenom, withdrawAmt.Amount.Int64())

	return &types.ActiveWithdrawalsResponse{
		Amount: withdrawAmount,
	}, nil
}

func (q QueryServer) DepositRecords(ctx context.Context, request *types.QueryDepositRecordRequest) (*types.QueryDepositRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) UndelegateRecords(ctx context.Context, request *types.QueryUndelegateRecordRequest) (*types.QueryUndelegateRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q QueryServer) WithdrawRecords(ctx context.Context, request *types.QueryWithdrawRecordRequest) (*types.QueryWithdrawRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}
