package keeper_test

import (
	"github.com/Carina-labs/nova/v2/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	pk                = ed25519.GenPrivKey().PubKey()
	managerKey        = sdk.AccAddress(pk.Address())
	osmosisOracleAddr = sdk.AccAddress(pk.Address())
	cosmosOracleAddr  = sdk.AccAddress(pk.Address())
)

func (suite *KeeperTestSuite) TestExportGenesis() {
	genesis := &types.GenesisState{
		Params: types.Params{
			OracleKeyManager: []string{managerKey.String()},
		},
		OracleAddressInfo: []types.OracleAddressInfo{
			{
				ZoneId:        "cosmos-test-chain",
				OracleAddress: []string{cosmosOracleAddr.String()},
			},
			{
				ZoneId:        "osmosis-test-chain",
				OracleAddress: []string{osmosisOracleAddr.String()},
			},
		},
		States: []types.ChainInfo{
			{
				Coin:            sdk.NewCoin("uatom", sdk.NewInt(100)),
				OperatorAddress: cosmosOracleAddr.String(),
				LastBlockHeight: 1,
				AppHash:         []byte{1, 2, 3, 4},
				ZoneId:          "cosmos-test-chain",
			},
			{
				Coin:            sdk.NewCoin("uosmo", sdk.NewInt(100)),
				OperatorAddress: osmosisOracleAddr.String(),
				LastBlockHeight: 10,
				AppHash:         []byte{1, 2, 3, 4, 5},
				ZoneId:          "osmosis-test-chain",
			},
		},
	}
	suite.App.OracleKeeper.InitGenesis(suite.Ctx, genesis)

	exported := suite.App.OracleKeeper.ExportGenesis(suite.Ctx)
	suite.Require().Equal(genesis, exported)
}
