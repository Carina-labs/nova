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

func (suite *KeeperTestSuite) TestInitGenesis() {
	genesis := types.NewGenesisState(types.Params{
		OracleOperators: []string{validOperator.String()},
	})

	err := genesis.Validate()
	suite.Require().NoError(err)

	genesis = types.NewGenesisState(types.Params{
		OracleOperators: []string{"invalid_addr"},
	})
	err = genesis.Validate()
	suite.Require().NotNil(err, "error expected but not found", err)
}

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
				OracleVersion:   10,
			},
			{
				Coin:            sdk.NewCoin("uosmo", sdk.NewInt(100)),
				OperatorAddress: validOperator.String(),
				LastBlockHeight: 10,
				AppHash:         []byte{1, 2, 3, 4, 5},
				ChainId:         "osmosis-test-chain",
				OracleVersion:   11,
			},
		},
	}
	suite.App.OracleKeeper.InitGenesis(suite.Ctx, genesis)

	exported := suite.App.OracleKeeper.ExportGenesis(suite.Ctx)
	suite.Require().Equal(genesis, exported)
}
