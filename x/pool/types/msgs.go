package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgCreatePool       = "createPool"
	TypeMsgModifyPoolWeight = "modifyPoolWeight"
)

var _ sdk.Msg = &MsgCreatePool{}
var _ sdk.Msg = &MsgSetPoolWeight{}

func NewMsgCreatePool(poolId string, poolContractAddress string) *MsgCreatePool {
	return &MsgCreatePool{
		PoolId:              poolId,
		PoolContractAddress: poolContractAddress,
	}
}

func (m MsgCreatePool) Route() string {
	return RouterKey
}

func (m MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (m MsgCreatePool) ValidateBasic() error {
	return nil
}

func (m MsgCreatePool) GetSignBytes() []byte {
	return nil
}

func (m MsgCreatePool) GetSigners() []sdk.AccAddress {
	return nil
}

func NewMsgModifyPoolWeight(poolId string, newWeight uint64, operator string) *MsgSetPoolWeight {
	return &MsgSetPoolWeight{
		PoolId:    poolId,
		NewWeight: newWeight,
		Operator:  operator,
	}
}

func (m MsgSetPoolWeight) Route() string {
	return RouterKey
}

func (m MsgSetPoolWeight) Type() string {
	return TypeMsgCreatePool
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
