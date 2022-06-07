package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (k Keeper) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData, base_denom string) error {
	//getstAsset
	stAsset := k.interTxKeeper.GetstDenomForBaseDenom(ctx, base_denom)

	amt := data.Amount + stAsset

	amount, err := sdk.ParseCoinsNormalized(amt)
	if err != nil {
		return err
	}

	senderAddr, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err

	}
	k.MintStTokenAndDistribute(ctx, senderAddr, amount)
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

func (h Hooks) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData, base_denom string) {
	h.k.AfterTransferEnd(ctx, data, base_denom)
}
