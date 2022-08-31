package types

import (
	"time"

	"github.com/Carina-labs/nova/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ghodss/yaml"
)

func NewParams() Params {
	return Params{}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		SnapshotTimestamp:  time.Unix(0, 0),
		ClaimableTimestamp: time.Unix(0, 0),
		AirdropDenom:       params.BaseCoinUnit,
		QuestsCount:        5,
		ControllerAddress:  sdk.AccAddress([]byte{0x1}).String(),
	}
}

func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p *Params) Validate() error {
	return nil
}
