package types_test

import (
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgClaimAirdropRequest_ValidateBasic(t *testing.T) {
	invalidAddr := "foobar"

	msg := &types.MsgClaimAirdropRequest{
		UserAddress: sdk.AccAddress([]byte{0x1}).String(),
		QuestType:   types.QuestType_QUEST_PROVIDE_LIQUIDITY,
	}
	err := msg.ValidateBasic()
	require.NoError(t, err)

	msg.UserAddress = invalidAddr
	err = msg.ValidateBasic()
	require.Error(t, err)
}

func TestMsgClaimAirdropRequest_GetSigners(t *testing.T) {
	userAddr := sdk.AccAddress([]byte{0x1})
	msg := &types.MsgClaimAirdropRequest{
		UserAddress: userAddr.String(),
		QuestType:   types.QuestType_QUEST_PROVIDE_LIQUIDITY,
	}
	signers := msg.GetSigners()
	require.Equal(t, signers, []sdk.AccAddress{userAddr})
}

func TestMsgMarkSocialQuestPerformedRequest_ValidateBasic(t *testing.T) {
	validUser := sdk.AccAddress([]byte{0x1})
	invalidAddr := "foobar"

	msg := &types.MsgMarkSocialQuestPerformedRequest{
		ControllerAddress: validUser.String(),
		UserAddresses:     []string{validUser.String()},
	}
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// check controller address is invalid
	msg.ControllerAddress = invalidAddr
	err = msg.ValidateBasic()
	require.Error(t, err)

	// check user addresses is invalid
	msg.ControllerAddress = validUser.String()
	msg.UserAddresses = []string{invalidAddr}
	err = msg.ValidateBasic()
	require.Error(t, err)
}

func TestMsgMarkSocialQuestPerformedRequest_GetSigners(t *testing.T) {
	controllerAddr := sdk.AccAddress([]byte{0x1})
	userAddr := sdk.AccAddress([]byte{0x2})

	msg := &types.MsgMarkSocialQuestPerformedRequest{
		ControllerAddress: controllerAddr.String(),
		UserAddresses:     []string{userAddr.String()},
	}
	signers := msg.GetSigners()
	require.Equal(t, signers, []sdk.AccAddress{controllerAddr})
}

func TestMsgMarkUserProvidedLiquidityRequest_ValidateBasic(t *testing.T) {
	validUser := sdk.AccAddress([]byte{0x1})
	invalidAddr := "foobar"

	msg := &types.MsgMarkUserProvidedLiquidityRequest{
		ControllerAddress: validUser.String(),
		UserAddresses:     []string{validUser.String()},
	}
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// check controller address is invalid
	msg.ControllerAddress = invalidAddr
	err = msg.ValidateBasic()
	require.Error(t, err)

	// check user addresses is invalid
	msg.ControllerAddress = validUser.String()
	msg.UserAddresses = []string{invalidAddr}
	err = msg.ValidateBasic()
	require.Error(t, err)
}

func TestMsgMarkUserProvidedLiquidityRequest_GetSigners(t *testing.T) {
	controllerAddr := sdk.AccAddress([]byte{0x1})
	userAddr := sdk.AccAddress([]byte{0x2})

	msg := &types.MsgMarkUserProvidedLiquidityRequest{
		ControllerAddress: controllerAddr.String(),
		UserAddresses:     []string{userAddr.String()},
	}
	signers := msg.GetSigners()
	require.Equal(t, signers, []sdk.AccAddress{controllerAddr})
}
