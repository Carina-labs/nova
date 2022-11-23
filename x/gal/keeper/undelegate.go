package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	icacontroltypes "github.com/Carina-labs/nova/x/icacontrol/types"
)

// GetUndelegateVersionStore returns the store that stores the UndelegateVersion data.
// The un-delegation task is periodically operated by the bot, so it stores the version for the last action.
func (k Keeper) GetUndelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
}

// SetUndelegateVersion sets the new un-delgate Version.
func (k Keeper) SetUndelegateVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {
	store := k.GetUndelegateVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetUndelegateVersion returns the latest undelegation version.
func (k Keeper) GetUndelegateVersion(ctx sdk.Context, zoneId string) types.VersionState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.VersionState
	k.cdc.MustUnmarshal(res, &record)

	return record
}

func (k Keeper) IsValidUndelegateVersion(ctx sdk.Context, zoneId string, version uint64) bool {
	//get undelegate version
	versionInfo := k.GetUndelegateVersion(ctx, zoneId)
	if versionInfo.ZoneId == "" {
		versionInfo.ZoneId = zoneId
		versionInfo.CurrentVersion = 0
		versionInfo.Record = make(map[uint64]*types.IBCTrace)
		versionInfo.Record[0] = &types.IBCTrace{
			Version: 0,
			State:   types.IcaPending,
		}

		k.SetUndelegateVersion(ctx, zoneId, versionInfo)
	}

	if versionInfo.CurrentVersion >= version && (versionInfo.Record[version].State == types.IcaPending || versionInfo.Record[version].State == types.IcaFail) {
		return true
	}

	return false
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
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "getAllUndelegateRecord")
	var undelegateInfo []*types.UndelegateRecord
	k.IterateUndelegatedRecords(ctx, zoneId, func(_ int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		undelegateInfo = append(undelegateInfo, undelegateRecord)
		return false
	})

	return undelegateInfo
}

// GetUndelegateAmount gets the information that corresponds to the zone during the de-delegation history.
func (k Keeper) GetUndelegateAmount(ctx sdk.Context, snDenom string, zone icacontroltypes.RegisteredZone, version uint64) (sdk.Coin, sdk.Int) {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "getUndelegateAmount")
	snAsset := sdk.NewCoin(snDenom, sdk.NewInt(0))
	wAsset := sdk.NewInt(0)

	k.IterateUndelegatedRecords(ctx, zone.ZoneId, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		for _, record := range undelegateRecord.Records {
			if record.OracleVersion < version && record.State == types.UndelegateRequestByUser {
				withdrawAsset, err := k.GetWithdrawAmt(ctx, *record.SnAssetAmount)
				if err != nil {
					return false
				}
				record.WithdrawAmount = withdrawAsset.Amount

				record.State = types.UndelegateRequestByIca
				wAsset = wAsset.Add(record.WithdrawAmount)
				snAsset = snAsset.Add(*record.SnAssetAmount)
			}
		}
		k.SetUndelegateRecord(ctx, undelegateRecord)
		return false
	})
	return snAsset, wAsset
}

func (k Keeper) GetReUndelegateAmount(ctx sdk.Context, snDenom string, zone icacontroltypes.RegisteredZone, version uint64) (sdk.Coin, sdk.Int) {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "getReUndelegateAmount")
	snAsset := sdk.NewCoin(snDenom, sdk.NewInt(0))
	wAsset := sdk.NewInt(0)

	k.IterateUndelegatedRecords(ctx, zone.ZoneId, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		for _, record := range undelegateRecord.Records {
			if record.OracleVersion < version && record.State == types.UndelegateRequestByIca {
				wAsset = wAsset.Add(record.WithdrawAmount)
				snAsset = snAsset.Add(*record.SnAssetAmount)
			}
		}
		return false
	})
	return snAsset, wAsset
}

// ChangeUndelegateState changes the status for recorded undelegation.
// UNDELEGATE_REQUEST_USER : Just requested undelegate by user. It is not in undelegate period.
// UNDELEGATE_REQUEST_ICA  : Requested by ICA, It is in undelegate period.
func (k Keeper) ChangeUndelegateState(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "changeUndelegateState")
	k.IterateUndelegatedRecords(ctx, zoneId, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		for _, record := range undelegateRecord.Records {
			record.State = state
		}
		k.SetUndelegateRecord(ctx, undelegateRecord)
		return false
	})
}

// GetWithdrawAmt is used for calculating the amount of coin user can withdraw
// after un-delegate. This function is executed when ICA un-delegate call executed,
// and calculate using the balance of user's share coin.
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (*sdk.Coin, error) {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "getWithdrawAmt")
	baseDenom := k.icaControlKeeper.GetBaseDenomForSnDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)

	oracleInfo, err := k.oracleKeeper.GetChainState(ctx, baseDenom)
	if err != nil {
		return nil, err
	}

	zoneInfo := k.icaControlKeeper.GetZoneForDenom(ctx, baseDenom)

	convOracleAmt, err := k.ConvertWAssetToSnAssetDecimal(oracleInfo.Coin.Amount.BigInt(), zoneInfo.Decimal, zoneInfo.BaseDenom)
	if err != nil {
		return nil, err
	}

	withdrawAmt := k.CalculateWithdrawAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), convOracleAmt.Amount.BigInt())
	wAsset, err := k.ConvertSnAssetToWAssetDecimal(withdrawAmt, zoneInfo.Decimal, baseDenom)
	if err != nil {
		return nil, err
	}

	return wAsset, nil
}

// SetUndelegateRecordVersion navigates undelegate records and updates version for records corresponding to zoneId and state.
func (k Keeper) SetUndelegateRecordVersion(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType, version uint64) bool {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "setUndelegateRecordVersion")
	k.IterateUndelegatedRecords(ctx, zoneId, func(index int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
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
		return false
	})

	return true
}

// DeleteUndelegateRecords deletes records corresponding to zoneId and state for undelegate records.
func (k Keeper) DeleteUndelegateRecords(ctx sdk.Context, zoneId string, state types.UndelegatedStatusType) {
	defer telemetry.MeasureSince(time.Now(), "gal", "undelegate", "deleteUndelegateRecords")
	k.IterateUndelegatedRecords(ctx, zoneId, func(_ int64, undelegateRecord *types.UndelegateRecord) (stop bool) {
		var recordItems []*types.UndelegateRecordContent
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
		return false
	})
}

func (k Keeper) HasMaxUndelegateEntries(undelegateRecords types.UndelegateRecord, maxEntries int64) bool {
	for _, record := range undelegateRecords.Records {
		if record.State == types.UndelegateRequestByUser {
			maxEntries -= 1
		}
		if maxEntries == 0 {
			return true
		}
	}
	return false
}

func (k Keeper) HasMaxDepositEntries(depositRecords types.DepositRecord, maxEntries int64) bool {
	for _, record := range depositRecords.Records {
		if record.State == types.DepositSuccess {
			maxEntries -= 1
		}
		if maxEntries == 0 {
			return true
		}
	}
	return false
}

// IterateUndelegatedRecords navigates de-delegation records.
func (k Keeper) IterateUndelegatedRecords(ctx sdk.Context, zoneId string, fn func(index int64, undelegateInfo *types.UndelegateRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
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
		var res types.UndelegateRecord
		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, &res)
		if stop {
			break
		}
		i++
	}
}
