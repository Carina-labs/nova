package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) ClaimAndMint(ctx sdk.Context, claimer string, amt sdk.Coin) {
	stAsset := k.interTxKeeper.GetstDenomForBaseDenom(ctx, amt.Denom)

	claimerAddr, err := sdk.AccAddressFromBech32(claimer)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	totalSharedToken := k.bankKeeper.GetSupply(ctx, stAsset)

	// alpha = user_deposit_amount / total_staked_amount
	alpha, err := k.calculateAlpha(ctx, amt.Denom, amt.Amount.Int64())
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	// minted_amount = alpha * total_share_token_supply
	mintAmt := sdk.NewInt(int64(alpha * float64(totalSharedToken.Amount.Int64())))
	err = k.MintShareTokens(ctx, claimerAddr, sdk.NewCoin(stAsset, mintAmt))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}
}
