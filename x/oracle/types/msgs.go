package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgUpdateChainState{}

func NewMsgUpdateChainState(signer sdk.AccAddress, chainId string, coin sdk.Coin, decimal, blockHeight uint64, appHash string) *MsgUpdateChainState {
	return &MsgUpdateChainState{
		Coin:        coin,
		Operator:    signer.String(),
		Decimal:     decimal,
		BlockHeight: blockHeight,
		AppHash:     appHash,
		ChainId:     chainId,
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
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	return err
}
