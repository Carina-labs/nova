package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	// Set airdrop info
	k.SetAirdropInfo(ctx, genState.AirdropInfo)

	// Set genesis state for airdrop
	for _, state := range genState.States {
		userAddr, err := sdk.AccAddressFromBech32(state.Recipient)
		if err != nil {
			panic(err)
		}

		if err = k.SetInitialUserState(ctx, userAddr, state); err != nil {
			panic(err)
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var genState types.GenesisState

	// get airdrop info
	info := k.GetAirdropInfo(ctx)
	genState.AirdropInfo = &info

	// iterate over all airdrop states
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyUserState)
	for ; iterator.Valid(); iterator.Next() {
		var userState types.UserState
		k.cdc.MustUnmarshal(iterator.Value(), &userState)
		genState.States = append(genState.States, &userState)
	}

	return &genState
}

// SetInitialUserState is called from InitGenesis function.
// it sets initial airdrop state for the user.
func (k Keeper) SetInitialUserState(ctx sdk.Context, userAddr sdk.AccAddress, state *types.UserState) error {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(userAddr.String())

	if store.Has(userKey) {
		ctx.Logger().Error("err: duplicated user was given when blockchain initializes genesis state", "user", userAddr)
		return fmt.Errorf("this user is already registered: %v", userAddr)
	}

	state.QuestStates = map[int32]*types.QuestState{
		int32(types.QuestType_QUEST_NOTHING_TO_DO):     types.EmptyQuestState(),
		int32(types.QuestType_QUEST_SOCIAL):            types.EmptyQuestState(),
		int32(types.QuestType_QUEST_PROVIDE_LIQUIDITY): types.EmptyQuestState(),
		int32(types.QuestType_QUEST_SN_ASSET_CLAIM):    types.EmptyQuestState(),
		int32(types.QuestType_QUEST_VOTE_ON_PROPOSALS): types.EmptyQuestState(),
	}

	// All user starts with the empty quests except for one quest.
	// This quest is the one that allows the user to claim 20% of the airdrop.
	state.QuestStates[int32(types.QuestType_QUEST_NOTHING_TO_DO)].State = types.QuestStateType_QUEST_STATE_CLAIMABLE
	state.QuestStates[int32(types.QuestType_QUEST_NOTHING_TO_DO)].AchievedAt = ctx.BlockTime()

	store.Set(userKey, k.cdc.MustMarshal(state))
	return nil
}
