package keeper

import (
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Deposit(ctx sdk.Context, deposit *types.MsgDeposit) error {
	// IBC transfer
	zoneInfo, ok := k.interTxKeeper.GetRegisteredZone(ctx, deposit.ZoneId)
	if !ok {
		return fmt.Errorf("can't find valid IBC zone, input zoneId: %s", deposit.ZoneId)
	}

	k.Logger(ctx).Info("ZoneInfo", "zoneInfo", zoneInfo)

	depositorAddr, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return err
	}

	record, err := k.GetRecordedDepositAmt(ctx, depositorAddr)
	if err == types.ErrNoDepositRecord {
		newRecord := &types.DepositRecordContent{
			ZoneId:        zoneInfo.ZoneId,
			Amount:        &deposit.Amount[0],
			IsTransferred: false,
		}
		if err := k.RecordDepositAmt(ctx, &types.DepositRecord{
			Address: deposit.Depositor,
			Records: []*types.DepositRecordContent{newRecord},
		}); err != nil {
			return err
		}
	} else {
		// append
		record.Records = append(record.Records, &types.DepositRecordContent{
			ZoneId:        zoneInfo.ZoneId,
			Amount:        &deposit.Amount[0],
			IsTransferred: false,
		})
		if err := k.RecordDepositAmt(ctx, record); err != nil {
			return err
		}
	}

	return k.TransferToTargetZone(ctx,
		zoneInfo.TransferConnectionInfo.PortId,
		zoneInfo.TransferConnectionInfo.ChannelId,
		deposit.Depositor,
		deposit.HostAddr,
		deposit.Amount[0])
}

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

// RecordDepositAmt write the amount of coin user deposit to the "DepositRecord" store.
func (k Keeper) RecordDepositAmt(ctx sdk.Context, msg *types.DepositRecord) error {
	store := k.getDepositRecordStore(ctx)

	// If no data in record, just set.
	if !store.Has([]byte(msg.Address)) {
		bz := k.cdc.MustMarshal(msg)

		store.Set([]byte(msg.Address), bz)
		return nil
	}

	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(msg.Address), bz)
	return nil
}

func (k Keeper) ReplaceDepositRecord(ctx sdk.Context, addr string, i int) error {
	store := k.getDepositRecordStore(ctx)

	var record types.DepositRecord
	k.cdc.MustUnmarshal(store.Get([]byte(addr)), &record)

	if len(record.Records) <= i {
		return fmt.Errorf("can't replace record")
	}

	record.Records = append(record.Records[:i], record.Records[i:]...)
	err := k.RecordDepositAmt(ctx, &record)
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
