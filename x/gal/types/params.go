package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeySnTokenDenoms = []byte("snTokenDenoms")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(snTokenDenoms []string) Params {
	return Params{
		SnTokenDenoms: snTokenDenoms,
	}
}

func DefaultParams() Params {
	return Params{
		SnTokenDenoms: []string{},
	}
}

func (Params) Validate() error {
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySnTokenDenoms, &p.SnTokenDenoms, validateSnTokenDenoms),
	}
}

func (p *Params) String() string {
	return ""
}

func validateSnTokenDenoms(i interface{}) error {
	return nil
}
