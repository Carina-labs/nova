package keeper

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

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

func (k Keeper) UpdateChainState(ctx sdk.Context, updateInfo *types.ChainInfo) error {
	params := k.GetParams(ctx)
	isValid := false
	for i := range params.OracleOperators {
		if params.OracleOperators[i] == updateInfo.OperatorAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("you are not valid oracle operator: %s", updateInfo.OperatorAddress)
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := updateInfo.Marshal()
	if err != nil {
		return err
	}

	store.Set([]byte(updateInfo.Coin.Denom), bz)
	return nil
}

func (k Keeper) GetChainState(ctx sdk.Context, chainDenom string) (*types.ChainInfo, error) {
	chainInfo := types.ChainInfo{}
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(chainDenom)) {
		return nil, fmt.Errorf("%s is not supported", chainDenom)
	}

	data := store.Get([]byte(chainDenom))
	if err := proto.Unmarshal(data, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}
