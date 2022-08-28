package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
)

// SetUndelegateRecord write undelegate record.
func (k Keeper) SetUndelegateRecord(ctx sdk.Context, record *types.UndelegateRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := k.cdc.MustMarshal(record)
	newStoreKey := record.ZoneId + record.Delegator
	store.Set([]byte(newStoreKey), bz)
}

// GetUndelegateRecord returns undelegate record by key.
func (k Keeper) GetUndelegateRecord(ctx sdk.Context, zoneId, delegator string) (result *types.UndelegateRecord, found bool) {
	key := zoneId + delegator
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := store.Get([]byte(key))
	if len(bz) == 0 {
		return nil, false
	}

	var record types.UndelegateRecord
	k.cdc.MustUnmarshal(bz, &record)
	return &record, true
}

// GetAllUndelegateRecord returns all undelegate record.
func (k Keeper) GetAllUndelegateRecord(ctx sdk.Context, zoneId string) []*types.UndelegateRecord {
	var undelegateInfo []*types.UndelegateRecord
	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId {
			undelegateInfo = append(undelegateInfo, undelegateRecord)
		}
		return false
	})

	return undelegateInfo
}

// GetUndelegateAmount returns the amount of undelegated coin.
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, snDenom string, zone ibcstakingtypes.RegisteredZone, version uint64, state types.UndelegatedState) (sdk.Coin, sdk.Int) {

	snAsset := sdk.Coin{
		Amount: sdk.NewInt(0),
		Denom:  snDenom,
	}

	wAsset := sdk.NewInt(0)

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zone.ZoneId {
			for _, record := range undelegateRecord.Records {
				if record.OracleVersion < version {
					withdrawAsset, err := k.GetWithdrawAmt(ctx, *record.SnAssetAmount)
					if err != nil {
						return false
					}
					record.WithdrawAmount = withdrawAsset.Amount

					record.State = state
					wAsset = wAsset.Add(record.WithdrawAmount)
					snAsset = snAsset.Add(*record.SnAssetAmount)
				}
			}
			k.SetUndelegateRecord(ctx, undelegateRecord)
		}
		return false
	})
	return snAsset, wAsset
}

// ChangeUndelegateState changes undelegate record.
// UNDELEGATE_REQUEST_USER : Just requested undelegate by user. It is not in undelegate period.
// UNDELEGATE_REQUEST_ICA  : Requested by ICA, It is in undelegate period.
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state types.UndelegatedState) {
	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId {
			for _, record := range undelegateRecord.Records {
				record.State = state
			}
			k.SetUndelegateRecord(ctx, undelegateRecord)
		}
		return false
	})
}

func (k Keeper) GetUndelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
}

func (k Keeper) SetUndelegateVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetUndelegateVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetUndelegateVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetUndelegateRecordVersion(ctx sdk.Context, zoneId string, state types.UndelegatedState, version uint64) bool {
	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId {
			isChanged := false
			for _, record := range undelegateRecord.Records {
				if record.State == state {
					isChanged = true
					record.UndelegateVersion = version
				}
			}
			if isChanged {
				k.SetUndelegateRecord(ctx, undelegateRecord)
			}
		}
		return false
	})

	return true
}

// DeleteUndelegateRecords removes undelegate record.
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state types.UndelegatedState) {
	var recordItems []*types.UndelegateRecordContent
	k.IterateUndelegatedRecords(ctx, func(_ int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zoneId {
			for _, record := range undelegateRecord.Records {
				if record.State != state {
					recordItems = append(recordItems, record)
				}
			}

			isDeleted := len(recordItems) < len(undelegateRecord.Records)
			if isDeleted {
				undelegateRecord.Records = recordItems
				k.SetUndelegateRecord(ctx, undelegateRecord)
			}
		}
		return false
	})
}

// IterateUndelegatedRecords iterate through zones
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, fn func(index int64, undelegateInfo *types.UndelegateRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}(iterator)
	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var res types.UndelegateRecord
		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, &res)
		if stop {
			break
		}
		i++
	}
}
