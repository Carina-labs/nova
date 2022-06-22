package keeper

import (
	"fmt"

	"github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterZone
func (k Keeper) RegisterZone(ctx sdk.Context, zone *types.RegisteredZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := k.cdc.MustMarshal(zone)
	store.Set([]byte(zone.ZoneId), bz)
}

// GetRegisteredZone
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

// DeleteRegisteredZone delete zone info
func (k Keeper) DeleteRegisteredZone(ctx sdk.Context, zone_name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	ctx.Logger().Error(fmt.Sprintf("Removing chain: %s", zone_name))
	store.Delete([]byte(zone_name))
}

// IterateRegisteredZones iterate through zones
func (k Keeper) IterateRegisteredZones(ctx sdk.Context, fn func(index int64, zoneInfo types.RegisteredZone) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			panic(fmt.Errorf("unexpectedly iterator closed: %v", err))
		}
	}(iterator)

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
func (k Keeper) GetRegisteredZoneForPortId(ctx sdk.Context, portId string) *types.RegisteredZone {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		portID := zoneInfo.IcaConnectionInfo.PortId
		if portID == portId {
			zone = &zoneInfo
			return true
		}
		return false
	})
	return zone
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

func (k Keeper) GetsnDenomForBaseDenom(ctx sdk.Context, denom string) string {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == denom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone.SnDenom
}

func (k Keeper) GetBaseDenomForSnDenom(ctx sdk.Context, snDenom string) string {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.SnDenom == snDenom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone.BaseDenom
}
