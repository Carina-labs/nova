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
func (k Keeper) SetAutoStakingVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetAutoStakingVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

// GetAutoStakingVersion returns version for autostaking corresponding to zone-id records.
func (k Keeper) GetAutoStakingVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAutoStakingVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
