# Keeper

---

## CreateLpIncentiveModuleAccount
```go
func (k Keeper) CreateLpIncentiveModuleAccount(ctx sdk.Context, amount sdk.Coin) {}
```

`CreateLpIncentiveModuleAccount` creates the module account for developer vesting.
Should only be called in initial genesis creation, never again.

## GetMinter
```go
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {}
```

`GetMinter` returns current minter.

## SetMinter
```go
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {}
```

`SetMinter` sets new minter.

## GetParams
```go
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {}
```

`GetParams` returns the total set of minting parameters.

## SetParams
```go
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {}
```

`SetParams` sets the total set of minting parameters.

## StakingTokenSupply
```go
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {}
```

`StakingTokenSupply` implements an alias call to the underlying staking keeper's StakingTokenSupply to be used in BeginBlocker.

## BondedRatio
```go
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {}
```

`BondedRatio` implements an alias call to the underlying staking keeper's BondedRatio to be used in BeginBlocker.

## MintCoins
```go
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {}
```

`MintCoins` implements an alias call to the underlying supply keeper's MintCoins to be used in BeginBlocker.

## AddCollectedFees
```go
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {}
```

`AddCollectedFees` implements an alias call to the underlying supply keeper's AddCollectedFees to be used in BeginBlocker.

## GetProportions
```go
func (k Keeper) GetProportions(mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {}
```

`GetProportions` gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.

## DistributeMintedCoin
```go
func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {}
```

`DistributeMintedCoin` implements distribution of minted coins from mint to external modules.