package keeper_test

import (
	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

func (suite *KeeperTestSuite) TestRegisterZone() {
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
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.args)

			_, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, tc.args.ZoneId)

			suite.Require().Equal(ok, tc.expect)
		})
	}
}

func (suite *KeeperTestSuite) TestGetRegisterZone() {
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
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)

			res, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, tc.zoneId)

			suite.Require().Equal(ok, tc.result)

			if ok {
				suite.Require().Equal(res.ZoneId, tc.zone.ZoneId)
				suite.Require().Equal(res.IcaConnectionInfo.ConnectionId, tc.zone.IcaConnectionInfo.ConnectionId)
				suite.Require().Equal(res.IcaAccount.ControllerAddress, tc.zone.IcaAccount.ControllerAddress)
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
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &tc.zone)
			suite.App.IcaControlKeeper.DeleteRegisteredZone(suite.Ctx, tc.zoneId)

			_, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, tc.zoneId)

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
			denom:  zoneInfo[0].BaseDenom,
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
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IcaControlKeeper.GetZoneForDenom(suite.Ctx, tc.denom)
			suite.Require().Equal(res, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetRegisteredZoneForPortId() {
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
			portId: icatypes.PortPrefix + zoneInfo[0].IcaConnectionInfo.PortId,
			result: &zoneInfo[0],
		},
		{
			name:   "should not get osmo for port id",
			zone:   &zoneInfo[1],
			portId: icatypes.PortPrefix + "testPortId",
			result: nil,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)

			res, _ := suite.App.IcaControlKeeper.GetRegisterZoneForPortId(suite.Ctx, tc.portId)
			suite.Require().Equal(res, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetsnDenomForBaseDenom() {
	zoneInfo := suite.setZone(2)

	tcs := []struct {
		name      string
		zone      *types.RegisteredZone
		baseDenom string
		result    string
	}{
		{
			name:      "should get snDenom for baseDenom",
			zone:      &zoneInfo[0],
			baseDenom: zoneInfo[0].BaseDenom,
			result:    zoneInfo[0].SnDenom,
		},
		{
			name:      "fail",
			zone:      &zoneInfo[1],
			baseDenom: "osmo",
			result:    "",
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)
			res := suite.App.IcaControlKeeper.GetsnDenomForBaseDenom(suite.Ctx, tc.baseDenom)
			suite.Require().Equal(res, tc.result)
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
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IcaControlKeeper.GetBaseDenomForSnDenom(suite.Ctx, tc.snDenom)

			suite.Require().Equal(res, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetIBCDenomForBaseDenom() {
	success := types.RegisteredZone{
		ZoneId:    "gaia",
		BaseDenom: "uatom",
	}

	fail := types.RegisteredZone{
		ZoneId:    "gaia1",
		BaseDenom: "osmo",
	}

	tcs := []struct {
		name     string
		zoneInfo types.RegisteredZone
		portId   string
		chanId   string
		expect   string
	}{
		{
			name:     "success",
			zoneInfo: success,
			portId:   "transfer",
			chanId:   "channel-0",
			expect:   "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
		},
		{
			name:     "path is nil",
			zoneInfo: fail,
			portId:   "",
			chanId:   "",
			expect:   "osmo",
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &tc.zoneInfo)
			res := suite.App.IcaControlKeeper.GetIBCHashDenom(tc.portId, tc.chanId, tc.zoneInfo.BaseDenom)

			suite.Require().Equal(res, tc.expect)
		})
	}
}

func (suite *KeeperTestSuite) TestGetRegisteredZoneForValidatorAddr() {
	zoneInfo := suite.setZone(2)

	tcs := []struct {
		name          string
		zone          *types.RegisteredZone
		result        *types.RegisteredZone
		validatorAddr string
	}{
		{
			name:          "should get zone for validator address",
			zone:          &zoneInfo[0],
			validatorAddr: zoneInfo[0].ValidatorAddress,
			result:        &zoneInfo[0],
		},
		{
			name:          "should not get zone for validator address",
			zone:          &zoneInfo[1],
			validatorAddr: "validator address",
			result:        nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, tc.zone)

			res := suite.App.IcaControlKeeper.GetRegisteredZoneForValidatorAddr(suite.Ctx, tc.validatorAddr)
			suite.Require().Equal(res, tc.result)
		})
	}
}
