package keeper

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// ClaimAndMintShareToken is used when user want to claim their share token.
// It calculates user's share and the amount of claimable share token.
func (k Keeper) ClaimAndMintShareToken(ctx sdk.Context, claimer sdk.AccAddress, asset sdk.Coin, transferPortId, transferChanId string) (sdk.Coin, error) {
	snDenom, err := k.GetsnDenomForIBCDenom(ctx, asset.Denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	baseDenom := k.ibcstakingKeeper.GetBaseDenomForSnDenom(ctx, snDenom)

	totalSnSupply := k.bankKeeper.GetSupply(ctx, snDenom)
	totalStakedAmount, err := k.GetTotalStakedForLazyMinting(ctx, baseDenom, transferPortId, transferChanId)
	if err != nil {
		return sdk.Coin{}, err
	}
	mintAmt := k.CalculateDepositAlpha(asset.Amount.BigInt(), totalSnSupply.Amount.BigInt(), totalStakedAmount.Amount.BigInt())
	err = k.MintShareTokens(ctx, claimer, sdk.NewCoin(snDenom, sdk.NewIntFromBigInt(mintAmt)))
	if err != nil {
		return sdk.Coin{}, err
	}

	// GetZoneInfo
	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)
	err = k.DeleteRecordedDepositItem(ctx, zoneInfo.ZoneId, claimer, DELEGATE_SUCCESS)
	if err != nil {
		return sdk.Coin{}, err
	}

	return sdk.NewInt64Coin(snDenom, mintAmt.Int64()), nil
}

// CalculateDepositAlpha calculates alpha value.
// Alpha = userDepositAmount / totalStakedAmount
// Delta = Alpha * totalShareTokenSupply
func (k Keeper) CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	if totalShareTokenSupply.Cmp(big.NewInt(0)) == 0 {
		totalShareTokenSupply = totalStakedAmount
	}
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}

// GetWithdrawAmt is used for calculating the amount of coin user can withdraw
// after un-delegate. This function is executed when ICA un-delegate call executed,
// and calculate using the balance of user's share coin.
func (k Keeper) GetWithdrawAmt(ctx sdk.Context, amt sdk.Coin) (sdk.Coin, error) {
	baseDenom := k.ibcstakingKeeper.GetBaseDenomForSnDenom(ctx, amt.Denom)
	totalSharedToken := k.bankKeeper.GetSupply(ctx, amt.Denom)
	totalStakedAmount, err := k.oracleKeeper.GetChainState(ctx, baseDenom)
	if err != nil {
		return sdk.Coin{}, err
	}

	withdrawAmt := k.CalculateWithdrawAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())
	return sdk.NewInt64Coin(baseDenom, withdrawAmt.Int64()), nil
}

// CalculateWithdrawAlpha calculates lambda value.
// Lambda = userWithdrawAmount / totalStakedAmount
// Delta = Lambda * totalShareTokenSupply
func (k Keeper) CalculateWithdrawAlpha(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	if totalShareTokenSupply.Cmp(big.NewInt(0)) == 0 {
		totalShareTokenSupply = totalStakedAmount
	}
	return res.Mul(burnedStTokenAmt, totalStakedAmount).Div(res, totalShareTokenSupply)
}

func (k Keeper) GetsnDenomForIBCDenom(ctx sdk.Context, ibcDenom string) (string, error) {
	denom, err := k.ibcTransferKeeper.DenomPathFromHash(ctx, ibcDenom)
	if err != nil {
		return "", err
	}

	denomTrace := transfertypes.ParseDenomTrace(denom)

	snDenom := k.ibcstakingKeeper.GetsnDenomForBaseDenom(ctx, denomTrace.BaseDenom)

	return snDenom, nil
}

func (k Keeper) GetTotalStakedForLazyMinting(ctx sdk.Context, denom, transferPortId, transferChanId string) (sdk.Coin, error) {
	zone := k.ibcstakingKeeper.GetZoneForDenom(ctx, denom)
	if zone == nil {
		return sdk.Coin{}, fmt.Errorf("cannot find zone denom : %s", denom)
	}

	chainInfo, err := k.oracleKeeper.GetChainState(ctx, denom)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("cannot find zone denom in oracle : %s", denom)
	}

	unMintedAmount, err := k.GetAllAmountNotMintShareToken(ctx, zone.ZoneId, transferPortId, transferChanId)
	if err != nil {
		return sdk.Coin{}, err
	}

	ibcDenom := k.ibcstakingKeeper.GetIBCHashDenom(ctx, transferPortId, transferChanId, denom)
	chainBalanceWithIbcDenom := sdk.Coin{
		Denom:  ibcDenom,
		Amount: chainInfo.Coin.Amount,
	}

	if chainBalanceWithIbcDenom.Sub(unMintedAmount).IsZero() {
		return unMintedAmount, nil
	}
	return chainBalanceWithIbcDenom.Sub(unMintedAmount), nil
}
