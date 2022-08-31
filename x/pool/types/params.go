package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	for _, operator := range p.Operators {
		if _, err := sdk.AccAddressFromBech32(operator); err != nil {
			return err
		}
	}
	return nil
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}
