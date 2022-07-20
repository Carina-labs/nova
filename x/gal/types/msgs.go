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
var _ sdk.Msg = &MsgGalUndelegate{}
var _ sdk.Msg = &MsgUndelegateRecord{}
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

func NewMsgUndelegate(zoneId, controllerAddr string) *MsgGalUndelegate {
	return &MsgGalUndelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr,
	}
}

func (msg MsgGalUndelegate) Route() string {
	return RouterKey
}

func (msg MsgGalUndelegate) Type() string {
	return TypeMsgDeposit
}

func (msg MsgGalUndelegate) ValidateBasic() error {
	if msg.ControllerAddress == "" {
		return errors.New("controller address is not null")
	}

	if msg.ZoneId == "" {
		return errors.New("zone id is not found")
	}
	return nil
}

func (msg MsgGalUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgGalUndelegate) GetSigners() []sdk.AccAddress {
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

func NewMsgWithdraw(zoneId string, withdrawer sdk.AccAddress, receiver, portId, chanId string) *MsgWithdraw {
	return &MsgWithdraw{
		ZoneId:            zoneId,
		Withdrawer:        withdrawer.String(),
		Recipient:         receiver,
		TransferPortId:    portId,
		TransferChannelId: chanId,
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
