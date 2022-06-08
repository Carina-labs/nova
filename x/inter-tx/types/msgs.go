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
)

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterZone(zone_name, chain_id, ica_connection_id, ica_owner_address, transfer_channel_id, transfer_connection_id, transfer_port_id,
	validator_address, base_denom string) *MsgRegisterZone {
	return &MsgRegisterZone{
		ZoneName: zone_name,
		ChainId:  chain_id,
		IcaInfo: &IcaConnectionInfo{
			ConnectionId: ica_connection_id,
			OwnerAddress: ica_owner_address,
		},
		TransferInfo: &TransferConnectionInfo{
			ConnectionId: transfer_connection_id,
			PortId:       transfer_port_id,
			ChannelId:    transfer_channel_id,
		},
		ValidatorAddress: validator_address,
		BaseDenom:        base_denom,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterZone) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneName) == "" {
		return errors.New("missing zone name")
	}

	if strings.TrimSpace(msg.ChainId) == "" {
		return errors.New("missing chain ID")
	}

	if strings.TrimSpace(msg.IcaInfo.ConnectionId) == "" {
		return errors.New("missing ICA connection ID")
	}

	if strings.TrimSpace(msg.IcaInfo.OwnerAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing ICA owner address")
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
	accAddr, err := sdk.AccAddressFromBech32(msg.IcaInfo.OwnerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

func NewMsgIcaDelegate(zone_name, sender, owner string, amount sdk.Coin) *MsgIcaDelegate {
	return &MsgIcaDelegate{
		ZoneName:      zone_name,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
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

func NewMsgIcaUnDelegate(zone_name, sender, owner string, amount sdk.Coin) *MsgIcaUndelegate {
	return &MsgIcaUndelegate{
		ZoneName:      zone_name,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaUndelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
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

func NewMsgIcaAutoStaking(zone_name, sender, owner string, amount sdk.Coin) *MsgIcaAutoStaking {
	return &MsgIcaAutoStaking{
		ZoneName:      zone_name,
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

func NewMsgIcaWithdraw(zone_name, sender, owner, receiver string, amount sdk.Coin) *MsgIcaWithdraw {
	return &MsgIcaWithdraw{
		ZoneName:        zone_name,
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

	//TODO: receiveAddress가 claim한 이력이 있는지 확인
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
