package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateCandidatePool creates a new candidate pool and stores its information.
func (k Keeper) CreateCandidatePool(ctx sdk.Context, pool *types.CandidatePool) error {
	store := k.getCandidatePoolStore(ctx)
	key := []byte(pool.PoolId)
	if store.Has(key) {
		return fmt.Errorf("poolincentive is already exist. poolincentive id : %s", pool.PoolId)
	}

	bz := k.cdc.MustMarshal(pool)
	store.Set(key, bz)
	return nil
}

// CreateIncentivePool creates a new incentive pool and stores its information.
func (k Keeper) CreateIncentivePool(ctx sdk.Context, pool *types.IncentivePool) error {
	store := k.getIncentivePoolStore(ctx)
	key := []byte(pool.PoolId)
	if store.Has(key) {
		return fmt.Errorf("poolincentive is already exist. poolincentive id : %s", pool.PoolId)
	}

	bz := k.cdc.MustMarshal(pool)
	store.Set(key, bz)
	return nil
}

// SetPoolWeight sets the weight of the intensive pool.
func (k Keeper) SetPoolWeight(ctx sdk.Context, poolId string, newWeight uint64) error {
	store := k.getIncentivePoolStore(ctx)
	key := []byte(poolId)
	res := store.Get(key)

	var pool types.IncentivePool
	k.cdc.MustUnmarshal(res, &pool)
	pool.Weight = newWeight

	bz, err := k.cdc.Marshal(&pool)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// GetTotalWeight calculate the value of total weight of all incentive pools.
func (k Keeper) GetTotalWeight(ctx sdk.Context) (result uint64) {
	k.IterateIncentivePools(ctx, func(i int64, pool *types.IncentivePool) bool {
		result += pool.Weight
		return false
	})

	return result
}

// FindCandidatePoolById searches for candidate pools based on poolId.
func (k Keeper) FindCandidatePoolById(ctx sdk.Context, poolId string) (*types.CandidatePool, error) {
	key := []byte(poolId)
	result := &types.CandidatePool{}
	bytes := k.getCandidatePoolStore(ctx).Get(key)
	if err := result.Unmarshal(bytes); err != nil {
		return nil, err
	}

	return result, nil
}

// FindIncentivePoolById searches for incentive pools based on poolId.
func (k Keeper) FindIncentivePoolById(ctx sdk.Context, poolId string) (*types.IncentivePool, error) {
	key := []byte(poolId)
	result := &types.IncentivePool{}
	bytes := k.getIncentivePoolStore(ctx).Get(key)
	if err := result.Unmarshal(bytes); err != nil {
		return nil, err
	}

	return result, nil
}

// IsIncentivePool searches if the entered poolId is an incentive pool.
func (k Keeper) IsIncentivePool(ctx sdk.Context, poolId string) bool {
	key := []byte(poolId)
	return k.getIncentivePoolStore(ctx).Has(key)
}

// IterateCandidatePools explores all candidate pools.
func (k Keeper) IterateCandidatePools(ctx sdk.Context, cb func(i int64, pool *types.CandidatePool) bool) {
	store := k.getCandidatePoolStore(ctx)
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
		value := &types.CandidatePool{}
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

// IterateIncentivePools explores all incentive pools.
func (k Keeper) IterateIncentivePools(ctx sdk.Context, cb func(i int64, pool *types.IncentivePool) bool) {
	store := k.getIncentivePoolStore(ctx)
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
		value := &types.IncentivePool{}
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

// ClearCandidatePools deletes all candidate pools.
// WARNING : Use this function only for testing!
func (k Keeper) ClearCandidatePools(ctx sdk.Context) {
	k.IterateCandidatePools(ctx, func(i int64, pool *types.CandidatePool) bool {
		key := []byte(pool.PoolId)
		k.getCandidatePoolStore(ctx).Delete(key)
		return false
	})
}

// ClearIncentivePools deletes all incentive pools.
// WARNING : Use this function only for testing!
func (k Keeper) ClearIncentivePools(ctx sdk.Context) {
	k.IterateIncentivePools(ctx, func(i int64, pool *types.IncentivePool) bool {
		key := []byte(pool.PoolId)
		k.getIncentivePoolStore(ctx).Delete(key)
		return false
	})
}
