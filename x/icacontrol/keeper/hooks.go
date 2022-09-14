package keeper

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Implements ICAHooks interface
var _ types.ICAHooks = Keeper{}

func (k Keeper) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	if k.hooks != nil {
		k.hooks.AfterDelegateEnd(ctx, delegateMsg)
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
