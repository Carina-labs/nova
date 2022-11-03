package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgCreateCandidatePool   = "createCandidatePool"
	TypeMsgCreateIncentivePool   = "createIncentivePool"
	TypeMsgSetPoolWeight         = "setPoolWeight"
	TypeMsgSetMultiplePoolWeight = "setMultiplePoolWeight"
)

var _ sdk.Msg = &MsgCreateCandidatePool{}
var _ sdk.Msg = &MsgSetPoolWeight{}
var _ sdk.Msg = &MsgCreateIncentivePool{}
var _ sdk.Msg = &MsgSetMultiplePoolWeight{}

func NewMsgCreateCandidatePool(poolId string, poolContractAddress string, creator sdk.AccAddress) *MsgCreateCandidatePool {
	return &MsgCreateCandidatePool{
		PoolId:              poolId,
		PoolContractAddress: poolContractAddress,
		Creator:             creator.String(),
	}
}

func (m MsgCreateCandidatePool) Route() string {
	return RouterKey
}

func (m MsgCreateCandidatePool) Type() string {
	return TypeMsgCreateCandidatePool
}

func (m MsgCreateCandidatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Creator); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.PoolContractAddress); err != nil {
		return err
	}

	return nil
}

func (m MsgCreateCandidatePool) GetSignBytes() []byte {
	return nil
}

func (m MsgCreateCandidatePool) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(m.Creator)
	return []sdk.AccAddress{creator}
}

func NewMsgSetPoolWeight(poolId string, newWeight uint64, operator sdk.AccAddress) *MsgSetPoolWeight {
	return &MsgSetPoolWeight{
		PoolId:    poolId,
		NewWeight: newWeight,
		Operator:  operator.String(),
	}
}

func (m MsgSetPoolWeight) Route() string {
	return RouterKey
}

func (m MsgSetPoolWeight) Type() string {
	return TypeMsgSetPoolWeight
}

func (m MsgSetPoolWeight) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return err
	}

	return nil
}

func (m MsgSetPoolWeight) GetSignBytes() []byte {
	return nil
}

func (m MsgSetPoolWeight) GetSigners() []sdk.AccAddress {
	controller, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{controller}
}

func NewMsgCreateIncentivePool(poolId string, poolContractAddress string, operator sdk.AccAddress) *MsgCreateIncentivePool {
	return &MsgCreateIncentivePool{
		PoolId:              poolId,
		PoolContractAddress: poolContractAddress,
		Operator:            operator.String(),
	}
}

func (m MsgCreateIncentivePool) Route() string {
	return RouterKey
}

func (m MsgCreateIncentivePool) Type() string {
	return TypeMsgCreateIncentivePool
}

func (m MsgCreateIncentivePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.PoolContractAddress); err != nil {
		return err
	}
	return nil
}

func (m MsgCreateIncentivePool) GetSignBytes() []byte {
	return nil
}

func (m MsgCreateIncentivePool) GetSigners() []sdk.AccAddress {
	controller, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{controller}
}

func NewMsgSetMultipleWeight(newPoolData []NewPoolWeight, operator sdk.AccAddress) *MsgSetMultiplePoolWeight {
	return &MsgSetMultiplePoolWeight{
		NewPoolData: newPoolData,
		Operator:    operator.String(),
	}
}

func (m MsgSetMultiplePoolWeight) Route() string {
	return RouterKey
}

func (m MsgSetMultiplePoolWeight) Type() string {
	return TypeMsgSetMultiplePoolWeight
}

func (m MsgSetMultiplePoolWeight) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return err
	}

	return nil
}

func (m MsgSetMultiplePoolWeight) GetSignBytes() []byte {
	return nil
}

func (m MsgSetMultiplePoolWeight) GetSigners() []sdk.AccAddress {
	controller, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{controller}
}
