package keeper

import (
	"github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//setRegesterZone
func (k Keeper) SetRegesterZone(ctx sdk.Context, zone types.MsgRegisterZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := k.cdc.MustMarshal(&zone)
	store.Set([]byte(zone.ZoneName), bz)
}

//GetRegisteredZone
func (k Keeper) GetRegisteredZone(ctx sdk.Context, zone_name string) (types.MsgRegisterZone, bool) {
	zone := types.MsgRegisterZone{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := store.Get([]byte(zone_name))

	if len(bz) == 0 {
		return zone, false
	}

	k.cdc.MustUnmarshal(bz, &zone)
	return zone, true
}

// func (k Keeper) SetICAConnectionInfo(ctx sdk.Context, connection_info types.MsgRegisterZone) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixConnectionInfo)
// 	bz := k.cdc.MustMarshal(&connection_info)
// 	store.Set([]byte(connection_info.ZoneName), bz)
// }

// func (k Keeper) getICAConnectionInfo(ctx sdk.Context, zone_name string) (types.ICAConnectionInfo, bool) {
// 	connectionInfo := types.ICAConnectionInfo{}
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixConnectionInfo)
// 	bz := store.Get([]byte(zone_name))

// 	if len(bz) == 0 {
// 		return connectionInfo, false
// 	}

// 	k.cdc.MustUnmarshal(bz, &connectionInfo)
// 	return connectionInfo, true

// }
