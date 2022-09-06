package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ibcstakingtypes "github.com/Carina-labs/nova/x/icacontrol/types"
)

// GetUndelegateVersionStore returns the store that stores the UndelegateVersion data.
// The un-delegation task is periodically operated by the bot, so it stores the version for the last action.
func (k Keeper) GetUndelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
}

// SetUndelegateVersion sets the new un-delgate Version.
func (k Keeper) SetUndelegateVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetUndelegateVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

// GetUndelegateVersion returns the latest undelegation version.
func (k Keeper) GetUndelegateVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetUndelegateRecord writes a record of the user's undelegation actions.
func (k Keeper) SetUndelegateRecord(ctx sdk.Context, record *types.UndelegateRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	bz := k.cdc.MustMarshal(record)
	newStoreKey := record.ZoneId + record.Delegator
	store.Set([]byte(newStoreKey), bz)
}

// GetUndelegateRecord returns the record corresponding to zoneId and delegator among the user's undelegation records.
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

// GetAllUndelegateRecord returns all undelegate records corresponding to zoneId.
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

// GetUndelegateAmount gets the information that corresponds to the zone during the de-delegation history.
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, snDenom string, zone ibcstakingtypes.RegisteredZone, version uint64) (sdk.Coin, sdk.Int) {
	snAsset := sdk.NewCoin(snDenom, sdk.NewInt(0))
	wAsset := sdk.NewInt(0)

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		if undelegateRecord.ZoneId == zone.ZoneId {
			for _, record := range undelegateRecord.Records {
				if record.OracleVersion < version && record.State == types.UndelegateRequestUser {
					withdrawAsset, err := k.GetWithdrawAmt(ctx, *record.SnAssetAmount)
					if err != nil {
						return false
					}
					record.WithdrawAmount = withdrawAsset.Amount

					record.State = types.UndelegateRequestIca
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

// ChangeUndelegateState changes the status for recorded undelegation.
// UNDELEGATE_REQUEST_USER : Just requested undelegate by user. It is not in undelegate period.
// UNDELEGATE_REQUEST_ICA  : Requested by ICA, It is in undelegate period.
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {
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

// GetWithdrawAmt is used for calculating the amount of coin user can withdraw
// after un-delegate. This function is executed when ICA un-delegate call executed,
// and calculate using the balance of user's share coin.
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (sdk.Coin, error) {
	baseDenom := k.ibcstakingKeeper.GetBaseDenomForSnDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)

	oracleInfo, err := k.oracleKeeper.GetChainState(ctx, baseDenom)
	if err != nil {
		return sdk.Coin{}, err
	}

	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)

	convOracleAmt := k.ConvertWAssetToSnAssetDecimal(oracleInfo.Coin.Amount.BigInt(), zoneInfo.Decimal, zoneInfo.BaseDenom)
	withdrawAmt := k.CalculateWithdrawAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), convOracleAmt.Amount.BigInt())
	wAsset := k.ConvertSnAssetToWAssetDecimal(withdrawAmt, zoneInfo.Decimal, baseDenom)

	return wAsset, nil
}

// SetUndelegateRecordVersion navigates undelegate records and updates version for records corresponding to zoneId and state.
func (k Keeper) SetUndelegateRecordVersion(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType, version uint64) bool {
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

// DeleteUndelegateRecords deletes records corresponding to zoneId and state for undelegate records.
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {
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

// IterateUndelegatedRecords navigates de-delegation records.
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
