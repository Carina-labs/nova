package types

import (
	"errors"
	time "time"

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

var _ sdk.Msg = &MsgDelegate{}
var _ sdk.Msg = &MsgUndelegate{}
var _ sdk.Msg = &MsgPendingUndelegate{}
var _ sdk.Msg = &MsgWithdraw{}
var _ sdk.Msg = &MsgClaimSnAsset{}
var _ sdk.Msg = &MsgIcaWithdraw{}

func NewMsgDeposit(zoneId string, depositor, claimer sdk.AccAddress, amount sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		ZoneId:    zoneId,
		Depositor: depositor.String(),
		Claimer:   claimer.String(),
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

	if !msg.Amount.IsPositive() {
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

func NewMsgDelegate(zoneId string, daomodifierAddr sdk.AccAddress) *MsgDelegate {
	return &MsgDelegate{
		ZoneId:            zoneId,
		ControllerAddress: daomodifierAddr.String(),
	}
}

func (msg MsgDelegate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ControllerAddress); err != nil {
		return err
	}

	if msg.ZoneId == "" {
		return errors.New("zone id is not found")
	}

	return nil
}

func (msg MsgDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	delegator, _ := sdk.AccAddressFromBech32(msg.ControllerAddress)
	return []sdk.AccAddress{delegator}
}

func NewMsgUndelegate(zoneId string, controllerAddr sdk.AccAddress) *MsgUndelegate {
	return &MsgUndelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
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

func NewMsgPendingUndelegate(zoneId string, delegator, withdrawAddr sdk.AccAddress, amount sdk.Coin) *MsgPendingUndelegate {
	return &MsgPendingUndelegate{
		ZoneId:     zoneId,
		Delegator:  delegator.String(),
		Withdrawer: withdrawAddr.String(),
		Amount:     amount,
	}
}

func (msg MsgPendingUndelegate) Route() string {
	return RouterKey
}

func (msg MsgPendingUndelegate) Type() string {
	return TypeMsgUndelegateRecord
}

func (msg MsgPendingUndelegate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Delegator); err != nil {
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

func (msg MsgPendingUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPendingUndelegate) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.Delegator)
	return []sdk.AccAddress{withdrawer}
}

func NewMsgWithdraw(zoneId string, withdrawer sdk.AccAddress) *MsgWithdraw {
	return &MsgWithdraw{
		ZoneId:     zoneId,
		Withdrawer: withdrawer.String(),
	}
}

func (msg MsgWithdraw) Route() string {
	return RouterKey
}

func (msg MsgWithdraw) Type() string {
	return TypeMsgWithdrawRecord
}

func (msg MsgWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Withdrawer)
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

func NewMsgClaimSnAsset(zoneId string, claimer sdk.AccAddress) *MsgClaimSnAsset {
	return &MsgClaimSnAsset{
		ZoneId:  zoneId,
		Claimer: claimer.String(),
	}
}

func (msg MsgClaimSnAsset) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return err
	}

	return nil
}

func (msg MsgClaimSnAsset) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgClaimSnAsset) GetSigners() []sdk.AccAddress {
	claimer, _ := sdk.AccAddressFromBech32(msg.Claimer)
	return []sdk.AccAddress{claimer}
}

func NewIcaWithdraw(zoneId string, controllerAddr sdk.AccAddress, portId, chanId string, blockTime time.Time) *MsgIcaWithdraw {
	return &MsgIcaWithdraw{
		ZoneId:               zoneId,
		ControllerAddress:    controllerAddr.String(),
		IcaTransferPortId:    portId,
		IcaTransferChannelId: chanId,
		ChainTime:            blockTime,
	}
}

func (msg MsgIcaWithdraw) Route() string {
	return RouterKey
}

func (msg MsgIcaWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ControllerAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	if msg.ChainTime.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ControllerAddress)
	}

	return nil
}

func (msg MsgIcaWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgIcaWithdraw) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.ControllerAddress)
	return []sdk.AccAddress{withdrawer}
}
