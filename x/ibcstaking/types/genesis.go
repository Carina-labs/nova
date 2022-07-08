package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// NewGenesisState creates a new GenesisState object
func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(gs GenesisState) error {
	for _, zoneRegisterAddr := range gs.Params.DaoModifiers {
		_, err := sdk.AccAddressFromBech32(zoneRegisterAddr)
		if err != nil {
			return err
		}
	}
	return nil
}
