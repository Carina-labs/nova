# Keeper

---

## Deposit

`deposit.go` is responsible for recording and managing the user's deposit behavior.

### SetDepositRecord

```go
func (k Keeper) SetDepositRecord(ctx sdk.Context, msg *types.DepositRecord) {}
```

`SetDepositRecord` stores the contents of the `DepositRecord` message entered as an input in the store.

### GetUserDepositRecord

```go
func (k Keeper) GetUserDepositRecord(ctx sdk.Context, zoneId string, depositor sdk.AccAddress) (result *types.DepositRecord, found bool) {}
```

`GetUserDepositRecord` gets the Deposit information from the store that corresponds to the `zondId` and the `claimer`
address.

### GetTotalDepositAmtForZoneId

```go
func (k Keeper) GetTotalDepositAmtForZoneId(ctx sdk.Context, zoneId, denom string, state types.DepositStatusType) sdk.Coin {}
```

`GetTotalDepositAmtForZoneId` returns the sum of all Deposit coins corresponding to a specified zoneId.

### GetTotalDepositAmtForUserAddr

```go
func (k Keeper) GetTotalDepositAmtForUserAddr(ctx sdk.Context, zoneId, userAddr, denom string) sdk.Coin {}
```

`GetTotalDepositAmtForUserAddr` returns the sum of the user's `address` entered as input and the deposit coin
corresponding to the coin `denom`.

### SetDepositOracleVersion

```go
func (k Keeper) SetDepositOracleVersion(ctx sdk.Context, zoneId string, state types.DepositStatusType, oracleVersion uint64) {}
```

`SetDepositOracleVersion` updates the Oracle version for recorded Deposit requests.
This action is required for the correct equity calculation.

### ChangeDepositState

```go
func (k Keeper) ChangeDepositState(ctx sdk.Context, zoneId, depositor string) {}
```

`ChangeDepositState` updates the deposit records corresponding to the preState to postState.
This operation runs in the hook after the remote deposit is run.

### DeleteRecordedDepositItem

```go
func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, zoneId string, depositor sdk.AccAddress, state types.DepositStatusType, amount sdk.Int) error {}
```

`DeleteRecordedDepositItem` deletes the records corresponding to state among the `depositor's` assets deposited in the
zone corresponding to `zoneId`.

### SetAssetInfo

```go
func (k Keeper) SetAssetInfo(ctx sdk.Context, assetInfo *types.AssetInfo) {}
```

`SetAssetInfo` stores information about snAsset that is not minted.

### GetAssetInfoForZoneId

```go
func (k Keeper) GetAssetInfoForZoneId(ctx sdk.Context, zoneId string) *types.AssetInfo {}
```

`GetAssetInfoForZoneId` returns not minted snAsset information according to the input zoneId.

### HasMaxDepositEntries

````go
func (k Keeper) HasMaxDepositEntries(depositRecords types.DepositRecord, maxEntries int64) bool
````

`HasMaxDepositEntries` checks that the maximum number of times a user can request deposit message during deployment has
been exceeded.

### IterateDepositRecord

```go
func (k Keeper) IterateDepositRecord(ctx sdk.Context, fn func (index int64, depositRecord types.DepositRecord) (stop bool)) {}
```

`IterateDepositRecord` navigates all deposit requests.

## Undelegate

`undelegate.go` is responsible for recording and managing the user's undelegation behavior.

### SetUndelegateRecord

```go
func (k Keeper) SetUndelegateRecord(ctx sdk.Context, record *types.UndelegateRecord) {}
```

`SetUndelegateRecord` writes a record of the user's undelegation actions.

### GetUndelegateRecord

```go
func (k Keeper) GetUndelegateRecord(ctx sdk.Context, zoneId, delegator string) (result *types.UndelegateRecord, found bool) {}
```

`GetUndelegateRecord` returns the record corresponding to zoneId and delegator among the user's undelegation records.

### GetAllUndelegateRecord

```go
func (k Keeper) GetAllUndelegateRecord(ctx sdk.Context, zoneId string) []*types.UndelegateRecord {}
```

`GetAllUndelegateRecord` returns all undelegate records corresponding to `zoneId`.

### GetUndelegateAmount

```go
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, snDenom string, zone icacontroltypes.RegisteredZone, version uint64, state types.UndelegatedStatusType) (sdk.Coin, sdk.Int) {}
```

`GetUndelegateAmount` gets the information that corresponds to the zone during the de-delegation history.

### GetReUndelegateAmount

```go
func (k Keeper) GetReUndelegateAmount(ctx sdk.Context, snDenom string, zone icacontroltypes.RegisteredZone, version uint64) (sdk.Coin, sdk.Int) {}
```

`GetReUndelegateAmount` is used to request again when the requested ica undelegate fails.

### ChangeUndelegateState

```go
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {}
```

`ChangeUndelegateState` changes the status for recorded undelegation.

**states**

`UNDELEGATE_REQUEST_USER` Just requested undelegate by user. It is not in undelegate period.

`UNDELEGATE_REQUEST_ICA` Requested by ICA, It is in undelegate period.

### GetWithdrawAmt

```go
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (*sdk.Coin, error) {}
```

`GetWithdrawAmt` 

### GetUndelegateVersionStore

```go
func (k Keeper) GetUndelegateVersionStore(ctx sdk.Context) prefix.Store {}
```

`GetUndelegateVersionStore` returns the store that stores the UndelegateVersion data.
The un-delegation task is periodically operated by the bot, so it stores the version for the last action.

### SetUndelegateVersion

```go
func (k Keeper) SetUndelegateVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {}
```

`SetUndelegateVersion` sets the new undelgate Version.

### GetUndelegateVersion

```go
func (k Keeper) GetUndelegateVersion(ctx sdk.Context, zoneId string) types.VersionState {}
```

`GetUndelegateVersion` returns the latest un-delegation version.

### SetUndelegateRecordVersion

```go
func (k Keeper) SetUndelegateRecordVersion(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType, version uint64) bool {}
```

`SetUndelegateRecordVersion` navigates undelegate records and updates version for records corresponding to zoneId and
state.

### DeleteUndelegateRecords

```go
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {}
```

`DeleteUndelegateRecords` deletes records corresponding to `zoneId` and state for undelegate records.

### HasMaxUndelegateEntries

```go
func (k Keeper) HasMaxUndelegateEntries(undelegateRecords types.UndelegateRecord, maxEntries int64) bool
```

`HasMaxUndelegateEntries` checks that the maximum number of times a user can request undelegate message during
deployment has been exceeded.

### IterateUndelegatedRecords

```go
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, fn func (index int64, undelegateInfo *types.UndelegateRecord) (stop bool)) {}
```

`IterateUndelegatedRecords` navigates de-delegation records.

