package keeper

import (
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleReplacePoolIncentivesProposal(ctx sdk.Context, proposal *types.ReplacePoolIncentivesProposal) error {
	k.ClearIncentivePools(ctx)

	for i := range proposal.NewIncentives {
		if err := k.CreateIncentivePool(ctx, &proposal.NewIncentives[i]); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) HandleUpdatePoolIncentivesProposal(ctx sdk.Context, proposal *types.UpdatePoolIncentivesProposal) error {
	for i := range proposal.UpdatedIncentives {
		if !k.IsIncentivePool(ctx, proposal.UpdatedIncentives[i].PoolId) {
			continue
		}

		if err := k.SetPoolWeight(ctx, proposal.UpdatedIncentives[i].PoolId, proposal.UpdatedIncentives[i].Weight); err != nil {
			return err
		}
	}

	return nil
}
