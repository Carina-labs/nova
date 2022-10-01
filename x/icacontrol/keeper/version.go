package keeper

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAutoStakingStore returns store for autostaking version.
func (k Keeper) GetAutoStakingVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAutoStakingVersion)
}

// SetAutoStakingVersion sets version for autostaking corresponding to zone-id records.
func (k Keeper) SetAutoStakingVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {
	store := k.GetAutoStakingVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetAutoStakingVersion returns version for autostaking corresponding to zone-id records.
func (k Keeper) GetAutoStakingVersion(ctx sdk.Context, zoneId string) types.VersionState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAutoStakingVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.VersionState
	k.cdc.MustUnmarshal(res, &record)

	return record
}
