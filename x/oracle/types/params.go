package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

var (
	KeyOracleManager = []byte("OracleKeyManagerAddress")
)

func NewParams(operators []string) Params {
	return Params{
		OracleKeyManager: operators,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		OracleKeyManager: []string{},
	}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOracleManager, &p.OracleKeyManager, validateOracleOperators),
	}
}

func (p *Params) Validate() error {
	if err := validateOracleOperators(p.OracleKeyManager); err != nil {
		return err
	}

	return nil
}

func validateOracleOperators(i interface{}) error {
	operators, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter: %T", i)
	}

	for i := range operators {
		val := operators[i]
		_, err := sdk.AccAddressFromBech32(val)

		if err != nil {
			return fmt.Errorf("invalid key manager address: %v", err)
		}
	}

	return nil
}
