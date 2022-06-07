package oracle

import (
	"encoding/json"

	"github.com/Carina-labs/nova/x/oracle/client/cli"
	"github.com/Carina-labs/nova/x/oracle/keeper"
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
	cdc codec.Codec
}

func NewAppModuleBasic(cdc codec.Codec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) ValidateGenesis(jsonCodec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, serveMux *runtime.ServeMux) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	//TODO implement me
	panic("implement me")
}

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
	}
}

func (a AppModule) Name() string {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) ValidateGenesis(jsonCodec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModule) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

func (a AppModule) InitGenesis(context sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) ExportGenesis(context sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) Route() sdk.Route {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) QuerierRoute() string {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) ConsensusVersion() uint64 {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) BeginBlock(context sdk.Context, block abci.RequestBeginBlock) {
	//TODO implement me
	panic("implement me")
}

func (a AppModule) EndBlock(context sdk.Context, block abci.RequestEndBlock) []abci.ValidatorUpdate {
	//TODO implement me
	panic("implement me")
}
