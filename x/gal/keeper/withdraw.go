package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getWithdrawRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
}

// SetWithdrawRecord stores the withdraw record.
func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record *types.WithdrawRecord) {
	store := k.getWithdrawRecordStore(ctx)
	bz := k.cdc.MustMarshal(record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

// GetWithdrawRecord returns withdraw record item by key.
func (k Keeper) GetWithdrawRecord(ctx sdk.Context, zoneId, withdrawer string) (result *types.WithdrawRecord, found bool) {
	store := k.getWithdrawRecordStore(ctx)
	keyBytes := []byte(zoneId + withdrawer)
	if !store.Has(keyBytes) {
		return nil, false
	}

	res := store.Get(keyBytes)

	var record types.WithdrawRecord
	k.cdc.MustUnmarshal(res, &record)
	return &record, true
}

// SetWithdrawRecords write multiple withdraw record.
func (k Keeper) SetWithdrawRecords(ctx sdk.Context, zoneId string, withdrawalTime time.Time) {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "setWithdrawRecords")
	k.IterateUndelegatedRecords(ctx, zoneId, func(index int64, undelegateInfo *types.UndelegateRecord) (stop bool) {
		for _, items := range undelegateInfo.Records {
			if items.State == types.UndelegateRequestByIca {
				withdrawRecord, found := k.GetWithdrawRecord(ctx, zoneId, items.Withdrawer)

				if !found {
					withdrawRecord = &types.WithdrawRecord{
						ZoneId:     zoneId,
						Withdrawer: items.Withdrawer,
					}
				}
				if withdrawRecord.Records == nil {
					withdrawRecord.Records = make(map[types.UndelegateVersion]*types.WithdrawRecordContent)
				}

				withdrawRecordContent, found := withdrawRecord.Records[items.UndelegateVersion]
				if !found {
					withdrawRecordContent = &types.WithdrawRecordContent{
						State:           types.WithdrawStatusRegistered,
						WithdrawVersion: items.UndelegateVersion,
						Amount:          items.WithdrawAmount,
						UnstakingAmount: items.SnAssetAmount,
						CompletionTime:  withdrawalTime,
					}
				} else {
					withdrawRecordContent.Amount = items.WithdrawAmount.Add(withdrawRecordContent.Amount)
					snAsset := items.SnAssetAmount.Add(*withdrawRecordContent.UnstakingAmount)
					withdrawRecordContent.UnstakingAmount = &snAsset
				}

				withdrawRecord.Records[items.UndelegateVersion] = withdrawRecordContent
				k.SetWithdrawRecord(ctx, withdrawRecord)
			}
		}
		return false
	})
}

// DeleteWithdrawRecord removes withdraw record.
func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context, withdraw *types.WithdrawRecord) {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "deleteWithdrawRecord")
	for key, record := range withdraw.Records {
		if record.State == types.WithdrawStatusTransferred {
			delete(withdraw.Records, key)
		}
	}
	k.SetWithdrawRecord(ctx, withdraw)
}

// GetWithdrawVersionStore returns store for Withdraw-version.
func (k Keeper) GetWithdrawVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawVersion)
}

// SetWithdrawVersion set withdraw version for zone id.
func (k Keeper) SetWithdrawVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {
	store := k.GetWithdrawVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetWithdrawVersion returns current withdraw-version.
func (k Keeper) GetWithdrawVersion(ctx sdk.Context, zoneId string) types.VersionState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.VersionState
	k.cdc.MustUnmarshal(res, &record)

	return record
}

func (k Keeper) IsValidWithdrawVersion(ctx sdk.Context, zoneId string, version uint64) bool {
	//get withdraw version
	versionInfo := k.GetWithdrawVersion(ctx, zoneId)
	if versionInfo.ZoneId == "" {
		versionInfo.ZoneId = zoneId
		versionInfo.CurrentVersion = 0
		versionInfo.Record = make(map[uint64]*types.IBCTrace)
		versionInfo.Record[0] = &types.IBCTrace{
			Version: 0,
			State:   types.IcaPending,
		}

		k.SetWithdrawVersion(ctx, zoneId, versionInfo)
	}

	if versionInfo.CurrentVersion >= version && (versionInfo.Record[version].State == types.IcaPending || versionInfo.Record[version].State == types.IcaFail) {
		return true
	}
	return false
}

// SetWithdrawRecordVersion set new version to withdraw record corresponding to zoneId and state.
func (k Keeper) SetWithdrawRecordVersion(ctx sdk.Context, zoneId string, state types.WithdrawStatusType, version uint64) {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "setWithdrawRecordVersion")
	k.IterateWithdrawRecords(ctx, zoneId, func(index int64, withdrawRecord *types.WithdrawRecord) (stop bool) {
		for key, record := range withdrawRecord.Records {
			if record.State == state {
				withdrawRecord.Records[key].WithdrawVersion = version
			}
		}
		k.SetWithdrawRecord(ctx, withdrawRecord)

		return false
	})
}

// GetWithdrawAmountForUser returns withdraw record corresponding to zone id and denom.
func (k Keeper) GetWithdrawAmountForUser(ctx sdk.Context, zoneId, denom string, withdrawer string) sdk.Coin {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "getWithdrawAmountForUser")
	amount := sdk.NewCoin(denom, sdk.ZeroInt())

	withdrawRecord, found := k.GetWithdrawRecord(ctx, zoneId, withdrawer)
	if !found {
		k.Logger(ctx).Error("withdraw record not found", "func", "GetWithdrawAmountForUser", "zoneId", zoneId, "address", withdrawer)
		return amount
	}

	for _, record := range withdrawRecord.Records {
		if record.State == types.WithdrawStatusTransferred {
			amount.Amount = amount.Amount.Add(record.Amount)
		}
	}

	return amount
}

// GetTotalWithdrawAmountForZoneId returns total withdraw amount corresponding to zone-id and denom.
func (k Keeper) GetTotalWithdrawAmountForZoneId(ctx sdk.Context, zoneId, denom string, blockTime time.Time) sdk.Coin {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "getTotalWithdrawAmountForZoneId")
	amount := sdk.NewCoin(denom, sdk.ZeroInt())

	k.IterateWithdrawRecords(ctx, zoneId, func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool) {
		for _, record := range withdrawInfo.Records {
			if record.CompletionTime.Before(blockTime) && record.State == types.WithdrawStatusRegistered {
				amount.Amount = amount.Amount.Add(record.Amount)
				record.State = types.WithdrawStatusTransferRequest
			}
		}
		k.SetWithdrawRecord(ctx, withdrawInfo)
		return false
	})
	return amount
}

func (k Keeper) GetTotalWithdrawAmountForFailCase(ctx sdk.Context, zoneId, denom string, blockTime time.Time) sdk.Coin {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "getTotalWithdrawAmountForFailCase")
	amount := sdk.NewCoin(denom, sdk.ZeroInt())

	k.IterateWithdrawRecords(ctx, zoneId, func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool) {
		for _, record := range withdrawInfo.Records {
			if record.CompletionTime.Before(blockTime) && record.State == types.WithdrawStatusTransferRequest {
				amount.Amount = amount.Amount.Add(record.Amount)
			}
		}
		return false
	})
	return amount
}

// ClaimWithdrawAsset is used when user want to claim their asset which is after undeleagted.
func (k Keeper) ClaimWithdrawAsset(ctx sdk.Context, withdrawer sdk.AccAddress, amt sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawer, sdk.NewCoins(amt))
	if err != nil {
		return err
	}

	return nil
}

// IsAbleToWithdraw returns if user can withdraw their asset.
// It refers nova ICA account. If ICA account's balance is greater than
// user withdraw request amount, this function returns true.
func (k Keeper) IsAbleToWithdraw(ctx sdk.Context, from sdk.AccAddress, amt sdk.Coin) bool {
	balance := k.bankKeeper.GetBalance(ctx, from, amt.Denom)
	return balance.Amount.BigInt().Cmp(amt.Amount.BigInt()) >= 0
}

// ChangeWithdrawState changes each withdraw states.
// WithdrawStatusRegistered : Withdrawal requests have been registered state of the user.
// The property of this condition is not carried over from chain host.
// WithdrawStatusTransferred : WithdrawStatusTransferred is a state in which assets are periodically transferred to the Supernova chain.
// Assets in this state can be withdrawn by the user.
func (k Keeper) ChangeWithdrawState(ctx sdk.Context, zoneId string, preState, postState types.WithdrawStatusType) {
	defer telemetry.MeasureSince(time.Now(), "gal", "withdraw", "changeWithdrawState")
	k.IterateWithdrawRecords(ctx, zoneId, func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool) {
		for _, record := range withdrawInfo.Records {
			if record.State == preState {
				record.State = postState
			}
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}

// IterateWithdrawRecords iterate all withdraw records.
func (k Keeper) IterateWithdrawRecords(ctx sdk.Context, zoneId string, fn func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool)) {
	store := k.getWithdrawRecordStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte(zoneId))
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}(iterator)
	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {

		res := types.WithdrawRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, &res)
		if stop {
			break
		}
		i++
	}
}
