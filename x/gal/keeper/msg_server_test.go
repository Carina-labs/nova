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
			OracleKeyManager: []string{
				baseOwnerAcc.String(),
			},
		},
		OracleAddressInfo: []oracletypes.OracleAddressInfo{
			{
				ZoneId:        zoneId,
				OracleAddress: []string{baseOwnerAcc.String()},
			},
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            sdk.NewCoin(baseDenom, sdk.NewInt(0)),
				ZoneId:          zoneId,
				OperatorAddress: baseOwnerAcc.String(),
			},
		},
	})
	trace := oracletypes.IBCTrace{
		Version: 1,
		Height:  uint64(suite.chainA.GetContext().BlockHeight()),
	}
	suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), zoneId, trace)
	suite.chainA.GetApp().IcaControlKeeper.InitGenesis(suite.chainA.GetContext(), &icacontroltypes.GenesisState{
		Params: icacontroltypes.Params{
			ControllerKeyManager: []string{
				baseOwnerAcc.String(),
			},
		},
		ControllerAddressInfo: []*icacontroltypes.ControllerAddressInfo{
			{
				ZoneId:            zoneId,
				ControllerAddress: []string{baseOwnerAcc.String()},
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

func (suite *KeeperTestSuite) transferRelay(ctx sdk.Context, fromChain, toChain *novatesting.TestChain, timeout bool) {
	em := ctx.EventManager()
	p, err := ibctesting.ParsePacketFromEvents(em.Events())
	suite.Require().NoError(err)
	fromChain.NextBlock()

	err = suite.transferPath.RelayPacket(p)
	if timeout {
		suite.Require().Error(err)
	} else {
		suite.Require().NoError(err)
	}
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
	depositor1 := suite.GenRandomAddress()

	suite.InitICA()
	_ = suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)

	baseIbcDenom := suite.chainA.GetApp().IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)
	invalidDenom := suite.chainA.GetApp().IcaControlKeeper.GetIBCHashDenom("channel", "port", baseDenom)

	record := types.DepositRecord{
		ZoneId:  zoneId,
		Claimer: depositor1.String(),
		Records: []*types.DepositRecordContent{
			{
				Depositor: depositor1.String(),
				State:     types.DepositSuccess,
			},
			{
				Depositor: depositor1.String(),
				State:     types.DepositSuccess,
			},
			{
				Depositor: depositor1.String(),
				State:     types.DepositSuccess,
			},
			{
				Depositor: depositor1.String(),
				State:     types.DepositSuccess,
			},
			{
				Depositor: depositor1.String(),
				State:     types.DepositSuccess,
			},
		},
	}
	suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &record)
	tcs := []struct {
		name      string
		msg       types.MsgDeposit
		zoneId    string
		denom     string
		depositor sdk.AccAddress
		result    types.DepositRecord
		resultAmt sdk.Coin
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
						State: types.DepositSuccess,
					},
				},
			},
			resultAmt: sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
			err:       false,
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
		{
			name: "fail case 4 - deposit requests exceeded",
			msg: types.MsgDeposit{
				ZoneId:    zoneId,
				Depositor: depositor1.String(),
				Claimer:   depositor1.String(),
				Amount:    sdk.NewCoin(baseIbcDenom, sdk.NewInt(10000)),
			},
			zoneId:    zoneId,
			denom:     baseDenom,
			depositor: depositor1,
			result:    types.DepositRecord{},
			err:       true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.setDenomTrace(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)

			ibcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)
			suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), ibcDenom, sdk.NewInt(10000), tc.depositor)

			escrowAddr := transfertypes.GetEscrowAddress(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), tc.denom, sdk.NewInt(10000), escrowAddr)

			ctxA := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.GetApp().GalKeeper)
			_, err := msgServer.Deposit(sdk.WrapSDKContext(ctxA), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				balance := suite.chainA.GetApp().BankKeeper.GetBalance(ctxA, depositor, ibcDenom)

				suite.transferRelay(ctxA, suite.chainA, suite.chainB, false)

				result, ok := suite.chainA.GetApp().GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)

				balance = suite.chainB.GetApp().BankKeeper.GetBalance(suite.chainB.GetContext(), hostAddr, tc.denom)
				suite.Require().Equal(tc.resultAmt, balance)
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
				Version:           0,
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
						DelegateVersion: 0,
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
				Version:           0,
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
				Version:           0,
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
				Version:           0,
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
				suite.Require().NoError(err)
				suite.icaRelay(excCtx)

				result, ok := suite.chainA.GetApp().GalKeeper.GetUserDepositRecord(suite.chainA.GetContext(), tc.zoneId, tc.depositor)
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestPendingUndelegate() {
	delegator1 := suite.GenRandomAddress()
	delegator2 := suite.GenRandomAddress()
	delegator3 := suite.GenRandomAddress()

	max_entries := suite.GenRandomAddress()

	record := &types.UndelegateRecord{
		ZoneId:    zoneId,
		Delegator: max_entries.String(),
		Records: []*types.UndelegateRecordContent{
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByUser,
			},
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByUser,
			},
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByIca,
			},
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByIca,
			},
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByIca,
			},
			{
				Withdrawer: max_entries.String(),
				SnAssetAmount: &sdk.Coin{
					Denom:  baseSnDenom,
					Amount: sdk.NewIntWithDecimal(0, 18),
				},
				WithdrawAmount: sdk.NewInt(0),
				State:          types.UndelegateRequestByIca,
			},
		},
	}

	suite.chainA.GetApp().GalKeeper.SetUndelegateRecord(suite.chainA.GetContext(), record)

	record = &types.UndelegateRecord{
		ZoneId:    zoneId,
		Delegator: delegator3.String(),
		Records: []*types.UndelegateRecordContent{
			{
				Withdrawer: delegator3.String(),
				State:      types.UndelegateRequestByUser,
			},
			{
				Withdrawer: delegator3.String(),
				State:      types.UndelegateRequestByUser,
			},
			{
				Withdrawer: delegator3.String(),
				State:      types.UndelegateRequestByUser,
			},
			{
				Withdrawer: delegator3.String(),
				State:      types.UndelegateRequestByUser,
			},
			{
				Withdrawer: delegator3.String(),
				State:      types.UndelegateRequestByUser,
			},
		},
	}

	suite.chainA.GetApp().GalKeeper.SetUndelegateRecord(suite.chainA.GetContext(), record)

	tcs := []struct {
		name           string
		zoneId         string
		delegator      sdk.AccAddress
		undelegateAddr sdk.AccAddress
		msg            types.MsgPendingUndelegate
		denom          string
		result         types.UndelegateRecord
		err            bool
	}{
		{
			name:           "success",
			zoneId:         zoneId,
			delegator:      delegator1,
			undelegateAddr: delegator1,
			msg: types.MsgPendingUndelegate{
				ZoneId:     zoneId,
				Delegator:  delegator1.String(),
				Withdrawer: delegator1.String(),
				Amount:     sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(10000, 18)),
			},
			denom: baseDenom,
			result: types.UndelegateRecord{
				ZoneId:    zoneId,
				Delegator: delegator1.String(),
				Records: []*types.UndelegateRecordContent{
					{
						Withdrawer: delegator1.String(),
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
			err: false,
		},
		{
			name:           "success - next sequence",
			zoneId:         zoneId,
			delegator:      max_entries,
			undelegateAddr: max_entries,
			msg: types.MsgPendingUndelegate{
				ZoneId:     zoneId,
				Delegator:  max_entries.String(),
				Withdrawer: max_entries.String(),
				Amount:     sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(10000, 18)),
			},
			denom: baseDenom,
			result: types.UndelegateRecord{
				ZoneId:    zoneId,
				Delegator: max_entries.String(),
				Records: []*types.UndelegateRecordContent{
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByUser,
					},
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByUser,
					},
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByIca,
					},
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByIca,
					},
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByIca,
					},
					{
						Withdrawer: max_entries.String(),
						SnAssetAmount: &sdk.Coin{
							Denom:  baseSnDenom,
							Amount: sdk.NewIntWithDecimal(0, 18),
						},
						WithdrawAmount: sdk.NewInt(0),
						State:          types.UndelegateRequestByIca,
					},
					{
						Withdrawer: max_entries.String(),
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
			err: false,
		},
		{
			name:           "fail case 1 - zone not found",
			zoneId:         "test1",
			delegator:      delegator2,
			undelegateAddr: delegator2,
			msg: types.MsgPendingUndelegate{
				ZoneId:     "test1",
				Delegator:  delegator2.String(),
				Withdrawer: delegator2.String(),
				Amount:     sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(10000, 18)),
			},
			denom:  baseDenom,
			result: types.UndelegateRecord{},
			err:    true,
		},
		{
			name:           "fail case 2 - snDenom is different from registered snDenom",
			zoneId:         zoneId,
			delegator:      delegator2,
			undelegateAddr: delegator2,
			msg: types.MsgPendingUndelegate{
				ZoneId:     zoneId,
				Delegator:  delegator2.String(),
				Withdrawer: delegator2.String(),
				Amount:     sdk.NewCoin("snFailDenom", sdk.NewIntWithDecimal(10000, 18)),
			},
			denom:  baseDenom,
			result: types.UndelegateRecord{},
			err:    true,
		},
		{
			name:           "fail case 3 - undelegate requests exceeded",
			zoneId:         zoneId,
			delegator:      delegator3,
			undelegateAddr: delegator3,
			msg: types.MsgPendingUndelegate{
				ZoneId:     zoneId,
				Delegator:  delegator3.String(),
				Withdrawer: delegator3.String(),
				Amount:     sdk.NewCoin("snFailDenom", sdk.NewIntWithDecimal(10000, 18)),
			},
			denom:  baseDenom,
			result: types.UndelegateRecord{},
			err:    true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.InitICA()
			suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), baseSnDenom, tc.msg.Amount.Amount, tc.delegator)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.PendingUndelegate(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				result, ok := suite.chainA.GetApp().GalKeeper.GetUndelegateRecord(suite.chainA.GetContext(), tc.zoneId, tc.delegator.String())
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUndelegate() {
	suite.InitICA()
	delegator := suite.GenRandomAddress()
	withdrawer := suite.GenRandomAddress()
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
		err               bool
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
						{
							Withdrawer: withdrawer.String(),
							SnAssetAmount: &sdk.Coin{
								Denom:  baseSnDenom,
								Amount: sdk.NewIntWithDecimal(1000, 18),
							},
							WithdrawAmount:    sdk.NewInt(0),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 0,
						},
						{
							Withdrawer: withdrawer.String(),
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
				Version:           0,
			},
			denom:         baseDenom,
			snDenom:       baseSnDenom,
			oracleVersion: 2,
			oracleAmount:  sdk.NewCoin(baseDenom, sdk.NewInt(21000)),
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
						{
							Withdrawer: withdrawer.String(),
							SnAssetAmount: &sdk.Coin{
								Denom:  baseSnDenom,
								Amount: sdk.NewIntWithDecimal(1000, 18),
							},
							WithdrawAmount:    sdk.NewInt(1000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 0,
						},
						{
							Withdrawer: withdrawer.String(),
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
			burnAmount: sdk.NewIntWithDecimal(21000, 18),
			result: []*types.WithdrawRecord{
				{
					ZoneId:     zoneId,
					Withdrawer: delegator.String(),
					Records: map[uint64]*types.WithdrawRecordContent{
						0: {
							Amount:          sdk.NewInt(10000),
							State:           types.WithdrawStatusRegistered,
							WithdrawVersion: 0,
						},
					},
				},
				{
					ZoneId:     zoneId,
					Withdrawer: withdrawer.String(),
					Records: map[uint64]*types.WithdrawRecordContent{
						0: {
							Amount:          sdk.NewInt(11000),
							State:           types.WithdrawStatusRegistered,
							WithdrawVersion: 0,
						},
					},
				},
			},
			err: false,
		},
		{
			name:      "fail case 1 - zone not found",
			zoneId:    "test",
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
				ZoneId:            "test",
				ControllerAddress: baseOwnerAcc.String(),
				Version:           1,
			},
			denom:            baseDenom,
			snDenom:          baseSnDenom,
			oracleVersion:    2,
			oracleAmount:     sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
			undelegateResult: []*types.UndelegateRecord(nil),
			burnAmount:       sdk.NewIntWithDecimal(0, 18),
			result:           []*types.WithdrawRecord{},
			err:              true,
		},
		{
			name:      "fail case 2 - oracle version",
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
							OracleVersion:     2,
							UndelegateVersion: 0,
						},
					},
				},
			},
			msg: types.MsgUndelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
				Version:           1,
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
							WithdrawAmount:    sdk.NewInt(0),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     2,
							UndelegateVersion: 0,
						},
					},
				},
			},
			burnAmount: sdk.NewIntWithDecimal(10000, 18),
			result:     []*types.WithdrawRecord{},
			err:        true,
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
				OperatorAddress: baseOwnerAcc.String(),
				ZoneId:          zoneId,
			}

			trace := oracletypes.IBCTrace{
				Version: tc.oracleVersion,
				Height:  uint64(suite.chainA.GetContext().BlockHeight()),
			}

			suite.chainA.GetApp().OracleKeeper.UpdateChainState(suite.chainA.GetContext(), chainInfo)
			suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), tc.zoneId, trace)

			for _, record := range tc.undelegateRecords {
				suite.chainA.GetApp().GalKeeper.SetUndelegateRecord(suite.chainA.GetContext(), record)
			}
			suite.mintCoinToModule(suite.chainA.GetContext(), suite.chainA.GetApp(), tc.snDenom, tc.burnAmount, types.ModuleName)

			excCtx := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err = msgServer.Undelegate(sdk.WrapSDKContext(excCtx), &tc.msg)

			undelegateResult := suite.chainA.GetApp().GalKeeper.GetAllUndelegateRecord(excCtx, tc.zoneId)

			if tc.err {
				suite.Require().Error(err)
				suite.Equal(tc.undelegateResult, undelegateResult)
			} else {
				suite.Require().NoError(err)
				suite.Equal(tc.undelegateResult, undelegateResult)
				suite.icaRelay(excCtx)

				for _, result := range tc.result {
					withdrawRecord, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), result.ZoneId, result.Withdrawer)
					suite.Require().True(ok)
					for key, value := range result.Records {
						suite.Require().Equal(withdrawRecord.Records[key].Amount, value.Amount)
						suite.Require().Equal(key, value.WithdrawVersion)
					}
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestIcaWithdraw() {
	withdrawAddr := suite.GenRandomAddress()

	suite.InitICA()
	_ = suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)

	icaPortId := suite.transferPath.EndpointB.ChannelConfig.PortID
	icaChanId := suite.transferPath.EndpointB.ChannelID
	tcs := []struct {
		name                 string
		zoneId               string
		withdrawAddr         sdk.AccAddress
		withdrawRecord       types.WithdrawRecord
		resultWithdrawRecord types.WithdrawRecord
		msg                  types.MsgIcaWithdraw
		denom                string
		result               bool
		err                  bool
	}{
		{
			name:         "success",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			resultWithdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferRequest,
						OracleVersion:   1,
						WithdrawVersion: 0,
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
				Version:              0,
			},
			denom:  baseDenom,
			result: true,
			err:    false,
		},
		{
			name:         "fail case 1 - sender address is not the controller address ",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			resultWithdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferRequest,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			msg: types.MsgIcaWithdraw{
				ZoneId:               zoneId,
				ControllerAddress:    withdrawAddr.String(),
				IcaTransferPortId:    icaPortId,
				IcaTransferChannelId: icaChanId,
				ChainTime:            suite.chainA.GetContext().BlockTime().Add(time.Hour),
				Version:              0,
			},
			denom:  baseDenom,
			result: true,
			err:    true,
		},
		{
			name:         "fail case 2 - ack result is fail",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			resultWithdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferRequest,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			msg: types.MsgIcaWithdraw{
				ZoneId:               zoneId,
				ControllerAddress:    baseOwnerAcc.String(),
				IcaTransferPortId:    "transfer",
				IcaTransferChannelId: "channel-1000",
				ChainTime:            suite.chainA.GetContext().BlockTime().Add(time.Hour),
				Version:              0,
			},
			denom:  baseDenom,
			result: false,
			err:    false,
		},
		{
			name:         "fail case 3 - time",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime().Add(time.Hour),
					},
				},
			},
			resultWithdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferRequest,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime().Add(time.Hour),
					},
				},
			},
			msg: types.MsgIcaWithdraw{
				ZoneId:               zoneId,
				ControllerAddress:    baseOwnerAcc.String(),
				IcaTransferPortId:    icaPortId,
				IcaTransferChannelId: icaChanId,
				ChainTime:            suite.chainA.GetContext().BlockTime(),
				Version:              0,
			},
			denom:  baseDenom,
			result: false,
			err:    true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			excCtxA := suite.chainA.GetContext()

			suite.chainA.GetApp().GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), &tc.withdrawRecord)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, sdk.NewInt(10000), hostAddr)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.IcaWithdraw(sdk.WrapSDKContext(excCtxA), &tc.msg)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				res := suite.icaRelay(excCtxA)

				var ack channeltypes.Acknowledgement
				err = channeltypes.SubModuleCdc.UnmarshalJSON(res, &ack)
				suite.Require().NoError(err)
				suite.Require().Equal(tc.result, ack.Success())

				resultRecords, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), tc.zoneId, tc.withdrawAddr.String())
				suite.Require().True(ok)
				suite.Require().Equal(tc.resultWithdrawRecord, *resultRecords)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestTransferWithdraw() {

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
		timestamp      uint64
		resultRecord   types.WithdrawRecord
		timeout        bool
		err            bool
	}{
		{
			name:         "success",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr,
			withdrawRecord: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferRequest,
						OracleVersion:   1,
						WithdrawVersion: 0,
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
					0: {
						Amount:          sdk.NewInt(10000),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 0,
						CompletionTime:  suite.chainB.GetContext().BlockTime(),
					},
				},
			},
			timestamp: uint64(suite.chainB.GetContext().BlockTime().Add(time.Hour).UnixNano()),
			timeout:   false,
			err:       false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// set version
			suite.chainA.GetApp().GalKeeper.IsValidWithdrawVersion(suite.chainA.GetContext(), tc.zoneId, 0)
			suite.chainA.GetApp().GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), &tc.withdrawRecord)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), tc.denom, tc.resultAmount.Amount, hostAddr)
			excCtxB := suite.chainB.GetContext()

			timeoutHeight := ibcclienttypes.Height{RevisionHeight: 0, RevisionNumber: 0}
			TimeoutTimestamp := tc.timestamp

			transferAmt := sdk.NewCoin(tc.denom, tc.resultAmount.Amount)
			err := suite.chainB.GetApp().TransferKeeper.SendTransfer(excCtxB, suite.transferPath.EndpointB.ChannelConfig.PortID, suite.transferPath.EndpointB.ChannelID, transferAmt, hostAddr, baseOwnerAcc.String(), timeoutHeight, TimeoutTimestamp)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.transferRelay(excCtxB, suite.chainB, suite.chainA, tc.timeout)

				resultRecords, ok := suite.chainA.GetApp().GalKeeper.GetWithdrawRecord(suite.chainA.GetContext(), tc.zoneId, tc.withdrawAddr.String())
				suite.Require().True(ok)
				suite.Require().Equal(tc.resultRecord, *resultRecords)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestWithdraw() {
	withdrawer1 := suite.GenRandomAddress()
	withdrawer2 := suite.GenRandomAddress()
	withdrawer3 := suite.GenRandomAddress()
	baseIbcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)

	tcs := []struct {
		name            string
		zoneId          string
		msg             types.MsgWithdraw
		WithdrawRecords []*types.WithdrawRecord
		withdrawAddr    []sdk.AccAddress
		denom           string
		result          sdk.Coins
		err             bool
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
			err:    false,
		},
		{
			name:   "success",
			zoneId: zoneId,
			msg: types.MsgWithdraw{
				ZoneId:     zoneId,
				Withdrawer: withdrawer2.String(),
			},
			WithdrawRecords: []*types.WithdrawRecord{
				{
					ZoneId:     zoneId,
					Withdrawer: withdrawer2.String(),
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
							State:           types.WithdrawStatusTransferred,
							OracleVersion:   1,
							WithdrawVersion: 2,
							CompletionTime:  time.Now(),
						},
						3: {
							Amount:          sdk.NewInt(10000),
							State:           types.WithdrawStatusRegistered,
							OracleVersion:   1,
							WithdrawVersion: 2,
							CompletionTime:  time.Now(),
						},
					},
				},
			},
			withdrawAddr: []sdk.AccAddress{
				withdrawer2,
			},
			denom:  baseIbcDenom,
			result: sdk.NewCoins(sdk.NewCoin(baseIbcDenom, sdk.NewInt(15000))),
			err:    false,
		},
		{
			name:   "success",
			zoneId: zoneId,
			msg: types.MsgWithdraw{
				ZoneId:     zoneId,
				Withdrawer: withdrawer3.String(),
			},
			WithdrawRecords: []*types.WithdrawRecord{
				{
					ZoneId:     zoneId,
					Withdrawer: withdrawer3.String(),
					Records: map[uint64]*types.WithdrawRecordContent{
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
				withdrawer3,
			},
			denom:  baseIbcDenom,
			result: sdk.NewCoins(sdk.NewCoin(baseIbcDenom, sdk.NewInt(0))),
			err:    true,
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

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				for i, addr := range tc.withdrawAddr {
					result := suite.chainA.GetApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, tc.denom)
					suite.Require().Equal(tc.result[i], result)
				}
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
		err           bool
	}{
		{
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
					{
						Depositor: claimer.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(2000),
							Denom:  baseIbcDenom,
						},
						State:           types.DelegateSuccess,
						OracleVersion:   1,
						DelegateVersion: 1,
					},
					{
						Depositor: claimer.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseIbcDenom,
						},
						State:           types.DelegateRequest,
						OracleVersion:   1,
						DelegateVersion: 1,
					},
					{
						Depositor: claimer.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  "denom",
						},
						State:           types.DelegateRequest,
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
			oracleAmount:  sdk.NewInt(12000),
			oracleVersion: 2,
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(12000, 18)),
			err:           false,
		},
		{
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
			oracleAmount:  sdk.NewInt(22000),
			oracleVersion: 2,
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(22000, 18)),
			err:           false,
		},
		{
			zoneId:  zoneId,
			claimer: claimer,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: claimer.String(),
				Records: []*types.DepositRecordContent(nil),
			},
			msg: types.MsgClaimSnAsset{
				ZoneId:  zoneId,
				Claimer: claimer.String(),
			},
			denom:         baseDenom,
			oracleAmount:  sdk.NewInt(22000),
			oracleVersion: 2,
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(22000, 18)),
			err:           true,
		},
	}

	for _, tc := range tcs {
		suite.InitICA()

		suite.setDenomTrace(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, tc.denom)
		suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &tc.depositRecord)

		chainInfo := oracletypes.ChainInfo{
			ZoneId:          tc.zoneId,
			OperatorAddress: baseOwnerAcc.String(),
			Coin:            sdk.NewCoin(tc.denom, tc.oracleAmount),
		}

		suite.chainA.GetApp().OracleKeeper.UpdateChainState(suite.chainA.GetContext(), &chainInfo)

		trace := oracletypes.IBCTrace{
			Version: tc.oracleVersion,
			Height:  uint64(suite.chainA.GetContext().BlockHeight()),
		}
		suite.chainA.GetApp().OracleKeeper.SetOracleVersion(suite.chainA.GetContext(), tc.zoneId, trace)

		msgServer := keeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
		_, err := msgServer.ClaimSnAsset(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)

		if tc.err {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
			result := suite.chainA.GetApp().BankKeeper.GetBalance(suite.chainA.GetContext(), claimer, baseSnDenom)
			suite.Require().Equal(tc.result, result)
		}
	}
}
