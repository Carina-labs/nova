package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// BankKeeper defines the contract needed to be fulfilled for banking and supply dependencies.
type BankKeeper interface {
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	IterateTotalSupply(ctx sdk.Context, cd func(sdk.Coin) bool)
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}
