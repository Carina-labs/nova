package keeper

import (
	"context"
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
		return nil, sdkerrors.Wrap(types.ErrNotFoundZone, request.ZoneId)
	}

	return &types.QueryZoneResponse{Zone: &zone}, nil
}

func (q *QueryServer) IcaGrant(goCtx context.Context, request *types.QueryIcaGrantRequest) (*types.QueryIcaGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	zone, ok := q.keeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZone, request.ZoneId)
	}

	grantInfo := q.keeper.GetAuthzGrant(ctx, zone.ZoneId)

	return &types.QueryIcaGrantResponse{
		Grants: &types.AuthzGrantInfo{
			ZoneId:    grantInfo.ZoneId,
			GrantInfo: grantInfo.GrantInfo,
		},
	}, nil
}

func (q QueryServer) AutoStakingVersion(goCtx context.Context, request *types.QueryAutoStakingVersion) (*types.QueryAutoStakingVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZone, request.ZoneId)
	}

	versionInfo := q.keeper.GetAutoStakingVersion(ctx, request.ZoneId)
	version := versionInfo.Record[request.Version]
	if versionInfo.Record[request.Version] == nil {
		version = &types.IBCTrace{
			Version: 0,
		}
	}
	return &types.QueryAutoStakingVersionResponse{
		VersionInfo: version,
	}, nil
}

func (q QueryServer) AutoStakingCurrentVersion(goCtx context.Context, request *types.QueryCurrentAutoStakingVersion) (*types.QueryCurrentAutoStakingVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := q.keeper.GetRegisteredZone(ctx, request.ZoneId)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNotFoundZone, request.ZoneId)
	}

	versionInfo := q.keeper.GetAutoStakingVersion(ctx, request.ZoneId)
	version := versionInfo.CurrentVersion
	if versionInfo.ZoneId == "" {
		version = 0
	}

	return &types.QueryCurrentAutoStakingVersionResponse{
		Version: version,
	}, nil
}
