package keeper_test

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (suite *KeeperTestSuite) TestDepositMsg() {
	type expectedSet struct {
		userBalance int64
		hostBalance int64
		snMinting   int64
	}
	tcs := []struct {
		name        string
		msg         types.MsgDeposit
		userBalance int64
		wantErr     bool
		expect      expectedSet
	}{
		{
			name: "valid test case 1",
			msg: types.MsgDeposit{
				Depositor: "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
				ZoneId:    hostId,
				HostAddr:  baseHostAcc.String(),
				Amount:    sdk.NewCoins(sdk.NewInt64Coin(hostIbcDenom, 1000)),
			},
			userBalance: 10000,
			wantErr:     false,
			expect: expectedSet{
				userBalance: 9000,
				hostBalance: 1000,
				snMinting:   100,
			},
		},
		//{
		//	name: "valid test case 2",
		//	msg: types.MsgDeposit{
		//		Depositor: "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
		//		ZoneId:    baseDenom,
		//		HostAddr:  baseHostAcc.String(),
		//		Amount:    sdk.NewCoins(sdk.NewInt64Coin(baseIbcDenom, 5000)),
		//	},
		//	userBalance: 10000,
		//	wantErr:     false,
		//	expect: expectedSet{
		//		userBalance: 5000,
		//		hostBalance: 5000,
		//		snMinting:   500,
		//	},
		//},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// setup validator for chainB
			validator := suite.chainB.App.StakingKeeper.GetValidators(suite.chainB.GetContext(), 1)[0]

			// setup ibc zone
			registerMsg := newBaseRegisteredZone()
			registerMsg.ValidatorAddress = validator.OperatorAddress
			trace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom(transferPort, transferChannel, hostBaseDenom))
			suite.chainA.App.IntertxKeeper.RegisterZone(suite.chainA.GetContext(), registerMsg)
			suite.chainA.App.TransferKeeper.InitGenesis(suite.chainA.GetContext(), transfertypes.GenesisState{
				PortId: transferPort,
				DenomTraces: []transfertypes.DenomTrace{
					trace,
				},
				Params: transfertypes.Params{
					SendEnabled:    true,
					ReceiveEnabled: true,
				},
			})
			suite.chainB.App.TransferKeeper.InitGenesis(suite.chainB.GetContext(), transfertypes.GenesisState{
				PortId:      transferPort,
				DenomTraces: []transfertypes.DenomTrace{},
				Params: transfertypes.Params{
					SendEnabled:    true,
					ReceiveEnabled: true,
				},
			})
			suite.chainB.App.ICAHostKeeper.SetParams(suite.chainB.GetContext(), icahosttypes.Params{})
			// TODO : should implement to call InitGenesis() in TransferKeeper of chainB

			// setup initial accounts for chainA, chainB
			suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &banktypes.GenesisState{
				Balances: []banktypes.Balance{
					{
						Address: tc.msg.Depositor,
						Coins: sdk.Coins{
							sdk.NewInt64Coin(hostIbcDenom, tc.userBalance),
						},
					},
				},
			})
			// initialize chainB's escrow address
			escrowAddr := transfertypes.GetEscrowAddress(transferPort, transferChannel)
			//suite.chainB.App.BankKeeper.InitGenesis(suite.chainB.GetContext(), &banktypes.GenesisState{
			//	Balances: []banktypes.Balance{
			//		{
			//			Address: tc.msg.HostAddr,
			//			Coins: sdk.Coins{
			//				sdk.NewInt64Coin(hostBaseDenom, 0),
			//			},
			//		},
			//		{
			//			Address: escrowAddr.String(),
			//			Coins: sdk.Coins{
			//				sdk.NewInt64Coin(hostBaseDenom, 10000000),
			//			},
			//		},
			//	},
			//})

			err := suite.chainB.App.BankKeeper.MintCoins(suite.chainB.GetContext(), "gal", sdk.Coins{sdk.NewInt64Coin(hostBaseDenom, 987654321)})
			suite.Require().NoError(err)
			err = suite.chainB.App.BankKeeper.SendCoinsFromModuleToAccount(suite.chainB.GetContext(), "gal", escrowAddr, sdk.Coins{sdk.NewInt64Coin(hostBaseDenom, 987654321)})
			suite.Require().NoError(err)

			// setup oracle
			suite.chainA.App.OracleKeeper.InitGenesis(suite.chainA.GetContext(), &oracletypes.GenesisState{
				Params: oracletypes.Params{
					OracleOperators: []string{
						baseHostAcc.String(),
					},
				},
				States: []oracletypes.ChainInfo{
					{
						Coin:            sdk.NewInt64Coin(hostBaseDenom, 100000_000000),
						Decimal:         6,
						OperatorAddress: baseHostAcc.String(),
						LastBlockHeight: 100,
						AppHash:         "",
						ChainId:         hostId,
						BlockProposer:   "",
					},
				},
			})

			userAcc, _ := sdk.AccAddressFromBech32(tc.msg.Depositor)
			hostAcc, _ := sdk.AccAddressFromBech32(tc.msg.HostAddr)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			executedCtx := suite.chainA.GetContext()
			goCtx := sdk.WrapSDKContext(executedCtx)

			st, _ := suite.chainB.App.StakingKeeper.GetValidator(suite.chainB.GetContext(), validator.GetOperator())
			previousStakedAmt := st.GetBondedTokens()
			fmt.Printf("previousStakedAmt: %d\n", previousStakedAmt.Int64())

			// EXECUTE
			_, err = msgServer.Deposit(goCtx, &tc.msg)
			suite.Require().NoError(err)

			em := executedCtx.EventManager()
			//for _, e := range em.Events() {
			//	fmt.Printf("type: %s, ", e.Type)
			//	for _, a := range e.Attributes {
			//		fmt.Printf("key: %s, value: %s", a.Key, a.Value)
			//	}
			//	fmt.Print("\n")
			//}

			// verify record
			record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), userAcc)
			suite.Require().NoError(err)
			p, err := ibctesting.ParsePacketFromEvents(em.Events())
			suite.Require().NoError(err)
			suite.chainA.NextBlock()

			err = suite.transferPath.RelayPacket(p)
			suite.Require().NoError(err)
			suite.chainB.NextBlock()

			fmt.Printf("hostIbcDenom: %s\n", hostIbcDenom)

			afterUserBalance := suite.chainA.App.BankKeeper.GetBalance(
				suite.chainA.GetContext(),
				userAcc,
				hostIbcDenom,
			)
			hostBalance := suite.chainB.App.BankKeeper.GetBalance(
				suite.chainB.GetContext(),
				hostAcc,
				hostBaseDenom,
			)
			fmt.Printf("hostBalance: %s\n", hostBalance.String())
			//
			//for _, e := range em.Events() {
			//	fmt.Printf("type: %s, ", e.Type)
			//	for _, a := range e.Attributes {
			//		fmt.Printf("[%s: %s] ", a.Key, a.Type)
			//	}
			//	fmt.Printf("\n")
			//}

			// verify transfer
			suite.Require().Equal(tc.expect.userBalance, afterUserBalance.Amount.Int64())
			suite.Require().Equal(tc.expect.hostBalance, hostBalance.Amount.Int64())

			// verify hook action
			record, err = suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), userAcc)
			suite.Require().NoError(err)

			// simulate delegation with bot
			bMsgServer := stakingkeeper.NewMsgServerImpl(*suite.chainB.App.StakingKeeper)
			_, e := bMsgServer.Delegate(sdk.WrapSDKContext(suite.chainB.GetContext()), &stakingtypes.MsgDelegate{
				DelegatorAddress: hostAcc.String(),
				ValidatorAddress: validator.OperatorAddress,
				Amount:           sdk.NewInt64Coin(hostBaseDenom, 1000),
			})
			suite.Require().NoError(e)

			// verify claim
			err = suite.chainA.App.BankKeeper.MintCoins(suite.chainA.GetContext(),
				types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(baseSnDenom, 10000_000000)))
			suite.Require().NoError(err)
			err = suite.chainA.App.GalKeeper.ClaimAndMintShareToken(suite.chainA.GetContext(), userAcc, *record.Records[0].Amount)
			snBalance := suite.chainA.App.BankKeeper.GetBalance(
				suite.chainA.GetContext(), userAcc, baseSnDenom)
			fmt.Printf("snBalance: %s\n", snBalance.String())
			suite.Require().True(sdk.NewInt64Coin(baseSnDenom, tc.expect.snMinting).IsEqual(snBalance))
			suite.Require().NoError(err)

			// Is staking correctly executed?
			st, _ = suite.chainB.App.StakingKeeper.GetValidator(suite.chainB.GetContext(), validator.GetOperator())
			suite.Require().Equal(previousStakedAmt.Int64()+tc.msg.Amount[0].Amount.Int64(), st.BondedTokens().Int64())
			// verify undelegate
			executedCtx = suite.chainA.GetContext()
			goCtx = sdk.WrapSDKContext(executedCtx)

			_, err = msgServer.UndelegateRecord(goCtx, &types.MsgUndelegateRecord{
				ZoneId:    hostId,
				Depositor: userAcc.String(),
				Amount:    snBalance,
			})

			executedCtx = suite.chainA.GetContext()
			goCtx = sdk.WrapSDKContext(executedCtx)
			_, err = msgServer.Undelegate(goCtx, &types.MsgUndelegate{
				ZoneId:            hostId,
				ControllerAddress: baseAcc.String(),
				HostAddress:       baseHostAcc.String(),
			})
			suite.Require().NoError(err)

			// relay ica packet
			em = executedCtx.EventManager()
			p, err = ibctesting.ParsePacketFromEvents(em.Events())
			fmt.Printf("packet: %s\n", p.String())
			suite.Require().NoError(err)
			suite.chainA.NextBlock()

			err = suite.icaPath.RelayPacket(p)

			suite.chainB.NextBlock()

			rs := suite.chainB.App.StakingKeeper.GetUnbondingDelegationsFromValidator(suite.chainB.GetContext(), validator.GetOperator())
			for _, r := range rs {
				fmt.Printf("result: %s\n", r.String())
			}
			//em = executedCtx.EventManager()
			//
			//suite.chainA.NextBlock()

			// verify user withdraw amount
			//executedCtx = suite.chainA.GetContext()
			//goCtx = sdk.WrapSDKContext(executedCtx)
			//_, err = msgServer.WithdrawRecord(goCtx, &types.MsgWithdrawRecord{
			//	ZoneId:     baseDenom,
			//	Withdrawer: userAcc.String(),
			//})
		})
	}
}

func (suite *KeeperTestSuite) TestWithdrawMsg() {

}

func ParseAddressToIbcAddress(destPort string, destChannel string, denom string) string {
	sourcePrefix := transfertypes.GetDenomPrefix(destPort, destChannel)
	prefixedDenom := sourcePrefix + denom
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	voucherDenom := denomTrace.IBCDenom()
	return voucherDenom
}
