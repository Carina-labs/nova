package keeper_test

import (
	"github.com/Carina-labs/nova/x/inter-tx/types"
)

func (suite *KeeperTestSuite) TestRegisterZoneInfo() {
	zoneInfo := suite.setZone(3)

	tcs := []struct {
		name   string
		args   *types.RegisteredZone
		expect bool
	}{
		{
			name:   "should set zone 1",
			args:   &zoneInfo[0],
			expect: true,
		},
		{
			name:   "should set zone 2",
			args:   &zoneInfo[1],
			expect: true,
		},
		{
			name:   "should set zone 3",
			args:   &zoneInfo[2],
			expect: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.args)

			_, ok := suite.App.IntertxKeeper.GetRegisteredZone(suite.Ctx, tc.args.ZoneId)

			suite.Require().Equal(ok, tc.expect)
		})
	}
}

func (suite *KeeperTestSuite) TestGetRegisterZoneInfo() {
	zoneInfo := suite.setZone(3)

	tcs := []struct {
		name   string
		zone   *types.RegisteredZone
		zoneId string
		result bool
	}{
		{
			name:   "should get zone",
			zone:   &zoneInfo[0],
			zoneId: "gaia0",
			result: true,
		},
		{
			name:   "should not get juno",
			zone:   &zoneInfo[1],
			zoneId: "juno",
			result: false,
		},
		{
			name:   "should not get test",
			zone:   &zoneInfo[2],
			zoneId: "test",
			result: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.zone)

			res, ok := suite.App.IntertxKeeper.GetRegisteredZone(suite.Ctx, tc.zoneId)

			suite.Require().Equal(ok, tc.result)

			if ok {
				suite.Require().Equal(res.ZoneId, tc.zone.ZoneId)
				suite.Require().Equal(res.TransferConnectionInfo.ChannelId, tc.zone.TransferConnectionInfo.ChannelId)
				suite.Require().Equal(res.TransferConnectionInfo.ConnectionId, tc.zone.TransferConnectionInfo.ConnectionId)
				suite.Require().Equal(res.TransferConnectionInfo.PortId, tc.zone.TransferConnectionInfo.PortId)
				suite.Require().Equal(res.IcaConnectionInfo.ConnectionId, tc.zone.IcaConnectionInfo.ConnectionId)
				suite.Require().Equal(res.IcaAccount.OwnerAddress, tc.zone.IcaAccount.OwnerAddress)
				suite.Require().Equal(res.ValidatorAddress, tc.zone.ValidatorAddress)
				suite.Require().Equal(res.SnDenom, tc.zone.SnDenom)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestDeleteRegisteredZone() {
	zoneInfo := suite.setZone(3)

	tcs := []struct {
		name   string
		zone   types.RegisteredZone
		zoneId string
		result bool
	}{
		{
			name:   "should not get gaia0 zone",
			zone:   zoneInfo[0],
			zoneId: "gaia0",
			result: false,
		},
		{
			name:   "should not get gaia1 zone",
			zone:   zoneInfo[1],
			zoneId: "gaia1",
			result: false,
		},
		{
			name:   "should not get gaia1 zone",
			zone:   zoneInfo[2],
			zoneId: "gaia2",
			result: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, &tc.zone)
			suite.App.IntertxKeeper.DeleteRegisteredZone(suite.Ctx, tc.zoneId)

			_, ok := suite.App.IntertxKeeper.GetRegisteredZone(suite.Ctx, tc.zoneId)

			suite.Require().Equal(ok, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetZoneForDenom() {
	zoneInfo := suite.setZone(2)

	tcs := []struct {
		name   string
		zone   *types.RegisteredZone
		denom  string
		result *types.RegisteredZone
	}{
		{
			name:   "should get gaia for denom",
			zone:   &zoneInfo[0],
			denom:  "atom",
			result: &zoneInfo[0],
		},
		{
			name:   "should not get osmo for denom",
			zone:   &zoneInfo[1],
			denom:  "osmo",
			result: nil,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IntertxKeeper.GetZoneForDenom(suite.Ctx, tc.denom)

			suite.Require().Equal(res, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) GetRegisteredZoneForPortId() {
	zoneInfo := suite.setZone(2)

	tcs := []struct {
		name   string
		zone   *types.RegisteredZone
		portId string
		result *types.RegisteredZone
	}{
		{
			name:   "should get gaia for port id",
			zone:   &zoneInfo[0],
			portId: zoneInfo[0].IcaConnectionInfo.PortId,
			result: &zoneInfo[0],
		},
		{
			name:   "should not get osmo for port id",
			zone:   &zoneInfo[1],
			portId: zoneInfo[0].IcaConnectionInfo.PortId,
			result: nil,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IntertxKeeper.GetZoneForDenom(suite.Ctx, tc.portId)

			suite.Require().Equal(res, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetsnDenomForBaseDenom() {
	zoneInfo := suite.setZone(1)

	tcs := []struct {
		name      string
		zone      *types.RegisteredZone
		baseDenom string
		result    *types.RegisteredZone
	}{
		{
			name:      "should get snDenom for baseDenom",
			zone:      &zoneInfo[0],
			baseDenom: zoneInfo[0].BaseDenom,
			result:    &zoneInfo[0],
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IntertxKeeper.GetsnDenomForBaseDenom(suite.Ctx, tc.baseDenom)

			suite.Require().Equal(res, tc.result.SnDenom)
		})
	}
}

func (suite *KeeperTestSuite) TestGetBaseDenomForSnDenom() {
	zoneInfo := suite.setZone(2)

	tcs := []struct {
		name    string
		zone    *types.RegisteredZone
		snDenom string
		result  string
	}{
		{
			name:    "should get baseDenom for snDenom",
			zone:    &zoneInfo[0],
			snDenom: zoneInfo[0].SnDenom,
			result:  zoneInfo[0].BaseDenom,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IntertxKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IntertxKeeper.GetBaseDenomForSnDenom(suite.Ctx, tc.snDenom)

			suite.Require().Equal(res, tc.result)
		})
	}
}
