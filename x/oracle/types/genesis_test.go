package types

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitGenesis(t *testing.T) {
	operatorAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	genesis := NewGenesisState(Params{
		OracleKeyManager: []string{operatorAddr.String()},
	}, []ChainInfo{})

	err := genesis.Validate()
	assert.NoError(t, err)

	genesis = NewGenesisState(Params{
		OracleKeyManager: []string{"invalid_addr"},
	}, []ChainInfo{})
	err = genesis.Validate()
	assert.Error(t, err, "error expected but not found: %v", err)
}
