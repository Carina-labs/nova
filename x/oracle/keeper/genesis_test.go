package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	pk := ed25519.GenPrivKey().PubKey()
	validOperator := sdk.AccAddress(pk.Address())

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
