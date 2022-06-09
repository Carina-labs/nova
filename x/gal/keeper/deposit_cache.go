package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func (k Keeper) getDepositCache(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}

func (k Keeper) CacheDepositAmt(ctx sdk.Context, depositor sdk.AccAddress, amt sdk.Coin) error {
	store := k.getDepositCache(ctx)
	data := make(map[string]string)

	data[types.KeyDenom] = amt.Denom
	data[types.KeyAmount] = amt.Amount.String()
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

	amt, err := strconv.ParseInt(result[types.KeyAmount], 10, 0)
	if err != nil {
		return nil, err
	}

	denom, ok := result[types.KeyDenom]
	if !ok {
		return nil, fmt.Errorf("denom is not registered")
	}

	return &types.QueryCachedDepositAmountResponse{
		Address: depositorStr,
		Denom:   denom,
		Amount:  amt,
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
