package app

import (
	"github.com/Carina-labs/nova/v2/x/gal"
	icacontrollkeeper "github.com/Carina-labs/nova/v2/x/icacontrol"
	icacontroltypes "github.com/Carina-labs/nova/v2/x/icacontrol/types"
	"github.com/Carina-labs/nova/v2/x/mint"
	minttypes "github.com/Carina-labs/nova/v2/x/mint/types"
	"github.com/Carina-labs/nova/v2/x/oracle"
	oracletypes "github.com/Carina-labs/nova/v2/x/oracle/types"
	"github.com/Carina-labs/nova/v2/x/poolincentive"
	pooltypes "github.com/Carina-labs/nova/v2/x/poolincentive/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/tendermint/spm/cosmoscmd"
)

func (app *NovaApp) GetModuleManager(
	encodingConfig cosmoscmd.EncodingConfig,
	skipGenesisInvariants bool,
	transferModule module.AppModule,
) []module.AppModule {
	appCodec := encodingConfig.Marshaler

	return []module.AppModule{
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, *app.AccountKeeper, nil),
		vesting.NewAppModule(*app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, *app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, *app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants),
		authzmodule.NewAppModule(appCodec, *app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, *app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, *app.MintKeeper, app.AccountKeeper, app.BankKeeper, app.PoolKeeper),
		slashing.NewAppModule(appCodec, *app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper),
		distr.NewAppModule(appCodec, *app.DistrKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper),
		staking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(*app.UpgradeKeeper),
		evidence.NewAppModule(*app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(*app.ParamsKeeper),
		transferModule,
		wasm.NewAppModule(appCodec, app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		gal.NewAppModule(appCodec, *app.GalKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, *app.OracleKeeper),
		ica.NewAppModule(app.ICAControllerKeeper, app.ICAHostKeeper),
		icacontrollkeeper.NewAppModule(appCodec, *app.IcaControlKeeper, app.AccountKeeper),
		poolincentive.NewAppModule(appCodec, *app.PoolKeeper),
	}
}

func GetOrderInitGenesis() []string {
	return []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		crisistypes.ModuleName,
		wasm.ModuleName,
		gal.ModuleName,
		oracletypes.ModuleName,
		icatypes.ModuleName,
		icacontroltypes.ModuleName,
		authz.ModuleName,
		pooltypes.ModuleName,
	}
}

func GetOrderBeginBlocker() []string {
	return []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		paramstypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		ibctransfertypes.ModuleName,
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		vestingtypes.ModuleName,
		wasm.ModuleName,
		gal.ModuleName,
		oracletypes.ModuleName,
		icatypes.ModuleName,
		icacontroltypes.ModuleName,
		authzkeeper.StoreKey,
		pooltypes.ModuleName,
	}
}

func GetOrderEndBlocker() []string {
	return []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		paramstypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		ibctransfertypes.ModuleName,
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		vestingtypes.ModuleName,
		wasm.ModuleName,
		gal.ModuleName,
		oracletypes.ModuleName,
		icatypes.ModuleName,
		icacontroltypes.ModuleName,
		authzkeeper.StoreKey,
		pooltypes.ModuleName,
	}
}
