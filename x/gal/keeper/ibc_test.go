package keeper_test

import (
	novatesting "github.com/Carina-labs/nova/testing"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	transferPort    = "transfer"
	transferChannel = "channel-0"
	icaConnection   = "connection-1"

	zoneId       = "baseZone"
	baseOwnerAcc = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	baseDenom    = "base"
	baseSnDenom  = "snbase"
)

func NewIbcTransferPath(chainA, chainB *novatesting.TestChain) *novatesting.Path {
	path := novatesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	return path
}

func newIcaPath(chainA, chainB *novatesting.TestChain) *novatesting.Path {
	path := novatesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = icatypes.PortID
	path.EndpointB.ChannelConfig.PortID = icatypes.PortID
	path.EndpointA.ChannelConfig.Order = ibcchanneltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = ibcchanneltypes.ORDERED

	version := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                "ics27-1",
		ControllerConnectionId: icaConnection,
		HostConnectionId:       icaConnection,
		Encoding:               "proto3",
		TxType:                 "sdk_multi_msg",
	}))
	path.EndpointA.ChannelConfig.Version = version
	path.EndpointB.ChannelConfig.Version = version
	return path
}

func setupIcaPath(path *novatesting.Path, owner string) error {
	if err := registerInterchainAccount(path.EndpointA, owner); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenTry(); err != nil {
		return err
	}

	if err := path.EndpointA.ChanOpenAck(); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenConfirm(); err != nil {
		return err
	}

	return nil
}

func registerInterchainAccount(e *novatesting.Endpoint, owner string) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelSeq := e.Chain.App.IBCKeeper.ChannelKeeper.GetNextChannelSequence(e.Chain.GetContext())

	if err := e.Chain.App.ICAControllerKeeper.RegisterInterchainAccount(
		e.Chain.GetContext(), e.ConnectionID, owner); err != nil {
		return err
	}

	e.Chain.NextBlock()
	e.ChannelID = ibcchanneltypes.FormatChannelIdentifier(channelSeq)
	e.ChannelConfig.PortID = portID
	return nil
}

// newBaseRegisteredZone returns a new zone info for testing purpose only
func newBaseRegisteredZone() *ibcstakingtypes.RegisteredZone {
	icaControllerPort := zoneId + "." + baseOwnerAcc.String()
	return &ibcstakingtypes.RegisteredZone{
		ZoneId: zoneId,
		IcaConnectionInfo: &ibcstakingtypes.IcaConnectionInfo{
			ConnectionId: icaConnection,
			PortId:       icaControllerPort,
		},
		IcaAccount: &ibcstakingtypes.IcaAccount{
			ControllerAddress: baseOwnerAcc.String(),
		},
		TransferInfo: &ibcstakingtypes.TransferConnectionInfo{
			PortId:    transferPort,
			ChannelId: transferChannel,
		},
		ValidatorAddress: "",
		BaseDenom:        baseDenom,
		SnDenom:          baseSnDenom,
	}
}
