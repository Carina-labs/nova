package keeper

import (
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAirdropInfo sets airdrop info.
func (k Keeper) SetAirdropInfo(ctx sdk.Context, info *types.AirdropInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetKeyAirdropInfo(), k.cdc.MustMarshal(info))
}

// GetAirdropInfo returns airdrop info
func (k Keeper) GetAirdropInfo(ctx sdk.Context) *types.AirdropInfo {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(types.GetKeyAirdropInfo()) {
		panic("airdrop info is missing")
	}

	bz := store.Get(types.GetKeyAirdropInfo())
	var info types.AirdropInfo
	k.cdc.MustUnmarshal(bz, &info)
	return &info
}

// DeleteAirdropInfo delete airdrop info
func (k Keeper) DeleteAirdropInfo(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyAirdropInfo())
}

// ValidQuestDate returns true if the current time is valid for the user to perform quests
func (k Keeper) ValidQuestDate(ctx sdk.Context) bool {
	info := k.GetAirdropInfo(ctx)
	return ctx.BlockTime().Before(info.AirdropEndTimestamp)
}

// ValidClaimableDate returns true if the current time is in airdrop period
func (k Keeper) ValidClaimableDate(ctx sdk.Context) bool {
	info := k.GetAirdropInfo(ctx)
	return ctx.BlockTime().After(info.AirdropStartTimestamp) && ctx.BlockTime().Before(info.AirdropEndTimestamp)
}
