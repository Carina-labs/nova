package keeper

import (
	"github.com/Carina-labs/nova/v2/x/poolincentive/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec

	paramSpace paramtypes.Subspace
}

// NewKeeper returns a new instance of poolincentive keeper.
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
	}
}

// Logger returns logger for poolincentive keeper.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) getIncentivePoolStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyIncentivePool)
}
