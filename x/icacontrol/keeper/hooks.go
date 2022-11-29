package keeper

import (
	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
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

func (k Keeper) AfterDelegateFail(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	if k.hooks != nil {
		k.hooks.AfterDelegateFail(ctx, delegateMsg)
	}
}

func (k Keeper) AfterUndelegateFail(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate) {
	if k.hooks != nil {
		k.hooks.AfterUndelegateFail(ctx, undelegateMsg)
	}
}

func (k Keeper) AfterIcaWithdrawFail(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	if k.hooks != nil {
		k.hooks.AfterIcaWithdrawFail(ctx, transferMsg)
	}
}
