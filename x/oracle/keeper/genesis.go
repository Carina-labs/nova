package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, oracleInfo := range genState.OracleAddressInfo {
		k.SetOracleAddress(ctx, oracleInfo.ZoneId, oracleInfo.OracleAddress)
	}

	for i := range genState.States {
		if err := k.UpdateChainState(ctx, &genState.States[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params:            k.GetParams(ctx),
		OracleAddressInfo: []types.OracleAddressInfo{},
		States:            []types.ChainInfo{},
	}

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyOracleAddr)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var oracleInfo types.OracleAddressInfo
		if err := proto.Unmarshal(iter.Value(), &oracleInfo); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.OracleAddressInfo = append(result.OracleAddressInfo, oracleInfo)
	}

	iter = sdk.KVStorePrefixIterator(store, types.KeyOracleChainState)
	defer func() {
		err := iter.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}()

	for ; iter.Valid(); iter.Next() {
		value := types.ChainInfo{}
		if err := proto.Unmarshal(iter.Value(), &value); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.States = append(result.States, types.ChainInfo{
			Coin:            value.Coin,
			OperatorAddress: value.OperatorAddress,
			LastBlockHeight: value.LastBlockHeight,
			AppHash:         value.AppHash,
			ZoneId:          value.ZoneId,
		})
	}

	return result
}
