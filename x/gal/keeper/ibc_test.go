package keeper_test

import (
	novatesting "github.com/Carina-labs/nova/testing"
	icacontroltypes "github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

var (
	transferPort    = "transfer"
	transferChannel = "channel-0"
	icaConnection   = "connection-1"

	zoneId       = "baseZone"
	baseOwnerAcc = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	baseDenom    = "stake"
	baseSnDenom  = "snstake"
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
func newBaseRegisteredZone() *icacontroltypes.RegisteredZone {
	icaControllerPort := zoneId + "." + baseOwnerAcc.String()
	return &icacontroltypes.RegisteredZone{
		ZoneId: zoneId,
		IcaConnectionInfo: &icacontroltypes.IcaConnectionInfo{
			ConnectionId: icaConnection,
			PortId:       icaControllerPort,
		},
		IcaAccount: &icacontroltypes.IcaAccount{
			ControllerAddress: baseOwnerAcc.String(),
		},
		TransferInfo: &icacontroltypes.TransferConnectionInfo{
			PortId:    transferPort,
			ChannelId: transferChannel,
		},
		BaseDenom:  baseDenom,
		SnDenom:    baseSnDenom,
		Decimal:    0,
		MaxEntries: 5,
	}
}
