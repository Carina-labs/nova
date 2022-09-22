package keeper

import (
	context "context"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type QueryServer struct {
	types.QueryServer
	keeper Keeper
}

func NewQueryServer(keeper Keeper) *QueryServer {
	return &QueryServer{
		keeper: keeper,
	}
}

func (q *QueryServer) AllZones(goCtx context.Context, request *types.QueryAllZonesRequest) (*types.QueryAllZonesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var zones []*types.RegisteredZone

	q.keeper.IterateRegisteredZones(ctx, func(index int64, zone types.RegisteredZone) (stop bool) {
		zones = append(zones, &zone)
		return false
	})

	return &types.QueryAllZonesResponse{Zones: zones}, nil
}

func (q *QueryServer) Zone(goCtx context.Context, request *types.QueryZoneRequest) (*types.QueryZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone, ok := q.keeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	return &types.QueryZoneResponse{Zone: &zone}, nil
}

func (q QueryServer) AutoStakingVersion(goCtx context.Context, request *types.QueryAutoStakingVersion) (*types.QueryAutoStakingVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZoneInfo, request.ZoneId)
	}

	version, height := q.keeper.GetAutoStakingVersion(ctx, request.ZoneId)

	return &types.QueryAutoStakingVersionResponse{
		Version: version,
		Height:  height,
	}, nil
}
