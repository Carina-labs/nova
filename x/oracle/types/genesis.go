package types

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: Params{},
	}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		States: []ChainInfo{},
	}
}

func (gs GenesisState) Validate() error {
	return nil
}
