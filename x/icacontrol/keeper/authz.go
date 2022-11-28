package keeper

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// set store
func (k Keeper) GetAuthzGrantStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAuthzGrantInfo)
}

//authz record store
func (k Keeper) SetAuthzGrant(ctx sdk.Context, grantInfo *types.AuthzGrantInfo) {
	store := k.GetAuthzGrantStore(ctx)
	bz := k.cdc.MustMarshal(grantInfo)
	store.Set([]byte(grantInfo.ZoneId), bz)
}

func (k Keeper) GetAuthzGrant(ctx sdk.Context, zoneId string) types.AuthzGrantInfo {
	store := k.GetAuthzGrantStore(ctx)
	bz := store.Get([]byte(zoneId))

	var authzGrantInfo types.AuthzGrantInfo
	k.cdc.MustUnmarshal(bz, &authzGrantInfo)

	return authzGrantInfo
}
