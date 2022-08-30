package pool

import (
	"github.com/Carina-labs/nova/x/pool/keeper"
	"github.com/Carina-labs/nova/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, genesisState *types.GenesisState) {
	keeper.SetParams(ctx, genesisState.Params)
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	return types.NewGenesisState(params)
}
