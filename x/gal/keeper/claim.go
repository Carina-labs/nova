package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ClaimAndMintShareToken is used when user want to claim their share token.
// It calculates user's share and the amount of claimable share token.
func (k Keeper) ClaimAndMintShareToken(ctx sdk.Context, claimer sdk.AccAddress, asset sdk.Coin) error {
	snAsset := k.interTxKeeper.GetsnDenomForBaseDenom(ctx, asset.Denom)
	baseDenom := k.interTxKeeper.GetBaseDenomForSnDenom(ctx, snAsset)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, snAsset)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, baseDenom)
	if err != nil {
		return err
	}

	mintAmt := k.CalculateAlpha(asset.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())

	err = k.MintShareTokens(ctx, claimer, sdk.NewCoin(snAsset, sdk.NewIntFromBigInt(mintAmt)))
	if err != nil {
		return err
	}

	return nil
}

// CalculateAlpha calculates alpha value.
// Alpha = userDepositAmount / totalStakedAmount
// Delta = Alpha * totalShareTokenSupply
func (k Keeper) CalculateAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}

// GetWithdrawAmt is used for calculating the amount of coin user can withdraw
// after un-delegate. This function is executed when ICA un-delegate call executed,
// and calculate using the balance of user's share coin.
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (sdk.Coin, error) {
	baseAsset := k.interTxKeeper.GetBaseDenomForSnDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, baseAsset)
	if err != nil {
		return sdk.Coin{}, err
	}

	withdrawAmt := k.CalculateLambda(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())

	return sdk.NewInt64Coin(baseAsset, withdrawAmt.Int64()), nil
}

// CalculateLambda calculates lambda value.
// Lambda = userWithdrawAmount / totalStakedAmount
// Delta = Lambda * totalShareTokenSupply
func (k Keeper) CalculateLambda(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(burnedStTokenAmt, totalStakedAmount).Div(res, totalShareTokenSupply)
}
