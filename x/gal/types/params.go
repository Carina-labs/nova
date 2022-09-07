package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

var (
	KeyWhiteListedTokenDenoms = []byte("whiteListedTokenDenoms")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(snAssetDenoms map[string]string) Params {
	return Params{
		WhiteListedTokenDenoms: snAssetDenoms,
	}
}

// Key관리
func DefaultParams() Params {
	return Params{
		WhiteListedTokenDenoms: map[string]string{
			"uatom": "statom",
			"ujuno": "stjuno",
			"uosmo": "stosmo",
		},
	}
}

func (Params) Validate() error {
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyWhiteListedTokenDenoms, &p.WhiteListedTokenDenoms, validateSnAssetDenoms),
	}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateSnAssetDenoms(i interface{}) error {
	return nil
}
