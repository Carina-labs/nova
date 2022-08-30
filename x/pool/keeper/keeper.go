package keeper

import (
	"github.com/Carina-labs/nova/x/pool/types"
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

// NewKeeper returns a new instance of pool keeper.
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
	}
}

// Logger returns logger for pool keeper.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// getPoolStore returns store saving pool information.
func (k Keeper) getPoolStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPool)
}

// isValidOperator checks if signer of msg is valid.
func (k Keeper) isValidOperator(msg *types.MsgSetPoolWeight) bool {
	controllerAcc, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return false
	}

	signers := msg.GetSigners()
	for _, signer := range signers {
		if controllerAcc.Equals(signer) {
			return true
		}
	}
	return false
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
