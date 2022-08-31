package keeper

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return nil
}
