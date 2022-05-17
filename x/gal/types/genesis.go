package types

// NewGenesisState returns a new GenesisState object.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

// DefaultGenesisState returns a default gal module genesis state.
func DefaultGenesisState() *GenesisState {

	genesisParam := DefaultParams()

	return &GenesisState{
		Params: genesisParam,
	}
}

// Validate performs basic validation of supply genesis data returning an
// error for any dailed validation criteria.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}
