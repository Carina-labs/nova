package keepers

import (
	airdropkeeper "github.com/Carina-labs/nova/x/airdrop/keeper"
	airdroptypes "github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/Carina-labs/nova/x/gal"
	galkeeper "github.com/Carina-labs/nova/x/gal/keeper"
	galtypes "github.com/Carina-labs/nova/x/gal/types"
	ibcstaking "github.com/Carina-labs/nova/x/ibcstaking"
	ibcstakingkeeper "github.com/Carina-labs/nova/x/ibcstaking/keeper"
	mintkeeper "github.com/Carina-labs/nova/x/mint/keeper"
	"github.com/Carina-labs/nova/x/oracle"
	oraclekeeper "github.com/Carina-labs/nova/x/oracle/keeper"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	"github.com/Carina-labs/nova/x/poolincentive"
	poolincentivekeeper "github.com/Carina-labs/nova/x/poolincentive/keeper"
	poolincentivetypes "github.com/Carina-labs/nova/x/poolincentive/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontroller "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icahost "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/keeper"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibcclient "github.com/cosmos/ibc-go/v3/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	minttypes "github.com/Carina-labs/nova/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

type AppKeepers struct {
	// keepers
	AccountKeeper       *authkeeper.AccountKeeper
	BankKeeper          *bankkeeper.BaseKeeper
	CapabilityKeeper    *capabilitykeeper.Keeper
	StakingKeeper       *stakingkeeper.Keeper
	SlashingKeeper      *slashingkeeper.Keeper
	MintKeeper          *mintkeeper.Keeper
	DistrKeeper         *distrkeeper.Keeper
	GovKeeper           *govkeeper.Keeper
	CrisisKeeper        *crisiskeeper.Keeper
	UpgradeKeeper       *upgradekeeper.Keeper
	ParamsKeeper        *paramskeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	ICAControllerKeeper *icacontrollerkeeper.Keeper
	ICAHostKeeper       *icahostkeeper.Keeper
	EvidenceKeeper      *evidencekeeper.Keeper
	TransferKeeper      *ibctransferkeeper.Keeper
	FeeGrantKeeper      *feegrantkeeper.Keeper
	WasmKeeper          *wasmkeeper.Keeper
	AuthzKeeper         *authzkeeper.Keeper

	// Supernova custom modules.
	GalKeeper        *galkeeper.Keeper
	IbcstakingKeeper *ibcstakingkeeper.Keeper
	OracleKeeper     *oraclekeeper.Keeper
	AirdropKeeper    *airdropkeeper.Keeper
	PoolKeeper       *poolincentivekeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedIbcstakingKeeper    capabilitykeeper.ScopedKeeper

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey
}

func (appKeepers *AppKeepers) InitNormalKeepers(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	maccPerms map[string][]string,
	wasmDir string,
	wasmConfig wasm.Config,
	wasmProposals []wasm.ProposalType,
	moduleAddrs map[string]bool,
) {
	// add keepers
	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec, appKeepers.keys[authtypes.StoreKey], appKeepers.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	appKeepers.AccountKeeper = &accountKeeper

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec, appKeepers.keys[banktypes.StoreKey], appKeepers.AccountKeeper, appKeepers.GetSubspace(banktypes.ModuleName), moduleAddrs,
	)
	appKeepers.BankKeeper = &bankKeeper

	authzKeeper := authzkeeper.NewKeeper(
		appKeepers.keys[authzkeeper.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
	)
	appKeepers.AuthzKeeper = &authzKeeper

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, appKeepers.keys[stakingtypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BankKeeper, appKeepers.GetSubspace(stakingtypes.ModuleName),
	)

	distrKeeper := distrkeeper.NewKeeper(
		appCodec, appKeepers.keys[distrtypes.StoreKey], appKeepers.GetSubspace(distrtypes.ModuleName), appKeepers.AccountKeeper, appKeepers.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, moduleAddrs,
	)
	appKeepers.DistrKeeper = &distrKeeper

	// Register Pool module.
	poolKeeper := poolincentivekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[poolincentivetypes.StoreKey],
		appKeepers.GetSubspace(poolincentivetypes.ModuleName),
	)
	appKeepers.PoolKeeper = &poolKeeper

	mintKeeper := mintkeeper.NewKeeper(
		appCodec, appKeepers.keys[minttypes.StoreKey], appKeepers.GetSubspace(minttypes.ModuleName), &stakingKeeper,
		appKeepers.AccountKeeper, appKeepers.BankKeeper, appKeepers.DistrKeeper, appKeepers.PoolKeeper, authtypes.FeeCollectorName,
	)
	appKeepers.MintKeeper = &mintKeeper

	slashingKeeper := slashingkeeper.NewKeeper(
		appCodec, appKeepers.keys[slashingtypes.StoreKey], &stakingKeeper, appKeepers.GetSubspace(slashingtypes.ModuleName),
	)
	appKeepers.SlashingKeeper = &slashingKeeper

	feeGrantKeeper := feegrantkeeper.NewKeeper(appCodec, appKeepers.keys[feegrant.StoreKey], appKeepers.AccountKeeper)
	appKeepers.FeeGrantKeeper = &feeGrantKeeper

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks

	appKeepers.StakingKeeper = stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(appKeepers.DistrKeeper.Hooks(), appKeepers.SlashingKeeper.Hooks()),
	)

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, appKeepers.keys[ibchost.StoreKey], appKeepers.GetSubspace(ibchost.ModuleName), appKeepers.StakingKeeper, appKeepers.UpgradeKeeper, appKeepers.ScopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(*appKeepers.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(*appKeepers.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(*appKeepers.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper)).
		AddRoute(poolincentivetypes.RouterKey, poolincentive.NewPoolIncentivesProposalHandler(*appKeepers.PoolKeeper))

	icaControllerKeeper := icacontrollerkeeper.NewKeeper(
		appCodec, appKeepers.keys[icacontrollertypes.StoreKey], appKeepers.GetSubspace(icacontrollertypes.SubModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedICAControllerKeeper, bApp.MsgServiceRouter(),
	)
	appKeepers.ICAControllerKeeper = &icaControllerKeeper

	icaHostKeeper := icahostkeeper.NewKeeper(
		appCodec, appKeepers.keys[icahosttypes.StoreKey], appKeepers.GetSubspace(icahosttypes.SubModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper, appKeepers.ScopedICAHostKeeper, bApp.MsgServiceRouter(),
	)
	appKeepers.ICAHostKeeper = &icaHostKeeper

	ibcstakingKeeper := ibcstakingkeeper.NewKeeper(appCodec, appKeepers.keys[ibcstakingtypes.StoreKey], appKeepers.AccountKeeper, *appKeepers.ICAControllerKeeper, appKeepers.ScopedIbcstakingKeeper, appKeepers.GetSubspace(ibcstakingtypes.ModuleName))

	// Create Transfer Keepers
	transferKeeper := ibctransferkeeper.NewKeeper(
		appCodec, appKeepers.keys[ibctransfertypes.StoreKey], appKeepers.GetSubspace(ibctransfertypes.ModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper, appKeepers.BankKeeper, appKeepers.ScopedTransferKeeper,
	)

	// Register OracleModule
	oracleKeeper := oraclekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[oracletypes.StoreKey],
		appKeepers.GetSubspace(oracletypes.ModuleName),
	)
	appKeepers.OracleKeeper = &oracleKeeper

	// Register AirdropKeeper
	airdropKeeper := airdropkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[airdroptypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)
	appKeepers.AirdropKeeper = &airdropKeeper

	// Register GAL module.
	galKeeper := galkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[galtypes.StoreKey],
		appKeepers.GetSubspace(galtypes.ModuleName),
		appKeepers.BankKeeper,
		ibcstakingKeeper,
		transferKeeper,
		oracleKeeper,
		appKeepers.AirdropKeeper,
	)
	appKeepers.GalKeeper = &galKeeper

	appKeepers.TransferKeeper = transferKeeper.SetHooks(
		ibctransfertypes.NewMultiTransferHooks(appKeepers.GalKeeper.Hooks()),
	)

	appKeepers.IbcstakingKeeper = ibcstakingKeeper.SetHooks(
		ibcstakingtypes.NewMultiICAHooks(appKeepers.GalKeeper.Hooks()),
	)

	ibcstakingIBCModule := ibcstaking.NewIBCModule(*appKeepers.IbcstakingKeeper)

	icaControllerIBCModule := icacontroller.NewIBCModule(*appKeepers.ICAControllerKeeper, ibcstakingIBCModule)
	icaHostIBCModule := icahost.NewIBCModule(*appKeepers.ICAHostKeeper)

	transferIBCModule := transfer.NewIBCModule(*appKeepers.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, appKeepers.keys[evidencetypes.StoreKey], appKeepers.StakingKeeper, appKeepers.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	appKeepers.EvidenceKeeper = evidenceKeeper

	supportedFeatures := "iterator,staking,stargate"

	wasmKeeper := wasm.NewKeeper(
		appCodec, appKeepers.keys[wasm.StoreKey], appKeepers.GetSubspace(wasm.ModuleName), appKeepers.AccountKeeper, appKeepers.BankKeeper,
		appKeepers.StakingKeeper, appKeepers.DistrKeeper, appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedWasmKeeper, appKeepers.TransferKeeper, bApp.MsgServiceRouter(), bApp.GRPCQueryRouter(),
		wasmDir, wasmConfig, supportedFeatures,
	)
	appKeepers.WasmKeeper = &wasmKeeper

	// register wasm gov proposal types
	if len(wasmProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(appKeepers.WasmKeeper, wasmProposals))
	}

	govKeeper := govkeeper.NewKeeper(
		appCodec, appKeepers.keys[govtypes.StoreKey], appKeepers.GetSubspace(govtypes.ModuleName), appKeepers.AccountKeeper, appKeepers.BankKeeper,
		&stakingKeeper, govRouter,
	)
	appKeepers.GovKeeper = govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(appKeepers.AirdropKeeper.Hooks()),
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule)
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)
	ibcRouter.AddRoute(ibcstakingtypes.ModuleName, icaControllerIBCModule)
	ibcRouter.AddRoute(wasm.ModuleName, wasm.NewIBCHandler(appKeepers.WasmKeeper, appKeepers.IBCKeeper.ChannelKeeper))
	appKeepers.IBCKeeper.SetRouter(ibcRouter)
}

func (appKeepers *AppKeepers) InitSpecialKeepers(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	wasmDir string,
	cdc *codec.LegacyAmino,
	invCheckPeriod uint,
	skipUpgradeHeights map[int64]bool,
	homePath string,
) {
	paramsKeeper := appKeepers.InitParamsKeeper(appCodec, cdc, appKeepers.keys[paramstypes.StoreKey], appKeepers.tkeys[paramstypes.TStoreKey])
	appKeepers.ParamsKeeper = &paramsKeeper

	// set the BaseApp's parameter store
	bApp.SetParamStore(appKeepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, appKeepers.keys[capabilitytypes.StoreKey], appKeepers.memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	appKeepers.ScopedWasmKeeper = appKeepers.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	appKeepers.ScopedICAControllerKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	appKeepers.ScopedICAHostKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	appKeepers.ScopedIbcstakingKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibcstakingtypes.ModuleName)

	upgradeKeeper := upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		appKeepers.keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		bApp,
	)
	appKeepers.UpgradeKeeper = &upgradeKeeper

	crisisKeeper := crisiskeeper.NewKeeper(
		appKeepers.GetSubspace(crisistypes.ModuleName), invCheckPeriod, appKeepers.BankKeeper, authtypes.FeeCollectorName,
	)
	appKeepers.CrisisKeeper = &crisisKeeper
}

// InitParamsKeepers init params keeper and its subspaces
func (appKeepers *AppKeepers) InitParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(ibcstakingtypes.ModuleName)
	paramsKeeper.Subspace(oracle.ModuleName)
	paramsKeeper.Subspace(gal.ModuleName)
	paramsKeeper.Subspace(poolincentivetypes.ModuleName)

	return paramsKeeper
}

func KVStoreKeys() []string {
	return []string{
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		ibchost.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		wasm.StoreKey,
		gal.StoreKey,
		icacontrollertypes.StoreKey,
		icahosttypes.StoreKey,
		ibcstakingtypes.StoreKey,
		authzkeeper.StoreKey,
		oracletypes.StoreKey,
		poolincentivetypes.StoreKey,
		airdroptypes.StoreKey,
	}
}
