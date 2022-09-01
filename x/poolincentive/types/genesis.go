package types

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
