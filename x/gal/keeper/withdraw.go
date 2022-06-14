package keeper

import (
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	WITHDRAW_REGISTER     = iota + 1
	WITHDRAW_REQUEST_USER = iota + 1
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

func (k Keeper) SetWithdrawRecords(ctx sdk.Context, zoneId string, state int64) {
	var withdrawRecords types.WithdrawRecord

	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId && undelegateInfo.State == state {
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

func (k Keeper) SetWithdrawTime(ctx sdk.Context, zoneId string, state int64, time time.Time) {
	k.IterateWithdrawdRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
		if withdrawInfo.ZoneId == zoneId && withdrawInfo.State == state {
			withdrawInfo.CompletionTime = time
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}

func (k Keeper) ClaimWithdrawAsset(ctx sdk.Context, withdrawer string, amt sdk.Coin) error {
	withdrawerAddr, err := sdk.AccAddressFromBech32(withdrawer)
	if err != nil {
		return err
	}

	// check record if user can withdraw asset
	enable, err := k.IsAbleToWithdraw(ctx, amt)
	if !enable {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawerAddr, sdk.NewCoins(amt))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) IsAbleToWithdraw(ctx sdk.Context, amt sdk.Coin) (bool, error) {
	goCtx := sdk.WrapSDKContext(ctx)
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	balance, err := k.bankKeeper.Balance(goCtx, &types2.QueryBalanceRequest{
		Address: moduleAddr.String(),
		Denom:   amt.Denom,
	})

	if err != nil {
		return false, fmt.Errorf("can't withdraw asset. Module have : %s, user request: %s",
			balance.Balance.Amount.String(), amt.Amount.String())
	}

	return balance.Balance.Amount.Int64() >= amt.Amount.Int64(), nil
}

// IterateWithdrawdRecords iterate
func (k Keeper) IterateWithdrawdRecords(ctx sdk.Context, fn func(index int64, withdrawInfo types.WithdrawRecord) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()
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
