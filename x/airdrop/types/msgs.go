package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgClaimAirdropRequest{}
var _ sdk.Msg = &MsgMarkSocialQuestPerformedRequest{}
var _ sdk.Msg = &MsgMarkUserProvidedLiquidityRequest{}
var _ sdk.Msg = &MsgImportAirdropDataRequest{}

func (m *MsgClaimAirdropRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.UserAddress)
	if err != nil {
		return err
	}

	switch m.QuestType {
	case QuestType_QUEST_NOTHING_TO_DO:
		return nil
	case QuestType_QUEST_SOCIAL:
		return nil
	case QuestType_QUEST_PROVIDE_LIQUIDITY:
		return nil
	case QuestType_QUEST_SN_ASSET_CLAIM:
		return nil
	case QuestType_QUEST_VOTE_ON_PROPOSALS:
		return nil
	default:
		return fmt.Errorf("invalid quest type: %v", m.QuestType)
	}
}

func (m *MsgClaimAirdropRequest) GetSigners() []sdk.AccAddress {
	userAddr, err := sdk.AccAddressFromBech32(m.UserAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{userAddr}
}

func (m *MsgMarkSocialQuestPerformedRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		return err
	}

	for _, addr := range m.UserAddresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return err
		}
	}

	return nil
}

func (m *MsgMarkSocialQuestPerformedRequest) GetSigners() []sdk.AccAddress {
	controllerAddr, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{controllerAddr}
}

func (m *MsgMarkUserProvidedLiquidityRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		return err
	}

	for _, addr := range m.UserAddresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return err
		}
	}

	return nil
}

func (m *MsgMarkUserProvidedLiquidityRequest) GetSigners() []sdk.AccAddress {
	controllerAddr, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{controllerAddr}
}

func (m *MsgImportAirdropDataRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		return err
	}

	if len(m.States) == 0 {
		return fmt.Errorf("user states cannot be nil")
	}

	return nil
}

func (m *MsgImportAirdropDataRequest) GetSigners() []sdk.AccAddress {
	controllerAddr, err := sdk.AccAddressFromBech32(m.ControllerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{controllerAddr}
}
