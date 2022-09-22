package keeper

import (
	"encoding/binary"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAutoStakingStore returns store for autostaking version.
func (k Keeper) GetAutoStakingVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAutoStakingVersion)
}

// SetAutoStakingVersion sets version for autostaking corresponding to zone-id records.
func (k Keeper) SetAutoStakingVersion(ctx sdk.Context, zoneId string, version uint64, height uint64) {
	store := k.GetAutoStakingVersionStore(ctx)
	key := zoneId
	v := make([]byte, 8)
	h := make([]byte, 8)

	binary.BigEndian.PutUint64(v, version)
	binary.BigEndian.PutUint64(h, height)

	bz := append(v, h...)
	store.Set([]byte(key), bz)
}

// GetAutoStakingVersion returns version for autostaking corresponding to zone-id records.
func (k Keeper) GetAutoStakingVersion(ctx sdk.Context, zoneId string) (version uint64, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAutoStakingVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0, 0
	}

	version = binary.BigEndian.Uint64(bz[:8])
	height = binary.BigEndian.Uint64(bz[8:])
	return version, height
}
