package keeper_test

import (
	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/types"
	types2 "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
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
}

func (suite *KeeperTestSuite) SetupTestOracle(msgs []*types2.MsgUpdateChainState) {
	for _, msg := range msgs {
		err := suite.App.OracleKeeper.UpdateChainState(suite.Ctx, msg)
		if err != nil {
			panic(err)
		}
	}
}

func (suite *KeeperTestSuite) setRandomState() {
	for _, acc := range suite.TestAccs {
		err := suite.App.GalKeeper.CacheDepositAmt(suite.Ctx, acc, sdk.NewInt64Coin("uosmo", 2))
		if err != nil {
			panic(err)
		}
	}
}

func (suite *KeeperTestSuite) TestGetShares() {
	tcs := []struct {
		expected int64
	}{
		{
			expected: 2,
		},
		{
			expected: 2,
		},
		{
			expected: 2,
		},
		{
			expected: 2,
		},
		{
			expected: 2,
		},
	}

	for i, tc := range tcs {
		shares, err := suite.App.GalKeeper.GetCachedDepositAmt(suite.Ctx, suite.TestAccs[i])
		suite.NoError(err)
		suite.Equal(tc.expected, shares.Amount)
	}
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

	sendAmountCoin := depositorBalance //deposit all aphoton balance of depositor
	errDeposit := appChainA.GalKeeper.DepositCoin(ctxA, accountAddrA, accountAddrB, sourcePort, sourceChannel, sendAmountCoin)
	require.NoError(suite.T(), errDeposit)

	suite.chainA.NextBlock()                                                            //must execute
	packet, errPacket := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events()) //parse packet
	require.NoError(suite.T(), errPacket)

	packetErr := path.RelayPacket(packet)
	suite.Require().NoError(packetErr) // relay committed

	voucherDenomTrace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom(sourcePort, sourceChannel, testDenom))
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
