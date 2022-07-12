package keeper_test

import (
	"fmt"
	"time"

	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingkeeper "github.com/Carina-labs/nova/x/ibcstaking/keeper"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

type oracleSet struct {
	operators []string
	state     sdk.Coin
	decimal   uint64
}

type icaConfig struct {
	icaHostAddress       string
	icaControllerAddress string
}

type initialSet struct {
	userAddress    string
	userBalance    int64
	nativeDenom    string
	depositAmount  int64
	snTokenBalance int64
	oracle         oracleSet
	withdrawAmount int64
}

type expectedSet struct {
	userBalance              int64
	hostBalance              int64
	snMinting                int64
	validatorBalance         int64
	withdrawAmount           int64
	afterWithdrawUserBalance int64
}

func (suite *KeeperTestSuite) prepare(initSet initialSet) {
	// prepare chainA balance
	ibcDenom := ParseAddressToIbcAddress(transferPort, transferChannel, hostBaseDenom)
	err := setFunds(suite.chainA, initSet.userAddress, sdk.NewInt64Coin(ibcDenom, initSet.userBalance))
	suite.Require().NoError(err)

	// prepare chainA's snToken state
	err = suite.chainA.App.BankKeeper.MintCoins(suite.chainA.GetContext(),
		types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(baseSnDenom, initSet.snTokenBalance)))

	// prepare chainB escrow
	escrowAddress := transfertypes.GetEscrowAddress(transferPort, transferChannel)
	err = setFunds(suite.chainB, escrowAddress.String(), sdk.NewInt64Coin(initSet.nativeDenom, 9999999999))

	// prepare transfer keeper
	trace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom(transferPort, transferChannel, initSet.nativeDenom))
	suite.chainA.App.TransferKeeper.InitGenesis(suite.chainA.GetContext(), transfertypes.GenesisState{
		PortId:      transferPort,
		DenomTraces: []transfertypes.DenomTrace{trace},
		Params:      transfertypes.Params{SendEnabled: true, ReceiveEnabled: true},
	})

	// prepare icaHostKeeper
	suite.chainB.App.ICAHostKeeper.SetParams(suite.chainB.GetContext(), icahosttypes.Params{
		HostEnabled:   true,
		AllowMessages: []string{undelegateMsgName, ibcTransferMsgName},
	})

	// prepare oracle
	suite.chainA.App.OracleKeeper.InitGenesis(suite.chainA.GetContext(), &oracletypes.GenesisState{
		Params: oracletypes.Params{
			OracleOperators: initSet.oracle.operators,
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            initSet.oracle.state,
				Decimal:         initSet.oracle.decimal,
				OperatorAddress: initSet.oracle.operators[0],
				LastBlockHeight: 100,
				AppHash:         "",
				ChainId:         hostId,
				BlockProposer:   "",
			},
		},
	})

	// prepare ibcstaking keeper
	suite.chainA.App.IbcstakingKeeper.SetParams(suite.chainA.GetContext(), ibcstakingtypes.Params{
		DaoModifiers: []string{
			suite.icaOwnerAddr.String(),
		},
	})
}

func (suite *KeeperTestSuite) TestGalAction() {
	tcs := []struct {
		name                    string
		initialUserBalance      int64
		initialValidatorBalance int64
		initSet                 initialSet
		wantErr                 bool
		expect                  expectedSet
	}{
		{
			name: "valid test case 1",
			initSet: initialSet{
				userBalance:    10000,
				userAddress:    "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
				nativeDenom:    "stake",
				depositAmount:  1000,
				snTokenBalance: 10000_000000,
				withdrawAmount: 1000,
				oracle: oracleSet{
					operators: []string{oracleOperatorAcc.String()},
					state:     sdk.NewInt64Coin(hostBaseDenom, 100000_000000),
				},
			},
			initialUserBalance: 10000,
			wantErr:            false,
			expect: expectedSet{
				userBalance:              9000,
				hostBalance:              1000,
				snMinting:                100,
				validatorBalance:         1001000,
				withdrawAmount:           1000,
				afterWithdrawUserBalance: 10000,
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			suite.prepare(tc.initSet)
			// setup
			validator := suite.chainB.App.StakingKeeper.GetValidators(suite.chainB.GetContext(), 1)[0]
			icaConf, err := setIbcZone(suite.chainA, suite.chainB, suite.icaOwnerAddr.String())
			suite.Require().NoError(err)
			userAcc, _ := sdk.AccAddressFromBech32(tc.initSet.userAddress)

			// execute
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			executedCtx := suite.chainA.GetContext()
			goCtx := sdk.WrapSDKContext(executedCtx)
			depositMsg := types.MsgDeposit{
				Depositor:         tc.initSet.userAddress,
				ZoneId:            hostId,
				HostAddr:          icaConf.icaHostAddress,
				TransferPortId:    transferPort,
				TransferChannelId: transferChannel,
				Amount:            sdk.NewInt64Coin(hostIbcDenom, 1000),
			}
			_, err = msgServer.Deposit(goCtx, &depositMsg)
			suite.Require().NoError(err)

			// relay packet
			em := executedCtx.EventManager()
			p, err := ibctesting.ParsePacketFromEvents(em.Events())
			suite.Require().NoError(err)
			suite.chainA.NextBlock()

			err = suite.transferPath.RelayPacket(p)
			suite.Require().NoError(err)
			suite.chainB.NextBlock()

			// Is IBC transfer correctly executed?
			suite.verifyTransferCorrectlyExecuted(tc.initSet, *icaConf, tc.expect)

			// simulate delegation with bot
			bMsgServer := stakingkeeper.NewMsgServerImpl(*suite.chainB.App.StakingKeeper)
			_, e := bMsgServer.Delegate(sdk.WrapSDKContext(suite.chainB.GetContext()), &stakingtypes.MsgDelegate{
				DelegatorAddress: icaConf.icaHostAddress,
				ValidatorAddress: validator.OperatorAddress,
				Amount:           sdk.NewInt64Coin(hostBaseDenom, 1000),
			})
			suite.Require().NoError(e)

			// claim share token
			record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), userAcc)
			_, err = suite.chainA.App.GalKeeper.ClaimAndMintShareToken(suite.chainA.GetContext(), userAcc, *record.Records[0].Amount)
			snBalance := suite.chainA.App.BankKeeper.GetBalance(
				suite.chainA.GetContext(), userAcc, baseSnDenom)
			suite.Require().True(sdk.NewInt64Coin(baseSnDenom, tc.expect.snMinting).IsEqual(snBalance))
			suite.Require().NoError(err)

			// Is staking correctly executed?
			suite.verifyStakingCorrectlyExecuted(validator.OperatorAddress, tc.expect.validatorBalance)

			// simulate undelegate
			_, err = msgServer.UndelegateRecord(sdk.WrapSDKContext(suite.chainA.GetContext()), &types.MsgUndelegateRecord{
				ZoneId:    hostId,
				Depositor: userAcc.String(),
				Amount:    snBalance,
			})

			// execute : undelegate
			executedCtx = suite.chainA.GetContext()
			goCtx = sdk.WrapSDKContext(executedCtx)
			_, err = msgServer.Undelegate(goCtx, &types.MsgUndelegate{
				ZoneId:            hostId,
				ControllerAddress: baseOwnerAcc.String(),
				HostAddress:       icaConf.icaHostAddress,
			})
			suite.Require().NoError(err)

			// relay ica packet
			em = executedCtx.EventManager()
			p, err = ibctesting.ParsePacketFromEvents(em.Events())
			suite.Require().NoError(suite.icaPath.EndpointA.UpdateClient())
			suite.Require().NoError(suite.icaPath.EndpointB.UpdateClient())
			res, err := suite.icaPath.EndpointB.RecvPacketWithResult(p)
			suite.Require().NoError(err)
			ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
			suite.Require().NoError(err)
			suite.Require().NotNil(ack)
			err = suite.icaPath.EndpointA.AcknowledgePacket(p, ack)
			suite.Require().NoError(err)

			// verify undelegate amount
			ubd := suite.chainB.App.StakingKeeper.GetUnbondingDelegationsFromValidator(suite.chainB.GetContext(), validator.GetOperator())
			suite.Require().Equal(icaConf.icaHostAddress, ubd[0].DelegatorAddress)

			// simulate unbond immediately, but assume there is no yield.
			icaHostAcc, _ := sdk.AccAddressFromBech32(icaConf.icaHostAddress)
			nextMonthCtx := suite.chainB.GetContext().WithBlockTime(time.Now().AddDate(0, 1, 0))
			unbondedCoins, err := suite.chainB.App.StakingKeeper.CompleteUnbonding(nextMonthCtx,
				icaHostAcc, validator.GetOperator())
			suite.Require().NoError(err)
			suite.Require().True(sdk.NewInt64Coin(hostBaseDenom, tc.initSet.depositAmount).IsEqual(unbondedCoins[0]))

			// verify withdraw record
			rc, err := suite.chainA.App.GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), hostId+tc.initSet.userAddress)
			suite.Require().NoError(err)
			suite.Require().NotNil(rc)
			suite.Require().Equal(tc.expect.withdrawAmount, rc.Amount.Amount.Int64())

			// simulation : bot requests transfer from host -> controller's gal module account.
			ibcStakingMsgServer := ibcstakingkeeper.NewMsgServerImpl(*suite.chainA.App.IbcstakingKeeper)
			ctx := suite.chainA.GetContext()
			_, err = ibcStakingMsgServer.IcaWithdraw(
				sdk.WrapSDKContext(ctx), &ibcstakingtypes.MsgIcaWithdraw{
					ZoneId:             hostId,
					HostAddress:        icaConf.icaHostAddress,
					DaomodifierAddress: suite.icaOwnerAddr.String(),
					ReceiverAddress:    suite.icaOwnerAddr.String(),
					TransferPortId:     transferPort,
					TransferChannelId:  transferChannel,
					Amount:             sdk.NewInt64Coin(hostBaseDenom, tc.expect.withdrawAmount),
				})
			suite.Require().NoError(err)
			p, e = ibctesting.ParsePacketFromEvents(ctx.EventManager().Events())
			suite.Require().NoError(suite.icaPath.EndpointA.UpdateClient())
			suite.Require().NoError(suite.icaPath.EndpointB.UpdateClient())
			r, err := suite.icaPath.EndpointB.RecvPacketWithResult(p)
			suite.Require().NotNil(r)
			suite.Require().NoError(err)

			p2, err := ibctesting.ParsePacketFromEvents(r.GetEvents())
			suite.Require().NoError(err)
			suite.Require().NotNil(p2)
			suite.Require().NoError(suite.icaPath.EndpointA.UpdateClient())
			suite.Require().NoError(suite.icaPath.EndpointB.UpdateClient())

			err = suite.transferPath.RelayPacket(p2)
			suite.chainA.NextBlock()
			suite.chainB.NextBlock()
			suite.Require().NoError(err)

			// execute : withdraw
			_, err = msgServer.Withdraw(sdk.WrapSDKContext(suite.chainA.GetContext()), &types.MsgWithdraw{
				ZoneId:            hostId,
				Withdrawer:        tc.initSet.userAddress,
				Recipient:         tc.initSet.userAddress,
				TransferPortId:    transferPort,
				TransferChannelId: transferChannel,
				Amount:            sdk.NewInt64Coin(hostBaseDenom, tc.initSet.withdrawAmount),
			})
			suite.Require().NoError(err)

			// verify user balance after withdraw
			afterUserBalance := suite.chainA.App.BankKeeper.GetBalance(suite.chainA.GetContext(), userAcc, hostIbcDenom)
			suite.Require().Equal(tc.expect.afterWithdrawUserBalance, afterUserBalance.Amount.Int64())
		})
	}
}

func (suite *KeeperTestSuite) TestMultiUserAction() {
	suite.SetupTest()
	suite.chainA.App.IbcstakingKeeper.SetParams(suite.chainA.GetContext(), ibcstakingtypes.Params{
		DaoModifiers: []string{
			suite.icaOwnerAddr.String(),
		},
	})

	trace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom(
		transferPort,
		transferChannel, "stake"))
	suite.chainA.App.TransferKeeper.InitGenesis(suite.chainA.GetContext(), transfertypes.GenesisState{
		PortId:      transferPort,
		DenomTraces: []transfertypes.DenomTrace{trace},
		Params:      transfertypes.Params{SendEnabled: true, ReceiveEnabled: true},
	})

	valAcc, err := sdk.ValAddressFromHex(suite.chainB.Vals.Validators[0].Address.String())
	validatorInfo, _ := suite.chainB.App.StakingKeeper.GetValidator(suite.chainB.GetContext(), valAcc)
	suite.chainA.App.OracleKeeper.InitGenesis(suite.chainA.GetContext(), &oracletypes.GenesisState{
		Params: oracletypes.Params{
			OracleOperators: []string{oracleOperatorAcc.String()},
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            sdk.NewInt64Coin("stake", validatorInfo.Tokens.Int64()),
				Decimal:         8,
				OperatorAddress: oracleOperatorAcc.String(),
			},
		},
	})

	// Mint sn tokens for test
	suite.chainA.App.BankKeeper.MintCoins(suite.chainA.GetContext(),
		types.ModuleName,
		sdk.NewCoins(sdk.NewInt64Coin("snstake", validatorInfo.Tokens.Int64())))

	user1 := suite.chainB.SenderAccounts[0].SenderAccount

	receiver1 := suite.chainA.SenderAccounts[0].SenderAccount
	receiver2 := suite.chainA.SenderAccounts[1].SenderAccount
	receiver3 := suite.chainA.SenderAccounts[2].SenderAccount

	// transfer stake B -> A
	tMsg1 := transfertypes.NewMsgTransfer(transferPort, transferChannel, sdk.NewInt64Coin("stake", 1000000000), user1.GetAddress().String(), receiver1.GetAddress().String(), clienttypes.NewHeight(0, 10000000), 0)
	tMsg2 := transfertypes.NewMsgTransfer(transferPort, transferChannel, sdk.NewInt64Coin("stake", 1000000000), user1.GetAddress().String(), receiver2.GetAddress().String(), clienttypes.NewHeight(0, 10000000), 0)
	tMsg3 := transfertypes.NewMsgTransfer(transferPort, transferChannel, sdk.NewInt64Coin("stake", 1000000000), user1.GetAddress().String(), receiver3.GetAddress().String(), clienttypes.NewHeight(0, 10000000), 0)

	res, err := suite.chainB.SendMsgs(tMsg1)
	suite.Require().NoError(err)
	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)
	err = suite.transferPath.RelayPacket(packet)
	suite.Require().NoError(err)

	res, err = suite.chainB.SendMsgs(tMsg2)
	suite.Require().NoError(err)
	packet, err = ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)
	err = suite.transferPath.RelayPacket(packet)
	suite.Require().NoError(err)

	res, err = suite.chainB.SendMsgs(tMsg3)
	suite.Require().NoError(err)
	packet, err = ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)
	err = suite.transferPath.RelayPacket(packet)
	suite.Require().NoError(err)

	//
	ibcDenom := trace.IBCDenom()

	icaConfig, err := setIbcZone(suite.chainA, suite.chainB, suite.icaOwnerAddr.String())
	suite.Require().NoError(err)

	err = simulateDeposit(suite, icaConfig, receiver1.GetAddress().String(), sdk.NewInt64Coin(ibcDenom, 1000000))
	suite.Require().NoError(err)
	err = simulateDeposit(suite, icaConfig, receiver2.GetAddress().String(), sdk.NewInt64Coin(ibcDenom, 970000))
	suite.Require().NoError(err)
	err = simulateDeposit(suite, icaConfig, receiver3.GetAddress().String(), sdk.NewInt64Coin(ibcDenom, 850000))
	suite.Require().NoError(err)

	record1, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), receiver1.GetAddress())
	record2, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), receiver2.GetAddress())
	record3, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), receiver3.GetAddress())
	fmt.Printf("record: %s\n", record1.String())
	fmt.Printf("record: %s\n", record2.String())
	fmt.Printf("record: %s\n", record3.String())

	suite.Require().NoError(err)
	hostAcc, err := sdk.AccAddressFromBech32(icaConfig.icaHostAddress)
	suite.Require().NoError(err)
	hostBalance := suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), hostAcc, "stake")
	fmt.Printf("host balance: %s\n", hostBalance)
	err = simulateIcaStaking(suite, icaConfig.icaHostAddress, valAcc.String(), hostBalance)
	suite.Require().NoError(err)

	err = suite.icaPath.EndpointB.UpdateClient()
	suite.Require().NoError(err)
	validatorInfo, ok := suite.chainB.App.StakingKeeper.GetValidator(suite.chainB.GetContext(), valAcc)
	suite.Require().True(ok)

	// update oracle
	validatorInfo, ok = suite.chainB.App.StakingKeeper.GetValidator(suite.chainB.GetContext(), valAcc)
	err = suite.chainA.App.OracleKeeper.UpdateChainState(suite.chainA.GetContext(), &oracletypes.ChainInfo{
		Coin:            sdk.NewInt64Coin("stake", validatorInfo.Tokens.Int64()),
		OperatorAddress: oracleOperatorAcc.String(),
		Decimal:         8,
	})

	// claim share token
	minted1, err := suite.chainA.App.GalKeeper.ClaimAndMintShareToken(suite.chainA.GetContext(), receiver1.GetAddress(), *record1.Records[0].Amount)
	shareTokenSupply := suite.chainA.App.BankKeeper.GetSupply(suite.chainA.GetContext(), "snstake")
	fmt.Printf("minted1: %s (total: %s)\n", minted1.String(), shareTokenSupply.String())
	suite.Require().NoError(err)
	minted2, err := suite.chainA.App.GalKeeper.ClaimAndMintShareToken(suite.chainA.GetContext(), receiver2.GetAddress(), *record2.Records[0].Amount)
	shareTokenSupply = suite.chainA.App.BankKeeper.GetSupply(suite.chainA.GetContext(), "snstake")
	fmt.Printf("minted2: %s (total: %s)\n", minted2.String(), shareTokenSupply.String())
	suite.Require().NoError(err)
	minted3, err := suite.chainA.App.GalKeeper.ClaimAndMintShareToken(suite.chainA.GetContext(), receiver3.GetAddress(), *record3.Records[0].Amount)
	shareTokenSupply = suite.chainA.App.BankKeeper.GetSupply(suite.chainA.GetContext(), "snstake")
	fmt.Printf("minted3: %s (total: %s)\n", minted3.String(), shareTokenSupply.String())
	suite.Require().NoError(err)

	hostBalance = suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), hostAcc, "stake")
	fmt.Printf("host balance: %s\n", hostBalance)

	// simulate undelegate without reward
	lambda := suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted1.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(int64(1_000_000), lambda.Int64())
	fmt.Printf("total sn token supply: %s\n", shareTokenSupply.String())
	fmt.Printf("lambda(1): %s\n", lambda.String())

	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted2.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(int64(970_000), lambda.Int64())
	fmt.Printf("lambda(2): %s\n", lambda.String())

	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted3.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(int64(850_000), lambda.Int64())
	fmt.Printf("lambda(3): %s\n", lambda.String())

	// Compound reward, new validator balance : 3838555
	// Expected the value of sn-token for each user
	// user 1 : 1,004,857
	// user 2 : 974,711
	// user 3 : 854,128
	fmt.Printf("validator updatae\n")
	validatorInfo.Tokens = sdk.NewInt(3838555)
	suite.chainB.App.StakingKeeper.SetValidator(suite.chainB.GetContext(), validatorInfo)

	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted1.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(1004857))
	fmt.Printf("lambda(1): %s\n", lambda.String())
	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted2.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(974711))
	fmt.Printf("lambda(2): %s\n", lambda.String())
	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted3.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(854128))
	fmt.Printf("lambda(3): %s\n", lambda.String())

	// Compound reward, new validator balance : 3838555
	// Expected the value of sn-token for each users
	// user 1 : 2,617,800
	// user 2 : 2,539,266
	// user 3 : 2,225,130
	fmt.Printf("validator updatae\n")
	validatorInfo.Tokens = sdk.NewInt(9_999_999)
	suite.chainB.App.StakingKeeper.SetValidator(suite.chainB.GetContext(), validatorInfo)
	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted1.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(2_617_800))
	fmt.Printf("lambda(1): %s\n", lambda.String())
	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted2.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(2_539_266))
	fmt.Printf("lambda(2): %s\n", lambda.String())
	lambda = suite.chainA.App.GalKeeper.CalculateWithdrawAlpha(minted3.Amount.BigInt(), shareTokenSupply.Amount.BigInt(), validatorInfo.Tokens.BigInt())
	suite.Require().Equal(lambda.Int64(), int64(2_225_130))
	fmt.Printf("lambda(3): %s\n", lambda.String())
}

func simulateDeposit(suite *KeeperTestSuite, icaConfig *icaConfig, sender string, amount sdk.Coin) error {
	galServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
	depositMsg := types.MsgDeposit{
		Depositor:         sender,
		ZoneId:            hostId,
		HostAddr:          icaConfig.icaHostAddress,
		Amount:            amount,
		TransferPortId:    "transfer",
		TransferChannelId: "channel-0",
	}
	ctx := suite.chainA.GetContext()
	_, err := galServer.Deposit(sdk.WrapSDKContext(ctx), &depositMsg)
	if err != nil {
		return err
	}

	suite.chainA.NextBlock()
	packet, err := ibctesting.ParsePacketFromEvents(ctx.EventManager().Events())
	if err != nil {
		return err
	}

	err = suite.transferPath.RelayPacket(packet)
	return err
}

func simulateIcaStaking(suite *KeeperTestSuite,
	delegator, validator string, amount sdk.Coin) error {
	owner := suite.icaOwnerAddr
	ctx := suite.chainA.GetContext()
	// check owner balance
	ownerBalance := suite.chainA.App.BankKeeper.GetBalance(ctx, owner, hostIbcDenom)
	fmt.Printf("owner balance : %s\n", ownerBalance.String())

	// check host balance
	ctx = suite.chainB.GetContext()
	delegateMsg := &stakingtypes.MsgDelegate{
		DelegatorAddress: delegator,
		ValidatorAddress: validator,
		Amount:           amount,
	}
	data, err := icatypes.SerializeCosmosTx(suite.chainA.App.AppCodec(), []sdk.Msg{delegateMsg})
	if err != nil {
		return err
	}

	icaPacketData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	params := icahosttypes.NewParams(true, []string{sdk.MsgTypeURL(delegateMsg)})
	suite.chainB.App.ICAHostKeeper.SetParams(suite.chainB.GetContext(), params)
	if err != nil {
		return err
	}

	chanCap, ok := suite.chainA.App.ScopedIBCKeeper.GetCapability(
		suite.chainA.GetContext(), host.ChannelCapabilityPath(suite.icaPath.EndpointA.ChannelConfig.PortID, suite.icaPath.EndpointA.ChannelID))
	suite.Require().True(ok)

	_, err = suite.chainA.App.ICAControllerKeeper.SendTx(suite.chainA.GetContext(), chanCap, "connection-1", suite.icaPath.EndpointA.ChannelConfig.PortID, icaPacketData, ^uint64(0))
	if err != nil {
		return err
	}

	suite.chainA.NextBlock()

	packetRelay := channeltypes.NewPacket(icaPacketData.GetBytes(), 1, suite.icaPath.EndpointA.ChannelConfig.PortID, suite.icaPath.EndpointA.ChannelID, suite.icaPath.EndpointB.ChannelConfig.PortID, suite.icaPath.EndpointB.ChannelID, clienttypes.ZeroHeight(), ^uint64(0))
	err = suite.icaPath.RelayPacket(packetRelay)
	return err
}

func (suite *KeeperTestSuite) verifyTransferCorrectlyExecuted(initSet initialSet, icaInfo icaConfig, set expectedSet) {
	userAcc, err := sdk.AccAddressFromBech32(initSet.userAddress)
	suite.Require().NoError(err)
	hostAcc, err := sdk.AccAddressFromBech32(icaInfo.icaHostAddress)
	suite.Require().NoError(err)

	ibcDenom := ParseAddressToIbcAddress(transferPort, transferChannel, initSet.nativeDenom)
	userBalance := suite.chainA.App.BankKeeper.GetBalance(suite.chainA.GetContext(), userAcc, ibcDenom)
	hostBalance := suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), hostAcc, initSet.nativeDenom)
	suite.Require().Equal(set.userBalance, userBalance.Amount.Int64())
	suite.Require().Equal(set.hostBalance, hostBalance.Amount.Int64())
}

func (suite *KeeperTestSuite) verifyStakingCorrectlyExecuted(validatorAddress string, expected int64) {
	val, err := sdk.ValAddressFromBech32(validatorAddress)
	suite.Require().NoError(err)

	validator, ok := suite.chainB.App.StakingKeeper.GetValidator(
		suite.chainB.GetContext(), val)
	suite.Require().True(ok)
	suite.Require().Equal(expected, validator.BondedTokens().Int64())
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

func setIbcZone(chainA *novatesting.TestChain, chainB *novatesting.TestChain, icaControllerOwnerAddress string) (*icaConfig, error) {
	// Setup ICA
	counterPartyValidator := chainB.App.StakingKeeper.GetValidators(chainB.GetContext(), 1)[0]
	icaControllerPort, err := icatypes.NewControllerPortID(icaControllerOwnerAddress)
	if err != nil {
		return nil, err
	}

	hostAddress, ok := chainB.App.ICAHostKeeper.GetInterchainAccountAddress(chainB.GetContext(), icaConnection, icaControllerPort)
	if !ok {
		return nil, fmt.Errorf("can't find ica host account")
	}

	registerMsg := newBaseRegisteredZone()
	registerMsg.ValidatorAddress = counterPartyValidator.OperatorAddress
	registerMsg.IcaAccount.HostAddress = hostAddress
	chainA.App.IbcstakingKeeper.RegisterZone(chainA.GetContext(), registerMsg)

	return &icaConfig{
		icaHostAddress: hostAddress,
	}, nil
}

func setFunds(chain *novatesting.TestChain, owner string, amt sdk.Coin) error {
	err := chain.App.BankKeeper.MintCoins(chain.GetContext(), types.ModuleName, sdk.Coins{amt})
	if err != nil {
		return err
	}

	acc, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}

	return chain.App.BankKeeper.SendCoinsFromModuleToAccount(chain.GetContext(), types.ModuleName, acc, sdk.Coins{amt})
}
