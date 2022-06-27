package keeper_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))

	suite.transferPath = newIbcTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.transferPath)

	suite.icaPath = newIcaPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.icaPath)
	suite.icaOwnerAddr = apptesting.CreateRandomAccounts(1)[0]
	err := setupIcaPath(suite.icaPath, suite.icaOwnerAddr.String())
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) SetIbcZone(zoneMsg []intertxtypes.RegisteredZone) {
	for _, msg := range zoneMsg {
		suite.App.IntertxKeeper.RegisterZone(suite.Ctx, &msg)
	}
}

func (suite *KeeperTestSuite) SetOracle(operators []sdk.Address, msg []oracletypes.ChainInfo) {
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

func newIbcTransferPath(chainA, chainB *novatesting.TestChain) *novatesting.Path {
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
		ControllerConnectionId: "connection-1",
		HostConnectionId:       "connection-1",
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
