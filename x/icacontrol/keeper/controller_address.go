package keeper

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// set store
func (k Keeper) GetControllerAddrStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyControllerAddress)
}

// set controller address
func (k Keeper) SetControllerAddr(ctx sdk.Context, zoneId string, controllerAddrs string) {
	store := k.GetControllerAddrStore(ctx)
	key := zoneId
	controllerInfo := types.ControllerAddressInfo{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddrs,
	}
	bz := k.cdc.MustMarshal(&controllerInfo)
	store.Set([]byte(key), bz)
}

// get controller address
func (k Keeper) GetControllerAddr(ctx sdk.Context, zoneId string) types.ControllerAddressInfo {
	store := k.GetControllerAddrStore(ctx)
	bz := store.Get([]byte(zoneId))

	var controllerInfo types.ControllerAddressInfo
	k.cdc.MustUnmarshal(bz, &controllerInfo)

	return controllerInfo
}

func (k Keeper) IsValidControllerAddr(ctx sdk.Context, zoneId, address string) bool {
	addrs := k.GetControllerAddr(ctx, zoneId)
	return address == addrs.ControllerAddress
}
