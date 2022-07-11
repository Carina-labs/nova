package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type UndelegatedState int64

const (
	UNDELEGATE_REQUEST_USER UndelegatedState = iota + 1
	UNDELEGATE_REQUEST_ICA
)

// GetUndelegateRecord returns undelegate record by key.
func (k Keeper) GetUndelegateRecord(ctx sdk.Context, key string) (types.UndelegateRecord, bool) {
	undelegateInfo := types.UndelegateRecord{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := store.Get([]byte(key))

	if len(bz) == 0 {
		return undelegateInfo, false
	}

	k.cdc.MustUnmarshal(bz, &undelegateInfo)
	return undelegateInfo, true
}

// GetAllUndelegateRecord returns all undelegate record.
func (k Keeper) GetAllUndelegateRecord(ctx sdk.Context, zoneId string) []types.UndelegateRecord {
	var undelegateInfo = []types.UndelegateRecord{}

	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId {
			undelegateInfo = append(undelegateInfo, undelegateRecord)
		}
		return false
	})

	return undelegateInfo
}

// GetUndelegateRecordsForZoneId returns undelegate coins by zone-id.
func (k Keeper) GetUndelegateRecordsForZoneId(ctx sdk.Context, zoneId string, state UndelegatedState) []types.UndelegateRecord {
	var undelegateInfo []types.UndelegateRecord

	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId && undelegateRecord.State == int64(state) {
			undelegateInfo = append(undelegateInfo, undelegateRecord)
		}
		return false
	})
	return undelegateInfo
}

// GetUndelegateAmount returns the amount of undelegated coin.
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, denom string, zoneId string, state UndelegatedState) sdk.Coin {
	amt := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  denom,
	}

	var result sdk.Coin

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId && undelegateInfo.State == int64(state) {
			result = amt.Add(*undelegateInfo.Amount)
			if !result.IsZero() {
				amt = result
			}
		}
		return false
	})

	return amt
}

// ChangeUndelegateState changes undelegate record.
// UNDELEGATE_REQUEST_USER : Just requested undelegate by user. It is not in undelegate period.
// UNDELEGATE_REQUEST_ICA  : Requested by ICA, It is in undelegate period.
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state UndelegatedState) {
	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId {
			undelegateInfo.State = int64(state)
			k.SetUndelegateRecord(ctx, undelegateInfo)
		}

		return false
	})
}

// SetUndelegateRecord write undelegate record.
func (k Keeper) SetUndelegateRecord(ctx sdk.Context, record types.UndelegateRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	newStoreKey := record.ZoneId + record.Delegator
	store.Set([]byte(newStoreKey), bz)
}

// DeleteUndelegateRecords removes undelegate record.
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state UndelegatedState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)

	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId && undelegateRecord.State == int64(state) {
			store.Delete([]byte(undelegateRecord.ZoneId + undelegateRecord.Delegator))
		}
		return false
	})
}

// IterateUndelegatedRecords iterate through zones
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, fn func(index int64, undelegateInfo types.UndelegateRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			return
		}
	}(iterator)
	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		res := types.UndelegateRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}
