package keeper_test

import (
	"fmt"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
)

var (
	key1 = secp256k1.GenPrivKey()
	acc1 = authtypes.NewBaseAccount(key1.PubKey().Address().Bytes(), key1.PubKey(), 0, 0)

	key2 = secp256k1.GenPrivKey()
	acc2 = authtypes.NewBaseAccount(key2.PubKey().Address().Bytes(), key2.PubKey(), 0, 0)

	key3    = secp256k1.GenPrivKey()
	acc3    = authtypes.NewBaseAccount(key3.PubKey().Address().Bytes(), key3.PubKey(), 0, 0)
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
	queryClient types.QueryClient

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
	suite.setRandomState()

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
	err := suite.SetupICAPath(suite.icaPath, acc2.Address)
	if err != nil {
		fmt.Printf("err ica path : %s\n", err.Error())
	}
	suite.NoError(err)
	println("Finish setup test")
}

func (suite *KeeperTestSuite) SetupTestOracle(
	operators []sdk.AccAddress,
	msgs []*oracletypes.ChainInfo) {
	for _, operator := range operators {
		suite.App.OracleKeeper.SetParams(suite.Ctx, oracletypes.Params{
			OracleOperators: []string{operator.String()},
		})
	}

	for _, msg := range msgs {
		err := suite.App.OracleKeeper.UpdateChainState(suite.Ctx, msg)
		if err != nil {
			panic(err)
		}
	}
}

func (suite *KeeperTestSuite) SetupTestIBCZone(zoneMsgs []intertxtypes.RegisteredZone) {
	for _, msg := range zoneMsgs {
		suite.App.IntertxKeeper.RegisterZone(suite.Ctx, &msg)
	}
}

func (suite *KeeperTestSuite) setRandomState() {
	//for _, acc := range suite.TestAccs {
	//	err := suite.App.GalKeeper.RecordDepositAmt(suite.Ctx, sdk.NewInt64Coin("uosmo", 2))
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

func (suite *KeeperTestSuite) TestSimulateDepositCoins() {
	// suite.SetupTest()
	operatorAddrs := []sdk.AccAddress{
		suite.GenRandomAddress(),
	}
	suite.SetupTestOracle(operatorAddrs, []*oracletypes.ChainInfo{
		{
			Coin:            sdk.NewInt64Coin("atom", 10000),
			OperatorAddress: operatorAddrs[0].String(),
			LastBlockHeight: 10000,
			Decimal:         6,
		},
		{
			Coin:            sdk.NewInt64Coin("osmo", 20000),
			OperatorAddress: operatorAddrs[0].String(),
			LastBlockHeight: 20000,
			Decimal:         6,
		},
	})

	ctxA, ctxB := suite.chainA.GetContext(), suite.chainB.GetContext()

	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc1)
	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc2)
	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc3)

	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc1)
	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc2)
	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc3)

	suite.chainA.App.BankKeeper.InitGenesis(ctxA,
		&banktypes.GenesisState{
			Supply: sdk.Coins{sdk.NewInt64Coin("nova", 30000000)},
			Balances: []banktypes.Balance{
				{
					Address: acc1.GetAddress().String(),
					Coins:   sdk.Coins{sdk.NewInt64Coin("nova", 10000000)},
				},
				{
					Address: acc2.GetAddress().String(),
					Coins:   sdk.Coins{sdk.NewInt64Coin("nova", 10000000)},
				},
				{
					Address: acc3.GetAddress().String(),
					Coins:   sdk.Coins{sdk.NewInt64Coin("nova", 10000000)},
				},
			},
		})
	ownerAddr := acc2

	suite.chainA.App.IntertxKeeper.RegisterZone(ctxA, &intertxtypes.RegisteredZone{
		ZoneId: suite.chainB.ChainID,
		IcaConnectionInfo: &intertxtypes.IcaConnectionInfo{
			ConnectionId: suite.icaPath.EndpointB.ConnectionID,
		},
		TransferConnectionInfo: &intertxtypes.TransferConnectionInfo{
			ConnectionId: suite.path.EndpointA.ConnectionID,
			PortId:       novatesting.TransferPort,
			ChannelId:    "channel-0",
		},
		IcaAccount: &intertxtypes.IcaAccount{
			OwnerAddress: ownerAddr.Address,
			HostAddress:  "",
			Balance:      sdk.Coin{},
		},
		ValidatorAddress: acc1.Address,
		BaseDenom:        "osmo",
		SnDenom:          "",
		StDenom:          "",
	})

	_, err := sdk.AccAddressFromHex(suite.chainA.SenderPrivKey.PubKey().Address().String())
	suite.Require().NoError(err)

	err = suite.chainA.App.GalKeeper.Deposit(ctxA, &types.MsgDeposit{
		Depositor: acc1.GetAddress().String(),
		ZoneId:    suite.chainB.ChainID,
		Amount:    sdk.NewCoins(sdk.NewInt64Coin("nova", 1000)),
	})

	suite.Require().NoError(err)
	suite.chainA.NextBlock()

	packet, errPacket := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events()) //parse packet
	require.NoError(suite.T(), errPacket)

	packetErr := suite.path.RelayPacket(packet)
	suite.Require().NoError(packetErr) // relay committed

	record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(ctxA, acc1.GetAddress())
	suite.NoError(err)
	fmt.Printf("record : %s\n", record.String())

	suite.chainB.App.BankKeeper.IterateAllBalances(ctxB, func(address sdk.AccAddress, coin sdk.Coin) bool {
		fmt.Printf("IterateAllBalances addr: %s, balance: %s\n", address.String(), coin.String())
		return false
	})

	//hostAddr := icatypes.GenerateAddress(suite.chainB.App.AccountKeeper.GetModuleAddress(icatypes.ModuleName), hostConnectionID, novatesting.TransferPort)
	//fmt.Printf("host address : %s", hostAddr.String())

	fmt.Printf("EndpointA connectionID : %s\n", suite.path.EndpointA.ConnectionID)
	fmt.Printf("EndpointB connectionID : %s\n", suite.path.EndpointB.ConnectionID)

	channels := suite.chainA.App.IBCKeeper.ChannelKeeper.GetAllChannels(ctxA)
	for _, channel := range channels {
		fmt.Printf("chain A channel : %s\n", channel.String())
	}

	channels = suite.chainB.App.IBCKeeper.ChannelKeeper.GetAllChannels(ctxB)
	for _, channel := range channels {
		fmt.Printf("chain B channel : %s\n", channel.String())
	}

	fmt.Printf("ownerAddr : %s\n", ownerAddr.Address)

	icaAddr, _ := suite.chainB.App.ICAControllerKeeper.GetInterchainAccountAddress(ctxB, suite.path.EndpointB.ConnectionID, ownerAddr.Address)
	fmt.Printf("icaAddr : %s\n", icaAddr)
}
