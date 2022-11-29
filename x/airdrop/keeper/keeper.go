package keeper

import (
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	bankKeeper types.BankKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
