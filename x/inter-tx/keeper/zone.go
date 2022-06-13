package keeper

import (
	"github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//setRegesterZone
func (k Keeper) SetRegesterZone(ctx sdk.Context, zone types.RegisteredZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := k.cdc.MustMarshal(&zone)
	store.Set([]byte(zone.ZoneName), bz)
}

//GetRegisteredZone
func (k Keeper) GetRegisteredZone(ctx sdk.Context, zone_name string) (types.RegisteredZone, bool) {
	zone := types.RegisteredZone{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := store.Get([]byte(zone_name))

	if len(bz) == 0 {
		return zone, false
	}

	k.cdc.MustUnmarshal(bz, &zone)
	return zone, true
}

// IterateRegisteredZones iterate through zones
func (k Keeper) IterateRegisteredZones(ctx sdk.Context, fn func(index int64, zoneInfo types.RegisteredZone) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()
	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {

		zone := types.RegisteredZone{}

		k.cdc.MustUnmarshal(iterator.Value(), &zone)
		stop := fn(i, zone)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) GetZoneForDenom(ctx sdk.Context, denom string) *types.RegisteredZone {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == denom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone
}

func (k Keeper) GetstDenomForBaseDenom(ctx sdk.Context, denom string) string {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == denom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone.StDenom
}

func (k Keeper) GetBaseDenomForStDenom(ctx sdk.Context, stDenom string) string {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.StDenom == stDenom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone.BaseDenom
}
