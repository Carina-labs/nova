package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	icacontroltypes "github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

// SetDepositRecord stores the deposit record which stores amount of the coin and user address.
func (k Keeper) SetDepositRecord(ctx sdk.Context, msg *types.DepositRecord) {
	store := k.getDepositRecordStore(ctx)
	key := msg.ZoneId + msg.Claimer
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(key), bz)
}

// GetUserDepositRecord returns the amount of coin user deposit by address.
func (k Keeper) GetUserDepositRecord(ctx sdk.Context, zoneId string, claimer sdk.AccAddress) (result *types.DepositRecord, found bool) {
	store := k.getDepositRecordStore(ctx)
	key := []byte(zoneId + claimer.String())
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
	totalDepositAmt := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  denom,
	}

	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == state && record.Amount.Denom == denom {
					totalDepositAmt = totalDepositAmt.Add(*record.Amount)
				}
			}
		}
		return false
	})

	return totalDepositAmt
}

// GetTotalDepositAmtForUserAddr returns the sum of the user's address entered as input
// and the deposit coin corresponding to the coin denom.
func (k Keeper) GetTotalDepositAmtForUserAddr(ctx sdk.Context, userAddr, denom string) sdk.Coin {
	totalDepositAmt := sdk.NewCoin(denom, sdk.NewInt(0))

	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.Claimer == userAddr {
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

// SetDepositOracleVersion updates the Oracle version for recorded Deposit requests.
// This action is required for the correct equity calculation.
func (k Keeper) SetDepositOracleVersion(ctx sdk.Context, zoneId string, state types.DepositStatusType, oracleVersion uint64) {
	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		isChanged := false
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == state && record.OracleVersion == 0 {
					record.OracleVersion = oracleVersion
					isChanged = true
				}
			}

			if isChanged {
				k.SetDepositRecord(ctx, &depositRecord)
			}
		}

		return false
	})

}

// ChangeDepositState updates the deposit records corresponding to the preState to postState.
// This operation runs in the hook after the remote deposit is run.
func (k Keeper) ChangeDepositState(ctx sdk.Context, zoneId string, preState, postState types.DepositStatusType) bool {
	isChanged := false

	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		stateCheck := false
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == preState {
					record.State = postState
					stateCheck = true
				}
			}
			if stateCheck {
				k.SetDepositRecord(ctx, &depositRecord)
				isChanged = true
			}
		}
		return false
	})

	if !isChanged {
		return isChanged
	}

	return true
}

// SetDelegateRecordVersion updates the deposit version performed by the bot for the state of the deposit records corresponding to zoneId.
func (k Keeper) SetDelegateRecordVersion(ctx sdk.Context, zoneId string, state types.DepositStatusType, version uint64) bool {
	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.ZoneId == zoneId {
			isChanged := false
			for _, record := range depositRecord.Records {
				if record.State == state {
					isChanged = true
					record.DelegateVersion = version
				}
			}
			if isChanged {
				k.SetDepositRecord(ctx, &depositRecord)
			}
		}
		return false
	})

	return true
}

// DeleteRecordedDepositItem deletes the records corresponding to state among the defositor's assets deposited in the zone corresponding to zoneId.
func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, zoneId string, depositor sdk.AccAddress, state types.DepositStatusType) error {
	record, found := k.GetUserDepositRecord(ctx, zoneId, depositor)
	if !found {
		return types.ErrNoDepositRecord
	}

	var recordItems []*types.DepositRecordContent
	for _, item := range record.Records {
		if item.State != state {
			recordItems = append(recordItems, item)
		}
	}

	isDeleted := len(recordItems) < len(record.Records)
	if isDeleted {
		record.Records = recordItems
		k.SetDepositRecord(ctx, record)
		return nil
	}

	return types.ErrNoDeleteRecord
}

// GetAllAmountNotMintShareToken returns the sum of assets that have not yet been issued by the user among the assets that have been deposited.
func (k Keeper) GetAllAmountNotMintShareToken(ctx sdk.Context, zone *icacontroltypes.RegisteredZone) (sdk.Coin, error) {
	ibcDenom := k.icaControlKeeper.GetIBCHashDenom(zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, zone.BaseDenom)

	res := sdk.NewInt64Coin(ibcDenom, 0)
	k.IterateDepositRecord(ctx, func(_ int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.ZoneId == zone.ZoneId {
			for _, record := range depositRecord.Records {
				if record.State == types.DelegateSuccess {
					res = res.Add(*record.Amount)
				}
			}
		}
		return false
	})

	return res, nil
}

// IterateDepositRecord navigates all deposit requests.
func (k Keeper) IterateDepositRecord(ctx sdk.Context, fn func(index int64, depositRecord types.DepositRecord) (stop bool)) {
	store := k.getDepositRecordStore(ctx)
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
		res := types.DepositRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}
