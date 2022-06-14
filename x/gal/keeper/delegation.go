package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MintShareTokens mints st-token(share token) regard with deposited token.
func (k Keeper) MintShareTokens(ctx sdk.Context, depositor sdk.Address, amt sdk.Coin) error {
	depositorAddr, err := sdk.AccAddressFromBech32(depositor.String())
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) BurnShareTokens(ctx sdk.Context, burner sdk.Address, amt sdk.Coin) error {
	burnerAddr, err := sdk.AccAddressFromBech32(burner.String())
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, burnerAddr, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetPairToken(ctx sdk.Context, denom string, shareTokenDenom string) {
	data := make(map[string]string)
	data[denom] = shareTokenDenom
	k.paramSpace.Set(ctx, types.KeyWhiteListedTokenDenoms, data)
}

func (k Keeper) Share(context context.Context, rq *types.QueryCacheDepositAmountRequest) (*types.QueryCachedDepositAmountResponse, error) {
	return nil, nil
}
