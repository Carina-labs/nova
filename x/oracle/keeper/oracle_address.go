package keeper

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// set oracle
func (k Keeper) GetOracleAddressStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleAddr)
}

// set oracle address
func (k Keeper) SetOracleAddress(ctx sdk.Context, zoneId string, oracleAddrs []string) {
	store := k.GetOracleAddressStore(ctx)
	key := zoneId
	oracleAddrInfo := types.OracleAddressInfo{
		ZoneId:        zoneId,
		OracleAddress: oracleAddrs,
	}
	bz := k.cdc.MustMarshal(&oracleAddrInfo)
	store.Set([]byte(key), bz)
}

// get oracle address
func (k Keeper) GetOracleAddress(ctx sdk.Context, zoneId string) types.OracleAddressInfo {
	store := k.GetOracleAddressStore(ctx)
	bz := store.Get([]byte(zoneId))

	var oracleAddrInfo types.OracleAddressInfo
	k.cdc.MustUnmarshal(bz, &oracleAddrInfo)

	return oracleAddrInfo
}

func (k Keeper) IsValidOracleAddress(ctx sdk.Context, zoneId, address string) bool {
	addrs := k.GetOracleAddress(ctx, zoneId)
	for i := range addrs.OracleAddress {
		if addrs.OracleAddress[i] == address {
			return true
		}
	}
	return false
}
