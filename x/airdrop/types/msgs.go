package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgClaimAirdropRequest{}
var _ sdk.Msg = &MsgMarkSocialQuestPerformedRequest{}
var _ sdk.Msg = &MsgMarkUserProvidedLiquidityRequest{}

func (m *MsgClaimAirdropRequest) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgClaimAirdropRequest) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMarkSocialQuestPerformedRequest) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMarkSocialQuestPerformedRequest) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMarkUserProvidedLiquidityRequest) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMarkUserProvidedLiquidityRequest) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}
