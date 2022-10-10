package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
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
	if err := msg.Coin.Validate(); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Operator); err != nil {
		return err
	}

	if strings.TrimSpace(msg.ZoneId) == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	if msg.BlockHeight < 0 {
		return ErrNegativeBlockHeight
	}

	return nil
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
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid from address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.OracleAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid oracle address (%s)", err)
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}
	return err
}
