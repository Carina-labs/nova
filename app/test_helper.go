package app

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestHelper struct {
	suite.Suite

	App *NovaApp
	Ctx sdk.Context
}

func Setup(isCheckTx bool) *NovaApp {
	db := dbm.NewMemDB()
	encodingConfig := cosmoscmd.MakeEncodingConfig(ModuleBasics)
	app := NewNovaApp(log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		5,
		encodingConfig,
		simapp.EmptyAppOptions{})

	if !isCheckTx {
		genesisState := NewDefaultGenesisState(encodingConfig.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			})
	}

	return app.(*NovaApp)
}
