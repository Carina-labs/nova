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

func (k Keeper) GetTotalWeight(ctx sdk.Context) (result uint64) {
	k.IteratePools(ctx, func(i int64, pool *types.Pool) bool {
		result += pool.Weight
		return false
	})

	return result
}

func (k Keeper) FindPoolById(ctx sdk.Context, poolId string) (result *types.Pool, ok bool) {
	k.IteratePools(ctx, func(i int64, pool *types.Pool) bool {
		if poolId == pool.PoolId {
			result = pool
			ok = true
			return true
		}
		return false
	})

	return result, ok
}

func (k Keeper) IteratePools(ctx sdk.Context, cb func(i int64, pool *types.Pool) bool) {
	store := k.getPoolStore(ctx)
	iterator := store.Iterator(nil, nil)
	defer func() {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err.Error()))
			return
		}
	}()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		value := &types.Pool{}
		err := value.Unmarshal(iterator.Value())
		if err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		stop := cb(i, value)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) ClearPools(ctx sdk.Context) {
	k.IteratePools(ctx, func(i int64, pool *types.Pool) bool {
		key := []byte(pool.PoolId)
		k.getPoolStore(ctx).Delete(key)
		return false
	})
}
