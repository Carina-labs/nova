package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	for _, chainInfo := range genState.States {
		if err := k.UpdateChainState(ctx, &types.MsgUpdateChainState{
			ChainDenom:    chainInfo.ChainDenom,
			StakedBalance: uint64(chainInfo.TotalStakedBalance),
			Decimal:       uint64(chainInfo.Decimal),
			BlockHeight:   uint64(chainInfo.LastBlockHeight),
		}); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		States: []types.ChainInfo{},
	}

	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer func() {
		err := iterator.Close()
		if err != nil {
			panic(fmt.Errorf("unexpectedly iterator closed: %v", err))
		}
	}()

	for ; iterator.Valid(); iterator.Next() {
		value := make(map[string]uint64)
		if err := json.Unmarshal(iterator.Value(), &value); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.States = append(result.States, types.ChainInfo{
			ChainDenom:         string(iterator.Key()),
			ValidatorAddress:   "", // TODO; what is validatorAddres ?
			LastBlockHeight:    int64(value["height"]),
			TotalStakedBalance: int64(value["balance"]),
			Decimal:            int64(value["decimal"]),
		})
	}

	return result
}
