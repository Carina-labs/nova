package keeper

import (
	"github.com/Carina-labs/novachain/x/gal/types"
	interTxKeeper "github.com/Carina-labs/novachain/x/inter-tx/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	types2 "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	types3 "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper defines a module interface that facilitates the transfer of coins between accounts.
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace

	bankKeeper        types.BankKeeper
	scopedKeeper      capabilitykeeper.ScopedKeeper
	interTxKeeper     interTxKeeper.Keeper
	ibcTransferKeeper transfer.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	interTxKeeper interTxKeeper.Keeper,
	ibcTransferKeeper transfer.Keeper) Keeper {
	return Keeper{
		cdc:               cdc,
		storeKey:          key,
		bankKeeper:        bankKeeper,
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

func (k Keeper) DepositNativeToken(ctx sdk.Context,
	depositor string,
	receiver string,
	sourcePort string,
	sourceChannel string,
	amt sdk.Coins) error {
	// wAtom -> [ GAL ] -> snAtom
	for _, coin := range amt {
		goCtx := sdk.WrapSDKContext(ctx)

		// IBC send token to target chain
		// testSourcePort := "transfer"
		// testSourceChannel := "channel-0"

		_, err := k.ibcTransferKeeper.Transfer(goCtx,
			&types2.MsgTransfer{
				SourcePort:    sourcePort,
				SourceChannel: sourceChannel,
				Token: sdk.Coin{
					Denom:  "",
					Amount: sdk.NewInt(0),
				},
				Sender:   depositor,
				Receiver: receiver,
				TimeoutHeight: types3.Height{
					RevisionHeight: 0,
					RevisionNumber: 0,
				},
				TimeoutTimestamp: 0,
			},
		)

		if err != nil {
			return err
		}

		// mint new sn token
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName,
			sdk.Coins{sdk.Coin{Denom: getPairSnToken(coin.Denom), Amount: coin.Amount}}); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) WithdrawNovaToken(ctx sdk.Context, withdrawer string, amt sdk.Coins) error {
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

func getPairSnToken(denom string) string {
	return ""
}
