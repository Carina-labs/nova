package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genesisState *types.GenesisState) {
	k.SetParams(ctx, genesisState.Params)
	for i := range genesisState.PoolInfo.Pools {
		if err := k.CreatePool(ctx, genesisState.PoolInfo.Pools[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params: k.GetParams(ctx),
		PoolInfo: types.PoolInfo{
			TotalWeight: 0,
			Pools:       []*types.Pool{},
		},
	}

	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer func() {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}()

	for ; iterator.Valid(); iterator.Next() {
		value := types.Pool{}
		if err := proto.Unmarshal(iterator.Value(), &value); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.PoolInfo.Pools = append(result.PoolInfo.Pools, &types.Pool{
			PoolId:              value.PoolId,
			PoolContractAddress: value.PoolContractAddress,
			Weight:              value.Weight,
		})

		result.PoolInfo.TotalWeight += value.Weight
	}

	return result
}
