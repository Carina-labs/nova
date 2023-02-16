package keeper_test

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"time"

	galtypes "github.com/Carina-labs/nova/x/gal/types"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"
)

func (suite *KeeperTestSuite) SetVersion(msgType string, zoneInfo *types.RegisteredZone) {
	versionInfo := oracletypes.IBCTrace{
		Version: uint64(0),
		Height:  1,
	}
	suite.App.OracleKeeper.SetOracleVersion(suite.Ctx, zoneInfo.ZoneId, versionInfo)

	switch msgType {
	case "delegate":
		versionInfo := galtypes.VersionState{
			CurrentVersion: 0,
			ZoneId:         zoneId,
			Record: map[uint64]*galtypes.IBCTrace{
				0: {
					Version: 0,
					State:   types.IcaRequest,
				},
			},
		}
		suite.App.GalKeeper.SetDelegateVersion(suite.Ctx, zoneInfo.ZoneId, versionInfo)
		break
	case "undelegate", "undelegate error":
		versionInfo := galtypes.VersionState{
			CurrentVersion: 0,
			ZoneId:         zoneId,
			Record: map[uint64]*galtypes.IBCTrace{
				0: {
					Version: 0,
					State:   types.IcaRequest,
				},
			},
		}
		suite.App.GalKeeper.SetUndelegateVersion(suite.Ctx, zoneInfo.ZoneId, versionInfo)
		break
	case "autostaking":
		versionInfo := types.VersionState{
			CurrentVersion: 0,
			ZoneId:         zoneId,
			Record: map[uint64]*types.IBCTrace{
				0: {
					Version: 0,
					State:   types.IcaRequest,
				},
			},
		}
		suite.App.IcaControlKeeper.SetAutoStakingVersion(suite.Ctx, zoneInfo.ZoneId, versionInfo)
		break
	case "transfer":
		versionInfo := galtypes.VersionState{
			CurrentVersion: 0,
			ZoneId:         zoneId,
			Record: map[uint64]*galtypes.IBCTrace{
				0: {
					Version: 0,
					State:   types.IcaRequest,
				},
			},
		}
		suite.App.GalKeeper.SetWithdrawVersion(suite.Ctx, zoneInfo.ZoneId, versionInfo)
		break
	}
	return
}

func (suite *KeeperTestSuite) SetMsgs(msgType string, zoneInfo *types.RegisteredZone) icatypes.InterchainAccountPacketData {
	var msgs []sdk.Msg
	switch msgType {
	case "delegate":
		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}
		msgs = append(msgs, delegateMsg)
		break
	case "undelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}

		msgs = append(msgs, undelegateMsg)
		break
	case "autostaking":
		distMsg := &distributiontype.MsgWithdrawDelegatorReward{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
		}
		msgs = append(msgs, distMsg)

		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
		}
		msgs = append(msgs, delegateMsg)
		break
	case "transfer":
		transferMsg := &transfertypes.MsgTransfer{
			SourcePort:    "transfer",
			SourceChannel: "channel-0",
			Token:         sdk.NewCoin("uatom", sdk.NewInt(1000)),
			Sender:        zoneInfo.IcaAccount.HostAddress,
			Receiver:      "receiver",
			TimeoutHeight: ibcclienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: uint64(suite.Ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		}

		msgs = append(msgs, transferMsg)
		break
	case "unkowntype":
		bankMsg := &banktypes.MsgSend{
			FromAddress: "from_address",
			ToAddress:   "to_address",
			Amount: sdk.NewCoins(
				sdk.NewCoin("uatom", sdk.NewInt(1000)),
			),
		}

		msgs = append(msgs, bankMsg)
		break
	case "authzGrant":
		spendLimit := sdk.NewCoins(sdk.NewCoin(baseDenom, sdk.NewInt(10000)))
		authMsg := &authz.MsgGrant{
			Granter: zoneInfo.IcaAccount.HostAddress,
			Grantee: zoneInfo.IcaAccount.ControllerAddress,
			Grant: authz.Grant{
				Expiration: time.Now(),
			},
		}
		authMsg.SetAuthorization(banktypes.NewSendAuthorization(spendLimit))
		msgs = append(msgs, authMsg)
		break
	default:
		return icatypes.InterchainAccountPacketData{
			Type: icatypes.EXECUTE_TX,
			Data: nil,
		}
	}
	data, err := icatypes.SerializeCosmosTx(suite.App.AppCodec(), msgs)
	suite.NoError(err)

	icapacket := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	return icapacket
}

func (suite *KeeperTestSuite) GetMsgs(msg string, zoneInfo *types.RegisteredZone) *sdk.MsgData {
	var data sdk.MsgData
	switch msg {
	case "delegate":
		res := &stakingtypes.MsgDelegateResponse{}
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgDelegate{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(10000)),
		})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "undelegate":
		res := &stakingtypes.MsgUndelegateResponse{
			CompletionTime: time.Unix(10000, 100000),
		}
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "transfer":
		res := &transfertypes.MsgTransferResponse{}
		data.MsgType = sdk.MsgTypeURL(&transfertypes.MsgTransfer{})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "autostaking":
		res := &distributiontype.MsgWithdrawDelegatorRewardResponse{}
		data.MsgType = sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{
			DelegatorAddress: zoneInfo.IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo.ValidatorAddress,
		})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "authzGrant":
		spendLimit := sdk.NewCoins(sdk.NewCoin(baseDenom, sdk.NewInt(10000)))
		res := &authz.MsgGrantResponse{}
		grantMsg := &authz.MsgGrant{
			Granter: zoneInfo.IcaAccount.HostAddress,
			Grantee: zoneInfo.IcaAccount.ControllerAddress,
			Grant: authz.Grant{
				Expiration: time.Now(),
			},
		}
		grantMsg.SetAuthorization(banktypes.NewSendAuthorization(spendLimit))
		data.MsgType = sdk.MsgTypeURL(grantMsg)

		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "errUndelegate":
		res := &stakingtypes.MsgUndelegateResponse{}
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	default:
		return &data
	}
}
func (suite *KeeperTestSuite) GetPacket(msg string, zone *types.RegisteredZone) channeltypes.Packet {
	var msgs []sdk.Msg

	switch msg {
	case "delegate":
		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: zone.IcaAccount.HostAddress,
			ValidatorAddress: zone.ValidatorAddress,
		}
		msgs = append(msgs, delegateMsg)
		break
	case "undelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: zone.IcaAccount.HostAddress,
			ValidatorAddress: zone.ValidatorAddress,
			Amount:           sdk.NewCoin(zone.BaseDenom, sdk.NewInt(10000)),
		}
		msgs = append(msgs, undelegateMsg)
		break
	case "errUndelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: zone.IcaAccount.HostAddress,
			ValidatorAddress: zone.ValidatorAddress,
			Amount:           sdk.NewCoin(zone.BaseDenom, sdk.NewInt(10000)),
		}
		msgs = append(msgs, undelegateMsg)
		break
	case "transfer":
		transferMsg := &transfertypes.MsgTransfer{
			SourcePort:       zone.TransferInfo.PortId,
			SourceChannel:    zone.TransferInfo.ChannelId,
			Token:            sdk.NewCoin(zone.BaseDenom, sdk.NewInt(10000)),
			Sender:           zone.IcaAccount.HostAddress,
			Receiver:         zone.IcaAccount.ControllerAddress,
			TimeoutHeight:    ibcclienttypes.NewHeight(0, 0),
			TimeoutTimestamp: uint64(suite.Ctx.BlockTime().UnixNano()),
		}
		msgs = append(msgs, transferMsg)
		break
	case "autostaking":
		autostakingMsg := &distributiontype.MsgWithdrawDelegatorReward{
			DelegatorAddress: zone.IcaAccount.HostAddress,
			ValidatorAddress: zone.ValidatorAddress,
		}
		msgs = append(msgs, autostakingMsg)
		break
	case "authzGrant":
		spendLimit := sdk.NewCoins(sdk.NewCoin(baseDenom, sdk.NewInt(10000)))
		authzMsg := &authz.MsgGrant{
			Granter: zone.IcaAccount.HostAddress,
			Grantee: zone.IcaAccount.ControllerAddress,
			Grant: authz.Grant{
				Expiration: time.Now(),
			},
		}
		authzMsg.SetAuthorization(banktypes.NewSendAuthorization(spendLimit))
		msgs = append(msgs, authzMsg)
		break
	}

	data, err := icatypes.SerializeCosmosTx(suite.App.AppCodec(), msgs)
	suite.NoError(err)

	icapacket := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         icatypes.PortPrefix + zone.IcaAccount.ControllerAddress,
		SourceChannel:      "channel-0",
		DestinationPort:    "icahost",
		DestinationChannel: "channel-0",
		Data:               icapacket.GetBytes(),
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(suite.Ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	}

	return packetData
}

func (suite *KeeperTestSuite) TestHandleMsgData() {
	zone := newBaseRegisteredZone()
	zone.IcaAccount.HostAddress = suite.GenRandomAddress().String()

	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, zone)

	tcs := []struct {
		name   string
		args   *sdk.MsgData
		packet channeltypes.Packet
		expect string
		err    error
	}{
		{
			name:   "delegate",
			args:   suite.GetMsgs("delegate", zone),
			packet: suite.GetPacket("delegate", zone),
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate",
			args:   suite.GetMsgs("undelegate", zone),
			packet: suite.GetPacket("undelegate", zone),
			expect: "completion_time:<seconds:10000 nanos:100000 > ",
			err:    nil,
		},
		{
			name:   "transfer",
			args:   suite.GetMsgs("transfer", zone),
			packet: suite.GetPacket("transfer", zone),
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate error",
			args:   suite.GetMsgs("errUndelegate", zone),
			packet: suite.GetPacket("errUndelegate", zone),
			expect: "",
			err:    errors.New("response cannot be nil"),
		},
		{
			name:   "autostaking",
			args:   suite.GetMsgs("autostaking", zone),
			packet: suite.GetPacket("autostaking", zone),
			expect: "",
			err:    nil,
		},
		{
			name:   "authzGrant",
			args:   suite.GetMsgs("authzGrant", zone),
			packet: suite.GetPacket("authzGrant", zone),
			expect: "",
			err:    nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetVersion(tc.name, zone)
			res, err := suite.App.IcaControlKeeper.HandleAckMsgData(suite.Ctx, tc.packet, tc.args)
			suite.NoError(err)

			if tc.err == nil {
				suite.Require().Equal(res, tc.expect)
			} else {
				suite.Require().Error(tc.err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleTimeoutPacket() {
	zone := newBaseRegisteredZone()
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, zone)

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         icatypes.PortPrefix + zone.IcaAccount.ControllerAddress,
		SourceChannel:      "channel-0",
		DestinationPort:    "icahost",
		DestinationChannel: "channel-0",
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(suite.Ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
	}

	tcs := []struct {
		name   string
		args   icatypes.InterchainAccountPacketData
		expect error
		err    bool
	}{
		{
			name:   "delegate",
			args:   suite.SetMsgs("delegate", zone),
			expect: nil,
			err:    false,
		},
		{
			name:   "undelegate",
			args:   suite.SetMsgs("undelegate", zone),
			expect: nil,
			err:    false,
		},
		{
			name:   "autostaking",
			args:   suite.SetMsgs("autostaking", zone),
			expect: nil,
			err:    false,
		},
		{
			name:   "transfer",
			args:   suite.SetMsgs("transfer", zone),
			expect: nil,
			err:    false,
		},
		{
			name:   "unkowntype",
			args:   suite.SetMsgs("unkowntype", zone),
			expect: types.ErrMsgNotFound,
			err:    false,
		},
		{
			name:   "nil",
			args:   suite.SetMsgs("nil", zone),
			expect: types.ErrInvalidMsg,
			err:    true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetVersion(tc.name, zone)
			packetData.Data = tc.args.GetBytes()
			err := suite.App.IcaControlKeeper.HandleAckFail(suite.Ctx, packetData)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
