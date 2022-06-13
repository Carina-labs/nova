package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
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

	mintAmt := k.CalculateMintAmount(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())

	err = k.MintShareTokens(ctx, claimerAddr, sdk.NewCoin(stAsset, sdk.NewIntFromBigInt(mintAmt)))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}
}

func (k Keeper) CalculateMintAmount(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}

// st coin -> wrapped coin
func (k Keeper) BurnShareTokenAndMintWrappedToken(ctx sdk.Context, burner string, amt sdk.Coin) {
	baseAsset := k.interTxKeeper.GetBaseDenomForStDenom(ctx, amt.Denom)

	burnerAddr, err := sdk.AccAddressFromBech32(burner)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, baseAsset)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	withdrawAmt := k.CalculateWithdrawAmount(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())

	err = k.BurnShareTokens(ctx, burnerAddr, amt)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(baseAsset, withdrawAmt.Int64())))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}
}

func (k Keeper) CalculateWithdrawAmount(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(burnedStTokenAmt, totalStakedAmount).Div(res, totalShareTokenSupply)
}
