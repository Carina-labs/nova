package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgUpdateChainState{}

func NewMsgUpdateChainState(chainDenom string,
	balance uint64,
	decimal uint64,
	blockHeight uint64,
	proof []byte) *MsgUpdateChainState {
	return &MsgUpdateChainState{
		ChainDenom:    chainDenom,
		StakedBalance: balance,
		Decimal:       decimal,
		BlockHeight:   blockHeight,
		Proof:         proof,
	}
}

func (msg MsgUpdateChainState) GetSigners() []sdk.AccAddress {
	return nil
}

func (msg MsgUpdateChainState) ValidateBasic() error {
	return nil
}
