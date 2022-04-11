package novachain_test

import (
	"testing"

	keepertest "github.com/Carina-labs/novachain/testutil/keeper"
	"github.com/Carina-labs/novachain/testutil/nullify"
	"github.com/Carina-labs/novachain/x/novachain"
	"github.com/Carina-labs/novachain/x/novachain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NovachainKeeper(t)
	novachain.InitGenesis(ctx, *k, genesisState)
	got := novachain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
