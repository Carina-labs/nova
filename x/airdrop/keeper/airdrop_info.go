package keeper

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAirdropInfo sets airdrop info
func (k Keeper) SetAirdropInfo(ctx sdk.Context, info *types.AirdropInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetKeyAirdropInfo(), k.cdc.MustMarshal(info))
}

// GetAirdropInfo returns airdrop info
func (k Keeper) GetAirdropInfo(ctx sdk.Context) types.AirdropInfo {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(types.GetKeyAirdropInfo()) {
		panic("airdrop info is missing")
	}

	bz := store.Get(types.GetKeyAirdropInfo())
	var info types.AirdropInfo
	k.cdc.MustUnmarshal(bz, &info)
	return info
}

// ValidAirdropDate returns true if we're in the airdrop period.
func (k Keeper) ValidAirdropDate(ctx sdk.Context) bool {
	info := k.GetAirdropInfo(ctx)
	return ctx.BlockTime().After(info.AirdropEndTimestamp)
}
