package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getWithdrawRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
}

// SetWithdrawRecord writes withdraw record.
func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record *types.WithdrawRecord) {
	store := k.getWithdrawRecordStore(ctx)
	bz := k.cdc.MustMarshal(record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

// GetWithdrawRecord returns withdraw record item by key.
func (k Keeper) GetWithdrawRecord(ctx sdk.Context, zoneId, withdrawer string) (result *types.WithdrawRecord, found bool) {
	store := k.getWithdrawRecordStore(ctx)
	keyBytes := []byte(zoneId + withdrawer)
	if !store.Has(keyBytes) {
		return nil, false
	}

	res := store.Get(keyBytes)

	var record types.WithdrawRecord
	k.cdc.MustUnmarshal(res, &record)
	return &record, true
}

// DeleteWithdrawRecord removes withdraw record.
func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context, withdraw types.WithdrawRecord) {
	store := k.getWithdrawRecordStore(ctx)
	store.Delete([]byte(withdraw.ZoneId + withdraw.Withdrawer))
}

func (k Keeper) GetWithdrawVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawVersion)
}

func (k Keeper) SetWithdrawVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetWithdrawVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetWithdrawVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetWithdrawRecordVersion(ctx sdk.Context, zoneId string, state types.WithdrawStatusType, version uint64) bool {
	k.IterateWithdrawRecords(ctx, func(index int64, withdrawRecord *types.WithdrawRecord) (stop bool) {
		if withdrawRecord.ZoneId == zoneId {
			isChanged := false
			for _, record := range withdrawRecord.Records {
				if record.State == state {
					isChanged = true
					record.WithdrawVersion = version
				}
			}
			if isChanged {
				k.SetWithdrawRecord(ctx, withdrawRecord)
			}
		}
		return false
	})

	return true
}

// SetWithdrawRecords write multiple withdraw record.
func (k Keeper) SetWithdrawRecords(ctx sdk.Context, zoneId string, time time.Time) {
	k.IterateUndelegatedRecords(ctx, func(index int64, undelegateInfo *types.UndelegateRecord) (stop bool) {
		if undelegateInfo.ZoneId == zoneId {
			for _, items := range undelegateInfo.Records {
				if items.State == types.UndelegateRequestIca {
					withdrawRecord, found := k.GetWithdrawRecord(ctx, zoneId, items.Withdrawer)
					if !found {
						withdrawRecord = &types.WithdrawRecord{
							ZoneId:     zoneId,
							Withdrawer: items.Withdrawer,
						}
						withdrawRecord.Records = make(map[uint64]*types.WithdrawRecordContent)
					}

					withdrawRecordContent, found := withdrawRecord.Records[items.UndelegateVersion]

					if !found {
						withdrawRecordContent = &types.WithdrawRecordContent{
							State:           types.WithdrawStatusRegistered,
							WithdrawVersion: items.UndelegateVersion,
							Amount:          items.WithdrawAmount,
							CompletionTime:  time,
						}

					} else {
						withdrawRecordContent.Amount = items.WithdrawAmount.Add(withdrawRecordContent.Amount)
					}
					withdrawRecord.Records[items.UndelegateVersion] = withdrawRecordContent

					k.SetWithdrawRecord(ctx, withdrawRecord)
				}
			}

		}
		return false
	})
}

func (k Keeper) GetWithdrawAmountForUser(ctx sdk.Context, zoneId, denom string, withdrawer string) sdk.Coin {
	amount := sdk.NewCoin(denom, sdk.ZeroInt())

	withdrawRecord, found := k.GetWithdrawRecord(ctx, zoneId, withdrawer)
	if !found {
		k.Logger(ctx).Error("withdraw record not found", "func", "GetWithdrawAmountForUser", "zoneId", zoneId, "address", withdrawer)
		return amount
	}

	for _, record := range withdrawRecord.Records {
		if record.State == types.WithdrawStatusTransferred {
			amount.Amount = amount.Amount.Add(record.Amount)
		}
	}

	return amount
}

func (k Keeper) GetTotalWithdrawAmountForZoneId(ctx sdk.Context, zoneId, denom string, blockTime time.Time) sdk.Coin {
	amount := sdk.NewCoin(denom, sdk.ZeroInt())

	k.IterateWithdrawRecords(ctx, func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool) {
		if withdrawInfo.ZoneId == zoneId {
			for _, record := range withdrawInfo.Records {
				if record.CompletionTime.Before(blockTime) && record.State == types.WithdrawStatusRegistered {
					amount.Amount = amount.Amount.Add(record.Amount)
				}
			}
		}
		return false
	})
	return amount
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
func (k Keeper) IterateWithdrawRecords(ctx sdk.Context, fn func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool)) {
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
		stop := fn(i, &res)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) ChangeWithdrawState(ctx sdk.Context, preState, postState types.WithdrawStatusType) {
	k.IterateWithdrawRecords(ctx, func(index int64, withdrawInfo *types.WithdrawRecord) (stop bool) {
		for _, record := range withdrawInfo.Records {
			if record.State == preState {
				record.State = postState
			}
			k.SetWithdrawRecord(ctx, withdrawInfo)
		}
		return false
	})
}
