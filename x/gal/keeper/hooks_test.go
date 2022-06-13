package keeper_test

import (
	"fmt"

	types3 "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (suite *KeeperTestSuite) TestAfterTransferEnd() {
	suite.SetupTestOracle([]*types.ChainInfo{
		{
			Coin:            sdk.NewInt64Coin("osmo", 1000000000),
			Decimal:         6,
			LastBlockHeight: 10000,
		},
	})

	suite.App.IntertxKeeper.SetRegesterZone(suite.Ctx,
		types3.RegisteredZone{
			ZoneName:               "Osmosis",
			IcaConnectionInfo:      nil,
			TransferConnectionInfo: nil,
			ValidatorAddress:       "test",
			BaseDenom:              "osmo",
			SnDenom:                "snOsmo",
			StDenom:                "stOsmo",
		},
	)

	suite.App.GalKeeper.SetPairToken(suite.Ctx, "osmo", "stOsmo")

	err := suite.App.BankKeeper.MintCoins(suite.Ctx, "gal", sdk.Coins{sdk.Coin{
		Denom:  "stOsmo",
		Amount: sdk.NewInt(10000),
	}})
	if err != nil {
		print(err.Error())
	}

	suite.App.GalKeeper.Hooks().AfterTransferEnd(suite.Ctx,
		types2.FungibleTokenPacketData{
			Denom:    "osmo",
			Amount:   "1000",
			Sender:   suite.TestAccs[0].String(),
			Receiver: suite.TestAccs[1].String(),
		},
		"osmo")

	minted := suite.App.BankKeeper.GetSupply(suite.Ctx, "stOsmo")
	print(fmt.Sprintf("%s", minted.Amount.String()))
}
