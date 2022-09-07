# Keeper

This document describes the keeper interface of the `ibcstaking` module.

---

## Zone

`zone.go` accesses information about the zone associated with Supernova.

### RegisterZone
```go
func (k Keeper) RegisterZone(ctx sdk.Context, zone *types.RegisteredZone) {}
```

`RegisterZone` stores metadata for the new zone.

### GetRegisteredZone
```go
func (k Keeper) GetRegisteredZone(ctx sdk.Context, zoneId string) (types.RegisteredZone, bool) {}
```

`GetRegisteredZone` gets information about the stored zone that fits the `zoneId`.

### DeleteRegisteredZone
```go
func (k Keeper) DeleteRegisteredZone(ctx sdk.Context, zoneId string) {}
```

`DeleteRegisteredZone` deletes zone information corresponding to `zoneId`.

### IterateRegisteredZones
```go
func (k Keeper) IterateRegisteredZones(ctx sdk.Context, fn func(index int64, zoneInfo types.RegisteredZone) (stop bool)) {}
```

`IterateRegisteredZones` navigates all registered zones.

### GetRegisteredZoneForValidatorAddr
```go
func (k Keeper) GetRegisteredZoneForValidatorAddr(ctx sdk.Context, validatorAddr string) *types.RegisteredZone {}
```

`GetRegisteredZoneForValidatorAddr` returns information about the correct zone using the validator address of the counterpart chain.

### GetZoneForDenom
```go
func (k Keeper) GetZoneForDenom(ctx sdk.Context, denom string) *types.RegisteredZone {}
```

`GetZoneForDenom` returns information about the zone that matches denom.

### GetRegisterZoneForPortId
```go
func (k Keeper) GetRegisterZoneForPortId(ctx sdk.Context, portId string) (*types.RegisteredZone, bool) {}
```

`GetRegisterZoneForPortId` returns the appropriate Zone information for portId.

### GetsnDenomForBaseDenom
```go
func (k Keeper) GetsnDenomForBaseDenom(ctx sdk.Context, baseDenom string) string {}
```

`GetsnDenomForBaseDenom` returns an appropriate pair of sn-Token denom for BaseDenom.
For example, uatom -> snatom.

### GetBaseDenomForSnDenom
```go
func (k Keeper) GetBaseDenomForSnDenom(ctx sdk.Context, snDenom string) string {}
```

`GetBaseDenomForSnDenom` returns an appropriate pair of BaseDenom for snDenom.
For example, snatom -> uatom.

### GetIBCHashDenom
```go
func (k Keeper) GetIBCHashDenom(ctx sdk.Context, portId, chanId, baseDenom string) string {}
```

`GetIBCHashDenom` uses baseDenom and portId and channelId to create the appropriate IBCdenom.

---

## Send Msgs
`send_msgs.go` sends an ICA message.

### SendIcaTx
```go
func (k Keeper) SendIcaTx(ctx sdk.Context, controllerId, connectionId string, msgs []sdk.Msg) error {}
```