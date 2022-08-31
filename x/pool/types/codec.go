package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCandidatePool{}, "supernova/pool/create-candidate-pool", nil)
	cdc.RegisterConcrete(&MsgCreateIncentivePool{}, "supernova/pool/create-incentive-pool", nil)
	cdc.RegisterConcrete(&MsgSetPoolWeight{}, "supernova/pool/set-pool-weight", nil)
	cdc.RegisterConcrete(&MsgSetMultiplePoolWeight{}, "supernova/pool/set-multiple-pool-weight", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateCandidatePool{},
		&MsgCreateIncentivePool{},
		&MsgSetPoolWeight{},
		&MsgSetMultiplePoolWeight{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
