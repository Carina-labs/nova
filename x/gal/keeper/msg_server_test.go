package keeper_test

import (
	"fmt"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	intertxkeeper "github.com/Carina-labs/nova/x/inter-tx/keeper"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"time"
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
				LastBlockHeight: 0,
				AppHash:         "",
				ChainId:         hostId,
				BlockProposer:   "",
			},
		},
	})

	// prepare inter-tx keeper
	suite.chainA.App.IntertxKeeper.SetParams(suite.chainA.GetContext(), intertxtypes.Params{
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
				Depositor: tc.initSet.userAddress,
				ZoneId:    hostId,
				HostAddr:  icaConf.icaHostAddress,
				Amount:    sdk.NewCoins(sdk.NewInt64Coin(hostIbcDenom, 1000)),
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
			interTxMsgServer := intertxkeeper.NewMsgServerImpl(*suite.chainA.App.IntertxKeeper)
			ctx := suite.chainA.GetContext()
			_, err = interTxMsgServer.IcaWithdraw(
				sdk.WrapSDKContext(ctx), &intertxtypes.MsgIcaWithdraw{
					ZoneName:        hostId,
					SenderAddress:   icaConf.icaHostAddress,
					OwnerAddress:    suite.icaOwnerAddr.String(),
					ReceiverAddress: suite.icaOwnerAddr.String(),
					Amount:          sdk.NewInt64Coin(hostBaseDenom, tc.expect.withdrawAmount),
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
				ZoneId:     hostId,
				Withdrawer: tc.initSet.userAddress,
				Recipient:  tc.initSet.userAddress,
				Amount:     sdk.NewInt64Coin(hostBaseDenom, tc.initSet.withdrawAmount),
			})
			suite.Require().NoError(err)

			// verify user balance after withdraw
			afterUserBalance := suite.chainA.App.BankKeeper.GetBalance(suite.chainA.GetContext(), userAcc, hostIbcDenom)
			suite.Require().Equal(tc.expect.afterWithdrawUserBalance, afterUserBalance.Amount.Int64())
		})
	}
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
	chainA.App.IntertxKeeper.RegisterZone(chainA.GetContext(), registerMsg)

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
