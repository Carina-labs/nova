package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	interTxKeeper "github.com/Carina-labs/nova/x/inter-tx/keeper"
	oraclekeeper "github.com/Carina-labs/nova/x/oracle/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	"math"

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
	oracleKeeper      oraclekeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	interTxKeeper interTxKeeper.Keeper,
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
		interTxKeeper:     interTxKeeper,
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

func (k Keeper) WithdrawCoin(ctx sdk.Context, withdrawer sdk.Address, amt sdk.Coins) error {
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

func (k Keeper) calculateAlpha(ctx sdk.Context, denom string, depositAmt int64) (float64, error) {
	res, err := k.oracleKeeper.GetChainState(ctx, denom)
	if err != nil {
		return 0, err
	}

	alpha := float64(depositAmt) / k.calculateCoinAmount(res.TotalStakedBalance, res.Decimal)
	return alpha, nil
}

func (k Keeper) calculateCoinAmount(amt uint64, decimal uint64) float64 {
	res := float64(amt) / math.Pow10(int(decimal))
	return res
}
