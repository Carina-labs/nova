package keeper_test

import (
	"testing"

	testkeeper "github.com/Carina-labs/novachain/testutil/keeper"
	"github.com/Carina-labs/novachain/x/novachain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NovachainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
