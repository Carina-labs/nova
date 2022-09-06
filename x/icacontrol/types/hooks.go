package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

type ICAHooks interface {
	AfterDelegateEnd(sdk.Context, stakingtypes.MsgDelegate)
	AfterUndelegateEnd(sdk.Context, stakingtypes.MsgUndelegate, *stakingtypes.MsgUndelegateResponse)
	AfterAutoStakingEnd()
	AfterWithdrawEnd(sdk.Context, transfertypes.MsgTransfer)
}

var _ ICAHooks = MultiICAHooks{}

type MultiICAHooks []ICAHooks

func NewMultiICAHooks(hooks ...ICAHooks) MultiICAHooks {
	return hooks
}

func (h MultiICAHooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	for i := range h {
		h[i].AfterDelegateEnd(ctx, delegateMsg)
	}
}

func (h MultiICAHooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, response *stakingtypes.MsgUndelegateResponse) {
	for i := range h {
		h[i].AfterUndelegateEnd(ctx, undelegateMsg, response)
	}
}

func (h MultiICAHooks) AfterAutoStakingEnd() {
	for i := range h {
		h[i].AfterAutoStakingEnd()
	}
}

func (h MultiICAHooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	for i := range h {
		h[i].AfterWithdrawEnd(ctx, transferMsg)
	}
}
