package keeper

import (
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WithdrawRegisterType int

const (
	WITHDRAW_REGISTER WithdrawRegisterType = iota + 1
	WITHDRAW_REQUEST_USER
)

func (k Keeper) getWithdrawRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
}

// GetWithdrawRecord returns withdraw record item by key.
func (k Keeper) GetWithdrawRecord(ctx sdk.Context, key string) (*types.WithdrawRecord, error) {
	store := k.getWithdrawRecordStore(ctx)
	keyBytes := []byte(key)
	if !store.Has(keyBytes) {
		return nil, types.ErrNoWithdrawRecord
	}

	res := store.Get(keyBytes)

	var withdrawRecord types.WithdrawRecord
	k.cdc.MustUnmarshal(res, &withdrawRecord)

	return &withdrawRecord, nil
}

// SetWithdrawRecord writes withdraw record.
func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record types.WithdrawRecord) {
	store := k.getWithdrawRecordStore(ctx)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

// DeleteWithdrawRecord removes withdraw record.
func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context, withdraw types.WithdrawRecord) {
	store := k.getWithdrawRecordStore(ctx)
	store.Delete([]byte(withdraw.ZoneId + withdraw.Withdrawer))
}

// SetWithdrawRecords write multiple withdraw record.
func (k Keeper) SetWithdrawRecords(ctx sdk.Context, zoneId string, state UndelegatedState) {
	var withdrawRecords []types.WithdrawRecord

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId && undelegateInfo.State == int64(state) {
			var withdrawRecord types.WithdrawRecord
			withdrawRecord.ZoneId = zoneId
			withdrawRecord.Withdrawer = undelegateInfo.Delegator
			amt, err := k.GetWithdrawAmt(ctx, *undelegateInfo.Amount)
			if err != nil {
				return true
			}
			withdrawRecord.Amount = &amt
			withdrawRecord.State = int64(WITHDRAW_REGISTER)
			withdrawRecords = append(withdrawRecords, withdrawRecord)
		}
		return false
	})

	if len(withdrawRecords) > 0 {
		for _, wr := range withdrawRecords {
			k.SetWithdrawRecord(ctx, wr)
		}
	}
}

// SetWithdrawTime writes the time undelegate finish.
func (k Keeper) SetWithdrawTime(ctx sdk.Context, zoneId string, state WithdrawRegisterType, time time.Time) {
	k.IterateWithdrawdRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
		if withdrawInfo.ZoneId == zoneId && withdrawInfo.State == int64(state) {
			withdrawInfo.CompletionTime = time
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}

// ClaimWithdrawAsset is used when user want to claim their asset which is after undeleagted.
func (k Keeper) ClaimWithdrawAsset(ctx sdk.Context, from sdk.AccAddress, withdrawer sdk.AccAddress, amt sdk.Coin) error {
	err := k.bankKeeper.SendCoins(ctx, from, withdrawer, sdk.NewCoins(amt))
	if err != nil {
		return err
	}

	return nil
}

// IsAbleToWithdraw returns if user can withdraw their asset.
// It refers nova ICA account. If ICA account's balance is greater than
// user withdraw request amount, this function returns true.
func (k Keeper) IsAbleToWithdraw(ctx sdk.Context, from sdk.AccAddress, amt sdk.Coin) bool {
	balance := k.bankKeeper.GetBalance(ctx, from, amt.Denom)
	return balance.Amount.BigInt().Cmp(amt.Amount.BigInt()) >= 0
}

// IterateWithdrawdRecords iterate
func (k Keeper) IterateWithdrawdRecords(ctx sdk.Context, fn func(index int64, withdrawInfo types.WithdrawRecord) (stop bool)) {
	store := k.getWithdrawRecordStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := (iterator).Close()
		if err != nil {
			panic(fmt.Errorf("unexpectedly iterator closed: %v", err))
		}
	}(iterator)
	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {

		res := types.WithdrawRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}
