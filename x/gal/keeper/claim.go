package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

const snAssetDecimal = 18

var (
	precisionMultipliers []*big.Int
)

func init() {
	precisionMultipliers = make([]*big.Int, snAssetDecimal+1)
	for i := 0; i <= snAssetDecimal; i++ {
		precisionMultipliers[i] = calcPrecisionMultiplier(int64(i))
	}
}

func precisionMultiplier(prec int64) *big.Int {
	if prec > snAssetDecimal {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", snAssetDecimal, prec))
	}
	return precisionMultipliers[prec]
}

func calcPrecisionMultiplier(prec int64) *big.Int {
	if prec > snAssetDecimal {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", snAssetDecimal, prec))
	}
	zerosToAdd := snAssetDecimal - prec
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(zerosToAdd), nil)
	return multiplier
}

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

	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)

	// convert decimal
	snAsset := k.ConvertWAssetToSnAssetDecimal(asset.Amount.BigInt(), zoneInfo.Decimal, snDenom)
	mintAmt := k.CalculateDepositAlpha(snAsset.Amount.BigInt(), totalSnSupply.Amount.BigInt(), totalStakedAmount.Amount.BigInt())
	mintCoin := sdk.NewCoin(zoneInfo.SnDenom, sdk.NewIntFromBigInt(mintAmt))

	err = k.MintShareTokens(ctx, claimer, mintCoin)
	if err != nil {
		return sdk.Coin{}, err
	}

	err = k.DeleteRecordedDepositItem(ctx, zoneInfo.ZoneId, claimer, DELEGATE_SUCCESS)
	if err != nil {
		return sdk.Coin{}, err
	}

	return mintCoin, nil
}

// TotalClaimableAssets returns the total amount of claimable snAsset.
func (k Keeper) TotalClaimableAssets(ctx sdk.Context, zone ibcstakingtypes.RegisteredZone, transferPortId string, transferChannelId string, claimer sdk.AccAddress) (*sdk.Coin, error) {
	ibcDenom := k.ibcstakingKeeper.GetIBCHashDenom(ctx, transferPortId, transferChannelId, zone.BaseDenom)
	result := sdk.NewCoin(ibcDenom, sdk.ZeroInt())

	oracleVersion := k.oracleKeeper.GetOracleVersion(ctx, zone.ZoneId)

	records, found := k.GetUserDepositRecord(ctx, zone.ZoneId, claimer)
	if !found {
		return nil, types.ErrNoDepositRecord
	}

	for _, record := range records.Records {
		if record.State == int64(DELEGATE_SUCCESS) && record.OracleVersion < oracleVersion {
			result = result.Add(*record.Amount)
		}
	}

	return &result, nil
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

	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)

	withdrawAmt := k.CalculateWithdrawAlpha(amt.Amount.BigInt(), totalSharedToken.Amount.BigInt(), totalStakedAmount.Coin.Amount.BigInt())
	wAsset := k.ConvertSnAssetToWAssetDecimal(withdrawAmt, zoneInfo.Decimal, baseDenom)

	return wAsset, nil
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
	chainBalanceWithIbcDenom := sdk.NewCoin(ibcDenom, chainInfo.Coin.Amount)

	if chainBalanceWithIbcDenom.Sub(unMintedAmount).IsZero() {
		return unMintedAmount, nil
	}
	return chainBalanceWithIbcDenom.Sub(unMintedAmount), nil
}

func (k Keeper) ConvertWAssetToSnAssetDecimal(amount *big.Int, decimal int64, denom string) sdk.Coin {
	convertDecimal := snAssetDecimal - decimal
	snAsset := new(big.Int).Mul(amount, precisionMultiplier(convertDecimal))
	return sdk.NewCoin(denom, sdk.NewIntFromBigInt(snAsset))
}

func (k Keeper) ConvertSnAssetToWAssetDecimal(amount *big.Int, decimal int64, denom string) sdk.Coin {
	convertDecimal := snAssetDecimal - decimal
	wAsset := new(big.Int).Quo(amount, precisionMultiplier(convertDecimal)).Int64()
	return sdk.NewInt64Coin(denom, wAsset)
}
