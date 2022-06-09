package keeper

import (
	"fmt"
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
	depositor, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}

	amt, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		h.k.Logger(ctx).Error(fmt.Sprintf("type casting error, %s", data.Amount))
	}

	err = h.k.CacheDepositAmt(ctx, depositor, sdk.NewCoin(data.Denom, amt))
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}
}
