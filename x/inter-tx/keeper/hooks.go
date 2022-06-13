package keeper

import (
	"github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// Implements ICAHooks interface
var _ types.ICAHooks = Keeper{}

func (k Keeper) AfterDelegateEnd() {
	if k.hooks != nil {
		k.hooks.AfterDelegateEnd()
	}
}

func (k Keeper) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, txHash string) {
	if k.hooks != nil {
		k.hooks.AfterUndelegateEnd(ctx, packet, txHash)
	}
}

func (k Keeper) AfterAutoStakingEnd() {
	if k.hooks != nil {
		k.hooks.AfterAutoStakingEnd()
	}
}

func (k Keeper) AfterWithdrawEnd() {
	if k.hooks != nil {
		k.hooks.AfterWithdrawEnd()
	}
}
