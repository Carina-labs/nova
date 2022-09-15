package keeper_test

import (
	"github.com/Carina-labs/nova/app"
	novatesting "github.com/Carina-labs/nova/testing"
	"github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	icacontroltypes "github.com/Carina-labs/nova/x/icacontrol/types"
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
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
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
	suite.chainA.GetApp().IcaControlKeeper.InitGenesis(suite.chainA.GetContext(), &icacontroltypes.GenesisState{
		Params: icacontroltypes.Params{
			ControllerAddress: []string{
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

func (suite *KeeperTestSuite) icaRelay(ctx sdk.Context) []byte {
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

	return ack
}

func (suite *KeeperTestSuite) transferRelay(ctx sdk.Context, fromChain, toChain *novatesting.TestChain) {
	em := ctx.EventManager()
	p, err := ibctesting.ParsePacketFromEvents(em.Events())
	suite.Require().NoError(err)
	fromChain.NextBlock()

	err = suite.transferPath.RelayPacket(p)
	suite.Require().NoError(err)
	toChain.NextBlock()
}

func (suite *KeeperTestSuite) setHostAddr(zoneId string) sdk.AccAddress {
	zone, ok := suite.chainA.GetApp().IcaControlKeeper.GetRegisteredZone(suite.chainA.GetContext(), zoneId)
	suite.Require().True(ok)

	hostAddrStr, ok := suite.chainB.GetApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), suite.icaPath.EndpointA.ConnectionID, suite.icaPath.EndpointA.ChannelConfig.PortID)
	suite.Require().True(ok)

	hostAddr, err := sdk.AccAddressFromBech32(hostAddrStr)
	suite.Require().NoError(err)

	zone.IcaAccount.HostAddress = hostAddrStr
	suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), &zone)

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
	zone, ok := suite.chainA.GetApp().IcaControlKeeper.GetRegisteredZone(suite.chainA.GetContext(), zoneId)
	suite.Require().True(ok)

	validatorAddr := suite.chainB.App.StakingKeeper.GetValidators(suite.chainB.GetContext(), 1)[0].OperatorAddress
	zone.ValidatorAddress = validatorAddr
	suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), &zone)

	return validatorAddr
}

func (suite *KeeperTestSuite) TestDeposit() {
	depositor := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)
	invalidDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom("channel", "port", baseDenom)

	tcs := []struct {
		name      string
		msg       types.MsgDeposit
		zoneId    string
		denom     string
		depositor sdk.AccAddress
		result    types.DepositRecord
		err       bool
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
			err: false,
		},
		{
			name: "fail case 1 - zone not found",
			msg: types.MsgDeposit{
				ZoneId:    "test",
				Depositor: depositor.String(),
				Claimer:   depositor.String(),
				Amount:    sdk.NewCoin(baseIbcDenom, sdk.NewInt(10000)),
			},
			zoneId:    "test",
			denom:     baseDenom,
			depositor: depositor,
			result:    types.DepositRecord{},
			err:       true,
		},
		{
			name: "fail case 2 - insufficient balance for deposit address",
			msg: types.MsgDeposit{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Claimer:   depositor.String(),
				Amount:    sdk.NewCoin(baseIbcDenom, sdk.NewInt(40000)),
			},
			zoneId:    zoneId,
			denom:     baseDenom,
			depositor: depositor,
			result:    types.DepositRecord{},
			err:       true,
		},
		{
			name: "fail case 3 - invalid denom",
			msg: types.MsgDeposit{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Claimer:   depositor.String(),
				Amount:    sdk.NewCoin(invalidDenom, sdk.NewInt(10000)),
			},
			zoneId:    zoneId,
			denom:     baseDenom,
			depositor: depositor,
			result:    types.DepositRecord{},
			err:       true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.setDenomTrace(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			ibcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), ibcDenom, sdk.NewInt(10000), tc.depositor)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				result, ok := suite.chainA.App.GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestDelegate() {
	suite.InitICA()

	depositor := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name          string
		depositRecord types.DepositRecord
		msg           types.MsgDelegate
		zoneId        string
		denom         string
		depositor     sdk.AccAddress
		result        types.DepositRecord
		err           bool
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
			err: false,
		},
		{
			name:      "fail case 1 - deposit record is not nil",
			depositor: depositor,
			zoneId:    zoneId,
			denom:     baseDenom,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{},
			},
			msg: types.MsgDelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
			},
			result: types.DepositRecord{},
			err:    true,
		},
		{
			name:      "fail case 2 - controller address is invalid",
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
				ControllerAddress: depositor.String(),
			},
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						State:     types.DepositSuccess,
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseIbcDenom,
						},
						OracleVersion:   1,
						DelegateVersion: 1,
					},
				},
			},
			err: true,
		},
		{
			name:      "fail case 3 - zone not found",
			depositor: depositor,
			zoneId:    zoneId,
			denom:     baseDenom,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{},
			},
			msg: types.MsgDelegate{
				ZoneId:            "test",
				ControllerAddress: baseOwnerAcc.String(),
			},
			result: types.DepositRecord{},
			err:    true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.setValidator()
			hostAddr := suite.setHostAddr(tc.zoneId)

			ibcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &tc.depositRecord)
			mintAmt := suite.chainA.GetApp().GalKeeper.GetTotalDepositAmtForZoneId(suite.chainA.GetContext(), tc.zoneId, ibcDenom, types.DepositSuccess)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, mintAmt.Amount, hostAddr)

			// delegate
			excCtx := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Delegate(sdk.WrapSDKContext(excCtx), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.icaRelay(excCtx)

				suite.Require().NoError(err)
				result, ok := suite.chainA.GetApp().GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)
			}
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
							State:             types.UndelegateRequestByUser,
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
							State:             types.UndelegateRequestByUser,
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
							State:             types.UndelegateRequestByIca,
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

func (suite *KeeperTestSuite) TestIcaWithdraw() {
	// IcaWithdraw가 성공적으로 진행 되었는지 확인 - ack의 err여부만 확인
	// 1. ICA에 성공(Transfer state 성공)
	// 2. ICA에 실패(Transfer state 실패)

	withdrawAddr := suite.GenRandomAddress()

	suite.InitICA()
	_ = suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)

	icaPortId := suite.transferPath.EndpointB.ChannelConfig.PortID
	icaChanId := suite.transferPath.EndpointB.ChannelID
	tcs := []struct {
		name           string
		zoneId         string
		withdrawAddr   sdk.AccAddress
		withdrawRecord types.WithdrawRecord
		msg            types.MsgIcaWithdraw
		denom          string
		result         bool
	}{
		{
			name:         "success",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 1,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			msg: types.MsgIcaWithdraw{
				ZoneId:               zoneId,
				ControllerAddress:    baseOwnerAcc.String(),
				IcaTransferPortId:    icaPortId,
				IcaTransferChannelId: icaChanId,
				ChainTime:            suite.chainA.GetContext().BlockTime().Add(time.Hour),
			},
			denom:  baseDenom,
			result: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			excCtxA := suite.chainA.GetContext()

			suite.chainA.GetApp().GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), &tc.withdrawRecord)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, sdk.NewInt(10000), hostAddr)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.IcaWithdraw(sdk.WrapSDKContext(excCtxA), &tc.msg)
			suite.Require().NoError(err)

			res := suite.icaRelay(excCtxA)

			var ack channeltypes.Acknowledgement
			err = channeltypes.SubModuleCdc.UnmarshalJSON(res, &ack)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.result, ack.Success())

			resultRecords, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), tc.zoneId, tc.withdrawAddr.String())
			suite.Require().True(ok)
			suite.Require().Equal(tc.withdrawRecord, *resultRecords)
		})
	}
}

// icaWithdraw 성공케이스
func (suite *KeeperTestSuite) TestTransferWithdraw() {

	// ICA withdraw는 당연히 성공 했고, chainB에서 transfer성공 여부
	// transfer - get timeout packet, get ack result is err
	withdrawAddr := suite.GenRandomAddress()

	suite.InitICA()
	_ = suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)

	portId := suite.transferPath.EndpointA.ChannelConfig.PortID
	chanId := suite.transferPath.EndpointA.ChannelID
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(portId, chanId, baseDenom)

	tcs := []struct {
		name           string
		zoneId         string
		withdrawAddr   sdk.AccAddress
		withdrawRecord types.WithdrawRecord
		denom          string
		resultAmount   sdk.Coin
		resultRecord   types.WithdrawRecord
	}{
		{
			name:         "success",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 1,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			denom:        baseDenom,
			resultAmount: sdk.NewCoin(baseIbcDenom, sdk.NewInt(10000)),
			resultRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {

			suite.chainA.GetApp().GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), &tc.withdrawRecord)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), tc.denom, tc.resultAmount.Amount, hostAddr)
			excCtxB := suite.chainB.GetContext()

			timeoutHeight := ibcclienttypes.Height{RevisionHeight: 0, RevisionNumber: 0}
			TimeoutTimestamp := uint64(excCtxB.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds())

			transferAmt := sdk.NewCoin(tc.denom, tc.resultAmount.Amount)
			err := suite.chainB.GetApp().TransferKeeper.SendTransfer(excCtxB, suite.transferPath.EndpointB.ChannelConfig.PortID, suite.transferPath.EndpointB.ChannelID, transferAmt, hostAddr, baseOwnerAcc.String(), timeoutHeight, TimeoutTimestamp)
			suite.Require().NoError(err)

			suite.transferRelay(excCtxB, suite.chainB, suite.chainA)

			resultRecords, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), tc.zoneId, tc.withdrawAddr.String())
			suite.Require().True(ok)
			suite.Require().Equal(tc.resultRecord, *resultRecords)
		})
	}
}

func (suite *KeeperTestSuite) TestWithdraw() {
	withdrawer1 := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

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
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

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
