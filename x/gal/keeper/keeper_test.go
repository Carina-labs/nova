package keeper_test

import (
	"fmt"
	types2 "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	types3 "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient

	coordinator *novatesting.Coordinator
	chainA      *novatesting.TestChain
	chainB      *novatesting.TestChain
	path        *novatesting.Path
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.setRandomState()

	//Coordinator is a testing struct which contains N TestChain's
	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))

	//setup path (chainA <===>chainB)
	path := novatesting.NewPath(suite.chainA, suite.chainB)
	path.EndpointA.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = novatesting.TransferPort
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	suite.coordinator.Setup(path)
	suite.path = path

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

func (suite *KeeperTestSuite) setRandomState() {
	//for _, acc := range suite.TestAccs {
	//	err := suite.App.GalKeeper.RecordDepositAmt(suite.Ctx, sdk.NewInt64Coin("uosmo", 2))
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

func (suite *KeeperTestSuite) TestDepositCoins() {
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

	key1 := secp256k1.GenPrivKey()
	acc1 := authtypes.NewBaseAccount(key1.PubKey().Address().Bytes(), key1.PubKey(), 0, 0)

	key2 := secp256k1.GenPrivKey()
	acc2 := authtypes.NewBaseAccount(key2.PubKey().Address().Bytes(), key2.PubKey(), 0, 0)

	key3 := secp256k1.GenPrivKey()
	acc3 := authtypes.NewBaseAccount(key3.PubKey().Address().Bytes(), key3.PubKey(), 0, 0)

	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc1)
	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc2)
	suite.chainA.App.AccountKeeper.SetAccount(ctxA, acc3)

	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc1)
	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc2)
	suite.chainB.App.AccountKeeper.SetAccount(ctxB, acc3)

	suite.chainA.App.BankKeeper.InitGenesis(ctxA,
		&types3.GenesisState{
			Supply: sdk.Coins{sdk.NewInt64Coin("nova", 30000000)},
			Balances: []types3.Balance{
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

	suite.chainA.App.IntertxKeeper.SetRegesterZone(ctxA, types2.RegisteredZone{
		ZoneName: suite.chainB.ChainID,
		IcaConnectionInfo: &types2.IcaConnectionInfo{
			ConnectionId: "connection-id",
			OwnerAddress: acc2.Address,
		},
		TransferConnectionInfo: &types2.TransferConnectionInfo{
			ConnectionId: "connection-id",
			PortId:       novatesting.TransferPort,
			ChannelId:    "channel-0",
		},
		ValidatorAddress: acc1.Address,
		BaseDenom:        "osmo",
		SnDenom:          "snOsmo",
		StDenom:          "stOsmo",
	})

	senderPrivAddr, _ := sdk.AccAddressFromHex(suite.chainA.SenderPrivKey.PubKey().Address().String())
	fmt.Printf("senderPrivAddr : %s\n", senderPrivAddr)

	err := suite.chainA.App.GalKeeper.DepositCoin(ctxA, types.MsgDeposit{
		Depositor: acc1.GetAddress().String(),
		ZoneId:    suite.chainB.ChainID,
		Amount:    sdk.NewCoins(sdk.NewInt64Coin("nova", 1000)),
	})

	if err != nil {
		fmt.Printf("err : %s", err.Error())
	}

	suite.chainA.NextBlock()

	packet, errPacket := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events()) //parse packet
	require.NoError(suite.T(), errPacket)

	packetErr := suite.path.RelayPacket(packet)
	suite.Require().NoError(packetErr) // relay committed

	record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(ctxA, acc1.GetAddress())
	suite.NoError(err)
	fmt.Printf(record.String())

	suite.chainB.App.BankKeeper.IterateAllBalances(ctxB, func(address sdk.AccAddress, coin sdk.Coin) bool {
		fmt.Printf("IterateAllBalances addr: %s, balance: %s\n", address.String(), coin.String())
		return false
	})
}
