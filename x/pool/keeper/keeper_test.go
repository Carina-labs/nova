package keeper_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/pool/keeper"
	"github.com/Carina-labs/nova/x/pool/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.PoolKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
