package keeper_test

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"
)

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
	suite.App.IntertxKeeper.RegisterZone(suite.Ctx, &zone[0])

	packetData := channeltypes.Packet{
		Sequence:           1,
		SourcePort:         "icacontroller-" + zone[0].IcaAccount.OwnerAddress,
		SourceChannel:      "channel-0",
		DestinationPort:    "icahost",
		DestinationChannel: "channel-0",
		Data:               []byte(""),
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
			res, err := suite.App.IntertxKeeper.HandleMsgData(suite.Ctx, tc.packet, tc.args)
			suite.NoError(err)

			if tc.err == nil {
				suite.Require().Equal(res, tc.expect)
			} else {
				suite.Require().Error(tc.err)
			}
		})
	}
}
