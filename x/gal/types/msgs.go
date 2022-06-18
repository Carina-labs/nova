package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDeposit          = "deposit"
	TypeMsgUndelegate       = "undelegate"
	TypeMsgWithdrawRecord   = "withdrawRecord"
	TypeMsgUndelegateRecord = "undelegateRecord"
)

var _ sdk.Msg = &MsgDeposit{}
var _ sdk.Msg = &MsgUndelegate{}
var _ sdk.Msg = &MsgUndelegateRecord{}
var _ sdk.Msg = &MsgWithdrawRecord{}

func NewMsgDeposit(fromAddr sdk.AccAddress, hostAddr string, amount sdk.Coins, zoneId string) *MsgDeposit {
	return &MsgDeposit{
		Depositor: fromAddr.String(),
		HostAddr:  hostAddr,
		ZoneId:    zoneId,
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

func NewMsgUndelegate(zoneId, controllerAddr, hostAddr string) *MsgUndelegate {
	return &MsgUndelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr,
		HostAddress:       hostAddr,
	}
}

func (msg MsgUndelegate) Route() string {
	return RouterKey
}

func (msg MsgUndelegate) Type() string {
	return TypeMsgDeposit
}

func (msg MsgUndelegate) ValidateBasic() error {
	if msg.ControllerAddress == "" {
		return errors.New("controller address is not null")
	}

	if msg.HostAddress == "" {
		return errors.New("host address is not null")
	}

	if msg.ZoneId == "" {
		return errors.New("zone id is not found")
	}
	return nil
}

func (msg MsgUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUndelegate) GetSigners() []sdk.AccAddress {
	depositor, _ := sdk.AccAddressFromBech32(msg.ControllerAddress)
	return []sdk.AccAddress{depositor}
}

func NewMsgUndelegateRecord(zoneId, Depositor string, amount sdk.Coin) *MsgUndelegateRecord {
	return &MsgUndelegateRecord{
		ZoneId:    zoneId,
		Depositor: Depositor,
		Amount:    amount,
	}
}

func (msg MsgUndelegateRecord) Route() string {
	return RouterKey
}

func (msg MsgUndelegateRecord) Type() string {
	return TypeMsgUndelegateRecord
}

func (msg MsgUndelegateRecord) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return err
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgUndelegateRecord) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUndelegateRecord) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.Depositor)
	return []sdk.AccAddress{withdrawer}
}

func NewMsgWithdrawRecord(zoneId string, toAddr sdk.AccAddress, amount sdk.Coin) *MsgWithdrawRecord {
	return &MsgWithdrawRecord{
		ZoneId:     zoneId,
		Withdrawer: toAddr.String(),
		Recipient:  "",
		Amount:     amount,
	}
}

func (msg MsgWithdrawRecord) Route() string {
	return RouterKey
}

func (msg MsgWithdrawRecord) Type() string {
	return TypeMsgWithdrawRecord
}

func (msg MsgWithdrawRecord) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
	// 	return err
	// }

	if msg.Withdrawer == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Withdrawer)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgWithdrawRecord) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawRecord) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.Withdrawer)
	return []sdk.AccAddress{withdrawer}
}
