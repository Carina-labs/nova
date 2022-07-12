package types

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgRegisterZone{}
	_ sdk.Msg = &MsgIcaDelegate{}
	_ sdk.Msg = &MsgIcaUndelegate{}
	_ sdk.Msg = &MsgIcaTransfer{}
	_ sdk.Msg = &MsgIcaAutoStaking{}
	_ sdk.Msg = &MsgRegisterHostAccount{}

	//modify
	_ sdk.Msg = &MsgDeleteRegisteredZone{}
	_ sdk.Msg = &MsgChangeRegisteredZoneInfo{}
)

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterZone(zoneId, icaConnectionId string, daomodifierAddress sdk.AccAddress, validatorAddress, baseDenom string) *MsgRegisterZone {
	return &MsgRegisterZone{
		ZoneId: zoneId,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: icaConnectionId,
			PortId:       "icacontroller-" + daomodifierAddress.String(),
		},
		IcaAccount: &IcaAccount{
			DaomodifierAddress: daomodifierAddress.String(),
		},
		ValidatorAddress: validatorAddress,
		BaseDenom:        baseDenom,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterZone) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneId) == "" {
		return errors.New("missing zone name")
	}

	if strings.TrimSpace(msg.IcaInfo.ConnectionId) == "" {
		return errors.New("missing ICA connection ID")
	}

	if strings.TrimSpace(msg.IcaInfo.PortId) == "" {
		return errors.New("missing ICA port ID")
	}

	_, err := sdk.AccAddressFromBech32(msg.IcaAccount.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
	}

	if strings.TrimSpace(msg.ValidatorAddress) == "" {
		return errors.New("missing validator address")
	}

	if strings.TrimSpace(msg.BaseDenom) == "" {
		return errors.New("missing denom")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterZone) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaAccount.DaomodifierAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaDelegate(zoneId, daomodifierAddr, hostAddr string, amount sdk.Coin) *MsgIcaDelegate {
	return &MsgIcaDelegate{
		ZoneId:             zoneId,
		DaomodifierAddress: daomodifierAddr,
		HostAddress:        hostAddr,
		Amount:             amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	if strings.TrimSpace(msg.HostAddress) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaDelegate) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaUnDelegate(zoneId, hostAddr string, daomodifierAddr sdk.AccAddress, amount sdk.Coin) *MsgIcaUndelegate {
	return &MsgIcaUndelegate{
		ZoneId:             zoneId,
		DaomodifierAddress: daomodifierAddr.String(),
		HostAddress:        hostAddr,
		Amount:             amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaUndelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	if strings.TrimSpace(msg.HostAddress) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaUndelegate) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaAutoStaking(zoneId, hostAddr, daomodifierAddr string, amount sdk.Coin) *MsgIcaAutoStaking {
	return &MsgIcaAutoStaking{
		ZoneId:             zoneId,
		HostAddress:        hostAddr,
		DaomodifierAddress: daomodifierAddr,
		Amount:             amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaAutoStaking) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneId) == "" {
		return errors.New("missing zone name")
	}

	_, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	if strings.TrimSpace(msg.HostAddress) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaAutoStaking) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaTransfer(zoneId, hostAddr string, daomodifierAddr sdk.AccAddress, receiver, portId, chanId string, amount sdk.Coin) *MsgIcaTransfer {
	return &MsgIcaTransfer{
		ZoneId:               zoneId,
		HostAddress:          hostAddr,
		DaomodifierAddress:   daomodifierAddr.String(),
		ReceiverAddress:      receiver,
		IcaTransferPortId:    portId,
		IcaTransferChannelId: chanId,
		Amount:               amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaTransfer) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneId) == "" {
		return errors.New("missing zone name")
	}

	_, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	if strings.TrimSpace(msg.HostAddress) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaTransfer) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgRegisterHostAccount(zoneId, hostAddr, daomodifierAddr string) *MsgRegisterHostAccount {
	return &MsgRegisterHostAccount{
		ZoneId: zoneId,
		AccountInfo: &IcaAccount{
			DaomodifierAddress: daomodifierAddr,
			HostAddress:        hostAddr,
		},
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterHostAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AccountInfo.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid modifier address")
	}

	if strings.TrimSpace(msg.AccountInfo.HostAddress) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterHostAccount) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.AccountInfo.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgDeleteRegisteredZone(zoneId, daomodifierAddr string) *MsgDeleteRegisteredZone {
	return &MsgDeleteRegisteredZone{
		ZoneId:             zoneId,
		DaomodifierAddress: daomodifierAddr,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteRegisteredZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteRegisteredZone) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgChangeZoneInfo(zoneId, icaConnectionId string, daomodifierAddress sdk.AccAddress, validatorAddress, baseDenom string) *MsgChangeRegisteredZoneInfo {
	return &MsgChangeRegisteredZoneInfo{
		ZoneId: zoneId,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: icaConnectionId,
			PortId:       "icacontroller-" + daomodifierAddress.String(),
		},
		IcaAccount: &IcaAccount{
			DaomodifierAddress: daomodifierAddress.String(),
		},
		ValidatorAddress: validatorAddress,
		BaseDenom:        baseDenom,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgChangeRegisteredZoneInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IcaAccount.DaomodifierAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid daomodifier address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgChangeRegisteredZoneInfo) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaAccount.DaomodifierAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}
