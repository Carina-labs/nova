package types

import (
	fmt "fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

var (
	_ sdk.Msg = &MsgRegisterZone{}
	_ sdk.Msg = &MsgIcaDelegate{}
	_ sdk.Msg = &MsgIcaUndelegate{}
	_ sdk.Msg = &MsgIcaAutoCompound{}
)

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterZone(zone_name, chain_id, connection_id, owner_address, validator_address, base_denom string) *MsgRegisterZone {
	return &MsgRegisterZone{
		ZoneName:         zone_name,
		ChainId:          chain_id,
		ConnectionId:     connection_id,
		OwnerAddress:     owner_address,
		ValidatorAddress: validator_address,
		BaseDenom:        base_denom,
		AuthzAddress:     "test",
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterZone) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneName) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing zone name")
	}

	if strings.TrimSpace(msg.ChainId) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing chain id")
	}

	if strings.TrimSpace(msg.ConnectionId) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing connection id")
	}

	if strings.TrimSpace(msg.OwnerAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}

	if strings.TrimSpace(msg.ValidatorAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}

	if strings.TrimSpace(msg.BaseDenom) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing denom")
	}

	if strings.TrimSpace(msg.AuthzAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing authz address")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterZone) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
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

// PackTxMsgAny marshals the sdk.Msg payload to a protobuf Any type
func PackTxMsgAny(sdkMsg sdk.Msg) (*codectypes.Any, error) {
	msg, ok := sdkMsg.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("can't proto marshal %T", sdkMsg)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	return any, nil
}

func NewMsgIcaAutoCompound(zone_name, sender, owner string, amount sdk.Coin) *MsgIcaAutoCompound {
	return &MsgIcaAutoCompound{
		ZoneName:      zone_name,
		SenderAddress: sender,
		OwnerAddress:  owner,
		Amount:        amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgIcaAutoCompound) ValidateBasic() error {
	if strings.TrimSpace(msg.ZoneName) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing zone name")
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
func (msg MsgIcaAutoCompound) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.OwnerAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}
