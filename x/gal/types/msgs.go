package types

import (
	"errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDeposit           = "deposit"
	TypeMsgDelegate          = "delegate"
	TypeMsgUndelegate        = "undelegate"
	TypeMsgPendingUndelegate = "pendingUndelegate"
	TypeMsgWithdrawRecord    = "withdrawRecord"
	TypeMsgClaim             = "claim"
	TypeMsgIcaWithdraw       = "icaWithdraw"
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

	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return err
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	if err := transfertypes.ValidateIBCDenom(msg.Amount.Denom); err != nil {
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

func NewMsgDelegate(zoneId string, version uint64, controllerAddr sdk.AccAddress) *MsgDelegate {
	return &MsgDelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
		Version:           version,
	}
}

func (msg MsgDelegate) Route() string {
	return RouterKey
}

func (msg MsgDelegate) Type() string {
	return TypeMsgDelegate
}

func (msg MsgDelegate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ControllerAddress); err != nil {
		return err
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
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

func NewMsgUndelegate(zoneId string, version uint64, controllerAddr sdk.AccAddress) *MsgUndelegate {
	return &MsgUndelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
		Version:           version,
	}
}

func (msg MsgUndelegate) Route() string {
	return RouterKey
}

func (msg MsgUndelegate) Type() string {
	return TypeMsgUndelegate
}

func (msg MsgUndelegate) ValidateBasic() error {
	if msg.ControllerAddress == "" {
		return errors.New("controller address is not null")
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	if msg.Version < 0 {
		return ErrNegativeVersion
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
	return TypeMsgPendingUndelegate
}

func (msg MsgPendingUndelegate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Delegator); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
		return err
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	if err := sdk.ValidateDenom(msg.Amount.Denom); err != nil {
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
	delegator, _ := sdk.AccAddressFromBech32(msg.Delegator)
	return []sdk.AccAddress{delegator}
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

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		return err
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

func (msg MsgClaimSnAsset) Route() string {
	return RouterKey
}

func (msg MsgClaimSnAsset) Type() string {
	return TypeMsgClaim
}

func (msg MsgClaimSnAsset) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return err
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
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

func NewMsgIcaWithdraw(zoneId string, controllerAddr sdk.AccAddress, portId, chanId string, blockTime time.Time, version uint64) *MsgIcaWithdraw {
	return &MsgIcaWithdraw{
		ZoneId:               zoneId,
		ControllerAddress:    controllerAddr.String(),
		IcaTransferPortId:    portId,
		IcaTransferChannelId: chanId,
		ChainTime:            blockTime,
		Version:              version,
	}
}

func (msg MsgIcaWithdraw) Route() string {
	return RouterKey
}

func (msg MsgIcaWithdraw) Type() string {
	return TypeMsgIcaWithdraw
}

func (msg MsgIcaWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ControllerAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ControllerAddress)
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrap(ErrNotFoundZoneInfo, "zoneId is not nil")
	}

	if msg.IcaTransferChannelId == "" {
		return sdkerrors.Wrap(ErrTransferInfoNotFound, "transfer channel id is not nil")
	}

	if msg.IcaTransferPortId == "" {
		return sdkerrors.Wrap(ErrTransferInfoNotFound, "transfer port id is not nil")
	}

	if msg.ChainTime.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ChainTime.String())
	}

	if msg.Version < 0 {
		return ErrNegativeVersion
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
