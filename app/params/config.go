package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	HumanCoinUnit = "nova"
	BaseCoinUnit  = "unova"
	OsmoExponent  = 6

	DefaultBondDenom = BaseCoinUnit
)

func RegisterDenoms() {
	err := sdk.RegisterDenom(HumanCoinUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, OsmoExponent))
	if err != nil {
		panic(err)
	}
}
