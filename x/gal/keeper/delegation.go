package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDelegateVersionStore returns store for delegation.
func (k Keeper) GetDelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
}

// SetDelegateVersion sets version for delegation corresponding to zone-id records.
func (k Keeper) SetDelegateVersion(ctx sdk.Context, zoneId string, trace types.IBCTrace) {
	store := k.GetDelegateVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetDelegateVersion returns version for delegation corresponding to zone-id records.
func (k Keeper) GetDelegateVersion(ctx sdk.Context, zoneId string) (version uint64, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.IBCTrace
	k.cdc.MustUnmarshal(res, &record)
	if res == nil {
		return 0, 0
	}

	return record.Version, record.Height
}
