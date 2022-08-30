package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{}
}

func DefaultParams() Params {
	return Params{}
}

func (Params) Validate() error {
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}
