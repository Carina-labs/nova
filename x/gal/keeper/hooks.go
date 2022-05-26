package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (k Keeper) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData) error {
	amt := data.Amount + data.Denom
	amount, err := sdk.ParseCoinsNormalized(amt)

	if err != nil {
		return err
	}

	k.MintStTokenAndDistribute(ctx, data.Sender, amount)
	return nil
}

// Hooks wrapper struct for gal keeper
type Hooks struct {
	k Keeper
}

var _ types.TransferHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData) {
	h.k.AfterTransferEnd(ctx, data)
}
