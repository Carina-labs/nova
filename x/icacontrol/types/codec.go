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
	cdc.RegisterConcrete(MsgRegisterZone{}, "icacontrol/MsgRegisterZone", nil)
	cdc.RegisterConcrete(MsgIcaDelegate{}, "icacontrol/MsgIcaDelegate", nil)
	cdc.RegisterConcrete(MsgIcaUndelegate{}, "icacontrol/MsgIcaUndelegate", nil)
	cdc.RegisterConcrete(MsgIcaTransfer{}, "icacontrol/MsgIcaTransfer", nil)
	cdc.RegisterConcrete(MsgIcaAutoStaking{}, "icacontrol/MsgIcaAutoStaking", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterZone{},
		&MsgIcaDelegate{},
		&MsgIcaUndelegate{},
		&MsgIcaTransfer{},
		&MsgIcaAutoStaking{},
	)
}
