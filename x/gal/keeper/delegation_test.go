package keeper_test

//
//import (
//	"github.com/Carina-labs/nova/x/oracle/types"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//)
//
//func (suite *KeeperTestSuite) TestDepositCoin() {
//	suite.SetupTestOracle([]*types.MsgUpdateChainState{
//		{
//			ChainDenom:    "cosmos",
//			StakedBalance: 1000000000,
//			Decimal:       6,
//			BlockHeight:   100000,
//		},
//	})
//
//	// TODO : fix test error in IBC transfer!
//	err := suite.App.GalKeeper.DepositCoin(
//		suite.Ctx,
//		suite.TestAccs[0],
//		suite.TestAccs[1],
//		"port-1",
//		"channel-1",
//		sdk.Coin{
//			Denom:  "atom",
//			Amount: sdk.NewInt(10000),
//		},
//	)
//	suite.NoError(err)
//}
