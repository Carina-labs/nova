package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WithdrawRegisterType int

const (
	WITHDRAW_REGISTER WithdrawRegisterType = iota + 1
	TRANSFER_SUCCESS
)

func (k Keeper) getWithdrawRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
}

// GetWithdrawRecord returns withdraw record item by key.
func (k Keeper) GetWithdrawRecord(ctx sdk.Context, zoneId, withdrawer string) (*types.WithdrawRecord, error) {
	store := k.getWithdrawRecordStore(ctx)
	keyBytes := []byte(zoneId + withdrawer)
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
			// 현재 오라클 버전으로 계산
			withdrawAmount, err := k.GetWithdrawAmt(ctx, *undelegateInfo.Amount)
			if err != nil {
				return true
			}
			withdrawRecord.Amount = &withdrawAmount
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
	k.IterateWithdrawRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
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

// IterateWithdrawRecords iterate
func (k Keeper) IterateWithdrawRecords(ctx sdk.Context, fn func(index int64, withdrawInfo types.WithdrawRecord) (stop bool)) {
	store := k.getWithdrawRecordStore(ctx)
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

		res := types.WithdrawRecord{}

		k.cdc.MustUnmarshal(iterator.Value(), &res)
		stop := fn(i, res)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) UndelegateHistory(goCtx context.Context, rq *types.QueryUndelegateHistoryRequest) (*types.QueryUndelegateHistoryResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(sdkCtx, rq.Denom)
	if zoneInfo == nil {
		return nil, fmt.Errorf("can't find registered zone for denom : %s", rq.Denom)
	}

	udInfo, ok := k.GetUndelegateRecord(sdkCtx, zoneInfo.ZoneId+rq.Address)
	if !ok {
		return nil, fmt.Errorf("there is no undelegate data for address: %s, denom: %s", rq.Address, rq.Denom)
	}

	return &types.QueryUndelegateHistoryResponse{
		Address: rq.Address,
		Amount:  sdk.NewCoins(*udInfo.Amount),
	}, nil
}

func (k Keeper) WithdrawHistory(goCtx context.Context, rq *types.QueryWithdrawHistoryRequest) (*types.QueryWithdrawHistoryResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(sdkCtx, rq.Denom)
	if zoneInfo == nil {
		return nil, fmt.Errorf("can't find registered zone for denom : %s", rq.Denom)
	}

	// TODO : why GetWithdrawRecord returns error, not bool?
	wdInfo, err := k.GetWithdrawRecord(sdkCtx, zoneInfo.ZoneId, rq.Address)
	if err != nil {
		return nil, err
	}

	return &types.QueryWithdrawHistoryResponse{
		Address: rq.Address,
		Amount:  sdk.NewCoins(*wdInfo.Amount),
	}, nil
}

func (k Keeper) GetTotalWithdrawAmountForZoneId(ctx sdk.Context, zoneId string, blockTime time.Time) sdk.Coin {
	amount := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
	}

	k.IterateWithdrawRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
		if amount.Denom == "" {
			amount.Denom = withdrawInfo.Amount.Denom
		}
		if withdrawInfo.ZoneId == zoneId && withdrawInfo.State == int64(WITHDRAW_REGISTER) {
			if withdrawInfo.CompletionTime.Before(blockTime) {
				amount = amount.AddAmount(withdrawInfo.Amount.Amount)
			}
		}
		return false
	})
	return amount
}

func (k Keeper) ChangeWithdrawState(ctx sdk.Context, zoneId string, preState, postState WithdrawRegisterType) {
	k.IterateWithdrawRecords(ctx, func(index int64, withdrawInfo types.WithdrawRecord) (stop bool) {
		if withdrawInfo.ZoneId == zoneId && withdrawInfo.State == int64(preState) {
			withdrawInfo.State = int64(postState)
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}
