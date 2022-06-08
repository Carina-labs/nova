package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	types2 "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (suite *KeeperTestSuite) TestAfterTransferEnd() {
	suite.SetupTestOracle([]*types.MsgUpdateChainState{
		{
			ChainDenom:    "uosmo",
			StakedBalance: 1000000000,
			Decimal:       6,
			BlockHeight:   10000,
		},
	})

	suite.App.GalKeeper.Hooks().AfterTransferEnd(suite.Ctx,
		types2.FungibleTokenPacketData{
			Denom:    "uosmo",
			Amount:   "1000",
			Sender:   "test",
			Receiver: "test",
		},
		"osmo")
}
