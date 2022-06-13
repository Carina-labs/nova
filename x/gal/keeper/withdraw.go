package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record types.WithdrawRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) GetWithdrawReceipt(ctx sdk.Context, key string) ([]byte, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawReceiptInfo)
	bk := []byte(key)
	if !store.Has(bk) {
		return nil, fmt.Errorf("The store does not have key %s", key)
	}

	return store.Get(bk), nil
}

func (k Keeper) SetWithdrawReceipt(ctx sdk.Context, receipt types.MsgWithdrawReceipt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawReceiptInfo)
	bz := k.cdc.MustMarshal(&receipt)
	store.Set([]byte(receipt.ZoneId+receipt.Withdrawer), bz)
}

func (k Keeper) DeleteWithdrawReceipt(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawReceiptInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		store.Delete(key)
	}
}

func (k Keeper) DeleteWithdrawReceiptItem(ctx sdk.Context, key string) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawReceiptInfo)
	bk := []byte(key)
	if !store.Has(bk) {
		return fmt.Errorf("The store does not have key %s", key)
	}

	store.Delete(bk)
	return nil
}
