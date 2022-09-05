package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
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

func (p *Params) Validate() error {
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
