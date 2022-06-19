package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgUpdateChainState{}

func (msg MsgUpdateChainState) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg MsgUpdateChainState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	return err
}
