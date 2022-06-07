package keeper_test

import (
	"github.com/Carina-labs/nova/app"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	app.KeeperTestHelper
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
}

func (suite *KeeperTestSuite) TestGetShares() {
	_, err := suite.App.GalKeeper.GetShare(suite.Ctx, suite.TestAccs[0])
	suite.NoError(err)
}
