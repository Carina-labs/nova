package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var _ govtypes.GovHooks = Hooks{}

type Hooks struct {
	keeper Keeper
}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// AfterProposalVote is executed from the gov module when a user votes on a proposal
func (h Hooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	h.keeper.PostProposalVote(ctx, proposalID, voterAddr)
}

// ------------------------------
// ignore these following hooks
// ------------------------------

func (h Hooks) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
}

func (h Hooks) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
}

func (h Hooks) AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64) {
}

func (h Hooks) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
}
