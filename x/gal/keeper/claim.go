package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func (k Keeper) ClaimAndMint(ctx sdk.Context, claimer string, amt sdk.Coin) {
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

	mintAmt := k.CalculateMintAmount(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), big.NewInt(int64(totalStakedAmount.TotalStakedBalance)))

	err = k.MintShareTokens(ctx, claimerAddr, sdk.NewCoin(stAsset, sdk.NewIntFromBigInt(mintAmt)))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}
}

func (k Keeper) CalculateMintAmount(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}
