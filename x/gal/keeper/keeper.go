package keeper

import (
	"github.com/Carina-labs/novachain/x/gal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper defines a module interface that facilitates the transfer of coins between accounts.
type Keeper struct {
	cdc                 codec.BinaryCodec
	storeKey            sdk.StoreKey
	paramSpace          paramtypes.Subspace
	bankKeeper          types.BankKeeper
	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	iaKeeper icacontrollerkeeper.Keeper) Keeper {
	return Keeper{
		cdc:                 cdc,
		storeKey:            key,
		bankKeeper:          bankKeeper,
		paramSpace:          paramSpace,
		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetParams sets the total set of gal parameters.
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.paramSpace.SetParamSet(ctx, params)
}

// GetParams returns total set of gal parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) DepositNativeToken(ctx sdk.Context, depositor string, amt sdk.Coins) {
	// wAtom -> [ GAL ] -> snAtom
	//for _, coin := range amt {
	//	// mint new sn token
	//	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{}}); err != nil {
	//
	//	}
	//
	//	// burn wrapped token
	//	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{}}); err != nil {
	//
	//	}
	//}
}

func (k Keeper) WithdrawNovaToken(ctx sdk.Context, withdrawer string, amt sdk.Coins) {
	// snAtom -> [GAL] -> wAtom
	//for _, coin := range amt {
	//	// burn sn token
	//	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{}}); err != nil {
	//
	//	}
	//
	//	// mint new w token
	//	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{}}); err != nil {
	//
	//	}
	//}
}
