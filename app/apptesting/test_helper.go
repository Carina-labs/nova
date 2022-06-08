package apptesting

import (
	"github.com/Carina-labs/nova/app"
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

	App *app.NovaApp
	Ctx sdk.Context
}

func Setup(isCheckTx bool) *app.NovaApp {
	db := dbm.NewMemDB()
	encodingConfig := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	novaApp := app.NewNovaApp(log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		encodingConfig,
		simapp.EmptyAppOptions{})

	if !isCheckTx {
		genesisState := app.NewDefaultGenesisState(encodingConfig.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}
		novaApp.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			})
	}

	return novaApp.(*app.NovaApp)
}
