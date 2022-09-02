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
func (k Keeper) GetUserDepositRecord(ctx sdk.Context, zoneId string, claimer sdk.AccAddress) (result *types.DepositRecord, found bool) {}
```
`GetUserDepositRecord` gets the Deposit information from the store that corresponds to the `zondId` and the `claimer` address.


### GetTotalDepositAmtForZoneId
```go
func (k Keeper) GetTotalDepositAmtForZoneId(ctx sdk.Context, zoneId, denom string, state types.DepositStatusType) sdk.Coin {}
```
`GetTotalDepositAmtForZoneId` returns the sum of all Deposit coins corresponding to a specified zoneId.


### GetTotalDepositAmtForUserAddr
```go
func (k Keeper) GetTotalDepositAmtForUserAddr(ctx sdk.Context, userAddr, denom string) sdk.Coin {}
```
`GetTotalDepositAmtForUserAddr` returns the sum of the user's `address` entered as input and the deposit coin corresponding to the coin `denom`.


### SetDepositOracleVersion
```go
func (k Keeper) SetDepositOracleVersion(ctx sdk.Context, zoneId string, state types.DepositStatusType, oracleVersion uint64) {}
```
`SetDepositOracleVersion` updates the Oracle version for recorded Deposit requests.
This action is required for the correct equity calculation.


### ChangeDepositState
```go
func (k Keeper) ChangeDepositState(ctx sdk.Context, zoneId string, preState, postState types.DepositStatusType) bool {}
```
`ChangeDepositState` updates the deposit records corresponding to the preState to postState.
This operation runs in the hook after the remote deposit is run.


### SetDelegateRecordVersion
```go
func (k Keeper) SetDelegateRecordVersion(ctx sdk.Context, zoneId string, state types.DepositStatusType, version uint64) bool {}
```
`SetDelegateRecordVersion` updates the deposit version performed by the bot for the state of the deposit records corresponding to zoneId.


### DeleteRecordedDepositItem
```go
func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, zoneId string, depositor sdk.AccAddress, state types.DepositStatusType) error {}
```
`DeleteRecordedDepositItem` deletes the records corresponding to state among the `depositor's` assets deposited in the zone corresponding to `zoneId`.


### GetAllAmountNotMintShareToken
```go
func (k Keeper) GetAllAmountNotMintShareToken(ctx sdk.Context, zone *ibcstakingtypes.RegisteredZone) (sdk.Coin, error) {}
```
`GetAllAmountNotMintShareToken` returns the sum of assets that have not yet been issued by the user among the assets that have been deposited.


### IterateDepositRecord
```go
func (k Keeper) IterateDepositRecord(ctx sdk.Context, fn func(index int64, depositRecord types.DepositRecord) (stop bool)) {}
```
`IterateDepositRecord` navigates all deposit requests.

---

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
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, snDenom string, zone ibcstakingtypes.RegisteredZone, version uint64, state types.UndelegatedStatusType) (sdk.Coin, sdk.Int) {}
```
`GetUndelegateAmount` gets the information that corresponds to the zone during the de-delegation history.


### ChangeUndelegateState
```go
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {}
```
`ChangeUndelegateState` changes the status for recorded undelegation.

**states**

`UNDELEGATE_REQUEST_USER` Just requested undelegate by user. It is not in undelegate period.

`UNDELEGATE_REQUEST_ICA` Requested by ICA, It is in undelegate period.


### GetUndelegateVersionStore
```go
func (k Keeper) GetUndelegateVersionStore(ctx sdk.Context) prefix.Store {}
```
`GetUndelegateVersionStore` returns the store that stores the UndelegateVersion data.
The un-delegation task is periodically operated by the bot, so it stores the version for the last action.


### SetUndelegateVersion
```go
func (k Keeper) SetUndelegateVersion(ctx sdk.Context, zoneId string, version uint64) {}
```
`SetUndelegateVersion` sets the new undelgate Version.


### GetUndelegateVersion
```go
func (k Keeper) GetUndelegateVersion(ctx sdk.Context, zoneId string) uint64 {}
```
`GetUndelegateVersion` returns the latest un-delegation version.


### SetUndelegateRecordVersion
```go
func (k Keeper) SetUndelegateRecordVersion(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType, version uint64) bool {}
```
`SetUndelegateRecordVersion` navigates undelegate records and updates version for records corresponding to zoneId and state.


### DeleteUndelegateRecords
```go
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {}
```
`DeleteUndelegateRecords` deletes records corresponding to `zoneId` and state for undelegate records.


### IterateUndelegatedRecords
```go
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, fn func(index int64, undelegateInfo *types.UndelegateRecord) (stop bool)) {}
```
`IterateUndelegatedRecords` navigates de-delegation records.

---

## Claim
`claim.go` is responsible for the logic of issuing the defosited asset as an snAsset.


### ClaimShareToken
```go
func (k Keeper) ClaimShareToken(ctx sdk.Context, zone *ibcstakingtypes.RegisteredZone, asset sdk.Coin) (sdk.Coin, error) {}
```
`ClaimShareToken` is used when user want to claim their share token.
It calculates user's share and the amount of claimable share token.


### MintTo
```go
func (k Keeper) MintTo(ctx sdk.Context, claimer sdk.AccAddress, mintCoin sdk.Coin) error {}
```
`MintTo` mints sn-token(share token) regard with deposited token to claimer.


### TotalClaimableAssets
```go
func (k Keeper) TotalClaimableAssets(ctx sdk.Context, zone ibcstakingtypes.RegisteredZone, claimer sdk.AccAddress) (*sdk.Coin, error) {}
```
`TotalClaimableAssets` returns the total amount of claimable snAsset.


### CalculateDepositAlpha
```go
func (k Keeper) CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {}
```
`CalculateDepositAlpha` calculates alpha value.
This function is used when calculating the amount of snAsset issued.

DepositAlpha = userDepositAmount / totalStakedAmount

Issued snAsset = Alpha * totalShareTokenSupply


### CalculateWithdrawAlpha
```go
func (k Keeper) CalculateWithdrawAlpha(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {}
```
`CalculateWithdrawAlpha` calculates lambda value.
This function is used when calculating the amount of native coin issued with burning snAsset.

WithdrawAlpha = userWithdrawAmount / totalStakedAmount

Issued coin = Lambda * totalShareTokenSupply


### GetSnDenomForIBCDenom
```go
func (k Keeper) GetSnDenomForIBCDenom(ctx sdk.Context, ibcDenom string) (string, error) {}
```
`GetSnDenomForIBCDenom` changes the `IBCDenom` to the appropriate `SnDenom`.


### GetTotalStakedForLazyMinting
```go
func (k Keeper) GetTotalStakedForLazyMinting(ctx sdk.Context, denom, transferPortId, transferChanId string) (sdk.Coin, error) {}
```
`GetTotalStakedForLazyMinting` returns the sum of coins delegated to the Host chain, which have not been issued snAsset.


### ConvertWAssetToSnAssetDecimal
```go
func (k Keeper) ConvertWAssetToSnAssetDecimal(amount *big.Int, decimal int64, denom string) sdk.Coin {}
```
`ConvertWAssetToSnAssetDecimal` changes the common coin to snAsset's denom and decimal.


### ConvertSnAssetToWAssetDecimal
```go
func (k Keeper) ConvertSnAssetToWAssetDecimal(amount *big.Int, decimal int64, denom string) sdk.Coin {}
```
`ConvertSnAssetToWAssetDecimal` changes snAsset to matching coin denim and decimal.
