package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// TODO : implements this!
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	// TODO : implements this!
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
