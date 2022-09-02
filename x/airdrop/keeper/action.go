package keeper

import (
	"fmt"

	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PostClaimedSnAsset is executed from the GAL module when a user claims an asset
func (k Keeper) PostClaimedSnAsset(ctx sdk.Context, userAddr sdk.AccAddress) {
	if !k.ValidQuestDate(ctx) {
		return
	}

	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(userAddr.String())

	// check if user is eligible
	if !k.IsEligible(ctx, userAddr) {
		return
	}

	bz := store.Get(userKey)
	var userState types.UserState
	k.cdc.MustUnmarshal(bz, &userState)

	// mark vote on proposal quest are performed
	quest := userState.QuestStates[int32(types.QuestType_QUEST_VOTE_ON_PROPOSALS)]
	if quest.State != types.QuestStateType_QUEST_STATE_NOT_STARTED || !quest.AchievedAt.IsZero() {
		return
	}

	quest.State = types.QuestStateType_QUEST_STATE_CLAIMABLE
	quest.AchievedAt = ctx.BlockTime()
	store.Set(userKey, k.cdc.MustMarshal(&userState))
}

// PostProposalVote is executed from the gov module when a user votes on a proposal
func (k Keeper) PostProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	if !k.ValidQuestDate(ctx) {
		return
	}

	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(voterAddr.String())

	// check if user is eligible
	if !k.IsEligible(ctx, voterAddr) {
		return
	}

	bz := store.Get(userKey)
	var userState types.UserState
	k.cdc.MustUnmarshal(bz, &userState)

	// mark vote on proposal quest are performed
	quest := userState.QuestStates[int32(types.QuestType_QUEST_VOTE_ON_PROPOSALS)]
	if quest.State != types.QuestStateType_QUEST_STATE_NOT_STARTED || !quest.AchievedAt.IsZero() {
		return
	}

	quest.State = types.QuestStateType_QUEST_STATE_CLAIMABLE
	quest.AchievedAt = ctx.BlockTime()
	store.Set(userKey, k.cdc.MustMarshal(&userState))
}

// MarkUserPerformedQuest marks a user performed some quest
// It fills user state with the achievement date
func (k Keeper) MarkUserPerformedQuest(ctx sdk.Context, userAddr sdk.AccAddress, questType types.QuestType) error {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(userAddr.String())

	// check if user is eligible
	if !k.IsEligible(ctx, userAddr) {
		return fmt.Errorf("this user is not eligible for airdrop: %v", userAddr)
	}

	bz := store.Get(userKey)
	var userState types.UserState
	k.cdc.MustUnmarshal(bz, &userState)

	// mark this quest as performed
	quest, ok := userState.QuestStates[int32(questType)]
	if !ok {
		ctx.Logger().Error("err: unsupported quest type was given", "questType", questType)
		panic("this quest type is not supported")
	}

	if quest.State != types.QuestStateType_QUEST_STATE_NOT_STARTED || !quest.AchievedAt.IsZero() {
		return fmt.Errorf("this user has already completed this quest: %v", userAddr)
	}

	quest.State = types.QuestStateType_QUEST_STATE_CLAIMABLE
	quest.AchievedAt = ctx.BlockTime()
	store.Set(userKey, k.cdc.MustMarshal(&userState))
	return nil
}
