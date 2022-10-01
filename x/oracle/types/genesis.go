package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewGenesisState(params Params, chainInfos []ChainInfo) *GenesisState {
	return &GenesisState{
		Params: params,
		States: chainInfos,
	}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		States: []ChainInfo{},
	}
}

func (gs GenesisState) Validate() error {
	for _, op := range gs.Params.OracleKeyManager {
		_, err := sdk.AccAddressFromBech32(op)
		if err != nil {
			return err
		}
	}

	return nil
}
