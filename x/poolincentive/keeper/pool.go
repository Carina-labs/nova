package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateCandidatePool creates a new candidate pool and stores its information.
func (k Keeper) CreateCandidatePool(ctx sdk.Context, pool *types.CandidatePool) error {
	if _, err := k.FindCandidatePoolById(ctx, pool.PoolId); err == nil {
		return fmt.Errorf("candidate pool id: %s is already exist", pool.PoolId)
	}

	candidatePoolInfo := k.GetCandidatePoolInfo(ctx)
	candidatePoolInfo.CandidatePools = append(candidatePoolInfo.CandidatePools, pool)
	k.SetCandidatePoolInfo(ctx, *candidatePoolInfo)

	return nil
}

// CreateIncentivePool creates a new incentive pool and stores its information.
func (k Keeper) CreateIncentivePool(ctx sdk.Context, pool *types.IncentivePool) error {
	if _, err := k.FindIncentivePoolById(ctx, pool.PoolId); err == nil {
		return fmt.Errorf("incentive pool id: %s is already exist", pool.PoolId)
	}

	incentivePoolInfo := k.GetIncentivePoolInfo(ctx)
	incentivePoolInfo.TotalWeight += pool.Weight
	incentivePoolInfo.IncentivePools = append(incentivePoolInfo.IncentivePools, pool)
	k.SetIncentivePoolInfo(ctx, *incentivePoolInfo)

	return nil
}

func (k Keeper) GetTotalWeight(ctx sdk.Context) uint64 {
	info := k.GetIncentivePoolInfo(ctx)
	return info.TotalWeight
}

func (k Keeper) GetCandidatePoolInfo(ctx sdk.Context) *types.CandidatePoolInfo {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCandidatePoolInfo)
	result := types.CandidatePoolInfo{}
	k.cdc.MustUnmarshal(bz, &result)

	return &result
}

func (k Keeper) GetIncentivePoolInfo(ctx sdk.Context) *types.IncentivePoolInfo {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyIncentivePoolInfo)
	result := types.IncentivePoolInfo{}
	k.cdc.MustUnmarshal(bz, &result)

	return &result
}

func (k Keeper) SetCandidatePoolInfo(ctx sdk.Context, candidatePoolInfo types.CandidatePoolInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&candidatePoolInfo)
	store.Set(types.KeyCandidatePoolInfo, bz)
}

func (k Keeper) SetIncentivePoolInfo(ctx sdk.Context, incentivePoolInfo types.IncentivePoolInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&incentivePoolInfo)
	store.Set(types.KeyIncentivePoolInfo, bz)
}

// FindCandidatePoolById searches for candidate pools based on poolId.
func (k Keeper) FindCandidatePoolById(ctx sdk.Context, poolId string) (*types.CandidatePool, error) {
	pools := k.GetCandidatePoolInfo(ctx).CandidatePools
	for i := range pools {
		if poolId == pools[i].PoolId {
			return pools[i], nil
		}
	}

	return nil, fmt.Errorf("cannot find candidate pool by id: %s", poolId)
}

// FindIncentivePoolById searches for incentive pools based on poolId.
func (k Keeper) FindIncentivePoolById(ctx sdk.Context, poolId string) (*types.IncentivePool, error) {
	pools := k.GetIncentivePoolInfo(ctx).IncentivePools
	for i := range pools {
		if poolId == pools[i].PoolId {
			return pools[i], nil
		}
	}

	return nil, fmt.Errorf("cannot find incentive pool by id: %s", poolId)
}

func (k Keeper) FindCandidatePoolByIdWithIndex(ctx sdk.Context, poolId string) (int, *types.CandidatePool, error) {
	pools := k.GetCandidatePoolInfo(ctx).CandidatePools
	for i := range pools {
		if poolId == pools[i].PoolId {
			return i, pools[i], nil
		}
	}

	return 0, nil, fmt.Errorf("cannot find candidate pool by id: %s", poolId)
}

func (k Keeper) FindIncentivePoolByIdWithIndex(ctx sdk.Context, poolId string) (int, *types.IncentivePool, error) {
	pools := k.GetIncentivePoolInfo(ctx).IncentivePools
	for i := range pools {
		if poolId == pools[i].PoolId {
			return i, pools[i], nil
		}
	}

	return 0, nil, fmt.Errorf("cannot find incentive pool by id: %s", poolId)
}

// IsIncentivePool searches if the entered poolId is an incentive pool.
func (k Keeper) IsIncentivePool(ctx sdk.Context, poolId string) bool {
	key := []byte(poolId)
	return k.getIncentivePoolStore(ctx).Has(key)
}

func (k Keeper) GetAllCandidatePool(ctx sdk.Context) []*types.CandidatePool {
	return k.GetCandidatePoolInfo(ctx).CandidatePools
}

func (k Keeper) GetAllIncentivePool(ctx sdk.Context) []*types.IncentivePool {
	return k.GetIncentivePoolInfo(ctx).IncentivePools
}
