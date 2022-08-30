package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreatePool stores information of pool to the store of keeper.
func (k Keeper) CreatePool(ctx sdk.Context, pool *types.Pool) error {
	store := k.getPoolStore(ctx)
	key := []byte(pool.PoolId)
	if store.Has(key) {
		return fmt.Errorf("pool is already exist. pool id : %s", pool.PoolId)
	}

	bz := k.cdc.MustMarshal(pool)
	store.Set(key, bz)
	return nil
}

// SetPoolWeight stores a new weight of pool.
func (k Keeper) SetPoolWeight(ctx sdk.Context, poolId string, newWeight uint64) error {
	store := k.getPoolStore(ctx)
	key := []byte(poolId)
	res := store.Get(key)

	var pool types.Pool
	k.cdc.MustUnmarshal(res, &pool)
	pool.Weight = newWeight

	bz, err := k.cdc.Marshal(&pool)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

func (k Keeper) GetTotalWeight(ctx sdk.Context) uint64 {
	store := k.getPoolStore(ctx)
	iterator := store.Iterator(nil, nil)
	defer func() {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err.Error()))
			return
		}
	}()

	var result uint64 = 0
	for ; iterator.Valid(); iterator.Next() {
		value := types.Pool{}
		result += value.Weight
	}
	
	return result
}