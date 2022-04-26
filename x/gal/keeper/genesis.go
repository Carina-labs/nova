package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// InitGenesis initializes the gal module's state from a given genesis state.
func (k BaseKeeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	// TODO : implements this!
}

// ExportGenesis returns the gal module's genesis state.
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	// TODO : implements this!
	return nil
}
