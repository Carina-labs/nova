package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgUpdateChainState{}

func NewMsgUpdateChainState(coin sdk.Coin,
	signer sdk.AccAddress,
	decimal uint64,
	blockHeight uint64) *MsgUpdateChainState {
	return &MsgUpdateChainState{
		Coin: coin,
		Operator:      signer.String(),
		Decimal:       decimal,
		BlockHeight:   blockHeight,
	}
}

func (msg MsgUpdateChainState) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg MsgUpdateChainState) ValidateBasic() error {
	return nil
}
