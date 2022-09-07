package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleReplacePoolIncentivesProposal(ctx sdk.Context, proposal *types.ReplacePoolIncentivesProposal) error {
	incentiveInfo := types.IncentivePoolInfo{
		TotalWeight:    0,
		IncentivePools: []*types.IncentivePool{},
	}

	for i, newPool := range proposal.NewIncentives {
		incentiveInfo.TotalWeight += newPool.Weight
		incentiveInfo.IncentivePools = append(incentiveInfo.IncentivePools, &proposal.NewIncentives[i])
	}

	k.SetIncentivePoolInfo(ctx, incentiveInfo)

	return nil
}

func (k Keeper) HandleUpdatePoolIncentivesProposal(ctx sdk.Context, proposal *types.UpdatePoolIncentivesProposal) error {
	incentiveInfo := k.GetIncentivePoolInfo(ctx)

	for i, updatePool := range proposal.UpdatedIncentives {
		j, pool, err := k.FindIncentivePoolByIdWithIndex(ctx, updatePool.PoolId)
		if err != nil {
			incentiveInfo.IncentivePools = append(incentiveInfo.IncentivePools, &proposal.UpdatedIncentives[i])
		} else {
			if incentiveInfo.IncentivePools[j].PoolContractAddress != updatePool.PoolContractAddress {
				return fmt.Errorf("contract address mismatch, input: %s", pool.PoolContractAddress)
			}

			incentiveInfo.IncentivePools[j].Weight = updatePool.Weight
		}
	}

	k.SetIncentivePoolInfo(ctx, *incentiveInfo)
	return nil
}
