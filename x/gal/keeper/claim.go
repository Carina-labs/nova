package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// wrapped coin -> st coin
func (k Keeper) ClaimAndMintShareToken(ctx sdk.Context, claimer string, amt sdk.Coin) {
	stAsset := k.interTxKeeper.GetstDenomForBaseDenom(ctx, amt.Denom)

	claimerAddr, err := sdk.AccAddressFromBech32(claimer)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	totalSharedToken := k.bankKeeper.GetSupply(ctx, stAsset)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, amt.Denom)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	mintAmt := k.CalculateAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())

	err = k.MintShareTokens(ctx, claimerAddr, sdk.NewCoin(stAsset, sdk.NewIntFromBigInt(mintAmt)))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}
}

func (k Keeper) GetShareTokenMintingAmt(ctx sdk.Context, amt sdk.Coin) (sdk.Coin, error) {
	stAsset := k.interTxKeeper.GetstDenomForBaseDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, stAsset)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, amt.Denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	mintAmt := k.CalculateAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())
	return sdk.NewInt64Coin(amt.Denom, mintAmt.Int64()), nil
}

func (k Keeper) CalculateAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}

// st coin -> wrapped coin
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (sdk.Coin, error) {
	baseAsset := k.interTxKeeper.GetBaseDenomForStDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, baseAsset)
	if err != nil {
		return sdk.Coin{}, err
	}

	withdrawAmt := k.CalculateLambda(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())
	return sdk.NewInt64Coin(baseAsset, withdrawAmt.Int64()), nil
}

func (k Keeper) CalculateLambda(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(burnedStTokenAmt, totalStakedAmount).Div(res, totalShareTokenSupply)
}
