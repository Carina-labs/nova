package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// Keeper defines a module interface that facilitates the transfer of coins between accounts.
type Keeper interface {
	InitGenesis(sdk.Context, *types.GenesisState)
	ExportGenesis(sdk.Context) *types.GenesisState
}

// BaseKeeper manages transfers between accounts. It implements the Keeper interface.
type BaseKeeper struct {
}

func NewBaseKeeper() BaseKeeper {
	// TODO : implements this!
	return BaseKeeper{}
}
