package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"strconv"
)

func (suite *KeeperTestSuite) TestHookAfterTransferEnd() {
	var (
		sender     sdk.AccAddress
		receiver   sdk.AccAddress
		sentAmount int64
	)

	testCases := []struct {
		name     string
		before   func()
		malleate func()
		exp      func()
	}{
		{
			"success",
			func() {
				sender = suite.GenRandomAddress()
				receiver = baseHostAcc
				sentAmount = 5000

				suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &banktypes.GenesisState{
					Balances: []banktypes.Balance{
						{sender.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, sentAmount)}},
					},
				})
			},
			func() {
				//record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), sender)
				//suite.Require().NoError(err)
				// suite.Require().Equal(record.IsTransferred, false)
			},
			func() {
				record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), sender)
				suite.Require().NoError(err)
				suite.Require().NotNil(record, "record doesn't exists")
				//suite.Require().Equal(sentAmount, record.Amount.Amount.Int64())
				//suite.Require().Equal(sender.String(), record.Address)
				//suite.Require().Equal(record.IsTransferred, true)
			},
		},
		{
			"hooks should not do anything",
			func() {
				sender = suite.GenRandomAddress()

				// receiver is an arbitrary address
				receiver = suite.GenRandomAddress()
				sentAmount = 5000

				suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &banktypes.GenesisState{
					Balances: []banktypes.Balance{
						{sender.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, sentAmount)}},
					},
				})
			},
			func() {
				//record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), sender)
				//suite.Require().NoError(err)
				//suite.Require().Equal(record.IsTransferred, false)
			},
			func() {
				//record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), sender)
				//suite.Require().NoError(err)
				//suite.Require().NotNil(record, "record doesn't exists")
				//suite.Require().Equal(sentAmount, record.Amount.Amount.Int64())
				//suite.Require().Equal(sender.String(), record.Address)
				//suite.Require().Equal(record.IsTransferred, false)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest()

			hooks := suite.chainA.App.GalKeeper.Hooks()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)

			tc.before()

			// register zone
			suite.chainA.App.IntertxKeeper.RegisterZone(suite.chainA.GetContext(), newBaseRegisteredZone())

			// should send deposit message to msg server
			_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.chainA.GetContext()), &types.MsgDeposit{
				Depositor: sender.String(),
				Amount:    sdk.Coins{sdk.NewInt64Coin(baseDenom, sentAmount)},
				HostAddr:  baseHostAcc.String(),
				ZoneId:    hostId,
			})
			suite.Require().NoError(err)

			tc.malleate()

			// after send deposit msg to msg_server hooks should execute.
			packet := ibctransfertypes.FungibleTokenPacketData{
				Denom:    baseDenom,
				Amount:   strconv.Itoa(int(sentAmount)),
				Sender:   sender.String(),
				Receiver: receiver.String(),
			}
			hooks.AfterTransferEnd(suite.chainA.GetContext(), packet, baseDenom)

			tc.exp()
		})
	}
}
