# Keeper

---

`airdrop_info.go` manages airdrop info.

## SetAirdropInfo
```go
func (k Keeper) SetAirdropInfo(ctx sdk.Context, info *types.AirdropInfo) {}
```

`SetAirdropInfo` sets airdrop info.

## GetAirdropInfo
```go
func (k Keeper) GetAirdropInfo(ctx sdk.Context) types.AirdropInfo {}
```

`GetAirdropInfo` returns airdrop info.

## ValidQuestDate
```go
func (k Keeper) ValidQuestDate(ctx sdk.Context) bool {}
```

`ValidQuestDate` returns true if the current time is valid for the user to perform quests

## ValidClaimableData
```go
func (k Keeper) ValidClaimableDate(ctx sdk.Context) bool {}
```
`ValidClaimableDate` returns true if the current time is in airdrop period

---

`airdrop_state.go` manasges airdrop state.

## SetUserState
```go
func (k Keeper) SetUserState(ctx sdk.Context, user sdk.AccAddress, state *types.UserState) error {}
```

`SetUserState` sets airdrop state for the user

## GetUserState
```go
func (k Keeper) GetUserState(ctx sdk.Context, user sdk.AccAddress) (*types.UserState, error) {}
```

`GetUserState` returns airdrop state for the user

## IsValidControllerAddr
```go
func (k Keeper) IsValidControllerAddr(ctx sdk.Context, addr sdk.AccAddress) bool {}
```

`IsValidControllerAddr` checks if the given address is a valid controller address

## IsEligible
```go
func (k Keeper) IsEligible(ctx sdk.Context, userAddr sdk.AccAddress) bool {}
```

`IsEligible` checks if the user is eligible for airdrop