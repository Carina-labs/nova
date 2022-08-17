package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) UpdateChainState(ctx sdk.Context, chainInfo *types.ChainInfo) error {
	if !k.IsValidOperator(ctx, chainInfo.OperatorAddress) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, chainInfo.OperatorAddress)
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(chainInfo)
	store.Set([]byte(chainInfo.Coin.Denom), bz)
	return nil
}

func (k Keeper) GetChainState(ctx sdk.Context, chainDenom string) (*types.ChainInfo, error) {
	chainInfo := types.ChainInfo{}
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(chainDenom)) {
		return nil, sdkerrors.Wrap(types.ErrNoSupportChain, fmt.Sprintf("chain %s", chainDenom))
	}

	data := store.Get([]byte(chainDenom))
	if err := proto.Unmarshal(data, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}

func (k Keeper) IsValidOperator(ctx sdk.Context, operatorAddress string) bool {
	params := k.GetParams(ctx)

	for i := range params.OracleOperators {
		if params.OracleOperators[i] == operatorAddress {
			return true
		}
	}

	return false
}

func (k Keeper) GetOracleVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
}

func (k Keeper) SetOracleVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetOracleVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetOracleVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
