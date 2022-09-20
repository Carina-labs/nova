package keeper

import (
	"errors"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	proto "github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcaccounttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

func (k *Keeper) HandleAckMsgData(ctx sdk.Context, packet channeltypes.Packet, msgData *sdk.MsgData) (string, error) {
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
			return "", types.ErrMsgNotFound
		}

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
			return "", types.ErrMsgNotFound
		}
		k.AfterUndelegateEnd(ctx, *undelegateMsg, msgResponse)
		return msgResponse.String(), nil
	case sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{}): // AutoStaking
		var data ibcaccounttypes.InterchainAccountPacketData
		err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data)
		if err != nil {
			return "", err
		}

		packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
		if err != nil {
			return "", err
		}

		rewardMsg, ok := packetData[0].(*distributiontype.MsgWithdrawDelegatorReward)
		if !ok {
			return "", types.ErrMsgNotFound
		}

		zone := k.GetRegisteredZoneForValidatorAddr(ctx, rewardMsg.ValidatorAddress)
		if zone == nil {
			return "", types.ErrMsgNotFound
		}

		version := k.GetAutoStakingVersion(ctx, zone.ZoneId)
		k.SetAutoStakingVersion(ctx, zone.ZoneId, version+1)
		return "", nil
	default:
		return "", nil
	}
}

func (k *Keeper) HandleTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
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

	switch packetData[0].(type) {
	case *stakingtypes.MsgDelegate:
		data := packetData[0].(*stakingtypes.MsgDelegate)

		//delegate fail event
		event := types.EventDelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			DelegatorAddress: data.DelegatorAddress,
			ValidatorAddress: data.ValidatorAddress,
			Amount:           data.Amount,
		}
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			return err
		}

	case *stakingtypes.MsgUndelegate:
		data := packetData[0].(*stakingtypes.MsgUndelegate)

		//undelegate fail event
		event := types.EventUndelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			DelegatorAddress: data.DelegatorAddress,
			ValidatorAddress: data.ValidatorAddress,
			Amount:           data.Amount,
		}

		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			return err
		}

	case *distributiontype.MsgWithdrawDelegatorReward:
		if len(packetData) != 2 {
			return types.ErrMsgNotFound
		}

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
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			return err
		}

	case *transfertypes.MsgTransfer:
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
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			return err
		}

	default:
		return types.ErrMsgNotFound
	}

	return nil
}
