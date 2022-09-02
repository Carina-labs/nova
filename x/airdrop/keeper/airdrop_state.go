package keeper

import (
	"fmt"
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// SetUserState sets airdrop state for the user
func (k Keeper) SetUserState(ctx sdk.Context, user sdk.AccAddress, state *types.UserState) error {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(user.String())

	if !k.IsEligible(ctx, user) {
		return fmt.Errorf("user is not eligible for airdrop: %v", user)
	}

	store.Set(userKey, k.cdc.MustMarshal(state))
	return nil
}

// GetUserState returns airdrop state for the user
func (k Keeper) GetUserState(ctx sdk.Context, user sdk.AccAddress) (*types.UserState, error) {
	store := ctx.KVStore(k.storeKey)
	userKey := types.GetKeyUserState(user.String())

	if !k.IsEligible(ctx, user) {
		return nil, fmt.Errorf("user is not eligible for airdrop: %v", user)
	}

	bz := store.Get(userKey)
	var userState types.UserState
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
	userKey := types.GetKeyUserState(userAddr.String())
	return store.Has(userKey)
}
