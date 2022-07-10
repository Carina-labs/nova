package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

// SetDepositAmt write the amount of coin user deposit to the "DepositRecord" store.
func (k Keeper) SetDepositAmt(ctx sdk.Context, msg *types.DepositRecord) {
	store := k.getDepositRecordStore(ctx)
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(msg.Address), bz)
}

func (k Keeper) MarkRecordTransfer(ctx sdk.Context, addr string, i int) error {
	store := k.getDepositRecordStore(ctx)

	var record types.DepositRecord
	k.cdc.MustUnmarshal(store.Get([]byte(addr)), &record)

	if len(record.Records) <= i {
		return types.ErrCanNotReplaceRecord
	}

	record.Records[i].IsTransferred = true
	k.SetDepositAmt(ctx, &record)

	return nil
}

// GetRecordedDepositAmt returns the amount of coin user deposit by address.
func (k Keeper) GetRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) (*types.DepositRecord, error) {
	store := k.getDepositRecordStore(ctx)
	depositorStr := depositor.String()
	key := []byte(depositorStr)
	if !store.Has(key) {
		return nil, types.ErrNoDepositRecord
	}

	res := store.Get(key)

	var msg types.DepositRecord
	k.cdc.MustUnmarshal(res, &msg)
	return &msg, nil
}

// ClearRecordedDepositAmt remove all data in "DepositRecord".
// It must be removed after staking in host chain.
func (k Keeper) ClearRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) error {
	store := k.getDepositRecordStore(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return sdkerrors.Wrap(types.ErrNoDepositRecord, fmt.Sprintf("account: %s", depositorStr))
	}

	store.Delete([]byte(depositorStr))
	return nil
}

func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, depositor sdk.AccAddress, amount sdk.Coin) error {
	record, err := k.GetRecordedDepositAmt(ctx, depositor)
	if err != nil {
		return err
	}

	recordItems := record.Records
	isDeleted := false
	for index, item := range record.Records {
		if item.IsTransferred && item.Amount.IsEqual(amount) {
			recordItems = removeIndex(recordItems, index)
			record.Records = recordItems
			isDeleted = true
			break
		}
	}

	if isDeleted {
		k.SetDepositAmt(ctx, record)
	}

	return nil
}

func (k Keeper) GetAllAmountNotMintShareToken(ctx sdk.Context, zoneId string) (sdk.Coin, error) {
	targetZoneInfo, ok := k.ibcstakingKeeper.GetRegisteredZone(ctx, zoneId)
	if !ok {
		return sdk.Coin{}, fmt.Errorf("cannot find zone id : %s", zoneId)
	}

	res := sdk.NewInt64Coin(targetZoneInfo.BaseDenom, 0)
	k.IterateDepositRecord(ctx, func(_ int64, depositRecord types.DepositRecord) (stop bool) {
		for _, record := range depositRecord.Records {
			if record.ZoneId == zoneId && record.IsTransferred == true {
				res = res.Add(*record.Amount)
			}
		}
		return false
	})

	return res, nil
}

func (k Keeper) IterateDepositRecord(ctx sdk.Context, fn func(index int64, depositRecord types.DepositRecord) (stop bool)) {
	store := k.getDepositRecordStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		iterator.Close()
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

func removeIndex(s []*types.DepositRecordContent, index int) []*types.DepositRecordContent {
	return append(s[:index], s[index+1:]...)
}
