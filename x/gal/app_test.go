package gal_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/gal/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestModuleParams(t *testing.T) {
	novaApp := apptesting.Setup(false)
	ctx := novaApp.NewContext(false, tmproto.Header{})
	novaApp.InitChainer(ctx, abcitypes.RequestInitChain{
		AppStateBytes: []byte("{}"),
		ChainId:       "novachain",
	},
	)
	moduleAddr := novaApp.AccountKeeper.GetModuleAddress(types.ModuleName)
	println(moduleAddr.String())

	moduleAddr2 := authtypes.NewModuleAddress(types.ModuleName)
	println(moduleAddr2.String())

	acc := novaApp.GalKeeper.GetParams(ctx)
	println(acc.String())
	require.NotNil(t, acc)
}
