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
func (k Keeper) SetDelegateVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {
	store := k.GetDelegateVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetDelegateVersion returns version for delegation corresponding to zone-id records.
func (k Keeper) GetDelegateVersion(ctx sdk.Context, zoneId string) types.VersionState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.VersionState
	k.cdc.MustUnmarshal(res, &record)

	return record
}

func (k Keeper) IsValidDelegateVersion(ctx sdk.Context, zoneId string, version uint64) bool {
	//get delegateState
	versionInfo := k.GetDelegateVersion(ctx, zoneId)
	if versionInfo.ZoneId == "" {
		versionInfo.ZoneId = zoneId
		versionInfo.CurrentVersion = 0
		versionInfo.Record = make(map[uint64]*types.IBCTrace)
		versionInfo.Record[0] = &types.IBCTrace{
			Version: 0,
			State:   types.IcaPending,
		}

		k.SetDelegateVersion(ctx, zoneId, versionInfo)
	}

	if versionInfo.CurrentVersion >= version && (versionInfo.Record[version].State == types.IcaPending || versionInfo.Record[version].State == types.IcaFail) {
		return true
	}
	return false
}
