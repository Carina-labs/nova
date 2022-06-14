package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

type ICAHooks interface {
	AfterDelegateEnd()
	AfterUndelegateEnd(sdk.Context, channeltypes.Packet, *stakingtypes.MsgUndelegateResponse)
	AfterAutoStakingEnd()
	AfterWithdrawEnd()
}

var _ ICAHooks = MultiICAHooks{}

type MultiICAHooks []ICAHooks

func NewMultiICAHooks(hooks ...ICAHooks) MultiICAHooks {
	return hooks
}

func (h MultiICAHooks) AfterDelegateEnd() {
	for i := range h {
		h[i].AfterDelegateEnd()
	}
}

func (h MultiICAHooks) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, response *stakingtypes.MsgUndelegateResponse) {
	for i := range h {
		h[i].AfterUndelegateEnd(ctx, packet, response)
	}
}

func (h MultiICAHooks) AfterAutoStakingEnd() {
	for i := range h {
		h[i].AfterAutoStakingEnd()
	}
}

func (h MultiICAHooks) AfterWithdrawEnd() {
	for i := range h {
		h[i].AfterWithdrawEnd()
	}
}
