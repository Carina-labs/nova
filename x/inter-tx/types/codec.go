package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgRegisterZone{}, "intertx/MsgRegisterZone", nil)
	cdc.RegisterConcrete(MsgIcaDelegate{}, "intertx/MsgIcaDelegate", nil)
	cdc.RegisterConcrete(MsgIcaUndelegate{}, "intertx/MsgIcaUndelegate", nil)
	cdc.RegisterConcrete(MsgIcaAutoCompound{}, "intertx/v1/MsgIcaAutoCompound", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterZone{},
		&MsgIcaDelegate{},
		&MsgIcaUndelegate{},
		&MsgIcaAutoCompound{},
	)
}
