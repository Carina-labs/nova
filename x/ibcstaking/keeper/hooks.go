package keeper

import (
	"github.com/Carina-labs/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
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

func (k Keeper) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	if k.hooks != nil {
		k.hooks.AfterWithdrawEnd(ctx, transferMsg)
	}
}
