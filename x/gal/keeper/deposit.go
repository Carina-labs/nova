package keeper

import (
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getDepositCache(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

func (k Keeper) RecordDepositAmt(ctx sdk.Context, msg types.DepositRecord) error {
	store := k.getDepositCache(ctx)
	bz := k.cdc.MustMarshal(&msg)
	store.Set([]byte(msg.Address), bz)
	return nil
}

func (k Keeper) GetRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) (*types.DepositRecord, error) {
	store := k.getDepositCache(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return nil, fmt.Errorf("depositor %s is not in state", depositor)
	}

	res := store.Get([]byte(depositorStr))

	var msg types.DepositRecord
	k.cdc.MustUnmarshal(res, &msg)
	return &msg, nil
}

func (k Keeper) ClearRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) error {
	store := k.getDepositCache(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return fmt.Errorf("depositor %s is not in state", depositor.String())
	}

	store.Delete([]byte(depositorStr))
	return nil
}
