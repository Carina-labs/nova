package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

var (
	KeyDaoModifierAddress = []byte("DaoModifierAddress")
	KeyCommissionInfo     = []byte("Commission")
)

func NewParams(daomodifierAddrs []string, commission []*CommissionInfo) Params {
	return Params{
		DaoModifiers: daomodifierAddrs,
		Commission:   commission,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		DaoModifiers: []string{},
		Commission:   []*CommissionInfo{},
	}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDaoModifierAddress, &p.DaoModifiers, validateDaoModifierAddress),
		paramtypes.NewParamSetPair(KeyCommissionInfo, &p.Commission, validateCommission),
	}
}

func (p *Params) Validate() error {
	if err := validateDaoModifierAddress(p.DaoModifiers); err != nil {
		return err
	}

	if err := validateCommission(p.Commission); err != nil {
		return err
	}

	return nil
}

func validateDaoModifierAddress(i interface{}) error {
	operators, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for i := range operators {
		val := operators[i]
		_, err := sdk.AccAddressFromBech32(val)

		if err != nil {
			return fmt.Errorf("invalid daomodifier address: %v", err)
		}
	}

	return nil
}

func validateCommission(i interface{}) error {
	params, ok := i.([]*CommissionInfo)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// ensure each commission info is only registered one time.
	registered := make(map[string]bool)
	for _, p := range params {
		if _, exists := registered[p.TreasuryAddress]; exists {
			return fmt.Errorf("duplicate treasury address parameter found: '%s'", p.TreasuryAddress)
		}

		if err := validateCommissionInfo(*p); err != nil {
			return err
		}
		registered[p.TreasuryAddress] = true
	}

	return nil
}

func validateCommissionInfo(i interface{}) error {
	param, ok := i.(CommissionInfo)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if param.CommissionRate.IsNegative() {
		return fmt.Errorf("commission rate cannot be negative: %s", param.CommissionRate)
	}

	if param.CommissionRate.GT(sdk.OneDec()) {
		return fmt.Errorf("commission rate is too large: %s", param.CommissionRate)
	}
	return nil
}
