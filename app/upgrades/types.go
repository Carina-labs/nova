package upgrades

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/Carina-labs/nova/v2/app/keepers"
)

type BaseAppParamManager interface {
	GetConsensusParams(ctx sdk.Context) *abci.ConsensusParams
	StoreConsensusParams(ctx sdk.Context, cp *abci.ConsensusParams)
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	UpgradeName string

	CreateUpgradeHandler func(*module.Manager, module.Configurator, BaseAppParamManager, *keepers.AppKeepers) upgradetypes.UpgradeHandler

	StoreUpgrades store.StoreUpgrades
}

type Fork struct {
	UpgradeName string

	UpgradeHeight int64

	BeginForkLogic func(ctx sdk.Context, keepers *keepers.AppKeepers)
}
