package keeper_test

import (
	"errors"
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"time"

	galtypes "github.com/Carina-labs/nova/x/gal/types"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcaccounttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"
)

func (suite *KeeperTestSuite) SetVersion(msgType string, zoneInfo types.RegisteredZone) {

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
func (suite *KeeperTestSuite) SetMsgs(msgType string, zoneInfo []types.RegisteredZone) ibcaccounttypes.InterchainAccountPacketData {
	var msgs []sdk.Msg
	switch msgType {
	case "delegate":
		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: "test",
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}

		msgs = append(msgs, delegateMsg)
		break
	case "undelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: "test",
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}

		msgs = append(msgs, undelegateMsg)
		break
	case "autostaking":
		distMsg := &distributiontype.MsgWithdrawDelegatorReward{
			DelegatorAddress: zoneInfo[0].IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
		}
		msgs = append(msgs, distMsg)

		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: zoneInfo[0].IcaAccount.HostAddress,
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
		}
		msgs = append(msgs, delegateMsg)
		break
	case "transfer":
		transferMsg := &transfertypes.MsgTransfer{
			SourcePort:    "transfer",
			SourceChannel: "channel-0",
			Token:         sdk.NewCoin("uatom", sdk.NewInt(1000)),
			Sender:        "test",
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
	default:
		return ibcaccounttypes.InterchainAccountPacketData{
			Type: ibcaccounttypes.EXECUTE_TX,
			Data: nil,
		}
	}

	data, err := ibcaccounttypes.SerializeCosmosTx(suite.App.AppCodec(), msgs)
	suite.NoError(err)

	icapacket := ibcaccounttypes.InterchainAccountPacketData{
		Type: ibcaccounttypes.EXECUTE_TX,
		Data: data,
	}

	return icapacket
}
func (suite *KeeperTestSuite) GetMsgs(msg string) *sdk.MsgData {
	var data sdk.MsgData
	switch msg {
	case "delegate":
		res := &stakingtypes.MsgDelegateResponse{}
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgDelegate{
			DelegatorAddress: "test",
			ValidatorAddress: "validator",
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
			DelegatorAddress: "test",
			ValidatorAddress: "validatorAddr",
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
			DelegatorAddress: "delegator",
			ValidatorAddress: "validator",
		})
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
func (suite *KeeperTestSuite) GetPacket(msg string, zone types.RegisteredZone) channeltypes.Packet {
	var msgs []sdk.Msg

	addr := suite.GenRandomAddress()
	switch msg {
	case "delegate":
		delegateMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: addr.String(),
			ValidatorAddress: zone.ValidatorAddress,
		}
		msgs = append(msgs, delegateMsg)
		break
	case "undelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: addr.String(),
			ValidatorAddress: zone.ValidatorAddress,
			Amount:           sdk.NewCoin(zone.BaseDenom, sdk.NewInt(10000)),
		}
		msgs = append(msgs, undelegateMsg)
		break
	case "errUndelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: addr.String(),
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
			DelegatorAddress: addr.String(),
			ValidatorAddress: zone.ValidatorAddress,
		}
		msgs = append(msgs, autostakingMsg)
		break
	}

	data, err := ibcaccounttypes.SerializeCosmosTx(suite.App.AppCodec(), msgs)
	suite.NoError(err)

	icapacket := ibcaccounttypes.InterchainAccountPacketData{
		Type: ibcaccounttypes.EXECUTE_TX,
		Data: data,
	}

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         "icacontroller-" + zone.IcaAccount.ControllerAddress,
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
	zone := suite.setZone(1)
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &zone[0])

	tcs := []struct {
		name   string
		args   *sdk.MsgData
		packet channeltypes.Packet
		expect string
		err    error
	}{
		{
			name:   "delegate",
			args:   suite.GetMsgs("delegate"),
			packet: suite.GetPacket("delegate", zone[0]),
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate",
			args:   suite.GetMsgs("undelegate"),
			packet: suite.GetPacket("undelegate", zone[0]),
			expect: "completion_time:<seconds:10000 nanos:100000 > ",
			err:    nil,
		},
		{
			name:   "transfer",
			args:   suite.GetMsgs("transfer"),
			packet: suite.GetPacket("transfer", zone[0]),
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate error",
			args:   suite.GetMsgs("errUndelegate"),
			packet: suite.GetPacket("errUndelegate", zone[0]),
			expect: "",
			err:    errors.New("response cannot be nil"),
		},
		{
			name:   "autostaking",
			args:   suite.GetMsgs("autostaking"),
			packet: suite.GetPacket("autostaking", zone[0]),
			expect: "",
			err:    nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetVersion(tc.name, zone[0])
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
	zone := suite.setZone(1)
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &zone[0])

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         "icacontroller-" + zone[0].IcaAccount.ControllerAddress,
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
		args   ibcaccounttypes.InterchainAccountPacketData
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
			err:    true,
		},
		{
			name:   "nil",
			args:   suite.SetMsgs("nil", zone),
			expect: types.ErrMsgNotNil,
			err:    true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetVersion(tc.name, zone[0])
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
