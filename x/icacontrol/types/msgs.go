package types

import (
	"errors"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/gogo/protobuf/proto"
	"strings"
	"time"
)

var (
	_ sdk.Msg = &MsgRegisterZone{}
	_ sdk.Msg = &MsgIcaDelegate{}
	_ sdk.Msg = &MsgIcaUndelegate{}
	_ sdk.Msg = &MsgIcaTransfer{}
	_ sdk.Msg = &MsgIcaAutoStaking{}
	_ sdk.Msg = &MsgIcaAuthzGrant{}
	_ sdk.Msg = &MsgIcaAuthzRevoke{}

	//modify
	_ sdk.Msg = &MsgDeleteRegisteredZone{}
	_ sdk.Msg = &MsgChangeRegisteredZone{}
)

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterZone(zoneId, icaConnectionId string, controllerAddr sdk.AccAddress, transferPortId, transferChanId string, validatorAddress, baseDenom string, decimal int64) *MsgRegisterZone {
	return &MsgRegisterZone{
		ZoneId: zoneId,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: icaConnectionId,
			PortId:       zoneId + "." + controllerAddr.String(),
		},
		IcaAccount: &IcaAccount{
			ControllerAddress: controllerAddr.String(),
		},
		TransferInfo: &TransferConnectionInfo{
			ChannelId: transferChanId,
			PortId:    transferPortId,
		},
		ValidatorAddress: validatorAddress,
		BaseDenom:        baseDenom,
		Decimal:          decimal,
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

	_, err := sdk.AccAddressFromBech32(msg.IcaAccount.ControllerAddress)
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
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaAccount.ControllerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaDelegate(zoneId string, controllerAddr sdk.AccAddress, hostAddr string, amount sdk.Coin) *MsgIcaDelegate {
	return &MsgIcaDelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
		HostAddress:       hostAddr,
		Amount:            amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
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
	accAddr, err := sdk.AccAddressFromBech32(msg.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaUnDelegate(zoneId, hostAddr string, controllerAddr sdk.AccAddress, amount sdk.Coin) *MsgIcaUndelegate {
	return &MsgIcaUndelegate{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
		HostAddress:       hostAddr,
		Amount:            amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaUndelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
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
	accAddr, err := sdk.AccAddressFromBech32(msg.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaAutoStaking(zoneId string, controllerAddr sdk.AccAddress, amount sdk.Coin) *MsgIcaAutoStaking {
	return &MsgIcaAutoStaking{
		ZoneId:             zoneId,
		DaomodifierAddress: controllerAddr.String(),
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
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

func NewMsgIcaTransfer(zoneId, hostAddr string, controllerAddr sdk.AccAddress, receiver, portId, chanId string, amount sdk.Coin) *MsgIcaTransfer {
	return &MsgIcaTransfer{
		ZoneId:               zoneId,
		HostAddress:          hostAddr,
		DaomodifierAddress:   controllerAddr.String(),
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
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

func NewMsgDeleteRegisteredZone(zoneId string, controllerAddr sdk.AccAddress) *MsgDeleteRegisteredZone {
	return &MsgDeleteRegisteredZone{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteRegisteredZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteRegisteredZone) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgChangeZoneInfo(zoneId, hostAddr string, controllerAddr sdk.AccAddress, icaConnectionId, transferPortId, transferChanId, validatorAddress, baseDenom string, decimal int64) *MsgChangeRegisteredZone {
	return &MsgChangeRegisteredZone{
		ZoneId: zoneId,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: icaConnectionId,
			PortId:       zoneId + "." + controllerAddr.String(),
		},
		IcaAccount: &IcaAccount{
			ControllerAddress: controllerAddr.String(),
			HostAddress:       hostAddr,
		},
		TransferInfo: &TransferConnectionInfo{
			PortId:    transferPortId,
			ChannelId: transferChanId,
		},
		ValidatorAddress: validatorAddress,
		BaseDenom:        baseDenom,
		Decimal:          decimal,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgChangeRegisteredZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IcaAccount.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgChangeRegisteredZone) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaAccount.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgAuthzGrant(zoneId, grantee string, granter sdk.AccAddress, authorization authz.Authorization, expiration time.Time) (*MsgIcaAuthzGrant, error) {
	m := &MsgIcaAuthzGrant{
		ZoneId:            zoneId,
		ControllerAddress: granter.String(),
		Grantee:           grantee,
		Grant:             authz.Grant{Expiration: expiration},
	}
	err := m.SetAuthorization(authorization)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// SetAuthorization converts Authorization to any and adds it to MsgGrant.Authorization.
func (msg *MsgIcaAuthzGrant) SetAuthorization(a authz.Authorization) error {
	m, ok := a.(proto.Message)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrPackAny, "can't proto marshal %T", m)
	}
	any, err := cdctypes.NewAnyWithValue(m)
	if err != nil {
		return err
	}
	msg.Grant.Authorization = any
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaAuthzGrant) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
	}

	if msg.Grantee == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "Grantee address is not nil")
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrapf(ErrZoneIdNotNil, "Grantee address is not nil")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaAuthzGrant) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgAuthzRevoke(zoneId, grantee, msgType string, granter sdk.AccAddress) *MsgIcaAuthzRevoke {
	return &MsgIcaAuthzRevoke{
		ZoneId:            zoneId,
		ControllerAddress: granter.String(),
		Grantee:           grantee,
		MsgTypeUrl:        msgType,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaAuthzRevoke) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address")
	}

	if msg.Grantee == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "Grantee address is not nil")
	}

	if msg.ZoneId == "" {
		return sdkerrors.Wrapf(ErrZoneIdNotNil, "Grantee address is not nil")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgIcaAuthzRevoke) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.ControllerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}
