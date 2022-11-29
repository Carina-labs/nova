package keeper

import (
	"fmt"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

// RegisterZone stores metadata for the new zone.
func (k Keeper) RegisterZone(ctx sdk.Context, zone *types.RegisteredZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := k.cdc.MustMarshal(zone)
	store.Set([]byte(zone.ZoneId), bz)
}

// GetRegisteredZone gets information about the stored zone that fits the zoneId.
func (k Keeper) GetRegisteredZone(ctx sdk.Context, zoneId string) (types.RegisteredZone, bool) {
	zone := types.RegisteredZone{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := store.Get([]byte(zoneId))

	if len(bz) == 0 {
		return zone, false
	}

	k.cdc.MustUnmarshal(bz, &zone)
	return zone, true
}

// DeleteRegisteredZone deletes zone information corresponding to zoneId.
func (k Keeper) DeleteRegisteredZone(ctx sdk.Context, zoneId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	ctx.Logger().Error(fmt.Sprintf("Removing chain: %s", zoneId))
	store.Delete([]byte(zoneId))
}

// IterateRegisteredZones navigates all registered zones.
func (k Keeper) IterateRegisteredZones(ctx sdk.Context, fn func(index int64, zoneInfo types.RegisteredZone) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
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

// GetRegisteredZoneForValidatorAddr returns information about the correct zone using the validator address of the counterparty chain.
func (k Keeper) GetRegisteredZoneForValidatorAddr(ctx sdk.Context, validatorAddr string) *types.RegisteredZone {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.ValidatorAddress == validatorAddr {
			zone = &zoneInfo
			return true
		}
		return false
	})
	return zone
}

// GetZoneForDenom returns information about the zone that matches denom.
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

// GetRegisterZoneForPortId returns the appropriate Zone information for portid.
func (k Keeper) GetRegisterZoneForPortId(ctx sdk.Context, portId string) (*types.RegisteredZone, bool) {
	var zone *types.RegisteredZone
	ok := false
	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if icatypes.PortPrefix+zoneInfo.IcaConnectionInfo.PortId == portId {
			zone = &zoneInfo
			ok = true
			return true
		}
		return false
	})

	return zone, ok
}

// GetRegisterZoneForHostAddr returns the appropriate Zone information for host address.
func (k Keeper) GetRegisterZoneForHostAddr(ctx sdk.Context, hostAddr string) (*types.RegisteredZone, bool) {
	var zone *types.RegisteredZone
	ok := false
	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.IcaAccount.HostAddress == hostAddr {
			zone = &zoneInfo
			ok = true
			return true
		}
		return false
	})

	return zone, ok
}

// GetsnDenomForBaseDenom returns an appropriate pair of sn-asset denom for BaseDenom.
func (k Keeper) GetsnDenomForBaseDenom(ctx sdk.Context, baseDenom string) string {
	var snDenom string

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == baseDenom {
			snDenom = zoneInfo.SnDenom
			return true
		}
		return false
	})

	return snDenom
}

// GetBaseDenomForSnDenom returns an appropriate pair of BaseDenom for snDenom.
func (k Keeper) GetBaseDenomForSnDenom(ctx sdk.Context, snDenom string) string {
	var baseDenom string

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.SnDenom == snDenom {
			baseDenom = zoneInfo.BaseDenom
			return true
		}
		return false
	})
	return baseDenom
}

func (k Keeper) DenomDuplicateCheck(ctx sdk.Context, baseDenom string) string {
	var zoneId string
	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == baseDenom {
			zoneId = zoneInfo.ZoneId
			return true
		}
		return false
	})

	return zoneId
}

// GetIBCHashDenom uses baseDenom and portId and channelId to create the appropriate IBCdenom.
func (k Keeper) GetIBCHashDenom(portId, chanId, baseDenom string) string {
	var path string

	if portId == "" || chanId == "" {
		path = ""
	} else {
		path = portId + "/" + chanId
	}

	denomTrace := transfertypes.DenomTrace{
		Path:      path,
		BaseDenom: baseDenom,
	}

	denomHash := denomTrace.IBCDenom()

	return denomHash
}
