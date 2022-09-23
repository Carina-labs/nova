package gal

import (
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the gal module's state from a given genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, zone := range genState.RecordInfo {
		k.SetDelegateVersion(ctx, zone.ZoneId, *zone.DelegateTrace)
		k.SetUndelegateVersion(ctx, zone.ZoneId, *zone.UndelegateTrace)
		k.SetWithdrawVersion(ctx, zone.ZoneId, *zone.WithdrawTrace)
	}
}

// ExportGenesis returns the gal module's genesis state.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesisState()
}
