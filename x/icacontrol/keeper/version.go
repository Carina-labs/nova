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

func (k Keeper) IsValidAutoStakingVersion(ctx sdk.Context, zoneId string, version uint64) bool {
	//get autostaking state
	versionInfo := k.GetAutoStakingVersion(ctx, zoneId)
	if versionInfo.ZoneId == "" {
		versionInfo.ZoneId = zoneId
		versionInfo.CurrentVersion = 0
		versionInfo.Record = make(map[uint64]*types.IBCTrace)
		versionInfo.Record[0] = &types.IBCTrace{
			Version: 0,
			State:   types.IcaPending,
		}

		k.SetAutoStakingVersion(ctx, zoneId, versionInfo)
	}

	currentVersion := versionInfo.Record[version]
	if versionInfo.CurrentVersion >= version && (currentVersion.State == types.IcaPending || currentVersion.State == types.IcaFail) {
		return true
	}
	return false
}
