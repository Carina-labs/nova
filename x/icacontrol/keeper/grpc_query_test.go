package keeper_test

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"strconv"
)

func (suite *KeeperTestSuite) TestAllZones() {

	queryClient := suite.queryClient
	ctx := suite.Ctx

	// zone is not registered
	res, err := queryClient.AllZones(ctx.Context(), &types.QueryAllZonesRequest{})
	suite.Require().NoError(err)

	exp := []*types.RegisteredZone(nil)
	suite.Require().Equal(exp, res.Zones)

	// two zones are registered
	zones := suite.setZone(2)
	for _, zone := range zones {
		suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &zone)
	}

	res, err = queryClient.AllZones(ctx.Context(), &types.QueryAllZonesRequest{})
	suite.Require().NoError(err)

	for i, zone := range res.Zones {
		expect, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, "gaia"+strconv.Itoa(i))
		suite.Require().True(ok)
		suite.Require().Equal(expect, *zone)
	}

}

func (suite *KeeperTestSuite) TestZone() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	// zone is not registered
	_, err := queryClient.Zone(ctx.Context(), &types.QueryZoneRequest{ZoneId: zoneId})
	suite.Require().Error(err)

	// set two zone
	zones := suite.setZone(2)
	for _, zone := range zones {
		suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &zone)
	}

	for i := range zones {
		res, err := queryClient.Zone(ctx.Context(), &types.QueryZoneRequest{ZoneId: "gaia" + strconv.Itoa(i)})
		suite.Require().NoError(err)

		exp, ok := suite.App.IcaControlKeeper.GetRegisteredZone(ctx, "gaia"+strconv.Itoa(i))
		suite.Require().True(ok)
		suite.Require().Equal(exp, *res.Zone)
	}

	// invalid zone id
	_, err = queryClient.Zone(ctx.Context(), &types.QueryZoneRequest{ZoneId: "invalid"})
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestAutoStakingVersion() {
	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	//set zone
	suite.App.IcaControlKeeper.RegisterZone(ctx, newBaseRegisteredZone())

	// query with invalid zone
	_, err := queryClient.AutoStakingVersion(ctx.Context(), &types.QueryAutoStakingVersion{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//version is zero
	exp := types.QueryAutoStakingVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 0,
		},
	}

	res, err := queryClient.AutoStakingVersion(ctx.Context(), &types.QueryAutoStakingVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	//version is 30
	exp = types.QueryAutoStakingVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 30,
			Height:  1,
			State:   types.IcaPending,
		},
	}

	versionInfo := types.VersionState{
		ZoneId:         zoneId,
		CurrentVersion: 1,
		Record: map[uint64]*types.IBCTrace{
			0: {
				Version: 30,
				Height:  uint64(ctx.BlockHeight()),
				State:   types.IcaPending,
			},
		},
	}

	//set autostaking version
	suite.App.IcaControlKeeper.SetAutoStakingVersion(ctx, zoneId, versionInfo)
	res, err = queryClient.AutoStakingVersion(ctx.Context(), &types.QueryAutoStakingVersion{
		ZoneId:  zoneId,
		Version: 0,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	currentVersion, err := queryClient.AutoStakingCurrentVersion(ctx.Context(), &types.QueryCurrentAutoStakingVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(currentVersion.Version, versionInfo.CurrentVersion)

}
