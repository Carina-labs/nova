package keeper

import (
	"fmt"

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
func (k Keeper) SetDepositAmt(ctx sdk.Context, msg *types.DepositRecord) error {
	store := k.getDepositRecordStore(ctx)
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(msg.Address), bz)
	return nil
}

func (k Keeper) MarkRecordTransfer(ctx sdk.Context, addr string, i int) error {
	store := k.getDepositRecordStore(ctx)

	var record types.DepositRecord
	k.cdc.MustUnmarshal(store.Get([]byte(addr)), &record)

	if len(record.Records) <= i {
		return fmt.Errorf("can't replace record")
	}

	record.Records[i].IsTransferred = true
	err := k.SetDepositAmt(ctx, &record)
	if err != nil {
		return err
	}

	return nil
}

// GetRecordedDepositAmt returns the amount of coin user deposit by address.
func (k Keeper) GetRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) (*types.DepositRecord, error) {
	store := k.getDepositRecordStore(ctx)
	key := []byte(depositor.String())
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
		return fmt.Errorf("depositor %s is not in state", depositor.String())
	}

	store.Delete([]byte(depositorStr))
	return nil
}
