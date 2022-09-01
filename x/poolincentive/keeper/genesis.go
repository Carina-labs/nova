package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genesisState *types.GenesisState) {
	k.SetParams(ctx, genesisState.Params)
	for i := range genesisState.CandidatePools {
		if err := k.CreateCandidatePool(ctx, &genesisState.CandidatePools[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}

	for i := range genesisState.IncentivePools {
		if err := k.CreateIncentivePool(ctx, &genesisState.IncentivePools[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params:         k.GetParams(ctx),
		CandidatePools: []types.CandidatePool{},
		IncentivePools: []types.IncentivePool{},
	}

	k.IterateCandidatePools(ctx, func(i int64, pool *types.CandidatePool) bool {
		result.CandidatePools = append(result.CandidatePools, *pool)
		return false
	})

	k.IterateIncentivePools(ctx, func(i int64, pool *types.IncentivePool) bool {
		result.IncentivePools = append(result.IncentivePools, *pool)
		return false
	})

	return result
}
