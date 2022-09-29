package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	pk            = ed25519.GenPrivKey().PubKey()
	validOperator = sdk.AccAddress(pk.Address())
)

func (suite *KeeperTestSuite) TestExportGenesis() {
	genesis := &types.GenesisState{
		Params: types.Params{
			OracleOperators: []string{validOperator.String()},
		},
		States: []types.ChainInfo{
			{
				Coin:            sdk.NewCoin("uatom", sdk.NewInt(100)),
				OperatorAddress: validOperator.String(),
				LastBlockHeight: 1,
				AppHash:         []byte{1, 2, 3, 4},
				ChainId:         "cosmos-test-chain",
			},
			{
				Coin:            sdk.NewCoin("uosmo", sdk.NewInt(100)),
				OperatorAddress: validOperator.String(),
				LastBlockHeight: 10,
				AppHash:         []byte{1, 2, 3, 4, 5},
				ChainId:         "osmosis-test-chain",
			},
		},
	}
	suite.App.OracleKeeper.InitGenesis(suite.Ctx, genesis)

	exported := suite.App.OracleKeeper.ExportGenesis(suite.Ctx)
	suite.Require().Equal(genesis, exported)
}
