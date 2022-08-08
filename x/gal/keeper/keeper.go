package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingKeeper "github.com/Carina-labs/nova/x/ibcstaking/keeper"
	oraclekeeper "github.com/Carina-labs/nova/x/oracle/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper defines a module interface that facilitates the transfer of coins between accounts.
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace

	bankKeeper        types.BankKeeper
	accountKeeper     types.AccountKeeper
	ibcstakingKeeper  ibcstakingKeeper.Keeper
	ibcTransferKeeper transfer.Keeper
	oracleKeeper      oraclekeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	ibcstakingKeeper ibcstakingKeeper.Keeper,
	ibcTransferKeeper transfer.Keeper,
	oracleKeeper oraclekeeper.Keeper) Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:               cdc,
		storeKey:          key,
		bankKeeper:        bankKeeper,
		accountKeeper:     accountKeeper,
		paramSpace:        paramSpace,
		ibcstakingKeeper:  ibcstakingKeeper,
		ibcTransferKeeper: ibcTransferKeeper,
		oracleKeeper:      oracleKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetParams sets the total set of gal parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams returns total set of gal parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}
