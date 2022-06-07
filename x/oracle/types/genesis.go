package types

func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		States: []ChainInfo{
			{
				ChainDenom:         "test",
				ValidatorAddress:   "test",
				LastBlockHeight:    1,
				TotalStakedBalance: 1,
				Decimal:            1,
			},
		},
	}
}

func (gs GenesisState) Validate() error {
	return nil
}
