package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDepositRecordInfo)
}

// SetDepositRecord stores the deposit record which stores amount of the coin and user address.
func (k Keeper) SetDepositRecord(ctx sdk.Context, msg *types.DepositRecord) {
	store := k.getDepositRecordStore(ctx)
	key := msg.ZoneId + msg.Depositor
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(key), bz)
}

// GetUserDepositRecord returns the amount of coin user deposit by address.
func (k Keeper) GetUserDepositRecord(ctx sdk.Context, zoneId string, depositor sdk.AccAddress) (result *types.DepositRecord, found bool) {
	store := k.getDepositRecordStore(ctx)
	key := []byte(zoneId + depositor.String())
	if !store.Has(key) {
		return nil, false
	}

	res := store.Get(key)
	var record types.DepositRecord
	k.cdc.MustUnmarshal(res, &record)
	return &record, true
}

// GetTotalDepositAmtForZoneId returns the sum of all Deposit coins corresponding to a specified zoneId.
func (k Keeper) GetTotalDepositAmtForZoneId(ctx sdk.Context, zoneId, denom string, state types.DepositStatusType) sdk.Coin {
	defer telemetry.MeasureSince(time.Now(), "gal", "deposit", "getTotalDepositAmtForZoneId")
	totalDepositAmt := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  denom,
	}

	k.IterateDepositRecord(ctx, zoneId, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		for _, record := range depositRecord.Records {
			if record.State == state && record.Amount.Denom == denom {
				totalDepositAmt = totalDepositAmt.Add(*record.Amount)
			}
		}
		return false
	})

	return totalDepositAmt
}

// GetTotalDepositAmtForUserAddr returns the sum of the user's address entered as input
// and the deposit coin corresponding to the coin denom.
func (k Keeper) GetTotalDepositAmtForUserAddr(ctx sdk.Context, zoneId, userAddr, denom string) sdk.Coin {
	defer telemetry.MeasureSince(time.Now(), "gal", "deposit", "getTotalDepositAmtForUserAddr")
	totalDepositAmt := sdk.NewCoin(denom, sdk.NewInt(0))

	k.IterateDepositRecord(ctx, zoneId, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.Depositor == userAddr {
			for _, record := range depositRecord.Records {
				if record.State == types.DepositSuccess && record.Amount.Denom == denom {
					totalDepositAmt = totalDepositAmt.Add(*record.Amount)
				}
			}
		}
		return false
	})

	return totalDepositAmt
}

// ChangeDepositState updates the deposit records corresponding to the preState to postState.
// This operation runs in the hook after the remote deposit is run.
func (k Keeper) ChangeDepositState(ctx sdk.Context, zoneId, depositor string) {
	defer telemetry.MeasureSince(time.Now(), "gal", "deposit", "changeDepositState")
	k.IterateDepositRecord(ctx, zoneId, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		stateCheck := false
		if depositRecord.Depositor != depositor {
			return false
		}
		for _, record := range depositRecord.Records {
			if record.State == types.DepositRequest {
				record.State = types.DepositSuccess
				stateCheck = true
			}
		}
		if stateCheck {
			k.SetDepositRecord(ctx, &depositRecord)
			return true
		}
		return false
	})
}

func (k Keeper) DeleteDepositRecords(ctx sdk.Context, zoneId string, state types.DepositStatusType) {
	defer telemetry.MeasureSince(time.Now(), "gal", "deposit", "deleteDepositRecords")
	k.IterateDepositRecord(ctx, zoneId, func(_ int64, depositRecord types.DepositRecord) (stop bool) {
		var recordItems []*types.DepositRecordContent
		for _, record := range depositRecord.Records {
			if record.State != state {
				recordItems = append(recordItems, record)
			}
		}

		isDeleted := len(recordItems) < len(depositRecord.Records)
		if isDeleted {
			depositRecord.Records = recordItems
			k.SetDepositRecord(ctx, &depositRecord)
		}
		return false
	})
}

func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, zoneId string, depositor sdk.AccAddress, state types.DepositStatusType, amount sdk.Int) error {
	defer telemetry.MeasureSince(time.Now(), "gal", "deposit", "deleteRecordedDepositItem")
	record, found := k.GetUserDepositRecord(ctx, zoneId, depositor)
	if !found {
		return types.ErrNoDepositRecord
	}

	var deleteRecord types.DepositRecordContent
	recordItems := record.Records
	for i, item := range record.Records {
		if item.State == state && item.Amount.Amount.Equal(amount) {
			recordItems = append(recordItems[:i], recordItems[i+1:]...)
			deleteRecord = *item
			break
		}
	}

	isDeleted := len(recordItems) < len(record.Records)
	if isDeleted {
		record.Records = recordItems
		k.SetDepositRecord(ctx, record)
		ctx.Logger().Info("DeleteRecordedDepositItem", "DeleteRecord", deleteRecord)
		return nil
	}

	ctx.Logger().Info("DeleteRecordedDepositItem", "Error", types.ErrNoDelegateRecord)
	return types.ErrNoDeleteRecord
}

func (k Keeper) getAssetInfo(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAssetInfo)
}

func (k Keeper) SetAssetInfo(ctx sdk.Context, assetInfo *types.AssetInfo) {
	store := k.getAssetInfo(ctx)
	fmt.Println(assetInfo)
	key := assetInfo.ZoneId
	bz := k.cdc.MustMarshal(assetInfo)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetAssetInfoForZoneId(ctx sdk.Context, zoneId string) *types.AssetInfo {
	store := k.getAssetInfo(ctx)
	key := []byte(zoneId)
	if !store.Has(key) {
		return nil
	}

	res := store.Get(key)
	var assetInfo types.AssetInfo
	k.cdc.MustUnmarshal(res, &assetInfo)
	return &assetInfo
}

// IterateDepositRecord navigates all deposit requests.
func (k Keeper) IterateDepositRecord(ctx sdk.Context, zoneId string, fn func(index int64, depositRecord types.DepositRecord) (stop bool)) {
	store := k.getDepositRecordStore(ctx)
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
		res := types.DepositRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}
