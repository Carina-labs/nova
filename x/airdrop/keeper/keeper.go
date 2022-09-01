package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"strings"
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

// SetAirdropState sets airdrop state for the user
func (k Keeper) SetAirdropState(ctx sdk.Context, user sdk.AccAddress, state *types.AirdropState) error {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyAirdropState(user.String())

	if !k.IsEligible(ctx, user) {
		return fmt.Errorf("user is not eligible for airdrop: %v", user)
	}

	store.Set(userKey, k.cdc.MustMarshal(state))
	return nil
}

// GetAirdropState returns airdrop state for the user
func (k Keeper) GetAirdropState(ctx sdk.Context, user sdk.AccAddress) (*types.AirdropState, error) {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyAirdropState(user.String())

	if !k.IsEligible(ctx, user) {
		return nil, fmt.Errorf("user is not eligible for airdrop: %v", user)
	}

	bz := store.Get(userKey)
	var userState types.AirdropState
	k.cdc.MustUnmarshal(bz, &userState)
	return &userState, nil
}

// IsValidControllerAddr checks if the given address is a valid controller address
func (k Keeper) IsValidControllerAddr(ctx sdk.Context, addr sdk.AccAddress) bool {
	info := k.GetAirdropInfo(ctx)
	return strings.TrimSpace(info.ControllerAddress) == strings.TrimSpace(addr.String())
}

// IsEligible checks if the user is eligible for airdrop
func (k Keeper) IsEligible(ctx sdk.Context, userAddr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyAirdropState(userAddr.String())
	return store.Has(userKey)
}
