package keeper_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
