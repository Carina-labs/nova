package keeper

import (
	"encoding/binary"
	"fmt"
	icacontrolkeeper "github.com/Carina-labs/nova/x/icacontrol/keeper"

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

	icaControlKeeper icacontrolkeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace, icaControlKeeper icacontrolkeeper.Keeper) Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc: cdc, storeKey: key, paramSpace: paramSpace, icaControlKeeper: icaControlKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// UpdateChainState updates the status of the zones stored in Oracle with a new status.
func (k Keeper) UpdateChainState(ctx sdk.Context, chainInfo *types.ChainInfo) error {
	if !k.IsValidOperator(ctx, chainInfo.OperatorAddress) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, chainInfo.OperatorAddress)
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(chainInfo)
	store.Set([]byte(chainInfo.Coin.Denom), bz)
	return nil
}

// GetChainState returns the status of the Zone stored in Oracle. This result is used to calculate the equity token.
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

// IsValidOperator verifies that the parameter address is the correct controller address.
func (k Keeper) IsValidOperator(ctx sdk.Context, operatorAddress string) bool {
	params := k.GetParams(ctx)

	for i := range params.OracleOperators {
		if params.OracleOperators[i] == operatorAddress {
			return true
		}
	}

	return false
}

func (k Keeper) oracleVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
}

func (k Keeper) SetOracleVersion(ctx sdk.Context, zoneId string, version uint64, height uint64) {
	store := k.oracleVersionStore(ctx)
	key := zoneId
	v := make([]byte, 8)
	h := make([]byte, 8)

	binary.BigEndian.PutUint64(v, version)
	binary.BigEndian.PutUint64(h, height)

	bz := append(v, h...)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetOracleVersion(ctx sdk.Context, zoneId string) (version uint64, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0, 0
	}

	version = binary.BigEndian.Uint64(bz[:8])
	height = binary.BigEndian.Uint64(bz[8:])
	return version, height
}
