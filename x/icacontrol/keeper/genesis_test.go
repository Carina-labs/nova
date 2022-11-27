package keeper_test

import "github.com/Carina-labs/nova/x/icacontrol/types"

func (suite *KeeperTestSuite) TestExportGenesis() {
	controllerKeyManager1 := suite.GenRandomAddress()
	controllerKeyManager2 := suite.GenRandomAddress()
	controllerAddr := suite.GetControllerAddr()

	genesis := &types.GenesisState{
		Params: types.Params{
			ControllerKeyManager: []string{
				controllerKeyManager1.String(),
				controllerKeyManager2.String(),
			},
		},
		ControllerAddressInfo: []*types.ControllerAddressInfo{
			{
				ZoneId:            zoneId,
				ControllerAddress: controllerAddr,
			},
		},
	}

	zoneInfo := &types.RegisteredZone{
		ZoneId: zoneId,
	}
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, zoneInfo)
	suite.App.IcaControlKeeper.InitGenesis(suite.Ctx, genesis)
	expected := suite.App.IcaControlKeeper.ExportGenesis(suite.Ctx)

	suite.Equal(genesis, expected)
}
