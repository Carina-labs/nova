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

	senderAddr, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		return err
	}

	// record
	record := &types.DepositRecord{
		ZoneId:        zoneInfo.ZoneId,
		Address:       senderAddr.String(),
		Amount:        &deposit.Amount[0],
		IsTransferred: false,
	}

	if err := k.RecordDepositAmt(ctx, record); err != nil {
		return err
	}

	err = k.TransferToTargetZone(ctx,
		zoneInfo.TransferConnectionInfo.PortId,
		zoneInfo.TransferConnectionInfo.ChannelId,
		senderAddr.String(),
		deposit.HostAddr,
		deposit.Amount[0])
	if err != nil {
		return err
	}

	return nil
}

// getDepositRecordStore returns "DepositRecord" store.
// It is used for finding the amount of coin user deposit.
func (k Keeper) getDepositRecordStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

// RecordDepositAmt write the amount of coin user deposit to the "DepositRecord" store.
func (k Keeper) RecordDepositAmt(ctx sdk.Context, msg *types.DepositRecord) error {
	store := k.getDepositRecordStore(ctx)
	bz := k.cdc.MustMarshal(msg)
	store.Set([]byte(msg.Address), bz)
	return nil
}

// GetRecordedDepositAmt returns the amount of coin user deposit by address.
func (k Keeper) GetRecordedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) (*types.DepositRecord, error) {
	store := k.getDepositRecordStore(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return nil, fmt.Errorf("depositor %s is not in state", depositor)
	}

	res := store.Get([]byte(depositorStr))

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
