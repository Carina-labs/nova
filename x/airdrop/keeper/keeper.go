package keeper

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) CreateModuleAccount(ctx sdk.Context, amount sdk.Coin) {
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	k.accountKeeper.SetModuleAccount(ctx, moduleAcc)

	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
}

func (k Keeper) BurnToken(ctx sdk.Context) {
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	denom := k.GetAirdropInfo(ctx).AirdropDenom
	amount := k.bankKeeper.GetBalance(ctx, moduleAccount.GetAddress(), denom)
	if amount.IsZero() {
		return
	}

	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		ctx.Logger().Error("burn token", "module", types.ModuleName, "err", err)
	}

	balance := k.bankKeeper.GetBalance(ctx, moduleAccount.GetAddress(), denom)
	ctx.Logger().Info("airdrop", "burn amount", amount, "airdrop module account balance", balance)
}
