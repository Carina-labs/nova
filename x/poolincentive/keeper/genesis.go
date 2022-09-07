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

	for i := range genesisState.IncentivePoolInfo.IncentivePools {
		if err := k.CreateIncentivePool(ctx, genesisState.IncentivePoolInfo.IncentivePools[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params:         k.GetParams(ctx),
		CandidatePools: []types.CandidatePool{},
		IncentivePoolInfo: types.IncentivePoolInfo{
			TotalWeight:    0,
			IncentivePools: []*types.IncentivePool{},
		},
	}

	k.IterateCandidatePools(ctx, func(i int64, pool *types.CandidatePool) bool {
		result.CandidatePools = append(result.CandidatePools, *pool)
		return false
	})

	k.IterateIncentivePools(ctx, func(i int64, pool *types.IncentivePool) bool {
		result.IncentivePoolInfo.IncentivePools = append(result.IncentivePoolInfo.IncentivePools, pool)
		return false
	})

	return result
}
