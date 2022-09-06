# Keeper

---

## CreateCandidatePool
```go
func (k Keeper) CreateCandidatePool(ctx sdk.Context, pool *types.CandidatePool) error {}
```

`CreateCandidatePool` creates a new candidate pool and stores its information.

## CreateIncentivePool
```go
func (k Keeper) CreateIncentivePool(ctx sdk.Context, pool *types.IncentivePool) error {}
```

`CreateCandidatePool` creates a new candidate pool and stores its information.

## SetPoolWeight
```go
func (k Keeper) SetPoolWeight(ctx sdk.Context, poolId string, newWeight uint64) error {}
```

`SetPoolWeight` sets the weight of the intensive pool.

## GetTotalWeight
```go
func (k Keeper) GetTotalWeight(ctx sdk.Context) (result uint64) {}
```

`GetTotalWeight` calculate the value of total weight of all incentive pools.

## FindCandidatePoolById
```go
func (k Keeper) FindCandidatePoolById(ctx sdk.Context, poolId string) (*types.CandidatePool, error) {}
```

`FindCandidatePoolById` searches for candidate pools based on poolId.

## FindIncentivePoolById
```go
func (k Keeper) FindIncentivePoolById(ctx sdk.Context, poolId string) (*types.IncentivePool, error) {}
```

`FindIncentivePoolById` searches for incentive pools based on poolId.

## IsIncentivePool
```go
func (k Keeper) IsIncentivePool(ctx sdk.Context, poolId string) bool {}
```

`IsIncentivePool` searches if the entered poolId is an incentive pool.

## IterateCandidatePools
```go
func (k Keeper) IterateCandidatePools(ctx sdk.Context, cb func(i int64, pool *types.CandidatePool) bool) {}
```

`IterateCandidatePools` explores all candidate pools.

## IterateIncentivePools
```go
func (k Keeper) IterateIncentivePools(ctx sdk.Context, cb func(i int64, pool *types.IncentivePool) bool) {}
```

`IterateIncentivePools` explores all incentive pools.

## ClearCandidatePools
```go
func (k Keeper) ClearCandidatePools(ctx sdk.Context) {}
```

`ClearCandidatePools` deletes all candidate pools

## ClearIncentivePools
```go
func (k Keeper) ClearIncentivePools(ctx sdk.Context) {}
```

`ClearIncentivePools` deletes all incentive pools.