package keeper_test

import (
	novatesting "github.com/Carina-labs/nova/testing"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	transferPort       = "transfer"
	transferChannel    = "channel-0"
	transferConnection = "connection-0"
	icaConnection      = "connection-1"

	baseDenom    = "unova"
	baseIbcDenom = ParseAddressToIbcAddress(transferPort, transferChannel, baseDenom)
	baseAcc      = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	baseHostAcc  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	baseOwnerAcc = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	baseSnDenom  = "snstake"

	oracleOperatorAcc = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	zoneId        = "cosmos-1"
	zoneBaseDenom = "stake"
	zoneIbcDeno   = ParseAddressToIbcAddress(transferPort, transferChannel, zoneBaseDenom)

	undelegateMsgName  = "/cosmos.staking.v1beta1.MsgUndelegate"
	ibcTransferMsgName = "/ibc.applications.transfer.v1.MsgTransfer"
)

func (suite *KeeperTestSuite) TestSetIbcZone(zoneMsg []ibcstakingtypes.RegisteredZone) {
	for _, msg := range zoneMsg {
		suite.App.IbcstakingKeeper.RegisterZone(suite.Ctx, &msg)
	}
}

func (suite *KeeperTestSuite) TestSetOracle(operators []sdk.Address, msg []oracletypes.ChainInfo) {
	for _, operator := range operators {
		suite.App.OracleKeeper.SetParams(suite.Ctx, oracletypes.Params{
			OracleOperators: []string{operator.String()},
		})
	}

	for _, m := range msg {
		err := suite.App.OracleKeeper.UpdateChainState(suite.Ctx, &m)
		suite.Require().NoError(err)
	}
}

func NewIbcTransferPath(chainA, chainB *novatesting.TestChain) *novatesting.Path {
	path := novatesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	return path
}

func NewIcaPath(chainA, chainB *novatesting.TestChain) *novatesting.Path {
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

func SetupIcaPath(path *novatesting.Path, owner string) error {
	if err := RegisterInterchainAccount(path.EndpointA, owner); err != nil {
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

func RegisterInterchainAccount(e *novatesting.Endpoint, owner string) error {
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
func NewBaseRegisteredZone() *ibcstakingtypes.RegisteredZone {
	icaControllerPort := zoneId + "." + baseOwnerAcc.String()
	return &ibcstakingtypes.RegisteredZone{
		ZoneId: zoneId,
		IcaConnectionInfo: &ibcstakingtypes.IcaConnectionInfo{
			ConnectionId: icaConnection,
			PortId:       icaControllerPort,
		},
		IcaAccount: &ibcstakingtypes.IcaAccount{
			ControllerAddress: baseOwnerAcc.String(),
			HostAddress:       baseHostAcc.String(),
		},
		TransferInfo: &ibcstakingtypes.TransferConnectionInfo{
			PortId:    transferPort,
			ChannelId: transferChannel,
		},
		ValidatorAddress: "",
		BaseDenom:        zoneBaseDenom,
		SnDenom:          baseSnDenom,
	}
}
