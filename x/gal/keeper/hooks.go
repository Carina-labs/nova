package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// Hooks wrapper struct for gal keeper
type Hooks struct {
	k Keeper
}

var _ types.TransferHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData, base_denom string) {
	stAsset := h.k.interTxKeeper.GetstDenomForBaseDenom(ctx, base_denom)

	amt := data.Amount + stAsset

	amount, err := sdk.ParseCoinsNormalized(amt)
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	} else {
		if err := h.k.MintStTokenAndDistribute(ctx, data.Sender, amount); err != nil {
			h.k.Logger(ctx).Error(err.Error())
		}
	}
}
