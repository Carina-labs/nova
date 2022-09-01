package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

var (
	KeyOperatorAddress = []byte("OperatorAddress")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(operators []string) Params {
	return Params{
		Operators: operators,
	}
}

func DefaultParams() Params {
	return Params{
		Operators: []string{},
	}
}

func (p *Params) Validate() error {
	if err := validateOperators(p.Operators); err != nil {
		return err
	}
	return nil
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOperatorAddress, &p.Operators, validateOperators),
	}
}

func validateOperators(i interface{}) error {
	operators, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter: %T", i)
	}

	for i := range operators {
		val := operators[i]
		_, err := sdk.AccAddressFromBech32(val)

		if err != nil {
			return fmt.Errorf("invalid operator address: %s", err)
		}
	}

	return nil
}
