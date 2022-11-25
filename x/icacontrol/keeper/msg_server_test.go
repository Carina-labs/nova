package keeper_test

import (
	"github.com/Carina-labs/nova/app"
	galtypes "github.com/Carina-labs/nova/x/gal/types"
	"github.com/Carina-labs/nova/x/icacontrol/keeper"
	"github.com/Carina-labs/nova/x/icacontrol/types"
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
	"time"
)

func bech32toValidatorAddresses(validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
}

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
	suite.chainA.GetApp().IcaControlKeeper.InitGenesis(suite.chainA.GetContext(), &types.GenesisState{
		Params: types.Params{
			ControllerKeyManager: []string{
				baseOwnerAcc.String(),
			},
		},
		ControllerAddressInfo: []*types.ControllerAddressInfo{
			{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
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

func (suite *KeeperTestSuite) setControllerAddr(address string) {
	var addresses []string
	addr1 := address
	addresses = append(addresses, addr1)
	params := types.Params{
		ControllerKeyManager: addresses,
	}
	suite.chainA.App.IcaControlKeeper.SetParams(suite.chainA.GetContext(), params)
}

func (suite *KeeperTestSuite) getGrantMsg(msg, zoneId, grantee string, controllerAddr sdk.AccAddress) types.MsgIcaAuthzGrant {
	var authorization authz.Authorization
	var allowed []sdk.ValAddress
	var denied []sdk.ValAddress
	var allowValidators []string
	var delegateLimit sdk.Coin
	var err error

	addr := suite.GenRandomAddress()

	switch msg {
	case "send":
		spendLimit := sdk.NewCoins(sdk.NewCoin(baseDenom, sdk.NewInt(10000)))
		authorization = banktypes.NewSendAuthorization(spendLimit)
		break
	case "generic":
		msgType := ""
		authorization = authz.NewGenericAuthorization(msgType)
		break
	case "delegate", "unbond", "redelegate", "undelegate":
		allowValidators = append(allowValidators, sdk.ValAddress(addr).String())
		allowed, err = bech32toValidatorAddresses(allowValidators)
		suite.Require().NoError(err)

		delegateLimit = sdk.NewCoin(baseDenom, sdk.NewInt(10000))
		break
	}

	switch msg {
	case "delegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, &delegateLimit)
		break
	case "undelegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, &delegateLimit)
		break
	case "redelegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, &delegateLimit)
		break
	}

	t := time.Now().AddDate(2, 0, 0).UTC()
	timeoutTimestamp := uint64(10)

	grantMsg, _ := types.NewMsgAuthzGrant(zoneId, grantee, controllerAddr, authorization, t, timeoutTimestamp)

	return *grantMsg
}

func (suite *KeeperTestSuite) getAuthzGrantMsg(msg string, grantee, hostAddr sdk.AccAddress) authz.MsgGrant {
	var authorization authz.Authorization
	var allowed []sdk.ValAddress
	var denied []sdk.ValAddress
	var allowValidators []string
	var delegateLimit sdk.Coin
	var err error

	addr := suite.GenRandomAddress()

	switch msg {
	case "send":
		spendLimit := sdk.NewCoins(sdk.NewCoin(baseDenom, sdk.NewInt(10000)))
		authorization = banktypes.NewSendAuthorization(spendLimit)
		break
	case "generic":
		msgType := ""
		authorization = authz.NewGenericAuthorization(msgType)
		break
	case "delegate", "unbond", "redelegate", "undelegate":
		allowValidators = append(allowValidators, sdk.ValAddress(addr).String())
		allowed, err = bech32toValidatorAddresses(allowValidators)
		suite.Require().NoError(err)

		delegateLimit = sdk.NewCoin(baseDenom, sdk.NewInt(10000))
		break
	}

	switch msg {
	case "delegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, &delegateLimit)
		break
	case "undelegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, &delegateLimit)
		break
	case "redelegate":
		authorization, _ = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, &delegateLimit)
		break
	}

	t := time.Now().AddDate(2, 0, 0).UTC()

	grantMsg, err := authz.NewMsgGrant(hostAddr, grantee, authorization, t)

	return *grantMsg
}

func (suite *KeeperTestSuite) mintCoin(ctx sdk.Context, app *app.NovaApp, denom string, amount sdk.Int, addr sdk.AccAddress) {
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	err := app.MintKeeper.MintCoins(ctx, coins)
	suite.Require().NoError(err)

	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins)
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

func (suite *KeeperTestSuite) mintCoinToModule(ctx sdk.Context, app *app.NovaApp, denom string, amount sdk.Int, module string) {
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	err := app.MintKeeper.MintCoins(ctx, coins)
	suite.Require().NoError(err)

	err = app.BankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, module, coins)
	suite.Require().NoError(err)
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

func (suite *KeeperTestSuite) setValidator() string {
	zone, ok := suite.chainA.GetApp().IcaControlKeeper.GetRegisteredZone(suite.chainA.GetContext(), zoneId)
	suite.Require().True(ok)

	validatorAddr := suite.chainB.App.StakingKeeper.GetValidators(suite.chainB.GetContext(), 1)[0].OperatorAddress
	zone.ValidatorAddress = validatorAddr
	suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), &zone)

	return validatorAddr
}

func (suite *KeeperTestSuite) TestChangeRegisteredZone() {
	suite.InitICA()
	zone := suite.setZone(3)
	hostAddr := suite.setHostAddr(zoneId)
	valAddr := suite.setValidator()

	tcs := []struct {
		name   string
		zoneId string
		zone   types.RegisteredZone
		msg    types.MsgChangeRegisteredZone
		result types.RegisteredZone
		err    bool
	}{
		{
			name:   "should set zone",
			zoneId: "osmo",
			zone:   zone[1],
			msg: types.MsgChangeRegisteredZone{
				ZoneId: "osmo",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        "uosmo",
				Decimal:          6,
			},
			result: types.RegisteredZone{
				ZoneId: "osmo",
				IcaConnectionInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				SnDenom:          "snuosmo",
				BaseDenom:        "uosmo",
				Decimal:          6,
			},
			err: false,
		},
		{
			name:   "fail case - denom is already exist",
			zoneId: zoneId,
			zone:   zone[0],
			msg: types.MsgChangeRegisteredZone{
				ZoneId: zoneId,
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        baseDenom,
				Decimal:          baseDecimal,
			},
			err: true,
		},
		{
			name:   "fail case - invalid connection id",
			zoneId: "juno",
			zone:   zone[0],
			msg: types.MsgChangeRegisteredZone{
				ZoneId: "juno",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: "connection-100",
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        "ujuno",
				Decimal:          baseDecimal,
			},
			err: true,
		},
		{
			name:   "fail case - invalid port id",
			zoneId: "juno",
			zone:   zone[0],
			msg: types.MsgChangeRegisteredZone{
				ZoneId: "juno",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       "portId",
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        "ujuno",
				Decimal:          baseDecimal,
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), &tc.zone)

			suite.chainA.GetApp().IcaControlKeeper.SetControllerAddr(suite.chainA.GetContext(), tc.zoneId, tc.msg.IcaAccount.ControllerAddress)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.GetApp().IcaControlKeeper)
			_, err := msgServer.ChangeRegisteredZone(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				result, ok := suite.chainA.GetApp().IcaControlKeeper.GetRegisteredZone(suite.chainA.GetContext(), tc.zoneId)
				suite.Require().True(ok)

				suite.Require().Equal(tc.result, result)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterZoneMsg() {
	suite.InitICA()
	hostAddr := suite.setHostAddr(zoneId)
	valAddr := suite.setValidator()

	tcs := []struct {
		name   string
		zoneId string
		msg    types.MsgRegisterZone
		result types.RegisteredZone
		err    bool
	}{
		{
			name:   "should set zone",
			zoneId: "osmo",
			msg: types.MsgRegisterZone{
				ZoneId: "osmo",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        "uosmo",
				Decimal:          6,
			},
			result: types.RegisteredZone{
				ZoneId: "osmo",
				IcaConnectionInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				SnDenom:          "snuosmo",
				BaseDenom:        "uosmo",
				Decimal:          6,
			},
			err: false,
		},
		{
			name:   "fail case - denom is already exist",
			zoneId: zoneId,
			msg: types.MsgRegisterZone{
				ZoneId: zoneId,
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				BaseDenom:        baseDenom,
				Decimal:          baseDecimal,
			},
			result: types.RegisteredZone{
				ZoneId: zoneId,
				IcaConnectionInfo: &types.IcaConnectionInfo{
					ConnectionId: suite.icaPath.EndpointA.ConnectionID,
					PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       hostAddr.String(),
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: suite.transferPath.EndpointA.ChannelID,
					PortId:    suite.transferPath.EndpointA.ChannelConfig.PortID,
				},
				ValidatorAddress: valAddr,
				SnDenom:          baseSnDenom,
				BaseDenom:        baseDenom,
				Decimal:          baseDecimal,
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {

			suite.chainA.GetApp().IcaControlKeeper.SetControllerAddr(suite.chainA.GetContext(), tc.zoneId, tc.msg.IcaAccount.ControllerAddress)

			msgServer := keeper.NewMsgServerImpl(suite.chainA.GetApp().IcaControlKeeper)
			_, err := msgServer.RegisterZone(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestIcaDelegate() {
	suite.InitICA()
	_ = suite.setValidator()
	randAddr := suite.GenRandomAddress()
	hostAddr := suite.setHostAddr(zoneId)
	ibcDenom := suite.chainA.App.IcaControlKeeper.GetIBCHashDenom(suite.transferPath.EndpointA.ChannelConfig.PortID, suite.transferPath.EndpointA.ChannelID, baseDenom)
	record := galtypes.DepositRecord{
		ZoneId:    zoneId,
		Depositor: randAddr.String(),
		Records: []*galtypes.DepositRecordContent{
			{
				Claimer: randAddr.String(),
				Amount: &sdk.Coin{
					Amount: sdk.NewInt(10000),
					Denom:  ibcDenom,
				},
				State: galtypes.DepositSuccess,
			},
		},
	}
	suite.chainA.GetApp().GalKeeper.SetDepositRecord(suite.chainA.GetContext(), &record)
	tcs := []struct {
		name string
		msg  types.MsgIcaDelegate
		err  bool
	}{
		{
			name: "success",
			msg: types.MsgIcaDelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			err: false,
		},
		{
			name: "fail case 1 - zone not found",
			msg: types.MsgIcaDelegate{
				ZoneId:            "test",
				ControllerAddress: baseOwnerAcc.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			err: true,
		},
		{
			name: "fail case 2 - invalid controller address",
			msg: types.MsgIcaDelegate{
				ZoneId:            zoneId,
				ControllerAddress: randAddr.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			err: true,
		},
		{
			name: "fail case 3 - invalid controller address",
			msg: types.MsgIcaDelegate{
				ZoneId:            zoneId,
				ControllerAddress: "test",
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.chainA.GetApp().GalKeeper.IsValidDelegateVersion(suite.chainA.GetContext(), zoneId, 0)
			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, sdk.NewInt(10000), hostAddr)
			exeCtxA := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.GetApp().IcaControlKeeper)
			_, err := msgServer.IcaDelegate(sdk.WrapSDKContext(exeCtxA), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.icaRelay(exeCtxA)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestIcaUndelegate() {
	suite.InitICA()
	valAddr := suite.setValidator()
	hostAddr := suite.setHostAddr(zoneId)
	randAddr := suite.GenRandomAddress()

	suite.chainA.GetApp().GalKeeper.IsValidUndelegateVersion(suite.chainA.GetContext(), zoneId, 0)
	suite.mintCoin(suite.chainA.GetContext(), suite.chainA.GetApp(), baseDenom, sdk.NewInt(10000), hostAddr)
	tcs := []struct {
		name           string
		msg            types.MsgIcaUndelegate
		delegateAmount sdk.Coin
		err            bool
	}{
		{
			name: "success",
			msg: types.MsgIcaUndelegate{
				ZoneId:            zoneId,
				ControllerAddress: baseOwnerAcc.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			delegateAmount: sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
			err:            false,
		},
		{
			name: "fail case 1 - zone not found",
			msg: types.MsgIcaUndelegate{
				ZoneId:            "test",
				ControllerAddress: baseOwnerAcc.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			delegateAmount: sdk.NewCoin(baseDenom, sdk.NewInt(0)),
			err:            true,
		},
		{
			name: "fail case 2 - invalid controller address",
			msg: types.MsgIcaUndelegate{
				ZoneId:            zoneId,
				ControllerAddress: randAddr.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			delegateAmount: sdk.NewCoin(baseDenom, sdk.NewInt(0)),
			err:            true,
		},
		{
			name: "fail case 3 - undelegate 요청 금액이 delegate 금액보다 많음",
			msg: types.MsgIcaUndelegate{
				ZoneId:            zoneId,
				ControllerAddress: randAddr.String(),
				Amount:            sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:  uint64(10),
			},
			delegateAmount: sdk.NewCoin(baseDenom, sdk.NewInt(1000)),
			err:            true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// delegate
			delegation := stakingtypes.MsgDelegate{
				DelegatorAddress: hostAddr.String(),
				ValidatorAddress: valAddr,
				Amount:           tc.delegateAmount,
			}

			suite.mintCoin(suite.chainB.GetContext(), suite.chainB.GetApp(), baseDenom, tc.delegateAmount.Amount, hostAddr)

			bmsgServer := stakingkeeper.NewMsgServerImpl(*suite.chainB.App.StakingKeeper)
			_, err := bmsgServer.Delegate(sdk.WrapSDKContext(suite.chainB.GetContext()), &delegation)
			suite.Require().NoError(err)

			exeCtxA := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.IcaControlKeeper)
			_, err = msgServer.IcaUndelegate(sdk.WrapSDKContext(exeCtxA), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.icaRelay(exeCtxA)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestIcaAutoStaking() {
}

func (suite *KeeperTestSuite) TestIcaTransfer() {
	suite.InitICA()
	randAddr := suite.GenRandomAddress()

	tcs := []struct {
		name string
		msg  types.MsgIcaTransfer
		err  bool
	}{
		{
			name: "success",
			msg: types.MsgIcaTransfer{
				ZoneId:               zoneId,
				ControllerAddress:    baseOwnerAcc.String(),
				ReceiverAddress:      baseOwnerAcc.String(),
				IcaTransferPortId:    suite.transferPath.EndpointB.ChannelConfig.PortID,
				IcaTransferChannelId: suite.transferPath.EndpointB.ChannelID,
				Amount:               sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:     uint64(10),
			},
			err: false,
		},
		{
			name: "fail case 1 - zone not found",
			msg: types.MsgIcaTransfer{
				ZoneId:               "test",
				ControllerAddress:    baseOwnerAcc.String(),
				ReceiverAddress:      baseOwnerAcc.String(),
				IcaTransferPortId:    suite.transferPath.EndpointB.ChannelConfig.PortID,
				IcaTransferChannelId: suite.transferPath.EndpointB.ChannelID,
				Amount:               sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:     uint64(10),
			},
			err: true,
		},
		{
			name: "fail case 2 - invalid controller address",
			msg: types.MsgIcaTransfer{
				ZoneId:               zoneId,
				ControllerAddress:    randAddr.String(),
				ReceiverAddress:      baseOwnerAcc.String(),
				IcaTransferPortId:    suite.transferPath.EndpointB.ChannelConfig.PortID,
				IcaTransferChannelId: suite.transferPath.EndpointB.ChannelID,
				Amount:               sdk.NewCoin(baseDenom, sdk.NewInt(10000)),
				TimeoutTimestamp:     uint64(10),
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.IcaControlKeeper)
			_, err := msgServer.IcaTransfer(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.msg)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAuthzGrant() {
	suite.InitICA()
	hostAddr := suite.setHostAddr(zoneId)
	granteeAddr1 := acc1.GetAddress()
	granteeAddr2 := acc2.GetAddress()
	granteeAddr3 := acc3.GetAddress()
	randAddr := suite.GenRandomAddress()

	suite.chainB.GetApp().AccountKeeper.SetAccount(suite.chainB.GetContext(), acc1)
	suite.chainB.GetApp().AccountKeeper.SetAccount(suite.chainB.GetContext(), acc2)
	suite.chainB.GetApp().AccountKeeper.SetAccount(suite.chainB.GetContext(), acc3)

	tcs := []struct {
		name    string
		granter sdk.AccAddress
		grantee sdk.AccAddress
		msg     types.MsgIcaAuthzGrant
		result  string
		err     bool
	}{
		{
			name:    "success - send",
			granter: hostAddr,
			grantee: granteeAddr1,
			msg:     suite.getGrantMsg("send", zoneId, granteeAddr1.String(), baseOwnerAcc),
			result:  sdk.MsgTypeURL(&banktypes.MsgSend{}),
			err:     false,
		},
		{
			name:    "success - delegate",
			granter: hostAddr,
			grantee: granteeAddr2,
			msg:     suite.getGrantMsg("delegate", zoneId, granteeAddr2.String(), baseOwnerAcc),
			result:  sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			err:     false,
		},
		{
			name:    "success - undelegate",
			granter: hostAddr,
			grantee: granteeAddr3,
			msg:     suite.getGrantMsg("undelegate", zoneId, granteeAddr3.String(), baseOwnerAcc),
			result:  sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			err:     false,
		},
		{
			name:    "fail case 1 - zone not found",
			granter: hostAddr,
			grantee: granteeAddr1,
			msg:     suite.getGrantMsg("delegate", "test", granteeAddr1.String(), baseOwnerAcc),
			err:     true,
		},
		{
			name:    "fail - invalid controller address",
			granter: hostAddr,
			grantee: granteeAddr1,
			msg:     suite.getGrantMsg("delegate", zoneId, granteeAddr1.String(), randAddr),
			err:     true,
		},
	}

	for _, tc := range tcs {
		exeCtxA := suite.chainA.GetContext()
		msgServer := keeper.NewMsgServerImpl(suite.chainA.App.IcaControlKeeper)
		_, err := msgServer.IcaAuthzGrant(sdk.WrapSDKContext(exeCtxA), &tc.msg)
		if tc.err {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)

			suite.icaRelay(exeCtxA)

			auth := suite.chainB.GetApp().AuthzKeeper.GetAuthorizations(suite.chainB.GetContext(), tc.grantee, tc.granter)
			suite.Equal(tc.result, auth[0].MsgTypeURL())
		}
	}
}

func (suite *KeeperTestSuite) TestIcaAuthzRevoke() {
	// zone not found
	suite.InitICA()
	hostAddr := suite.setHostAddr(zoneId)
	granteeAddr := acc1.GetAddress()
	randAddr := suite.GenRandomAddress()

	suite.chainB.GetApp().AccountKeeper.SetAccount(suite.chainB.GetContext(), acc1)

	tcs := []struct {
		name    string
		granter sdk.AccAddress
		grantee sdk.AccAddress
		msg     types.MsgIcaAuthzRevoke
		err     bool
	}{
		{
			name:    "send",
			granter: hostAddr,
			grantee: granteeAddr,
			msg: types.MsgIcaAuthzRevoke{
				ZoneId:            zoneId,
				Grantee:           granteeAddr.String(),
				ControllerAddress: baseOwnerAcc.String(),
				MsgTypeUrl:        sdk.MsgTypeURL(&banktypes.MsgSend{}),
				TimeoutTimestamp:  uint64(10),
			},
			err: false,
		},
		{
			name:    "delegate",
			granter: hostAddr,
			grantee: granteeAddr,
			msg: types.MsgIcaAuthzRevoke{
				ZoneId:            zoneId,
				Grantee:           granteeAddr.String(),
				ControllerAddress: baseOwnerAcc.String(),
				MsgTypeUrl:        sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
				TimeoutTimestamp:  uint64(10),
			},
			err: false,
		},
		{
			name:    "undelegate",
			granter: hostAddr,
			grantee: granteeAddr,
			msg: types.MsgIcaAuthzRevoke{
				ZoneId:            zoneId,
				Grantee:           granteeAddr.String(),
				ControllerAddress: baseOwnerAcc.String(),
				MsgTypeUrl:        sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
				TimeoutTimestamp:  uint64(10),
			},
			err: false,
		},
		{
			name:    "fail case 1 - zone not found",
			granter: hostAddr,
			grantee: granteeAddr,
			msg: types.MsgIcaAuthzRevoke{
				ZoneId:            "test",
				Grantee:           granteeAddr.String(),
				ControllerAddress: baseOwnerAcc.String(),
				MsgTypeUrl:        sdk.MsgTypeURL(&banktypes.MsgSend{}),
				TimeoutTimestamp:  uint64(10),
			},
			err: true,
		},
		{
			name:    "fail case 2 - invalid controller address",
			granter: hostAddr,
			grantee: granteeAddr,
			msg: types.MsgIcaAuthzRevoke{
				ZoneId:            zoneId,
				Grantee:           granteeAddr.String(),
				ControllerAddress: randAddr.String(),
				MsgTypeUrl:        sdk.MsgTypeURL(&banktypes.MsgSend{}),
				TimeoutTimestamp:  uint64(10),
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			if !tc.err {
				sendMsg := suite.getAuthzGrantMsg(tc.name, granteeAddr, hostAddr)
				suite.chainB.GetApp().AuthzKeeper.Grant(sdk.WrapSDKContext(suite.chainB.GetContext()), &sendMsg)
			}

			exeCtxA := suite.chainA.GetContext()
			msgServer := keeper.NewMsgServerImpl(suite.chainA.App.IcaControlKeeper)
			_, err := msgServer.IcaAuthzRevoke(sdk.WrapSDKContext(exeCtxA), &tc.msg)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.icaRelay(exeCtxA)

				auths := suite.chainB.App.AuthzKeeper.GetAuthorizations(suite.chainB.GetContext(), tc.grantee, tc.granter)
				suite.Require().Nil(auths)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterControllerAddress() {
	managerAddr := suite.GenRandomAddress()
	controllerAddr := suite.GenRandomAddress()

	// set params
	suite.chainA.GetApp().IcaControlKeeper.SetParams(suite.chainA.GetContext(), types.Params{ControllerKeyManager: []string{managerAddr.String()}})

	msg := types.MsgRegisterControllerAddr{
		ZoneId:            zoneId,
		ControllerAddress: controllerAddr.String(),
		FromAddress:       managerAddr.String(),
	}

	exeCtxA := suite.chainA.GetContext()
	msgServer := keeper.NewMsgServerImpl(suite.chainA.App.IcaControlKeeper)
	_, err := msgServer.RegisterControllerAddress(sdk.WrapSDKContext(exeCtxA), &msg)
	suite.Require().NoError(err)

	controllerInfo := suite.chainA.GetApp().IcaControlKeeper.GetControllerAddr(suite.chainA.GetContext(), zoneId)

	suite.Require().Equal(controllerInfo.ControllerAddress, controllerAddr.String())
}
