package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	UNDELEGATE_REQUEST_USER int64 = iota + 1
	UNDELEGATE_REQUEST_ICA  int64 = iota + 1
)

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

// undelegate 조회
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

func (k Keeper) GetUndelegateRecordsForZoneId(ctx sdk.Context, zoneId string, state int64) []types.UndelegateRecord {
	var undelegateInfo = []types.UndelegateRecord{}

	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId && undelegateRecord.State == state {
			undelegateInfo = append(undelegateInfo, undelegateRecord)
		}
		return false
	})
	return undelegateInfo
}

func (k Keeper) GetUndelegateAmount(ctx sdk.Context, denom string, zoneId string, state int64) *sdk.Coin {
	amt := &sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  denom,
	}

	var result sdk.Coin

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId && undelegateInfo.State == state {
			result = amt.Add(*undelegateInfo.Amount)
			if !result.IsZero() {
				amt = &result
			}
		}
		return false
	})

	return amt
}

// change undelegate status
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state int64) {
	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId {
			undelegateInfo.State = state
			k.SetUndelegateRecord(ctx, undelegateInfo)
		}

		return false
	})
}

func (k Keeper) SetUndelegateRecord(ctx sdk.Context, record types.UndelegateRecord) {
	// key : zoneId + delegator
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Delegator), bz)
}

func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)

	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId && undelegateRecord.State == state {
			store.Delete([]byte(undelegateRecord.ZoneId + undelegateRecord.Delegator))
		}
		return false
	})
}

// IterateUndelegatedRecords iterate through zones
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, fn func(index int64, undelegateInfo types.UndelegateRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()
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
