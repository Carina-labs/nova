package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateChainState{}
var _ sdk.Msg = &MsgRegisterOracleAddr{}

func NewMsgUpdateChainState(signer sdk.AccAddress, chainId string, coin sdk.Coin, blockHeight int64, appHash []byte) *MsgUpdateChainState {
	return &MsgUpdateChainState{
		Coin:        coin,
		Operator:    signer.String(),
		BlockHeight: blockHeight,
		AppHash:     appHash,
		ZoneId:      chainId,
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

func NewMsgRegisterOracleAddr(zoneId string, oracleAddr, fromAddr sdk.AccAddress) *MsgRegisterOracleAddr {
	return &MsgRegisterOracleAddr{
		ZoneId:        zoneId,
		OracleAddress: oracleAddr.String(),
		FromAddress:   fromAddr.String(),
	}
}

func (msg MsgRegisterOracleAddr) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg MsgRegisterOracleAddr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	_, err = sdk.AccAddressFromBech32(msg.OracleAddress)

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}
	return err
}
