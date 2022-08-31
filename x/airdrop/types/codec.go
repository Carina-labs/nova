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

// RegisterCodec registers the airdrop module's types on the given LegacyAmino codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimAirdropRequest{}, "supernova/airdrop/claim-airdrop", nil)
	cdc.RegisterConcrete(&MsgMarkUserProvidedLiquidityRequest{}, "supernova/airdrop/mark-user-provided-liquidity", nil)
	cdc.RegisterConcrete(&MsgMarkSocialQuestPerformedRequest{}, "supernova/airdrop/mark-user-performed-social", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgClaimAirdropRequest{},
		&MsgMarkUserProvidedLiquidityRequest{},
		&MsgMarkSocialQuestPerformedRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
