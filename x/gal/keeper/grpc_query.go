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

func (q QueryServer) Params(goCtx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q QueryServer) EstimateSnAsset(goCtx context.Context, request *types.QueryEstimateSnAssetRequest) (*types.QueryEstimateSnAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	snDenom, err := q.keeper.GetSnDenomForIBCDenom(ctx, request.Amount.Denom)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, err.Error())
	}

	zone, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, "zone not found")
	}

	if zone.SnDenom != snDenom {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "given amount is not supported")
	}

	totalSnSupply := q.keeper.bankKeeper.GetSupply(ctx, snDenom)
	totalStakedAmount, err := q.keeper.GetTotalStakedForLazyMinting(ctx, zone.BaseDenom, zone.TransferInfo.PortId, zone.TransferInfo.ChannelId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "failed to get total staked amount")
	}

	// convert decimal
	snAsset := q.keeper.ConvertWAssetToSnAssetDecimal(request.Amount.Amount.BigInt(), zone.Decimal, snDenom)
	mintAmt := q.keeper.CalculateDepositAlpha(snAsset.Amount.BigInt(), totalSnSupply.Amount.BigInt(), totalStakedAmount.Amount.BigInt())

	return &types.QueryEstimateSnAssetResponse{
		Amount: sdk.NewCoin(snDenom, sdk.NewIntFromBigInt(mintAmt)),
	}, nil
}

func (q QueryServer) ClaimableAmount(goCtx context.Context, request *types.QueryClaimableAmountRequest) (*types.QueryClaimableAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zone, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)

	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, "zone not found")
	}

	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}

	assets, err := q.keeper.TotalClaimableAssets(ctx, zone, zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, addr)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "failed to get total claimable assets: %s", err.Error())
	}

	return &types.QueryClaimableAmountResponse{
		Amount: *assets,
	}, nil
}

func (q QueryServer) PendingWithdrawals(goCtx context.Context, request *types.QueryPendingWithdrawalsRequest) (*types.QueryPendingWithdrawalsResponse, error) {
	// return sum of all withdraw-able assets with WithdrawStatus_Registered status
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	ibcDenom := q.keeper.ibcstakingKeeper.GetIBCHashDenom(ctx, zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	amount := sdk.NewCoin(ibcDenom, sdk.ZeroInt())

	// if the user has no withdraw-able assets (when transfer success record doesn't exist), return 0
	withdrawRecord, found := q.keeper.GetWithdrawRecord(ctx, request.ZoneId, request.Address)

	// if found is false, withdrawRecord variable is nil
	if !found {
		ctx.Logger().Debug("failed to find withdraw record", "request", request)
		return &types.QueryPendingWithdrawalsResponse{
			Amount: amount,
		}, nil
	}

	for _, record := range withdrawRecord.Records {
		if record.State == int64(WithdrawStatus_Registered) {
			amount.Amount = amount.Amount.Add(record.Amount)
		}
	}

	return &types.QueryPendingWithdrawalsResponse{
		Amount: amount,
	}, nil
}

func (q QueryServer) ActiveWithdrawals(goCtx context.Context, request *types.QueryActiveWithdrawalsRequest) (*types.QueryActiveWithdrawalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.ibcstakingKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		ctx.Logger().Error("zone_id not found", "zone_id", request.ZoneId, "module", types.ModuleName)
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	// sum of all pending withdrawals
	// if the user has no pending withdrawals (when transfer success record doesn't exist), return 0
	withdrawAmt := q.keeper.GetWithdrawAmountForUser(ctx, zoneInfo.ZoneId, zoneInfo.BaseDenom, request.Address)
	ibcDenom := q.keeper.ibcstakingKeeper.GetIBCHashDenom(ctx, zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	withdrawAmount := sdk.NewInt64Coin(ibcDenom, withdrawAmt.Amount.Int64())

	return &types.QueryActiveWithdrawalsResponse{
		Amount: withdrawAmount,
	}, nil
}

func (q QueryServer) DepositRecords(goCtx context.Context, request *types.QueryDepositRecordRequest) (*types.QueryDepositRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}

	records, found := q.keeper.GetUserDepositRecord(ctx, request.ZoneId, user)
	if !found {
		return nil, types.ErrNoDepositRecord
	}

	return &types.QueryDepositRecordResponse{
		DepositRecord: records,
	}, nil
}

func (q QueryServer) UndelegateRecords(goCtx context.Context, request *types.QueryUndelegateRecordRequest) (*types.QueryUndelegateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result, found := q.keeper.GetUndelegateRecord(ctx, request.ZoneId, request.Address)
	if !found {
		return nil, types.ErrNoUndelegateRecord
	}

	return &types.QueryUndelegateRecordResponse{
		UndelegateRecord: result,
	}, nil
}

func (q QueryServer) WithdrawRecords(goCtx context.Context, request *types.QueryWithdrawRecordRequest) (*types.QueryWithdrawRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result, found := q.keeper.GetWithdrawRecord(ctx, request.ZoneId, request.Address)
	if !found {
		return nil, types.ErrNoWithdrawRecord
	}

	return &types.QueryWithdrawRecordResponse{
		WithdrawRecord: result,
	}, nil
}
