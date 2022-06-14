package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) GetWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record types.WithdrawRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) ClaimWithdrawAsset(ctx sdk.Context, withdrawer string, amt sdk.Coin) error {
	withdrawerAddr, err := sdk.AccAddressFromBech32(withdrawer)
	if err != nil {
		return err
	}

	// check record if user can withdraw asset
	enable, err := k.IsAbleToWithdraw(ctx, amt)
	if !enable {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawerAddr, sdk.NewCoins(amt))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) IsAbleToWithdraw(ctx sdk.Context, amt sdk.Coin) (bool, error) {
	goCtx := sdk.WrapSDKContext(ctx)
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	balance, err := k.bankKeeper.Balance(goCtx, &types2.QueryBalanceRequest{
		Address: moduleAddr.String(),
		Denom:   amt.Denom,
	})

	if err != nil {
		return false, fmt.Errorf("can't withdraw asset. Module have : %s, user request: %s",
			balance.Balance.Amount.String(), amt.Amount.String())
	}

	return balance.Balance.Amount.Int64() >= amt.Amount.Int64(), nil
}
