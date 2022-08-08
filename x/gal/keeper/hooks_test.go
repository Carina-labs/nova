package keeper_test

// import (
// 	"fmt"
// 	"strconv"

// 	oracletypes "github.com/Carina-labs/nova/x/oracle/types"

// 	"github.com/Carina-labs/nova/x/gal/keeper"
// 	"github.com/Carina-labs/nova/x/gal/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
// )

// func (suite *KeeperTestSuite) TestHookAfterTransferEnd() {
// 	var (
// 		sender     sdk.AccAddress
// 		receiver   sdk.AccAddress
// 		sentAmount int64
// 		baseDenom  string
// 	)

// 	testCases := []struct {
// 		name     string
// 		before   func()
// 		malleate func()
// 		exp      func()
// 	}{
// 		{
// 			"success",
// 			func() {
// 				sender = suite.GenRandomAddress()
// 				receiver = baseHostAcc
// 				sentAmount = 5000
// 				baseDenom = "stake"

// 				suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &banktypes.GenesisState{
// 					Balances: []banktypes.Balance{
// 						{sender.String(), sdk.Coins{sdk.NewInt64Coin("stake", sentAmount)}},
// 					},
// 				})
// 			},
// 			func() {},
// 			func() {
// 				zoneInfo := suite.chainA.App.IbcstakingKeeper.GetZoneForDenom(suite.chainA.GetContext(), baseDenom)
// 				record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), zoneInfo.ZoneId, sender)
// 				suite.Require().NoError(err)
// 				suite.Require().NotNil(record, "record doesn't exists")
// 			},
// 		},
// 		{
// 			"hooks should not do anything",
// 			func() {
// 				sender = suite.GenRandomAddress()

// 				// receiver is an arbitrary address
// 				receiver = suite.GenRandomAddress()
// 				sentAmount = 5000

// 				suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &banktypes.GenesisState{
// 					Balances: []banktypes.Balance{
// 						{sender.String(), sdk.Coins{sdk.NewInt64Coin("stake", sentAmount)}},
// 					},
// 				})
// 			},
// 			func() {},
// 			func() {},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc

// 		suite.Run(tc.name, func() {
// 			suite.SetupTest()

// 			hooks := suite.chainA.App.GalKeeper.Hooks()
// 			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)

// 			tc.before()

// 			// register zone
// 			operator := suite.GenRandomAddress()
// 			zone := newBaseRegisteredZone()
// 			suite.chainA.App.IbcstakingKeeper.RegisterZone(suite.chainA.GetContext(), zone)
// 			suite.chainA.App.OracleKeeper.InitGenesis(suite.chainA.GetContext(), &oracletypes.GenesisState{
// 				Params: oracletypes.Params{
// 					OracleOperators: []string{operator.String()},
// 				},
// 				States: []oracletypes.ChainInfo{
// 					{
// 						ChainId:         zone.ZoneId,
// 						OperatorAddress: operator.String(),
// 						LastBlockHeight: 1000,
// 						Coin:            sdk.NewInt64Coin("stake", 1000000),
// 					},
// 				},
// 			})

// 			// ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, "stake")
// 			fmt.Println("ibcdenom", hostIbcDenom)
// 			// should send deposit message to msg server
// 			_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.chainA.GetContext()), &types.MsgDeposit{
// 				Depositor:         sender.String(),
// 				Claimer:           sender.String(),
// 				Amount:            sdk.NewInt64Coin(hostIbcDenom, sentAmount),
// 				TransferPortId:    transferPort,
// 				TransferChannelId: transferChannel,
// 				ZoneId:            hostId,
// 			})
// 			fmt.Println("err", err)
// 			suite.Require().NoError(err)

// 			tc.malleate()

// 			// after send deposit msg to msg_server hooks should execute.
// 			packet := ibctransfertypes.FungibleTokenPacketData{
// 				Denom:    hostIbcDenom,
// 				Amount:   strconv.Itoa(int(sentAmount)),
// 				Sender:   sender.String(),
// 				Receiver: receiver.String(),
// 			}
// 			hooks.AfterTransferEnd(suite.chainA.GetContext(), packet, "stake")

// 			tc.exp()
// 		})
// 	}
// }
