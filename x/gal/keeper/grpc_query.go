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
	zone, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, "zone not found")
	}

	amt, ok := sdk.NewIntFromString(request.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "failed to get amount")
	}

	estimateAsset, err := q.keeper.ClaimShareToken(ctx, &zone, sdk.NewCoin(request.Denom, amt))
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "failed to get amount")
	}

	return &types.QueryEstimateSnAssetResponse{
		Amount: *estimateAsset,
	}, nil
}

func (q QueryServer) ClaimableAmount(goCtx context.Context, request *types.QueryClaimableAmountRequest) (*types.QueryClaimableAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zone, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)

	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}

	assets, err := q.keeper.TotalClaimableAssets(ctx, zone, addr)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "failed to get total claimable assets: %s", err.Error())
	}

	return &types.QueryClaimableAmountResponse{
		Amount: *assets,
	}, nil
}

func (q QueryServer) DepositAmount(goCtx context.Context, request *types.QueryDepositAmountRequest) (*types.QueryDepositAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	ibcDenom := q.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	depositAmount := q.keeper.GetTotalDepositAmtForUserAddr(ctx, request.ZoneId, request.Address, ibcDenom)

	claimer, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}
	delegateAmount := q.keeper.GetTotalDelegateAmtForUser(ctx, request.ZoneId, ibcDenom, claimer, types.DelegateRequest)
	result := depositAmount.Add(delegateAmount)

	return &types.QueryDepositAmountResponse{
		Amount: result,
	}, nil
}

func (q QueryServer) PendingWithdrawals(goCtx context.Context, request *types.QueryPendingWithdrawalsRequest) (*types.QueryPendingWithdrawalsResponse, error) {
	// return sum of all withdraw-able assets with WithdrawStatus_Registered status
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	ibcDenom := q.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	amount := sdk.NewCoin(ibcDenom, sdk.ZeroInt())

	undelegateRecord, found := q.keeper.GetUndelegateRecord(ctx, request.ZoneId, request.Address)
	if found {
		for _, record := range undelegateRecord.Records {
			amount.Amount = amount.Amount.Add(record.WithdrawAmount)
		}
	} else {
		ctx.Logger().Debug("failed to find undelegate record", "request", request)
	}

	// if the user has no withdraw-able assets (when transfer success record doesn't exist), return 0
	withdrawRecord, found := q.keeper.GetWithdrawRecord(ctx, request.ZoneId, request.Address)

	// if found is false, withdrawRecord variable is nil
	if found {
		for _, record := range withdrawRecord.Records {
			if record.State == types.WithdrawStatusRegistered {
				amount.Amount = amount.Amount.Add(record.Amount)
			}
		}
	} else {
		ctx.Logger().Debug("failed to find withdraw record", "request", request)
	}

	return &types.QueryPendingWithdrawalsResponse{
		Amount: amount,
	}, nil
}

func (q QueryServer) ActiveWithdrawals(goCtx context.Context, request *types.QueryActiveWithdrawalsRequest) (*types.QueryActiveWithdrawalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrNotFoundZoneInfo, "zone id: %s", request.ZoneId)
	}

	// sum of all pending withdrawals
	// if the user has no pending withdrawals (when transfer success record doesn't exist), return 0
	withdrawAmt := q.keeper.GetWithdrawAmountForUser(ctx, zoneInfo.ZoneId, zoneInfo.BaseDenom, request.Address)
	ibcDenom := q.keeper.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
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

func (q QueryServer) DelegateRecords(goCtx context.Context, request *types.QueryDelegateRecordRequest) (*types.QueryDelegateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}

	records, found := q.keeper.GetUserDelegateRecord(ctx, request.ZoneId, user)
	if !found {
		return nil, types.ErrNoDelegateRecord
	}

	return &types.QueryDelegateRecordResponse{
		DelegateRecord: records,
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

func (q QueryServer) DelegateVersion(goCtx context.Context, request *types.QueryDelegateVersion) (*types.QueryDelegateVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetDelegateVersion(ctx, request.ZoneId)
	version := versionInfo.Record[request.Version]

	if versionInfo.Record[request.Version] == nil {
		version = &types.IBCTrace{
			Version: 0,
		}
	}

	return &types.QueryDelegateVersionResponse{
		VersionInfo: version,
	}, nil
}

func (q QueryServer) UndelegateVersion(goCtx context.Context, request *types.QueryUndelegateVersion) (*types.QueryUndelegateVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetUndelegateVersion(ctx, request.ZoneId)
	version := versionInfo.Record[request.Version]

	if versionInfo.Record[request.Version] == nil {
		version = &types.IBCTrace{
			Version: 0,
		}
	}

	return &types.QueryUndelegateVersionResponse{
		VersionInfo: version,
	}, nil
}

func (q QueryServer) WithdrawVersion(goCtx context.Context, request *types.QueryWithdrawVersion) (*types.QueryWithdrawVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetWithdrawVersion(ctx, request.ZoneId)
	version := versionInfo.Record[request.Version]

	if versionInfo.Record[request.Version] == nil {
		version = &types.IBCTrace{
			Version: 0,
		}
	}

	return &types.QueryWithdrawVersionResponse{
		VersionInfo: version,
	}, nil
}

func (q QueryServer) DelegateCurrentVersion(goCtx context.Context, request *types.QueryCurrentDelegateVersion) (*types.QueryCurrentDelegateVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetDelegateVersion(ctx, request.ZoneId)
	version := versionInfo.CurrentVersion

	if versionInfo.ZoneId == "" {
		version = 0
	}

	return &types.QueryCurrentDelegateVersionResponse{
		Version: version,
	}, nil
}

func (q QueryServer) UndelegateCurrentVersion(goCtx context.Context, request *types.QueryCurrentUndelegateVersion) (*types.QueryCurrentUndelegateVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetUndelegateVersion(ctx, request.ZoneId)
	version := versionInfo.CurrentVersion

	if versionInfo.ZoneId == "" {
		version = 0
	}

	return &types.QueryCurrentUndelegateVersionResponse{
		Version: version,
	}, nil
}

func (q QueryServer) WithdrawCurrentVersion(goCtx context.Context, request *types.QueryCurrentWithdrawVersion) (*types.QueryCurrentWithdrawVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.icaControlKeeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	versionInfo := q.keeper.GetWithdrawVersion(ctx, request.ZoneId)
	version := versionInfo.CurrentVersion

	if versionInfo.ZoneId == "" {
		version = 0
	}

	return &types.QueryCurrentWithdrawVersionResponse{
		Version: version,
	}, nil
}
