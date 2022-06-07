package keeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	interTxKeeper "github.com/Carina-labs/nova/x/inter-tx/keeper"
	oraclekeeper "github.com/Carina-labs/nova/x/oracle/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"

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
	oracleKeeper      oraclekeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	interTxKeeper interTxKeeper.Keeper,
	ibcTransferKeeper transfer.Keeper,
	oracleKeeper oraclekeeper.Keeper) Keeper {

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

func (k Keeper) WithdrawCoin(ctx sdk.Context, withdrawer sdk.Address, amt sdk.Coins) error {
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

func (k Keeper) SetShare(ctx sdk.Context, depositor sdk.AccAddress, shares float64) error {
	store := k.getShareStore(ctx)
	data := make(map[string]interface{})
	data[types.KeyDepositor] = depositor.String()
	data[types.KeyShares] = shares
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	store.Set([]byte(depositor), bytes)
	return nil
}

func (k Keeper) GetShare(ctx sdk.Context, depositor sdk.AccAddress) (*types.QuerySharesResponse, error) {
	store := k.getShareStore(ctx)
	if !store.Has([]byte(depositor)) {
		return nil, errors.New(fmt.Sprintf("Depositor %s is not in state...", depositor))
	}

	result := make(map[string]interface{})
	err := json.Unmarshal(store.Get([]byte(depositor.String())), &result)
	if err != nil {
		return nil, err
	}

	shares, ok := result[types.KeyShares].(float32)
	if !ok {
		// TODO : fix error msg
		return nil, errors.New(fmt.Sprintf("Convert fail"))
	}

	return &types.QuerySharesResponse{
		Address: depositor.String(),
		Shares:  shares,
	}, nil
}

func (k Keeper) getShareStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyShare)
}
