package keeper

import (
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

func (k Keeper) ChainStateStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleChainState)
}

// UpdateChainState updates the status of the zones stored in Oracle with a new status.
func (k Keeper) UpdateChainState(ctx sdk.Context, chainInfo *types.ChainInfo) error {
	if !k.IsValidOracleAddress(ctx, chainInfo.ZoneId, chainInfo.OperatorAddress) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, chainInfo.OperatorAddress)
	}

	store := k.ChainStateStore(ctx)
	bz := k.cdc.MustMarshal(chainInfo)
	store.Set([]byte(chainInfo.Coin.Denom), bz)
	return nil
}

// GetChainState returns the status of the Zone stored in Oracle. This result is used to calculate the equity token.
func (k Keeper) GetChainState(ctx sdk.Context, chainDenom string) (*types.ChainInfo, error) {
	chainInfo := types.ChainInfo{}
	store := k.ChainStateStore(ctx)
	if !store.Has([]byte(chainDenom)) {
		return nil, sdkerrors.Wrap(types.ErrNoSupportChain, fmt.Sprintf("chain %s", chainDenom))
	}

	data := store.Get([]byte(chainDenom))
	if err := proto.Unmarshal(data, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}

// IsValidOracleKeyManager verifies that the parameter address is the correct key manager address.
func (k Keeper) IsValidOracleKeyManager(ctx sdk.Context, oracleAddress string) bool {
	params := k.GetParams(ctx)

	for i := range params.OracleKeyManager {
		if params.OracleKeyManager[i] == oracleAddress {
			return true
		}
	}

	return false
}

func (k Keeper) oracleVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
}

func (k Keeper) SetOracleVersion(ctx sdk.Context, zoneId string, trace types.IBCTrace) {
	store := k.oracleVersionStore(ctx)
	key := zoneId
	bz := k.cdc.MustMarshal(&trace)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetOracleVersion(ctx sdk.Context, zoneId string) (version uint64, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyOracleVersion)
	key := []byte(zoneId)
	res := store.Get(key)

	var record types.IBCTrace
	k.cdc.MustUnmarshal(res, &record)
	if res == nil {
		return 0, 0
	}

	return record.Version, record.Height
}
