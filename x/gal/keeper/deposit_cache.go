package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getDepositCache(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

func (k Keeper) CacheDepositAmt(ctx sdk.Context, depositor sdk.AccAddress, amt sdk.Coin) error {
	store := k.getDepositCache(ctx)
	data := make(map[string]string)
	data[types.KeyCoin] = amt.String()
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	store.Set([]byte(depositor.String()), bytes)
	return nil
}

func (k Keeper) GetCachedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) (*types.QueryCachedDepositAmountResponse, error) {
	store := k.getDepositCache(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return nil, fmt.Errorf("depositor %s is not in state", depositor)
	}

	result := make(map[string]string)
	err := json.Unmarshal(store.Get([]byte(depositorStr)), &result)
	if err != nil {
		return nil, err
	}

	coinStr, ok := result[types.KeyCoin]
	if !ok {
		return nil, fmt.Errorf("coin is not exist")
	}

	coin, err := sdk.ParseCoinNormalized(coinStr)
	if err != nil {
		return nil, err
	}

	return &types.QueryCachedDepositAmountResponse{
		Address: depositorStr,
		Amount:  coin,
	}, nil
}

func (k Keeper) ClearCachedDepositAmt(ctx sdk.Context, depositor sdk.AccAddress) error {
	store := k.getDepositCache(ctx)
	depositorStr := depositor.String()
	if !store.Has([]byte(depositorStr)) {
		return fmt.Errorf("depositor %s is not in state", depositor.String())
	}

	store.Delete([]byte(depositorStr))
	return nil
}
