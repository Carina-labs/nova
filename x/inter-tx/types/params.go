package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

var (
	KeyZoneRegisterAddress = []byte("ZoneRegisterAddress")
)

func NewParams(operators []string) Params {
	return Params{
		ZoneRegisterAddress: operators,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		ZoneRegisterAddress: []string{},
	}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyZoneRegisterAddress, &p.ZoneRegisterAddress, validateZoneRegisterAddress),
	}
}

func (p *Params) Validate() error {
	if err := validateZoneRegisterAddress(p.ZoneRegisterAddress); err != nil {
		return err
	}

	return nil
}

func validateZoneRegisterAddress(i interface{}) error {
	operators, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter: %T", i)
	}

	for i := range operators {
		val := operators[i]
		_, err := sdk.AccAddressFromBech32(val)

		if err != nil {
			return fmt.Errorf("invalid operator address: %v", err)
		}
	}

	return nil
}
