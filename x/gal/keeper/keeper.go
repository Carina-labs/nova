package keeper

import (
	"context"

	"github.com/Carina-labs/novachain/x/gal/types"
	interTxKeeper "github.com/Carina-labs/novachain/x/inter-tx/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
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
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams returns total set of gal parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) DepositCoin(ctx sdk.Context,
	depositor string,
	receiver string,
	sourcePort string,
	sourceChannel string,
	amt sdk.Coins) error {
	//wAtom -> [ GAL ] -> snAtom
	for _, coin := range amt {
		goCtx := sdk.WrapSDKContext(ctx)

		_, err := k.ibcTransferKeeper.Transfer(goCtx,
			&transfertypes.MsgTransfer{
				SourcePort:    sourcePort,
				SourceChannel: sourceChannel,
				Token:         coin,
				Sender:        depositor,
				Receiver:      receiver,
				TimeoutHeight: ibcclienttypes.Height{
					RevisionHeight: 100,
					RevisionNumber: 0,
				},
				TimeoutTimestamp: 0,
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) UnStaking(ctx sdk.Context) {

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

func (k Keeper) MintStTokenAndDistribute(ctx sdk.Context, depositor string, amt sdk.Coins) error {
	depositorAddr, err := sdk.AccAddressFromBech32(depositor)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, amt); err != nil {
		return err
	}

	return nil
}

func (k Keeper) CalculateShares(ctx sdk.Context, depositor string, coin sdk.Coin) (float64, error) {
	totalSupply := k.bankKeeper.GetSupply(ctx, coin.Denom)
	depositorBalance, err := k.bankKeeper.Balance(ctx.Context(), &banktypes.QueryBalanceRequest{
		Address: depositor,
		Denom:   coin.Denom,
	})

	if err != nil {
		return 0, err
	}

	shares := float64(depositorBalance.Balance.Amount.Int64()/totalSupply.Amount.Int64()) * 100
	return shares, nil
}

func (k Keeper) getPairSnToken(ctx sdk.Context, denom string) (stTokenDenom string) {
	k.paramSpace.Get(ctx, types.KeyWhiteListedTokenDenoms, &stTokenDenom)
	return
}

func (k Keeper) Share(context context.Context, rq *types.QuerySharesRequest) (*types.QuerySharesResponse, error) {
	return nil, nil
}
