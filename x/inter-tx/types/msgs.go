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
	_ sdk.Msg = &MsgIcaWithdraw{}
	_ sdk.Msg = &MsgIcaAutoStaking{}
	_ sdk.Msg = &MsgRegisterHostAccount{}
)

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterZone(zoneId, chainId, icaConnectionId, icaOwnerAddr, transferChannelId, transferConnectionId, transferPortId,
	validatorAddress, baseDenom string) *MsgRegisterZone {
	return &MsgRegisterZone{
		ZoneId: zoneId,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: icaConnectionId,
			PortId:       icaOwnerAddr,
		},
		TransferInfo: &TransferConnectionInfo{
			ConnectionId: transferConnectionId,
			PortId:       "icacontroller-" + icaOwnerAddr,
			ChannelId:    transferChannelId,
		},
		IcaAccount: &IcaAccount{
			OwnerAddress: icaOwnerAddr,
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

	_, err := sdk.AccAddressFromBech32(msg.IcaAccount.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address (%s)", err)
	}

	if strings.TrimSpace(msg.TransferInfo.ConnectionId) == "" {
		return errors.New("missing IBC transfer connection ID")
	}

	if strings.TrimSpace(msg.TransferInfo.PortId) == "" {
		return errors.New("missing IBC transfer port ID")
	}

	if strings.TrimSpace(msg.TransferInfo.ChannelId) == "" {
		return errors.New("missing IBC transfer channel ID")
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
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaAccount.OwnerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaDelegate(zoneName, sender, owner string, amount sdk.Coin) *MsgIcaDelegate {
	return &MsgIcaDelegate{
		ZoneName:      zoneName,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address (%s)", err)
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
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaUnDelegate(zoneName, sender, owner string, amount sdk.Coin) *MsgIcaUndelegate {
	return &MsgIcaUndelegate{
		ZoneName:      zoneName,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaUndelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address (%s)", err)
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
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaAutoStaking(zoneName, sender, owner string, amount sdk.Coin) *MsgIcaAutoStaking {
	return &MsgIcaAutoStaking{
		ZoneName:      zoneName,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaAutoStaking) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneName) == "" {
		return errors.New("missing zone name")
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
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaWithdraw(zoneName, sender, owner, receiver string, amount sdk.Coin) *MsgIcaWithdraw {
	return &MsgIcaWithdraw{
		ZoneName:        zoneName,
		SenderAddress:   sender,
		OwnerAddress:    owner,
		ReceiverAddress: receiver,
		Amount:          amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaWithdraw) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneName) == "" {
		return errors.New("missing zone name")
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
func (msg MsgIcaWithdraw) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgRegisterHostAccount(ownerAddr, hostAddr string) *MsgRegisterHostAccount {
	return &MsgRegisterHostAccount{
		AccountInfo: &IcaAccount{
			OwnerAddress: ownerAddr,
			HostAddress:  hostAddr,
		},
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterHostAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AccountInfo.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid controller address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.AccountInfo.HostAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid host address (%s)", err)
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterHostAccount) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.AccountInfo.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}
