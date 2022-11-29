package keeper_test

import (
	"github.com/Carina-labs/nova/v2/app/apptesting"
	"github.com/Carina-labs/nova/v2/x/poolincentive/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
