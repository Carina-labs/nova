package keeper

import (
	"encoding/binary"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDelegateVersionStore returns store for delegation.
func (k Keeper) GetDelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
}

// SetDelegateVersion sets version for delegation corresponding to zone-id records.
func (k Keeper) SetDelegateVersion(ctx sdk.Context, zoneId string, version uint64, height uint64) {
	store := k.GetDelegateVersionStore(ctx)
	key := zoneId
	v := make([]byte, 8)
	h := make([]byte, 8)

	binary.BigEndian.PutUint64(v, version)
	binary.BigEndian.PutUint64(h, height)

	bz := append(v, h...)
	store.Set([]byte(key), bz)
}

// GetDelegateVersion returns version for delegation corresponding to zone-id records.
func (k Keeper) GetDelegateVersion(ctx sdk.Context, zoneId string) (version uint64, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0, 0
	}

	version = binary.BigEndian.Uint64(bz[:8])
	height = binary.BigEndian.Uint64(bz[8:])
	return version, height
}
