package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/v2/x/gal/types"
	icacontrolkeeper "github.com/Carina-labs/nova/v2/x/icacontrol/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		multiplier, err := calcPrecisionMultiplier(int64(i))
		if err != nil {
			continue
		}
		precisionMultipliers[i] = multiplier
	}
}

func precisionMultiplier(prec int64) (*big.Int, error) {
	if prec > snAssetDecimal {
		return nil, fmt.Errorf("too much precision, maximum %v, provided %v", snAssetDecimal, prec)
	}

	return precisionMultipliers[prec], nil
}

func calcPrecisionMultiplier(prec int64) (*big.Int, error) {
	if prec > snAssetDecimal {
		return nil, fmt.Errorf("too much precision, maximum %v, provided %v", snAssetDecimal, prec)
	}

	zerosToAdd := snAssetDecimal - prec
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(zerosToAdd), nil)
	return multiplier, nil
}

// ClaimShareToken is used when user want to claim their share token.
// It calculates user's share and the amount of claimable share token.
func (k Keeper) ClaimShareToken(ctx sdk.Context, zone *icacontrolkeeper.RegisteredZone, asset sdk.Coin) (*sdk.Coin, error) {
	snDenom, err := k.GetSnDenomForIBCDenom(ctx, asset.Denom)
	if err != nil {
		return nil, err
	}

	baseDenom := k.icaControlKeeper.GetBaseDenomForSnDenom(ctx, snDenom)
	totalSnSupply := k.bankKeeper.GetSupply(ctx, snDenom)
	totalStakedAmount, err := k.GetTotalStakedForLazyMinting(ctx, baseDenom, zone.TransferInfo.PortId, zone.TransferInfo.ChannelId)
	if err != nil {
		return nil, err
	}

	// convert decimal
	snAsset, err := k.ConvertWAssetToSnAssetDecimal(asset.Amount.BigInt(), zone.Decimal, snDenom)
	if err != nil {
		return nil, err
	}
	convertTotalStakedAmount, err := k.ConvertWAssetToSnAssetDecimal(totalStakedAmount.Amount.BigInt(), zone.Decimal, baseDenom)
	if err != nil {
		return nil, err
	}
	mintAmt := k.CalculateDepositAlpha(snAsset.Amount.BigInt(), totalSnSupply.Amount.BigInt(), convertTotalStakedAmount.Amount.BigInt())

	return &sdk.Coin{Denom: snDenom, Amount: sdk.NewIntFromBigInt(mintAmt)}, nil
}

// MintTo mints sn-asset(share token) regard with deposited token to claimer.
func (k Keeper) MintTo(ctx sdk.Context, claimer sdk.AccAddress, mintCoin sdk.Coin) error {
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintCoin)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, sdk.NewCoins(mintCoin)); err != nil {
		return err
	}

	return nil
}

// TotalClaimableAssets returns the total amount of claimable snAsset.
func (k Keeper) TotalClaimableAssets(ctx sdk.Context, zone icacontrolkeeper.RegisteredZone, claimer sdk.AccAddress) (*sdk.Coin, error) {
	ibcDenom := k.icaControlKeeper.GetIBCHashDenom(zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, zone.BaseDenom)
	result := sdk.NewCoin(ibcDenom, sdk.ZeroInt())

	oracleVersion, _ := k.oracleKeeper.GetOracleVersion(ctx, zone.ZoneId)

	records, found := k.GetUserDelegateRecord(ctx, zone.ZoneId, claimer)
	if !found {
		return nil, types.ErrNoDelegateRecord
	}

	for _, record := range records.Records {
		if record.State == types.DelegateSuccess && record.OracleVersion < oracleVersion {
			result = result.Add(*record.Amount)
		}
	}

	amt, err := k.ClaimShareToken(ctx, &zone, result)
	if err != nil {
		return nil, err
	}

	return amt, nil
}

// CalculateDepositAlpha calculates alpha value.
// DepositAlpha = userDepositAmount / totalStakedAmount
// Issued snAsset = Alpha * totalShareTokenSupply
func (k Keeper) CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	if totalShareTokenSupply.Cmp(big.NewInt(0)) == 0 {
		totalShareTokenSupply = totalStakedAmount
	}
	return res.Mul(userDepositAmt, totalShareTokenSupply).Div(res, totalStakedAmount)
}

// CalculateWithdrawAlpha calculates lambda value.
// WithdrawAlpha = userWithdrawAmount / totalStakedAmount
// Issued coin = Lambda * totalShareTokenSupply
func (k Keeper) CalculateWithdrawAlpha(burnedStTokenAmt, totalShareTokenSupply, totalStakedAmount *big.Int) *big.Int {
	res := new(big.Int)
	if totalShareTokenSupply.Cmp(big.NewInt(0)) == 0 {
		totalShareTokenSupply = totalStakedAmount
	}
	return res.Mul(burnedStTokenAmt, totalStakedAmount).Div(res, totalShareTokenSupply)
}

// GetSnDenomForIBCDenom changes the IBCDenom to the appropriate SnDenom.
func (k Keeper) GetSnDenomForIBCDenom(ctx sdk.Context, ibcDenom string) (string, error) {
	err := transfertypes.ValidateIBCDenom(ibcDenom)
	if err != nil {
		return "", err
	}

	denom, err := k.ibcTransferKeeper.DenomPathFromHash(ctx, ibcDenom)
	if err != nil {
		return "", err
	}

	denomTrace := transfertypes.ParseDenomTrace(denom)

	snDenom := k.icaControlKeeper.GetsnDenomForBaseDenom(ctx, denomTrace.BaseDenom)

	return snDenom, nil
}

// GetTotalStakedForLazyMinting returns the sum of coins delegated to the Host chain, which have not been issued snAsset.
func (k Keeper) GetTotalStakedForLazyMinting(ctx sdk.Context, denom, transferPortId, transferChanId string) (sdk.Coin, error) {
	zone := k.icaControlKeeper.GetZoneForDenom(ctx, denom)
	if zone == nil {
		return sdk.Coin{}, fmt.Errorf("cannot find zone denom : %s", denom)
	}

	chainInfo, err := k.oracleKeeper.GetChainState(ctx, denom)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("cannot find zone denom in oracle : %s", denom)
	}

	unMintedAmount, err := k.GetAllAmountNotMintShareToken(ctx, zone)
	if err != nil {
		return sdk.Coin{}, err
	}

	ibcDenom := k.icaControlKeeper.GetIBCHashDenom(transferPortId, transferChanId, denom)
	chainBalanceWithIbcDenom := sdk.NewCoin(ibcDenom, chainInfo.Coin.Amount)
	if chainBalanceWithIbcDenom.Sub(unMintedAmount).IsZero() {
		return unMintedAmount, nil
	}
	return chainBalanceWithIbcDenom.Sub(unMintedAmount), nil
}

// ConvertWAssetToSnAssetDecimal changes the common coin to snAsset's denom and decimal.
func (k Keeper) ConvertWAssetToSnAssetDecimal(amount *big.Int, decimal int64, denom string) (*sdk.Coin, error) {
	convertDecimal := snAssetDecimal - decimal
	zeroMultiplier, err := precisionMultiplier(0)
	if err != nil {
		return nil, err
	}
	asset := new(big.Int).Mul(amount, zeroMultiplier)

	convertDecimalMultiplier, err := precisionMultiplier(convertDecimal)
	snAsset := new(big.Int).Quo(asset, convertDecimalMultiplier)
	if err != nil {
		return nil, err
	}

	return &sdk.Coin{
		Denom: denom, Amount: sdk.NewIntFromBigInt(snAsset),
	}, nil
}

// ConvertSnAssetToWAssetDecimal changes snAsset to matching coin denom and decimal.
func (k Keeper) ConvertSnAssetToWAssetDecimal(amount *big.Int, decimal int64, denom string) (*sdk.Coin, error) {
	decimalMultiplier, err := precisionMultiplier(decimal)
	if err != nil {
		return nil, err
	}

	wAsset := new(big.Int).Quo(amount, decimalMultiplier)
	return &sdk.Coin{Denom: denom, Amount: sdk.NewIntFromBigInt(wAsset)}, nil
}

func (k Keeper) CheckDecimal(amount sdk.Coin, decimal int64) error {
	convertAsset, err := k.ConvertSnAssetToWAssetDecimal(amount.Amount.BigInt(), decimal, amount.Denom)
	if err != nil {
		return err
	}
	if convertAsset.Amount.IsZero() || convertAsset.Amount.IsNil() {
		minAsset, err := k.ConvertWAssetToSnAssetDecimal(big.NewInt(1), decimal, amount.Denom)
		if err != nil {
			return err
		}
		return sdkerrors.Wrapf(types.ErrConvertWAssetIsZero, "Input amount is %s, minimum input is %s", amount.String(), minAsset.String())
	}
	return nil
}
