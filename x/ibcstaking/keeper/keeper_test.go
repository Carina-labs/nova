package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/stretchr/testify/suite"

	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
)

var (
	key1 = secp256k1.GenPrivKey()
	acc1 = authtypes.NewBaseAccount(key1.PubKey().Address().Bytes(), key1.PubKey(), 0, 0)

	version = string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                "ics27-1",
		ControllerConnectionId: "connection-1",
		HostConnectionId:       "connection-1",
		Encoding:               "proto3",
		TxType:                 "sdk_multi_msg",
	}))
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	coordinator *novatesting.Coordinator
	chainA      *novatesting.TestChain
	chainB      *novatesting.TestChain
	ctxA        sdk.Context
	ctxB        sdk.Context
	path        *novatesting.Path
	icaPath     *novatesting.Path
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}

func (suite *KeeperTestSuite) setZone(num int) []ibcstakingtypes.RegisteredZone {

	addr := make([]sdk.AccAddress, 0)
	zones := make([]ibcstakingtypes.RegisteredZone, 0)

	for i := 0; i < num; i++ {
		addr = append(addr, suite.GenRandomAddress())
		zones = append(zones, ibcstakingtypes.RegisteredZone{
			ZoneId: "gaia" + strconv.Itoa(i),
			IcaConnectionInfo: &ibcstakingtypes.IcaConnectionInfo{
				ConnectionId: "connection-" + strconv.Itoa(i),
				PortId:       "icacontroller-" + addr[i].String(),
			},
			IcaAccount: &ibcstakingtypes.IcaAccount{
				DaomodifierAddress: addr[i].String(),
				HostAddress:        addr[i].String(),
			},
			ValidatorAddress: sdk.ValAddress(addr[i]).String(),
			BaseDenom:        "atom",
			SnDenom:          "snatom",
		})
	}

	return zones
}

func NewICAPAth(chainA, chainB *novatesting.TestChain) *novatesting.Path {
	path := novatesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = icatypes.PortID
	path.EndpointB.ChannelConfig.PortID = icatypes.PortID
	path.EndpointA.ChannelConfig.Order = ibcchanneltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = ibcchanneltypes.ORDERED
	path.EndpointA.ChannelConfig.Version = version
	path.EndpointB.ChannelConfig.Version = version
	return path
}

// SetupICAPath invokes the InterchainAccounts entrypoint and subsequent channel handshake handlers
func (suite *KeeperTestSuite) SetupICAPath(path *novatesting.Path, owner string) error {
	if err := suite.RegisterInterchainAccount(path.EndpointA, owner); err != nil {
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

// RegisterInterchainAccount is a helper function for starting the channel handshake
func (suite *KeeperTestSuite) RegisterInterchainAccount(endpoint *novatesting.Endpoint, owner string) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelSequence := endpoint.Chain.App.IBCKeeper.ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := endpoint.Chain.App.ICAControllerKeeper.RegisterInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, owner); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = ibcchanneltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID
	return nil
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	//Coordinator is a testing struct which contains N TestChain's
	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))
	suite.ctxA = suite.chainA.GetContext()
	suite.ctxB = suite.chainB.GetContext()

	//setup path (chainA <===>chainB)
	path := novatesting.NewPath(suite.chainA, suite.chainB)
	path.EndpointA.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	suite.coordinator.Setup(path)
	suite.path = path

	suite.icaPath = NewICAPAth(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.icaPath)
	err := suite.SetupICAPath(suite.icaPath, acc1.Address)
	if err != nil {
		fmt.Printf("err ica path : %s\n", err.Error())
	}
	suite.NoError(err)
	println("Finish setup test")
}

func (suite *KeeperTestSuite) SetupTestIBCZone(zoneMsgs []ibcstakingtypes.RegisteredZone) {
	for _, msg := range zoneMsgs {
		suite.App.IbcstakingKeeper.RegisterZone(suite.Ctx, &msg)
	}
}

func (suite *KeeperTestSuite) TestIsValidZoneRegisterAddress() {
	var addresses []string

	addr1 := suite.GenRandomAddress().String()
	addr2 := suite.GenRandomAddress().String()
	addr3 := suite.GenRandomAddress().String()

	addresses = append(addresses, addr1)
	addresses = append(addresses, addr2)

	params := ibcstakingtypes.Params{
		DaoModifiers: addresses,
	}

	suite.App.IbcstakingKeeper.SetParams(suite.Ctx, params)

	tcs := []struct {
		name   string
		addr   string
		expect bool
	}{
		{
			name:   "success",
			addr:   addr1,
			expect: true,
		},
		{
			name:   "success",
			addr:   addr2,
			expect: true,
		},
		{
			name:   "nil address",
			addr:   "",
			expect: false,
		},
		{
			name:   "random address",
			addr:   addr3,
			expect: false,
		},
		{
			name:   "invalid address",
			addr:   "test",
			expect: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			ok := suite.App.IbcstakingKeeper.IsValidDaoModifier(suite.Ctx, tc.addr)

			suite.Require().Equal(ok, tc.expect)
		})
	}
}