package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

func (g GenesisState) Validate() error {
	return nil
}

func ValidateGenesis(gs GenesisState) error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func (ip IncentivePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(ip.PoolContractAddress); err != nil {
		return err
	}

	return nil
}
