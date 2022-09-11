package keeper_test

import (
	"github.com/Carina-labs/nova/app"
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	minttypes "github.com/Carina-labs/nova/x/mint/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	//"github.com/tendermint/tendermint/types/time"

	//"github.com/tendermint/tendermint/types/time"
	"time"
)

// set genesis param
func (suite *KeeperTestSuite) InitICA() {
	suite.chainA.GetApp().OracleKeeper.InitGenesis(suite.chainA.GetContext(), &oracletypes.GenesisState{
		Params: oracletypes.Params{
			OracleOperators: []string{
				baseOwnerAcc.String(),
			},
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            sdk.NewCoin(baseDenom, sdk.NewInt(0)),
				ChainId:         zoneId,
				OperatorAddress: baseOwnerAcc.String(),
				OracleVersion:   1,
			},
		},
	})
	suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), zoneId, 1)
	suite.chainA.GetApp().IbcstakingKeeper.InitGenesis(suite.chainA.GetContext(), &ibcstakingtypes.GenesisState{
		Params: ibcstakingtypes.Params{
			DaoModifiers: []string{
				baseOwnerAcc.String(),
			},
		},
	})
	suite.chainB.GetApp().ICAHostKeeper.SetParams(suite.chainB.GetContext(), icahosttypes.Params{
		HostEnabled: true,
		AllowMessages: []string{
			sdk.MsgTypeURL(&banktypes.MsgSend{}),
			sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
			sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
			sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
			sdk.MsgTypeURL(&distrtypes.MsgWithdrawDelegatorReward{}),
			sdk.MsgTypeURL(&distrtypes.MsgSetWithdrawAddress{}),
			sdk.MsgTypeURL(&distrtypes.MsgWithdrawValidatorCommission{}),
			sdk.MsgTypeURL(&distrtypes.MsgFundCommunityPool{}),
			sdk.MsgTypeURL(&govtypes.MsgVote{}),
			sdk.MsgTypeURL(&authz.MsgExec{}),
			sdk.MsgTypeURL(&authz.MsgGrant{}),
			sdk.MsgTypeURL(&authz.MsgRevoke{}),
			sdk.MsgTypeURL(&transfertypes.MsgTransfer{}),
		},
	})
}

func (suite *KeeperTestSuite) mintCoin(ctx sdk.Context, app *app.NovaApp, denom string, amount sdk.Int, addr sdk.AccAddress) {
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	err := app.MintKeeper.MintCoins(ctx, coins)
	suite.Require().NoError(err)

	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) mintCoinToModule(ctx sdk.Context, app *app.NovaApp, denom string, amount sdk.Int, module string) {
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	err := app.MintKeeper.MintCoins(ctx, coins)
	suite.Require().NoError(err)

	err = app.BankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, module, coins)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) icaRelay(ctx sdk.Context) {
	em := ctx.EventManager()
	packet, err := ibctesting.ParsePacketFromEvents(em.Events())

	// 클라이언트 업데이트
	suite.Require().NoError(suite.icaPath.EndpointA.UpdateClient())
	suite.Require().NoError(suite.icaPath.EndpointB.UpdateClient())

	res, err := suite.icaPath.EndpointB.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
	suite.NoError(err)
	suite.Require().NotNil(ack)

	err = suite.icaPath.EndpointA.AcknowledgePacket(packet, ack)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) setHostAddr(zoneId string) sdk.AccAddress {
	zone, ok := suite.chainA.GetApp().IbcstakingKeeper.GetRegisteredZone(suite.chainA.GetContext(), zoneId)
	suite.Require().True(ok)

	hostAddrStr, ok := suite.chainB.GetApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), suite.icaPath.EndpointA.ConnectionID, suite.icaPath.EndpointA.ChannelConfig.PortID)
	suite.Require().True(ok)

	hostAddr, err := sdk.AccAddressFromBech32(hostAddrStr)
	suite.Require().NoError(err)

	zone.IcaAccount.HostAddress = hostAddrStr
	suite.chainA.GetApp().IbcstakingKeeper.RegisterZone(suite.chainA.GetContext(), &zone)

	return hostAddr
}

func (suite *KeeperTestSuite) setDenomTrace(portId, chanId, denom string) transfertypes.DenomTrace {
	denomTrace := transfertypes.DenomTrace{
		Path:      portId + "/" + chanId,
		BaseDenom: denom,
	}

	suite.chainA.GetApp().TransferKeeper.SetDenomTrace(suite.chainA.GetContext(), denomTrace)
	return denomTrace
}

func (suite *KeeperTestSuite) setValidator() string {
	zone, ok := suite.chainA.GetApp().IbcstakingKeeper.GetRegisteredZone(suite.chainA.GetContext(), zoneId)
	suite.Require().True(ok)

	validatorAddr := suite.chainB.App.StakingKeeper.GetValidators(suite.chainB.GetContext(), 1)[0].OperatorAddress
	zone.ValidatorAddress = validatorAddr
	suite.chainA.GetApp().IbcstakingKeeper.RegisterZone(suite.chainA.GetContext(), &zone)

	return validatorAddr
}

func (suite *KeeperTestSuite) TestDeposit() {
	depositor := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name      string
		msg       types.MsgDeposit
		zoneId    string
		denom     string
		depositor sdk.AccAddress
		result    types.DepositRecord
	}{
		{
			name: "success",
			msg: types.MsgDeposit{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Claimer:   depositor.String(),
				Amount:    sdk.NewCoin(baseIbcDenom, sdk.NewInt(10000)),
			},
			zoneId:    zoneId,
			denom:     baseDenom,
			depositor: depositor,
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Denom:  baseIbcDenom,
							Amount: sdk.NewInt(10000),
						},
						State: types.DepositRequest,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.setDenomTrace(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			ibcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), ibcDenom, sdk.NewInt(10000000), tc.depositor)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			suite.Require().NoError(err)

			result, ok := suite.chainA.App.GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result, *result)
		})
	}
}

func (suite *KeeperTestSuite) TestDelegate() {
	suite.InitICA()

	depositor := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name          string
		depositRecord types.DepositRecord
		msg           types.MsgDelegate
		zoneId        string
		denom         string
		depositor     sdk.AccAddress
		result        types.DepositRecord
	}{
		{
			name:      "success",
			depositor: depositor,
			zoneId:    zoneId,
			denom:     baseDenom,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseIbcDenom,
						},
						State:         types.DepositSuccess,
						OracleVersion: 0,
					},
				},
			},
			msg: types.MsgDelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
			},
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						State:     types.DelegateSuccess,
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseIbcDenom,
						},
						OracleVersion:   1,
						DelegateVersion: 1,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.setValidator()
			hostAddr := suite.setHostAddr(tc.zoneId)

			ibcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &tc.depositRecord)
			mintAmt := suite.chainA.GetApp().GalKeeper.GetTotalDepositAmtForZoneId(suite.chainA.GetContext(), tc.zoneId, ibcDenom, types.DepositSuccess)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, mintAmt.Amount, hostAddr)

			// delegate
			excCtx := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Delegate(sdk.WrapSDKContext(excCtx), &tc.msg)
			suite.Require().NoError(err)

			suite.icaRelay(excCtx)

			result, ok := suite.chainA.GetApp().GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result, *result)
		})
	}
}

func (suite *KeeperTestSuite) TestPendingUndelegate() {
	delegator := suite.GenRandomAddress()

	tcs := []struct {
		name           string
		zoneId         string
		delegator      sdk.AccAddress
		undelegateAddr sdk.AccAddress
		msg            types.MsgPendingUndelegate
		denom          string
		result         []*types.UndelegateRecord
	}{
		{
			name:           "success",
			zoneId:         zoneId,
			delegator:      delegator,
			undelegateAddr: delegator,
			msg: types.MsgPendingUndelegate{
				ZoneId:     zoneId,
				Delegator:  delegator.String(),
				Withdrawer: delegator.String(),
				Amount:     sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(10000, 18)),
			},
			denom: baseDenom,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								Denom:  baseSnDenom,
								Amount: sdk.NewIntWithDecimal(10000, 18),
							},
							WithdrawAmount:    sdk.NewInt(0),
							State:             types.UndelegateRequestUser,
							OracleVersion:     1,
							UndelegateVersion: 0,
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.InitICA()

			suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), baseSnDenom, tc.msg.Amount.Amount, tc.delegator)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.PendingUndelegate(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			suite.Require().NoError(err)

			result := suite.chainA.GetApp().GalKeeper.GetAllUndelegateRecord(suite.chainA.GetContext(), tc.zoneId)
			suite.Require().Equal(tc.result, result)
		})
	}
}

func (suite *KeeperTestSuite) TestUndelegate() {
	suite.InitICA()
	delegator := suite.GenRandomAddress()
	valAddr := suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)

	tcs := []struct {
		name              string
		zoneId            string
		delegator         sdk.AccAddress
		undelegateRecords []*types.UndelegateRecord
		msg               types.MsgUndelegate
		denom             string
		snDenom           string
		oracleVersion     uint64
		oracleAmount      sdk.Coin
		burnAmount        sdk.Int
		undelegateResult  []*types.UndelegateRecord
		result            []*types.WithdrawRecord
	}{
		{
			name:      "success",
			zoneId:    zoneId,
			delegator: delegator,
			undelegateRecords: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								Denom:  baseSnDenom,
								Amount: sdk.NewIntWithDecimal(10000, 18),
							},
							WithdrawAmount:    sdk.NewInt(0),
							State:             types.UndelegateRequestUser,
							OracleVersion:     1,
							UndelegateVersion: 0,
						},
					},
				},
			},
			msg: types.MsgUndelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
			},
			denom:         baseDenom,
			snDenom:       baseSnDenom,
			oracleVersion: 2,
			oracleAmount:  sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
			undelegateResult: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								Denom:  baseSnDenom,
								Amount: sdk.NewIntWithDecimal(10000, 18),
							},
							WithdrawAmount:    sdk.NewInt(10000),
							State:             types.UndelegateRequestIca,
							OracleVersion:     1,
							UndelegateVersion: 0,
						},
					},
				},
			},
			burnAmount: sdk.NewIntWithDecimal(10000, 18),
			result: []*types.WithdrawRecord{
				{
					ZoneId:     zoneId,
					Withdrawer: delegator.String(),
					Records: map[uint64]*types.WithdrawRecordContent{
						1: {
							Amount:          sdk.NewInt(10000),
							State:           types.WithdrawStatusRegistered,
							WithdrawVersion: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {

			delegation := stakingtypes.MsgDelegate{
				DelegatorAddress: hostAddr.String(),
				ValidatorAddress: valAddr,
				Amount:           tc.oracleAmount,
			}

			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), tc.denom, tc.oracleAmount.Amount, hostAddr)

			bmsgServer := stakingkeeper.NewMsgServerImpl(*suite.chainB.App.StakingKeeper)
			_, err := bmsgServer.Delegate(sdk.WrapSDKContext(suite.chainB.GetContext()), &delegation)
			suite.Require().NoError(err)

			chainInfo := &oracletypes.ChainInfo{
				Coin:            tc.oracleAmount,
				OracleVersion:   2,
				OperatorAddress: baseOwnerAcc.String(),
			}

			suite.chainA.GetApp().OracleKeeper.UpdateChainState(suite.chainA.GetContext(), chainInfo)
			suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), tc.zoneId, tc.oracleVersion)

			for _, record := range tc.undelegateRecords {
				suite.chainA.GetApp().GalKeeper.SetUndelegateRecord(suite.chainA.GetContext(), record)
				suite.mintCoinToModule(suite.chainA.GetContext(), suite.chainA.GetApp(), tc.snDenom, tc.burnAmount, types.ModuleName)
			}

			excCtx := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err = msgServer.Undelegate(sdk.WrapSDKContext(excCtx), &tc.msg)
			suite.Require().NoError(err)

			undelegateResult := suite.chainA.GetApp().GalKeeper.GetAllUndelegateRecord(excCtx, tc.zoneId)
			suite.Equal(tc.undelegateResult, undelegateResult)

			suite.icaRelay(excCtx)

			for _, result := range tc.result {
				withdrawRecord, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), result.ZoneId, result.Withdrawer)
				suite.Require().True(ok)
				for key, value := range result.Records {
					suite.Require().Equal(withdrawRecord.Records[key].Amount, value.Amount)
					suite.Require().Equal(withdrawRecord.Records[key].WithdrawVersion, value.WithdrawVersion)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestWithdraw() {
	withdrawer1 := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name            string
		zoneId          string
		msg             types.MsgWithdraw
		WithdrawRecords []*types.WithdrawRecord
		withdrawAddr    []sdk.AccAddress
		denom           string
		result          sdk.Coins
	}{
		{
			name:   "success",
			zoneId: zoneId,
			msg: types.MsgWithdraw{
				ZoneId:     zoneId,
				Withdrawer: withdrawer1.String(),
			},
			WithdrawRecords: []*types.WithdrawRecord{
				{
					ZoneId:     zoneId,
					Withdrawer: withdrawer1.String(),
					Records: map[uint64]*types.WithdrawRecordContent{
						1: {
							Amount:          sdk.NewInt(10000),
							State:           types.WithdrawStatusTransferred,
							OracleVersion:   1,
							WithdrawVersion: 1,
							CompletionTime:  time.Now(),
						},
						2: {
							Amount:          sdk.NewInt(5000),
							State:           types.WithdrawStatusRegistered,
							OracleVersion:   1,
							WithdrawVersion: 2,
							CompletionTime:  time.Now(),
						},
					},
				},
			},
			withdrawAddr: []sdk.AccAddress{
				withdrawer1,
			},
			denom:  baseIbcDenom,
			result: sdk.NewCoins(sdk.NewCoin(baseIbcDenom, sdk.NewInt(10000))),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, record := range tc.WithdrawRecords {
				suite.chainA.GetApp().GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), record)
				for _, withdrawRecords := range record.Records {
					suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), baseIbcDenom, withdrawRecords.Amount, baseOwnerAcc)
				}
			}

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Withdraw(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			suite.Require().NoError(err)

			for i, addr := range tc.withdrawAddr {
				result := suite.chainA.GetApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, tc.denom)
				suite.Require().Equal(tc.result[i], result)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestClaimSnAsset() {
	claimer := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name          string
		zoneId        string
		claimer       sdk.AccAddress
		depositRecord types.DepositRecord
		msg           types.MsgClaimSnAsset
		denom         string
		oracleVersion uint64
		oracleAmount  sdk.Int
		result        sdk.Coin
	}{
		{
			name:    "success",
			zoneId:  zoneId,
			claimer: claimer,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: claimer.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: claimer.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseIbcDenom,
						},
						State:           types.DelegateSuccess,
						OracleVersion:   1,
						DelegateVersion: 1,
					},
				},
			},
			msg: types.MsgClaimSnAsset{
				ZoneId:  zoneId,
				Claimer: claimer.String(),
			},
			denom:         baseDenom,
			oracleAmount:  sdk.NewInt(10000),
			oracleVersion: 2,
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(10000, 18)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.InitICA()

			suite.setDenomTrace(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)
			suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &tc.depositRecord)

			chainInfo := oracletypes.ChainInfo{
				ChainId:         tc.zoneId,
				OperatorAddress: baseOwnerAcc.String(),
				Coin:            sdk.NewCoin(tc.denom, tc.oracleAmount),
				OracleVersion:   tc.oracleVersion,
			}
			suite.chainA.GetApp().OracleKeeper.UpdateChainState(suite.chainA.GetContext(), &chainInfo)
			suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), tc.zoneId, tc.oracleVersion)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.ClaimSnAsset(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			suite.Require().NoError(err)

			result := suite.chainA.GetApp().BankKeeper.GetBalance(suite.chainA.GetContext(), claimer, baseSnDenom)
			suite.Require().Equal(tc.result, result)
		})
	}
}
