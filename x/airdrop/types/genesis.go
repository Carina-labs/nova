package types

func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

func (gs GenesisState) Validate() error {
	return nil
}
