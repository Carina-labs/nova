package keeper_test

import (
	novatesting "github.com/Carina-labs/novachain/testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *novatesting.Coordinator
	chainA      *novatesting.TestChain
	chainB      *novatesting.TestChain
	path        *novatesting.Path
}

func (suite *KeeperTestSuite) SetupTest() {
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
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestDepositCoins() {
	appChainA, ctxA := suite.chainA.App, suite.chainA.GetContext() //source channel
	appChainB, ctxB := suite.chainB.App, suite.chainB.GetContext() //dest channel
	path := suite.path

	accountAddrA := suite.chainA.SenderAccount.GetAddress() //depositor account
	accountAddrB := suite.chainB.SenderAccount.GetAddress() //receiver account

	//get path info
	sourcePort := path.EndpointA.ChannelConfig.PortID
	sourceChannel := path.EndpointA.ChannelID

	testDenom := "aphoton"

	depositorBalance := appChainA.BankKeeper.GetBalance(ctxA, accountAddrA, testDenom)
	receiverBalance := appChainB.BankKeeper.GetBalance(ctxB, accountAddrB, testDenom)

	println("================== test deposit ==================")
	println("[before deposit]")
	println("sender : ", depositorBalance.String())
	println("receiver : ", receiverBalance.String())

	sendAmountCoins := sdk.NewCoins(depositorBalance) //deposit all aphoton balance of depositor
	errDeposit := appChainA.GalKeeper.DepositCoin(ctxA, accountAddrA.String(), accountAddrB.String(), sourcePort, sourceChannel, sendAmountCoins)
	require.NoError(suite.T(), errDeposit)

	suite.chainA.NextBlock()                                                            //must execute
	packet, errPacket := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events()) //parse packet
	require.NoError(suite.T(), errPacket)

	packetErr := path.RelayPacket(packet)
	suite.Require().NoError(packetErr) // relay committed

	voucherDenomTrace := types.ParseDenomTrace(types.GetPrefixedDenom(sourcePort, sourceChannel, testDenom))
	IBCTestDemnom := voucherDenomTrace.IBCDenom()

	depositorBalanceAfterDeposit := appChainA.BankKeeper.GetBalance(ctxA, accountAddrA, testDenom)
	depositedBalance := appChainB.BankKeeper.GetBalance(ctxB, accountAddrB, IBCTestDemnom)

	suite.Require().Equal(depositedBalance.Amount, depositorBalance.Amount)
	suite.Require().Zero(depositorBalanceAfterDeposit.Amount.Int64())

	println("send", depositorBalance.String())
	println(testDenom, "==========>", IBCTestDemnom)
	println("[after deposit]")
	println("sender : ", depositorBalanceAfterDeposit.String())
	println("receiver : ", depositedBalance.String())
}
