package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCandidatePool{}, "nova/MsgCreateCandidatePool", nil)
	cdc.RegisterConcrete(&MsgCreateIncentivePool{}, "nova/MsgCreateIncentivePool", nil)
	cdc.RegisterConcrete(&MsgSetPoolWeight{}, "nova/MsgSetPoolWeight", nil)
	cdc.RegisterConcrete(&MsgSetMultiplePoolWeight{}, "nova/MsgSetMultiplePoolWeight", nil)
	cdc.RegisterConcrete(&UpdatePoolIncentivesProposal{}, "nova/UpdatePoolIncentivesProposal", nil)
	cdc.RegisterConcrete(&ReplacePoolIncentivesProposal{}, "nova/ReplacePoolIncentivesProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateCandidatePool{},
		&MsgCreateIncentivePool{},
		&MsgSetPoolWeight{},
		&MsgSetMultiplePoolWeight{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdatePoolIncentivesProposal{},
		&ReplacePoolIncentivesProposal{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
