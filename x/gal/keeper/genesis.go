package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/v2/x/gal/types"
	icacontroltypes "github.com/Carina-labs/nova/v2/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// InitGenesis initializes the gal module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, deposit := range genState.DepositRecord {
		k.SetDepositRecord(ctx, deposit)
	}
	for _, delegate := range genState.DelegateRecord {
		k.SetDelegateRecord(ctx, delegate)
	}
	for _, undelegate := range genState.UndelegateRecord {
		k.SetUndelegateRecord(ctx, undelegate)
	}
	for _, withdraw := range genState.WithdrawRecord {
		k.SetWithdrawRecord(ctx, withdraw)
	}

	for _, zone := range genState.RecordInfo {
		k.SetDelegateVersion(ctx, zone.DelegateTrace.ZoneId, *zone.DelegateTrace)
		k.SetUndelegateVersion(ctx, zone.UndelegateTrace.ZoneId, *zone.UndelegateTrace)
		k.SetWithdrawVersion(ctx, zone.WithdrawTrace.ZoneId, *zone.WithdrawTrace)
	}
}

// ExportGenesis returns the gal module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params:           k.GetParams(ctx),
		DepositRecord:    []*types.DepositRecord{},
		DelegateRecord:   []*types.DelegateRecord{},
		UndelegateRecord: []*types.UndelegateRecord{},
		WithdrawRecord:   []*types.WithdrawRecord{},
		RecordInfo:       []*types.RecordInfo{},
	}

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyDepositRecordInfo)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var depositRecord types.DepositRecord
		if err := proto.Unmarshal(iter.Value(), &depositRecord); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.DepositRecord = append(result.DepositRecord, &depositRecord)
	}

	iter = sdk.KVStorePrefixIterator(store, types.KeyDelegateRecordInfo)
	for ; iter.Valid(); iter.Next() {
		var delegateRecord types.DelegateRecord
		if err := proto.Unmarshal(iter.Value(), &delegateRecord); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.DelegateRecord = append(result.DelegateRecord, &delegateRecord)
	}

	iter = sdk.KVStorePrefixIterator(store, types.KeyUndelegateRecordInfo)
	for ; iter.Valid(); iter.Next() {
		var undelegateRecord types.UndelegateRecord
		if err := proto.Unmarshal(iter.Value(), &undelegateRecord); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.UndelegateRecord = append(result.UndelegateRecord, &undelegateRecord)
	}

	iter = sdk.KVStorePrefixIterator(store, types.KeyWithdrawRecordInfo)
	for ; iter.Valid(); iter.Next() {
		var withdrawRecord types.WithdrawRecord
		if err := proto.Unmarshal(iter.Value(), &withdrawRecord); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.WithdrawRecord = append(result.WithdrawRecord, &withdrawRecord)
	}

	var versionRecords []*types.RecordInfo
	k.icaControlKeeper.IterateRegisteredZones(ctx, func(index int64, zoneInfo icacontroltypes.RegisteredZone) (stop bool) {
		zoneId := zoneInfo.ZoneId
		delegateVersion := k.GetDelegateVersion(ctx, zoneId)
		undelegateVersion := k.GetUndelegateVersion(ctx, zoneId)
		withdrawVersion := k.GetWithdrawVersion(ctx, zoneId)

		versionRecord := types.RecordInfo{
			DelegateTrace:   &delegateVersion,
			UndelegateTrace: &undelegateVersion,
			WithdrawTrace:   &withdrawVersion,
		}
		versionRecords = append(versionRecords, &versionRecord)
		return false
	})

	return result
}
