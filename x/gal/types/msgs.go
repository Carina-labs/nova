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
var _ sdk.Msg = &MsgPendingUndelegateRecord{}
var _ sdk.Msg = &MsgWithdraw{}
var _ sdk.Msg = &MsgClaimSnAsset{}
var _ sdk.Msg = &MsgPendingWithdraw{}

func NewMsgDeposit(zoneId string, depositor sdk.AccAddress, amount sdk.Coin, portId, chanId string) *MsgDeposit {
	return &MsgDeposit{
		ZoneId:            zoneId,
		Depositor:         depositor.String(),
		Amount:            amount,
		TransferPortId:    portId,
		TransferChannelId: chanId,
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

	if msg.TransferChannelId == "" {
		return errors.New("transfer channel id is not null")
	}

	if msg.TransferPortId == "" {
		return errors.New("transfer port id is not null")
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

func NewMsgDelegate(zoneId string, daomodifierAddr sdk.AccAddress, transferPortId, transferChanId string) *MsgDelegate {
	return &MsgDelegate{
		ZoneId:            zoneId,
		ControllerAddress: daomodifierAddr.String(),
		TransferPortId:    transferPortId,
		TransferChannelId: transferChanId,
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

func NewMsgPendingUndelegateRecord(zoneId, Depositor string, amount sdk.Coin) *MsgPendingUndelegateRecord {
	return &MsgPendingUndelegateRecord{
		ZoneId:    zoneId,
		Depositor: Depositor,
		Amount:    amount,
	}
}

func (msg MsgPendingUndelegateRecord) Route() string {
	return RouterKey
}

func (msg MsgPendingUndelegateRecord) Type() string {
	return TypeMsgUndelegateRecord
}

func (msg MsgPendingUndelegateRecord) ValidateBasic() error {
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

func (msg MsgPendingUndelegateRecord) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPendingUndelegateRecord) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.Depositor)
	return []sdk.AccAddress{withdrawer}
}

func NewMsgWithdraw(zoneId string, withdrawer sdk.AccAddress, receiver, portId, chanId string) *MsgWithdraw {
	return &MsgWithdraw{
		ZoneId:               zoneId,
		Withdrawer:           withdrawer.String(),
		Recipient:            receiver,
		IcaTransferPortId:    portId,
		IcaTransferChannelId: chanId,
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

func NewMsgClaimSnAsset(zoneId string, claimer sdk.AccAddress, transferPortId, transferChanId string) *MsgClaimSnAsset {
	return &MsgClaimSnAsset{
		ZoneId:            zoneId,
		Claimer:           claimer.String(),
		TransferPortId:    transferPortId,
		TransferChannelId: transferChanId,
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

func NewMsgPendingWithdraw(zoneId string, daomodifierAddr sdk.AccAddress, portId, chanId string, blockTime time.Time) *MsgPendingWithdraw {
	return &MsgPendingWithdraw{
		ZoneId:             zoneId,
		DaomodifierAddress: daomodifierAddr.String(),
		TransferPortId:     portId,
		TransferChannelId:  chanId,
		ChainTime:          blockTime,
	}
}

func (msg MsgPendingWithdraw) Route() string {
	return RouterKey
}

func (msg MsgPendingWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DaomodifierAddress)
	}

	if msg.ChainTime.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.DaomodifierAddress)
	}

	return nil
}

func (msg MsgPendingWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPendingWithdraw) GetSigners() []sdk.AccAddress {
	withdrawer, _ := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	return []sdk.AccAddress{withdrawer}
}
