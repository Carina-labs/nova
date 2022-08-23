package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient

	coordinator  *novatesting.Coordinator
	chainA       *novatesting.TestChain
	chainB       *novatesting.TestChain
	ctxA         sdk.Context
	ctxB         sdk.Context
	transferPath *novatesting.Path
	icaPath      *novatesting.Path
	icaOwnerAddr sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	//register_zone
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)

	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))

	suite.transferPath = NewIbcTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.transferPath)

	suite.icaPath = NewIcaPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.icaPath)
	suite.icaOwnerAddr = baseOwnerAcc

	suite.App.IbcstakingKeeper.RegisterZone(suite.Ctx, NewBaseRegisteredZone())
	suite.chainA.GetApp().IbcstakingKeeper.RegisterZone(suite.chainA.GetContext(), NewBaseRegisteredZone())

	err := SetupIcaPath(suite.icaPath, zoneId+"."+suite.icaOwnerAddr.String())
	suite.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
