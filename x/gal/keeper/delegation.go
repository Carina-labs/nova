package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getDelegateRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateRecordInfo)
}

func (k Keeper) SetDelegateRecord(ctx sdk.Context, msg *types.DelegateRecord) {
	store := k.getDelegateRecordStore(ctx)
	key := msg.ZoneId + msg.Claimer
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(key), bz)
}

func (k Keeper) SetDelegateRecords(ctx sdk.Context, zoneId string) {
	k.IterateDepositRecord(ctx, zoneId, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		for _, record := range depositRecord.Records {
			claimer, err := sdk.AccAddressFromBech32(record.Claimer)
			if err != nil {
				return
			}

			if record.State == types.DepositSuccess {
				delegateRecord, found := k.GetUserDelegateRecord(ctx, zoneId, claimer)
				if !found {
					delegateRecord = &types.DelegateRecord{
						ZoneId:  zoneId,
						Claimer: record.Claimer,
					}
				}
				if delegateRecord.Records == nil {
					delegateRecord.Records = make(map[types.DelegateVersion]*types.DelegateRecordContent)
				}

				versionInfo := k.GetDelegateVersion(ctx, zoneId)

				delegateRecordContent, found := delegateRecord.Records[versionInfo.CurrentVersion]
				if !found {
					delegateRecordContent = &types.DelegateRecordContent{
						Amount: record.Amount,
						State:  types.DelegateRequest,
					}
				} else {
					*delegateRecordContent.Amount = record.Amount.Add(*delegateRecordContent.Amount)
				}
				delegateRecord.Records[versionInfo.CurrentVersion] = delegateRecordContent
				k.SetDelegateRecord(ctx, delegateRecord)
			}
		}
		return false
	})
}

func (k Keeper) GetUserDelegateRecord(ctx sdk.Context, zoneId string, claimer sdk.AccAddress) (result *types.DelegateRecord, found bool) {
	store := k.getDelegateRecordStore(ctx)
	key := []byte(zoneId + claimer.String())
	if !store.Has(key) {
		return nil, false
	}

	res := store.Get(key)
	var record types.DelegateRecord
	k.cdc.MustUnmarshal(res, &record)
	return &record, true
}

// SetDelegateOracleVersion updates the Oracle version for recorded delegate requests.
// This action is required for the correct equity calculation.
func (k Keeper) SetDelegateOracleVersion(ctx sdk.Context, zoneId string, version types.DelegateVersion, oracleVersion uint64) {
	k.IterateDelegateRecord(ctx, zoneId, func(index int64, delegateRecord types.DelegateRecord) (stop bool) {
		_, found := delegateRecord.Records[version]
		if !found {
			return false
		}
		delegateRecord.Records[version].OracleVersion = oracleVersion
		k.SetDelegateRecord(ctx, &delegateRecord)
		return false
	})
}

func (k Keeper) ChangeDelegateState(ctx sdk.Context, zoneId string, version types.DelegateVersion) {
	k.IterateDelegateRecord(ctx, zoneId, func(index int64, delegateRecord types.DelegateRecord) (stop bool) {
		_, found := delegateRecord.Records[version]
		if !found {
			return false
		}

		delegateRecord.Records[version].State = types.DelegateSuccess
		k.SetDelegateRecord(ctx, &delegateRecord)
		return false
	})
}

func (k Keeper) GetTotalDelegateAmtForZoneId(ctx sdk.Context, zoneId, denom string, version types.DelegateVersion, state types.DelegateStatusType) sdk.Coin {
	totalDelegateAmt := sdk.NewCoin(denom, sdk.NewInt(0))

	k.IterateDelegateRecord(ctx, zoneId, func(index int64, delegateRecord types.DelegateRecord) (stop bool) {
		record, found := delegateRecord.Records[version]
		if !found {
			return false
		}

		if record.State == state && record.Amount.Denom == denom {
			totalDelegateAmt = totalDelegateAmt.Add(*record.Amount)
		}

		return false
	})

	return totalDelegateAmt
}

func (k Keeper) GetTotalDelegateAmtForUser(ctx sdk.Context, zoneId, denom string, userAddr sdk.AccAddress, state types.DelegateStatusType) sdk.Coin {
	totalDelegateAmt := sdk.NewCoin(denom, sdk.NewInt(0))

	record, found := k.GetUserDelegateRecord(ctx, zoneId, userAddr)
	if !found {
		return sdk.NewCoin(denom, sdk.NewInt(0))
	}

	for _, item := range record.Records {
		if item.State == state {
			totalDelegateAmt = totalDelegateAmt.Add(*item.Amount)
		}
	}

	return totalDelegateAmt
}

func (k Keeper) DeleteDelegateRecord(ctx sdk.Context, delegate *types.DelegateRecord) {
	for key, record := range delegate.Records {
		if record.State == types.DelegateSuccess {
			delete(delegate.Records, key)
		}
	}
	k.SetDelegateRecord(ctx, delegate)
}

// GetDelegateVersionStore returns store for delegation.
func (k Keeper) GetDelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
}

// SetDelegateVersion sets version for delegation corresponding to zone-id records.
func (k Keeper) SetDelegateVersion(ctx sdk.Context, zoneId string, trace types.VersionState) {
	store := k.GetDelegateVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

// GetDelegateVersion returns version for delegation corresponding to zone-id records.
func (k Keeper) GetDelegateVersion(ctx sdk.Context, zoneId string) types.VersionState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.VersionState
	k.cdc.MustUnmarshal(res, &record)

	return record
}

func (k Keeper) IsValidDelegateVersion(ctx sdk.Context, zoneId string, version uint64) bool {
	//get delegateState
	versionInfo := k.GetDelegateVersion(ctx, zoneId)
	if versionInfo.ZoneId == "" {
		versionInfo.ZoneId = zoneId
		versionInfo.CurrentVersion = 0
		versionInfo.Record = make(map[uint64]*types.IBCTrace)
		versionInfo.Record[0] = &types.IBCTrace{
			Version: 0,
			State:   types.IcaPending,
		}

		k.SetDelegateVersion(ctx, zoneId, versionInfo)
	}

	if versionInfo.CurrentVersion >= version && (versionInfo.Record[version].State == types.IcaPending || versionInfo.Record[version].State == types.IcaFail) {
		return true
	}
	return false
}

// IterateDelegateRecord navigates all delegate requests.
func (k Keeper) IterateDelegateRecord(ctx sdk.Context, zoneId string, fn func(index int64, delegateRecord types.DelegateRecord) (stop bool)) {
	store := k.getDelegateRecordStore(ctx)
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
		res := types.DelegateRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}
