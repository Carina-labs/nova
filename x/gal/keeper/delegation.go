package keeper

import (
	"encoding/binary"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MintShareTokens mints sn-token(share token) regard with deposited token.
func (k Keeper) MintShareTokens(ctx sdk.Context, depositor sdk.AccAddress, amt sdk.Coin) error {
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

// BurnShareTokens burns share token.
func (k Keeper) BurnShareTokens(ctx sdk.Context, burner sdk.Address, amt sdk.Coin) error {
	burnerAddr, err := sdk.AccAddressFromBech32(burner.String())
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, burnerAddr, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetDelegateVersionStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
}

func (k Keeper) SetDelegateVersion(ctx sdk.Context, zoneId string, version uint64) {
	store := k.GetDelegateVersionStore(ctx)
	key := zoneId
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, version)
	store.Set([]byte(key), bz)
}

func (k Keeper) GetDelegateVersion(ctx sdk.Context, zoneId string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyDelegateVersion)
	key := []byte(zoneId)
	bz := store.Get(key)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
