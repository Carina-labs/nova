## Keeper

---

## UpdateChainState
```go
func (k Keeper) UpdateChainState(ctx sdk.Context, chainInfo *types.ChainInfo) error {}
```

`UpdateChainState` updates the status of the zones stored in Oracle with a new status.

## GetChainState
```go
func (k Keeper) GetChainState(ctx sdk.Context, chainDenom string) (*types.ChainInfo, error) {}
```

`GetChainState` returns the status of the Zone stored in Oracle. This result is used to calculate the equity token.

## IsValidOperator
```go
func (k Keeper) IsValidOperator(ctx sdk.Context, operatorAddress string) bool {}
```

`IsValidOperator` verifies that the parameter address is the correct controller address.