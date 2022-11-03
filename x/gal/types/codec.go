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

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgDeposit{}, "gal/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "gal/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgClaimSnAsset{}, "gal/MsgClaimSnAsset", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "gal/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "gal/MsgUndelegate", nil)
	cdc.RegisterConcrete(MsgPendingUndelegate{}, "gal/MsgPendingUndelegate", nil)
	cdc.RegisterConcrete(MsgIcaWithdraw{}, "gal/MsgIcaWithdraw", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgDelegate{},
		&MsgClaimSnAsset{},
		&MsgWithdraw{},
		&MsgUndelegate{},
		&MsgPendingUndelegate{},
		&MsgIcaWithdraw{},
	)
}
