package keeper

import (
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	WITHDRAW_REGISTER     = iota + 1
	WITHDRAW_REQUEST_USER = iota + 1
)

// GetWithdrawRecord returns withdraw record item by key.
func (k Keeper) GetWithdrawRecord(ctx sdk.Context, key string) (types.WithdrawRecord, bool) {
	withdrawInfo := types.WithdrawRecord{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	bz := store.Get([]byte(key))

	if len(bz) == 0 {
		return withdrawInfo, false
	}

	k.cdc.MustUnmarshal(bz, &withdrawInfo)

	return withdrawInfo, true
}

// SetWithdrawRecord writes withdraw record.
func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record types.WithdrawRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

// DeleteWithdrawRecord removes withdraw record.
func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context, withdraw types.WithdrawRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyUndelegateRecordInfo)
	store.Delete([]byte(withdraw.ZoneId + withdraw.Withdrawer))
}

// SetWithdrawRecords write multiple withdraw record.
func (k Keeper) SetWithdrawRecords(ctx sdk.Context, zoneId string, state UndelegatedState) {
	var withdrawRecords types.WithdrawRecord

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId && undelegateInfo.State == int64(state) {
			withdrawRecords.ZoneId = zoneId
			withdrawRecords.Withdrawer = undelegateInfo.Delegator
			amt, err := k.GetWithdrawAmt(ctx, *undelegateInfo.Amount)
			if err != nil {
				return true
			}
			withdrawRecords.Amount = &amt
			withdrawRecords.State = WITHDRAW_REGISTER
		}
		return false
	})

}

// SetWithdrawTime writes the time undelegate finish.
func (k Keeper) SetWithdrawTime(ctx sdk.Context, zoneId string, state UndelegatedState, time time.Time) {
	k.IterateWithdrawdRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
		if withdrawInfo.ZoneId == zoneId && withdrawInfo.State == int64(state) {
			withdrawInfo.CompletionTime = time
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}

// ClaimWithdrawAsset is used when user want to claim their asset which is after undeleagted.
func (k Keeper) ClaimWithdrawAsset(ctx sdk.Context, withdrawer string, amt sdk.Coin) error {
	withdrawerAddr, err := sdk.AccAddressFromBech32(withdrawer)
	if err != nil {
		return err
	}

	// check record if user can withdraw asset
	enable := k.IsAbleToWithdraw(ctx, amt)
	if !enable {
		return fmt.Errorf("not enough balance to withdraw")
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawerAddr, sdk.NewCoins(amt))
	if err != nil {
		return err
	}

	return nil
}

// IsAbleToWithdraw returns if user can withdraw their asset.
// It refers nova ICA account. If ICA account's balance is greater than
// user withdraw request amount, this function returns true.
func (k Keeper) IsAbleToWithdraw(ctx sdk.Context, amt sdk.Coin) bool {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddr, amt.Denom)
	return balance.Amount.BigInt().Cmp(amt.Amount.BigInt()) >= 0
}

// IterateWithdrawdRecords iterate
func (k Keeper) IterateWithdrawdRecords(ctx sdk.Context, fn func(index int64, withdrawInfo types.WithdrawRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
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
