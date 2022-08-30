package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreatePool(ctx sdk.Context, poolId string, poolContractAddr string) error {
	store := k.getPoolStore(ctx)
	key := []byte(poolId)
	if store.Has(key) {
		return fmt.Errorf("pool is already exist. pool id : %s", poolId)
	}

	return nil
}

func (k Keeper) SetPoolWeight(ctx sdk.Context, poolId string, newWeight uint64) error {
	store := k.getPoolStore(ctx)
	key := []byte(poolId)
	res := store.Get(key)

	var pool types.Pool
	k.cdc.MustUnmarshal(res, &pool)
	pool.Weight = newWeight

	bz, err := k.cdc.Marshal(&pool)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}
