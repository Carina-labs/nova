package keeper

import (
	"encoding/json"
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc: cdc, storeKey: key, paramSpace: paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) UpdateChainState(ctx sdk.Context, updateInfo *types.MsgUpdateChainState) error {
	if err := updateInfo.ValidateBasic(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	data := make(map[string]interface{})
	data["balance"] = updateInfo.StakedBalance
	data["decimal"] = updateInfo.Decimal
	data["height"] = updateInfo.BlockHeight
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	store.Set([]byte(updateInfo.ChainDenom), bytes)
	return nil
}

func (k Keeper) GetChainState(chainDenom string) *types.QueryStateResponse {
	return &types.QueryStateResponse{}
}
