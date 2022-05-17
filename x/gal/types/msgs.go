package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDeposit  = "deposit"
	TypeMsgWithdraw = "withdraw"
)

var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amount sdk.Coins) *MsgDeposit {
	return &MsgDeposit{
		Depositor: fromAddr.String(),
		Receiver:  toAddr.String(),
		Amount:    amount,
	}
}

func (msg MsgDeposit) Route() string {
	return RouterKey
}

func (msg MsgDeposit) Type() string {
	return TypeMsgDeposit
}

func (msg MsgDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return err
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	depositor, _ := sdk.AccAddressFromBech32(msg.Depositor)
	return []sdk.AccAddress{depositor}
}

var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(toAddr sdk.AccAddress, amount sdk.Coins) *MsgWithdraw {
	return &MsgWithdraw{
		Withdrawer: toAddr.String(),
		Amount:     amount,
	}
}

func (msg MsgWithdraw) Route() string {
	return RouterKey
}

func (msg MsgWithdraw) Type() string {
	return TypeMsgWithdraw
}

func (msg MsgWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
		return err
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.Withdrawer)
	return []sdk.AccAddress{withdrawer}
}
