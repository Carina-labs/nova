package v2_test

import (
	"fmt"
	"github.com/Carina-labs/nova/v2/app/apptesting"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

const dummyUpgradeHeight = 5

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) TestMigration() {

	testCases := []struct {
		name         string
		pre_upgrade  func()
		upgrade      func()
		post_upgrade func()
	}{
		{
			"Test that the upgrade succeeds",
			func() {
				// query icahost params
				params := suite.App.ICAHostKeeper.GetParams(suite.Ctx)
				fmt.Println(params.AllowMessages)
				suite.Require().Nil(params.AllowMessages)
			},
			func() {
				suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight - 1)
				plan := upgradetypes.Plan{Name: "v2", Height: dummyUpgradeHeight}
				err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
				suite.Require().NoError(err)
				plan, exists := suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
				suite.Require().True(exists)

				suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight)
				suite.Require().NotPanics(func() {
					beginBlockRequest := abci.RequestBeginBlock{}
					suite.App.BeginBlocker(suite.Ctx, beginBlockRequest)
				})
			},
			func() {
				params := suite.App.ICAHostKeeper.GetParams(suite.Ctx)
				fmt.Println(params.AllowMessages)
				suite.Require().NotNil(params.AllowMessages)

			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			// creating pools before upgrade
			tc.pre_upgrade()

			// run upgrade
			tc.upgrade()

			// check that pool migration has been successfully done, did not break state
			tc.post_upgrade()
		})
	}
}
