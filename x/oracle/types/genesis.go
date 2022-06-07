package types

func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		States: []ChainInfo{},
	}
}

func (gs GenesisState) Validate() error {
	return nil
}
