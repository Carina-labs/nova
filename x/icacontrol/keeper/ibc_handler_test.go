package keeper_test

import (
	"errors"
	"time"

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

	case "undelegate":
		undelegateMsg := &stakingtypes.MsgUndelegate{
			DelegatorAddress: "test",
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}

		msgs = append(msgs, undelegateMsg)
	case "autostaking":
		autostakingMsg := &stakingtypes.MsgDelegate{
			DelegatorAddress: "test",
			ValidatorAddress: zoneInfo[0].ValidatorAddress,
			Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
		}

		msgs = append(msgs, autostakingMsg)
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

	case "unkowntype":
		bankMsg := &banktypes.MsgSend{
			FromAddress: "from_address",
			ToAddress:   "to_address",
			Amount: sdk.NewCoins(
				sdk.NewCoin("uatom", sdk.NewInt(1000)),
			),
		}

		msgs = append(msgs, bankMsg)
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
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgDelegate{})
		resMsg, err := proto.Marshal(res)
		suite.NoError(err)
		data.Data = resMsg
		return &data
	case "undelegate":
		res := &stakingtypes.MsgUndelegateResponse{
			CompletionTime: time.Unix(10000, 100000),
		}
		data.MsgType = sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{})
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
	case "bank":
		res := &banktypes.MsgSendResponse{}
		data.MsgType = sdk.MsgTypeURL(&banktypes.MsgSend{})
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
func (suite *KeeperTestSuite) TestHandleMsgData() {
	zone := suite.setZone(1)
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &zone[0])

	var msgs []sdk.Msg
	undelegateMsg := &stakingtypes.MsgUndelegate{
		DelegatorAddress: "test",
		ValidatorAddress: zone[0].ValidatorAddress,
		Amount:           sdk.NewCoin("uatom", sdk.NewInt(1000)),
	}

	msgs = append(msgs, undelegateMsg)

	data, err := ibcaccounttypes.SerializeCosmosTx(suite.App.AppCodec(), msgs)
	suite.NoError(err)

	icapacket := ibcaccounttypes.InterchainAccountPacketData{
		Type: ibcaccounttypes.EXECUTE_TX,
		Data: data,
	}

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         "icacontroller-" + zone[0].IcaAccount.ControllerAddress,
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
			packet: packetData,
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate",
			args:   suite.GetMsgs("undelegate"),
			packet: packetData,
			expect: "completion_time:<seconds:10000 nanos:100000 > ",
			err:    nil,
		},
		{
			name:   "transfer",
			args:   suite.GetMsgs("transfer"),
			packet: packetData,
			expect: "",
			err:    nil,
		},
		{
			name:   "undelegate error",
			args:   suite.GetMsgs("errUndelegate"),
			packet: packetData,
			expect: "",
			err:    errors.New("response cannot be nil"),
		},
		{
			name:   "bank",
			args:   suite.GetMsgs("bank"),
			packet: packetData,
			expect: "",
			err:    nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
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
			packetData.Data = tc.args.GetBytes()
			err := suite.App.IcaControlKeeper.HandleTimeoutPacket(suite.Ctx, packetData)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
