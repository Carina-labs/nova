package keeper

import (
	"fmt"
	"strings"

	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	QuestStateNotStarted = int32(types.QuestStateType_QUEST_STATE_NOT_STARTED)
	QuestStateClaimable  = int32(types.QuestStateType_QUEST_STATE_CLAIMABLE)
	QuestStateClaimed    = int32(types.QuestStateType_QUEST_STATE_CLAIMED)

	QuestTypeNothingToDo      = int32(types.QuestType_QUEST_NOTHING_TO_DO)
	QuestTypeSocial           = int32(types.QuestType_QUEST_SOCIAL)
	QuestTypeSnAssetClaim     = int32(types.QuestType_QUEST_SN_ASSET_CLAIM)
	QuestTypeProvideLiquidity = int32(types.QuestType_QUEST_PROVIDE_LIQUIDITY)
	QuestTypeVoteOnProposals  = int32(types.QuestType_QUEST_VOTE_ON_PROPOSALS)
)

// SetUserState sets airdrop state for the user
func (k Keeper) SetUserState(ctx sdk.Context, user sdk.AccAddress, state *types.UserState) error {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(user.String())

	store.Set(userKey, k.cdc.MustMarshal(state))
	return nil
}

// GetUserState returns airdrop state for the user
func (k Keeper) GetUserState(ctx sdk.Context, user sdk.AccAddress) (*types.UserState, error) {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(user.String())

	if !k.IsEligible(ctx, user) {
		return nil, fmt.Errorf("user is not eligible for airdrop: %v", user)
	}

	bz := store.Get(userKey)
	var userState types.UserState
	k.cdc.MustUnmarshal(bz, &userState)
	return &userState, nil
}

// IsValidControllerAddr checks if the given address is a valid controller address
func (k Keeper) isValidControllerAddr(ctx sdk.Context, addr sdk.AccAddress) bool {
	info := k.GetAirdropInfo(ctx)
	return strings.TrimSpace(info.ControllerAddress) == strings.TrimSpace(addr.String())
}

// IsEligible checks if the user is eligible for airdrop
func (k Keeper) IsEligible(ctx sdk.Context, userAddr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(userAddr.String())
	return store.Has(userKey)
}

func (k Keeper) CalcClaimableAmount(ctx sdk.Context, userAddr sdk.AccAddress) (claimable sdk.Int, claimed sdk.Int, claimedCnt int32) {
	info := k.GetAirdropInfo(ctx)
	userState, err := k.GetUserState(ctx, userAddr)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), 0
	}

	getTotalClaimed := func() (claimedCnt int32, claimed sdk.Int) {
		claimedCnt = 0
		claimed = sdk.ZeroInt()

		for _, v := range userState.QuestStates {
			if !v.ClaimedAt.IsZero() {
				amt, ok := sdk.NewIntFromString(v.ClaimedAmount)
				if !ok {
					k.Logger(ctx).Error("failed to parse claimed amount, this is not intended..", "amount", v.ClaimedAmount)
					panic("invalid claimed amount")
				}

				claimedCnt += 1
				claimed = claimed.Add(amt)
			}
		}

		return claimedCnt, claimed
	}

	cnt, claimedAmt := getTotalClaimed()
	total, ok := sdk.NewIntFromString(userState.TotalAmount)
	if !ok {
		panic("invalid total amount")
	}

	onlyOneRemained := cnt == info.QuestsCount-1
	if onlyOneRemained {
		// if only one quest remained, claim all remaining amount
		return total.Sub(claimedAmt), claimedAmt, cnt
	}

	// if not, claim fixed rate amount
	return total.QuoRaw(int64(info.QuestsCount)), claimedAmt, cnt
}
