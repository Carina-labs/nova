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

// GetAllAmountNotMintShareToken returns the sum of assets that have not yet been issued by the user among the assets that have been deposited.
func (k Keeper) GetAllAmountNotMintShareToken(ctx sdk.Context, zone *icacontroltypes.RegisteredZone) (sdk.Coin, error) {
	ibcDenom := k.icaControlKeeper.GetIBCHashDenom(zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, zone.BaseDenom)

	res := sdk.NewInt64Coin(ibcDenom, 0)
	k.IterateDelegateRecord(ctx, zone.ZoneId, func(_ int64, delegateRecord types.DelegateRecord) (stop bool) {
		for _, record := range delegateRecord.Records {
			if record.State == types.DelegateSuccess {
				res = res.Add(*record.Amount)
			}
		}
		return false
	})

	return res, nil
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
