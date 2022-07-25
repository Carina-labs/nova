package keeper

import (
	"errors"

	"github.com/Carina-labs/nova/x/ibcstaking/types"
	proto "github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcaccounttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

func (k *Keeper) HandleMsgData(ctx sdk.Context, packet channeltypes.Packet, msgData *sdk.MsgData) (string, error) {
	switch msgData.MsgType {
	case sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}): // delegate
		var data ibcaccounttypes.InterchainAccountPacketData
		if err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
			return "", sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal packet data: %s", err.Error())
		}
		packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
		if err != nil {
			return "", err
		}

		delegateMsg, ok := packetData[0].(*stakingtypes.MsgDelegate)
		if !ok {
			return "", err
		}

		ctx.Logger().Info("DelegateHandler", "delegateMsg", delegateMsg)
		k.AfterDelegateEnd(ctx, *delegateMsg)
		return "", nil
	case sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}): // undelegate
		msgResponse := &stakingtypes.MsgUndelegateResponse{}
		if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
			return "", sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal send response message: %s", err.Error())
		}
		if msgResponse.String() == "" {
			return "", errors.New("response cannot be nil")
		}

		var data ibcaccounttypes.InterchainAccountPacketData
		if err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
			return "", sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal packet data: %s", err.Error())
		}
		packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
		if err != nil {
			return "", err
		}

		undelegateMsg, ok := packetData[0].(*stakingtypes.MsgUndelegate)
		if !ok {
			return "", err
		}
		k.AfterUndelegateEnd(ctx, *undelegateMsg, msgResponse)
		return msgResponse.String(), nil
	case sdk.MsgTypeURL(&transfertypes.MsgTransfer{}): // withdraw(transfer)
		var data ibcaccounttypes.InterchainAccountPacketData
		err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data)
		if err != nil {
			return "", err
		}

		packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
		if err != nil {
			return "", err
		}

		transferMsg, ok := packetData[0].(*transfertypes.MsgTransfer)
		if !ok {
			return "", err
		}

		k.AfterWithdrawEnd(ctx, *transferMsg)
		return "", nil
	default:
		return "", nil
	}
}

func (k *Keeper) HandleTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	// packet을 해부
	var data ibcaccounttypes.InterchainAccountPacketData
	if err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal packet data: %s", err.Error())
	}

	if data.Data == nil {
		return types.ErrMsgNotNil
	}

	packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
	if err != nil {
		return err
	}

	msgType := CheckPacketType(packetData[0])
	// switch true {
	// case packetData[0].(*stakingtypes.MsgDelegate).Type() == (sdk.MsgTypeURL(&stakingtypes.MsgDelegate{})):
	// 	data := packetData[0].(*stakingtypes.MsgDelegate)

	// 	//delegate fail event
	// 	event := types.EventDelegateFail{
	// 		MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
	// 		DelegatorAddress: data.DelegatorAddress,
	// 		ValidatorAddress: data.ValidatorAddress,
	// 		Amount:           &data.Amount,
	// 	}
	// 	ctx.EventManager().EmitTypedEvent(&event)
	// case packetData[0].(*stakingtypes.MsgUndelegate).Type() == sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}):
	// 	data := packetData[0].(*stakingtypes.MsgUndelegate)

	// 	//undelegate fail event
	// 	event := types.EventUndelegateFail{
	// 		MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
	// 		DelegatorAddress: data.DelegatorAddress,
	// 		ValidatorAddress: data.ValidatorAddress,
	// 		Amount:           &data.Amount,
	// 	}
	// 	ctx.EventManager().EmitTypedEvent(&event)
	// case packetData[0].(*distributiontype.MsgWithdrawDelegatorReward).Type() == sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{}):
	// 	if _, ok := packetData[0].(*stakingtypes.MsgDelegate); !ok {
	// 		return ErrMsgNotFound
	// 	}
	// 	data := packetData[0].(*stakingtypes.MsgDelegate)

	// 	//delegate fail event
	// 	event := types.EventAutostakingFail{
	// 		MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
	// 		DelegatorAddress: data.DelegatorAddress,
	// 		ValidatorAddress: data.ValidatorAddress,
	// 		Amount:           &data.Amount,
	// 	}
	// 	ctx.EventManager().EmitTypedEvent(&event)
	// case packetData[0].(*transfertypes.MsgTransfer).Type() == sdk.MsgTypeURL(&transfertypes.MsgTransfer{}):
	// 	data := packetData[0].(*transfertypes.MsgTransfer)

	// 	//delegate fail event
	// 	event := types.EventTransferFail{
	// 		MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
	// 		SourcePort:       data.SourcePort,
	// 		SourceChannel:    data.SourceChannel,
	// 		Token:            &data.Token,
	// 		Sender:           data.Sender,
	// 		Receiver:         data.Receiver,
	// 		TimeoutHeight:    data.TimeoutHeight.String(),
	// 		TimeoutTimestamp: data.TimeoutTimestamp,
	// 	}
	// 	ctx.EventManager().EmitTypedEvent(&event)

	// default:
	// 	return types.ErrMsgNotFound
	// }

	switch msgType {
	case sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}):
		data := packetData[0].(*stakingtypes.MsgDelegate)

		//delegate fail event
		event := types.EventDelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			DelegatorAddress: data.DelegatorAddress,
			ValidatorAddress: data.ValidatorAddress,
			Amount:           data.Amount,
		}
		ctx.EventManager().EmitTypedEvent(&event)

	case sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}):
		data := packetData[0].(*stakingtypes.MsgUndelegate)

		//undelegate fail event
		event := types.EventUndelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			DelegatorAddress: data.DelegatorAddress,
			ValidatorAddress: data.ValidatorAddress,
			Amount:           data.Amount,
		}

		ctx.EventManager().EmitTypedEvent(&event)

	case sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{}):
		if _, ok := packetData[1].(*stakingtypes.MsgDelegate); !ok {
			return types.ErrMsgNotFound
		}
		data := packetData[1].(*stakingtypes.MsgDelegate)

		//delegate fail event
		event := types.EventAutostakingFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			DelegatorAddress: data.DelegatorAddress,
			ValidatorAddress: data.ValidatorAddress,
			Amount:           data.Amount,
		}
		ctx.EventManager().EmitTypedEvent(&event)

	case sdk.MsgTypeURL(&transfertypes.MsgTransfer{}):
		data := packetData[0].(*transfertypes.MsgTransfer)

		//delegate fail event
		event := types.EventTransferFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&transfertypes.MsgTransfer{}),
			SourcePort:       data.SourcePort,
			SourceChannel:    data.SourceChannel,
			Token:            data.Token,
			Sender:           data.Sender,
			Receiver:         data.Receiver,
			TimeoutHeight:    data.TimeoutHeight.String(),
			TimeoutTimestamp: data.TimeoutTimestamp,
		}
		ctx.EventManager().EmitTypedEvent(&event)

	default:
		return types.ErrMsgNotFound
	}

	return nil
}

func CheckPacketType(data sdk.Msg) string {
	// delegate
	if _, ok := data.(*stakingtypes.MsgDelegate); ok {
		return sdk.MsgTypeURL(&stakingtypes.MsgDelegate{})
	}

	// undelegate
	if _, ok := data.(*stakingtypes.MsgUndelegate); ok {
		return sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{})
	}

	// autostaking
	if _, ok := data.(*distributiontype.MsgWithdrawDelegatorReward); ok {
		return sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{})
	}

	// withdraw
	if _, ok := data.(*transfertypes.MsgTransfer); ok {
		return sdk.MsgTypeURL(&transfertypes.MsgTransfer{})
	}

	return ""
}
