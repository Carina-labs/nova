package keeper_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	key1 = secp256k1.GenPrivKey()
	acc1 = authtypes.NewBaseAccount(key1.PubKey().Address().Bytes(), key1.PubKey(), 0, 0)

	version = string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                "ics27-1",
		ControllerConnectionId: "connection-1",
		HostConnectionId:       "connection-1",
		Encoding:               "proto3",
		TxType:                 "sdk_multi_msg",
	}))
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

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)

	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))

	suite.transferPath = NewIbcTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.transferPath)
	suite.icaPath = newIcaPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.icaPath)
	suite.icaOwnerAddr = baseOwnerAcc

	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, newBaseRegisteredZone())
	suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), newBaseRegisteredZone())

	err := setupIcaPath(suite.icaPath, zoneId+"."+suite.icaOwnerAddr.String())
	suite.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) GetControllerAddr() string {
	return acc1.Address
}
