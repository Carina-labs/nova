package keeper

import (
	"github.com/Carina-labs/novachain/x/gal/types"
	interTxKeeper "github.com/Carina-labs/novachain/x/inter-tx/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
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
	scopedKeeper      capabilitykeeper.ScopedKeeper
	interTxKeeper     interTxKeeper.Keeper
	ibcTransferKeeper transfer.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	interTxKeeper interTxKeeper.Keeper,
	ibcTransferKeeper transfer.Keeper) Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:               cdc,
		storeKey:          key,
		bankKeeper:        bankKeeper,
		accountKeeper:     accountKeeper,
		paramSpace:        paramSpace,
		interTxKeeper:     interTxKeeper,
		ibcTransferKeeper: ibcTransferKeeper,
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

func (k Keeper) WithdrawCoin(ctx sdk.Context, withdrawer string, amt sdk.Coins) error {
	// snAtom -> [GAL] -> wAtom
	for _, coin := range amt {
		// burn sn token
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName,
			sdk.Coins{sdk.Coin{Denom: coin.Denom, Amount: coin.Amount}}); err != nil {
			return err
		}

		// mint new w token
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{}}); err != nil {
			return err
		}
	}

	return nil
}
