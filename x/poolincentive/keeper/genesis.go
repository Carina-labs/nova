package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genesisState *types.GenesisState) {
	k.SetParams(ctx, genesisState.Params)
	for i := range genesisState.CandidatePoolInfo.CandidatePools {
		if err := k.CreateCandidatePool(ctx, genesisState.CandidatePoolInfo.CandidatePools[i]); err != nil {
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
		Params: k.GetParams(ctx),
		CandidatePoolInfo: types.CandidatePoolInfo{
			CandidatePools: []*types.CandidatePool{},
		},
		IncentivePoolInfo: types.IncentivePoolInfo{
			TotalWeight:    0,
			IncentivePools: []*types.IncentivePool{},
		},
	}

	candidatePools := k.GetAllCandidatePool(ctx)
	for i := range candidatePools {
		result.CandidatePoolInfo.CandidatePools = append(result.CandidatePoolInfo.CandidatePools, candidatePools[i])
	}

	incentivePools := k.GetAllIncentivePool(ctx)
	for i := range incentivePools {
		result.IncentivePoolInfo.IncentivePools = append(result.IncentivePoolInfo.IncentivePools, incentivePools[i])
	}

	return result
}
