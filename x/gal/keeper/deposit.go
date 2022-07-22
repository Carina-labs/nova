package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DepositState int64

const (
	DEPOSIT_REQUEST DepositState = iota + 1
	DEPOSIT_SUCCESS
	DELEGATE_REQUEST
	DELEGATE_SUCCESS
)

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

// SetDepositAmt write the amount of coin user deposit to the "DepositRecord" store.
func (k Keeper) SetDepositAmt(ctx sdk.Context, msg *types.DepositRecord) {
	store := k.getDepositRecordStore(ctx)
	key := msg.ZoneId + msg.Address
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(key), bz)
}

// func (k Keeper) MarkRecordTransfer(ctx sdk.Context, zoneId, addr string, i int) error {
// 	store := k.getDepositRecordStore(ctx)
// 	key := zoneId + addr

// 	var record types.DepositRecord
// 	k.cdc.MustUnmarshal(store.Get([]byte(key)), &record)

// 	if len(record.Records) <= i {
// 		return types.ErrCanNotReplaceRecord
// 	}

// 	record.Records[i].IsTransferred = true
// 	k.SetDepositAmt(ctx, &record)

// 	return nil
// }

// GetRecordedDepositAmt returns the amount of coin user deposit by address.
func (k Keeper) GetRecordedDepositAmt(ctx sdk.Context, zoneId string, depositor sdk.AccAddress) (*types.DepositRecord, error) {
	store := k.getDepositRecordStore(ctx)
	depositorStr := depositor.String()
	key := []byte(zoneId + depositorStr)
	if !store.Has(key) {
		return nil, types.ErrNoDepositRecord
	}

	res := store.Get(key)

	var msg types.DepositRecord
	k.cdc.MustUnmarshal(res, &msg)
	return &msg, nil
}

// Iterate로 depositRecord.content 반복하면서 해당 zone, state의 records 조회
// func (k Keeper) GetAllRecordedDepositAmt() {

// }

// delegate를 위한 TotlaDepositAmt로, 해당 zone에서 deposit 완료된 금액을 받아서 반환하는 함수
func (k Keeper) GetTotalDepositAmtForZoneId(ctx sdk.Context, zoneId, denom string, state DepositState) sdk.Coin {
	// zoneId에 해당하면서 state가 transfer인 친구들의 합
	totalDepositAmt := sdk.Coin{
		Amount: sdk.NewIntFromUint64(0),
		Denom:  denom,
	}
	// var check bool
	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		// check = false
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == int64(state) {
					totalDepositAmt = totalDepositAmt.Add(*record.Amount)
				}
			}
		}
		return false
	})

	if totalDepositAmt.Amount.IsZero() {
		return totalDepositAmt
	}

	return totalDepositAmt
}

func (k Keeper) SetBlockHeight(ctx sdk.Context, zoneId string, state DepositState, blockHeight uint64) {
	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		isChanged := false
		// check = false
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == int64(state) {
					record.BlockHeight = blockHeight
					isChanged = true
				}
			}

			if isChanged {
				k.SetDepositAmt(ctx, &depositRecord)
			}
		}

		return false
	})

}

// state 변경
func (k Keeper) ChangeDepositState(ctx sdk.Context, zoneId string, preState, postState DepositState) bool {
	isChanged := false

	k.IterateDepositRecord(ctx, func(index int64, depositRecord types.DepositRecord) (stop bool) {
		stateCheck := false
		if depositRecord.ZoneId == zoneId {
			for _, record := range depositRecord.Records {
				if record.State == int64(preState) {
					record.State = int64(postState)
					stateCheck = true
				}
			}
			if stateCheck {
				k.SetDepositAmt(ctx, &depositRecord)
				isChanged = true
			}
		}
		return false
	})

	if !isChanged {
		return isChanged
	}

	return true
}

// ClearRecordedDepositAmt remove all data in "DepositRecord".
// It must be removed after staking in host chain.
func (k Keeper) ClearRecordedDepositAmt(ctx sdk.Context, zoneId string, depositor sdk.AccAddress) error {
	key := zoneId + depositor.String()
	store := k.getDepositRecordStore(ctx)
	if !store.Has([]byte(key)) {
		return sdkerrors.Wrap(types.ErrNoDepositRecord, fmt.Sprintf("account: %s", depositor.String()))
	}

	store.Delete([]byte(key))
	return nil
}

func (k Keeper) DeleteRecordedDepositItem(ctx sdk.Context, zoneId string, depositor sdk.AccAddress, state DepositState) error {
	record, err := k.GetRecordedDepositAmt(ctx, zoneId, depositor)
	if err != nil {
		return err
	}

	recordItems := record.Records

	isDeleted := false
	for index, item := range record.Records {
		if item.State == int64(state) {
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

	// TODO : channel information should be contained in zone.
	ibcDenom := k.ibcstakingKeeper.GetIBCHashDenom(ctx,
		"transfer", "channel-0", targetZoneInfo.BaseDenom)
	res := sdk.NewInt64Coin(ibcDenom, 0)
	k.IterateDepositRecord(ctx, func(_ int64, depositRecord types.DepositRecord) (stop bool) {
		if depositRecord.ZoneId != zoneId {
			return false
		}
		for _, record := range depositRecord.Records {
			if record.IsTransferred {
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

func removeIndex(s []*types.DepositRecordContent, index int) []*types.DepositRecordContent {
	return append(s[:index], s[index+1:]...)
}
