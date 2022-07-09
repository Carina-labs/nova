package keeper

import (
	"github.com/Carina-labs/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Implements ICAHooks interface
var _ types.ICAHooks = Keeper{}

func (k Keeper) AfterDelegateEnd() {
	if k.hooks != nil {
		k.hooks.AfterDelegateEnd()
	}
}

func (k Keeper) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, response *stakingtypes.MsgUndelegateResponse) {
	if k.hooks != nil {
		k.hooks.AfterUndelegateEnd(ctx, undelegateMsg, response)
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
