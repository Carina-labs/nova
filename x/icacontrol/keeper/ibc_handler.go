package keeper

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/gogo/protobuf/proto"

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

		versionInfo := k.GetAutoStakingVersion(ctx, zone.ZoneId)
		currentVersion := versionInfo.CurrentVersion

		versionInfo.Record[currentVersion] = &types.IBCTrace{
			Version: currentVersion,
			Height:  uint64(ctx.BlockHeight()),
			State:   types.IcaSuccess,
		}

		// set withdraw version
		nextVersion := versionInfo.CurrentVersion + 1
		versionInfo.CurrentVersion = nextVersion
		versionInfo.Record[nextVersion] = &types.IBCTrace{
			Version: nextVersion,
			State:   types.IcaPending,
		}

		k.SetAutoStakingVersion(ctx, zone.ZoneId, versionInfo)
		return "", nil
	default:
		return "", nil
	}
}

func (k *Keeper) HandleAckFail(ctx sdk.Context, packet channeltypes.Packet) error {
	var data ibcaccounttypes.InterchainAccountPacketData
	if err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal packet data: %s", err.Error())
	}

	if data.Data == nil {
		return types.ErrInvalidMsg
	}

	packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data)
	if err != nil {
		return err
	}

	switch packetData[0].(type) {
	case *stakingtypes.MsgDelegate:
		msgData := packetData[0].(*stakingtypes.MsgDelegate)
		ctx.Logger().Error("HandleAckFail", "MsgType", "MsgDelegate", "Data", msgData)

		//delegate fail event
		event := types.EventDelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			DelegatorAddress: msgData.DelegatorAddress,
			ValidatorAddress: msgData.ValidatorAddress,
			Amount:           msgData.Amount,
		}
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			ctx.Logger().Error("HandleAckFail", "MsgType", "MsgDelegate", "EventError", err)
		}

		k.AfterDelegateFail(ctx, *msgData)
	case *stakingtypes.MsgUndelegate:
		msgData := packetData[0].(*stakingtypes.MsgUndelegate)
		ctx.Logger().Error("HandleAckFail", "MsgType", "MsgUndelegate", "Data", msgData)

		//undelegate fail event
		event := types.EventUndelegateFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			DelegatorAddress: msgData.DelegatorAddress,
			ValidatorAddress: msgData.ValidatorAddress,
			Amount:           msgData.Amount,
		}

		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			ctx.Logger().Error("HandleAckFail", "MsgType", "MsgUndelegate", "EventError", err)
		}

		k.AfterUndelegateFail(ctx, *msgData)
	case *distributiontype.MsgWithdrawDelegatorReward:
		if len(packetData) != 2 {
			return types.ErrMsgNotFound
		}

		if _, ok := packetData[1].(*stakingtypes.MsgDelegate); !ok {
			return types.ErrMsgNotFound
		}

		msgData := packetData[1].(*stakingtypes.MsgDelegate)
		ctx.Logger().Error("HandleAckFail", "MsgType", "MsgWithdrawDelegatorReward", "Data", msgData)

		//delegate fail event
		event := types.EventAutostakingFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
			DelegatorAddress: msgData.DelegatorAddress,
			ValidatorAddress: msgData.ValidatorAddress,
			Amount:           msgData.Amount,
		}
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			ctx.Logger().Error("HandleAckFail", "MsgType", "MsgWithdrawDelegatorReward", "EventError", err)
		}

		zoneInfo := k.GetRegisteredZoneForValidatorAddr(ctx, msgData.ValidatorAddress)
		versionInfo := k.GetAutoStakingVersion(ctx, zoneInfo.ZoneId)
		currentVersion := versionInfo.CurrentVersion

		versionInfo.Record[currentVersion] = &types.IBCTrace{
			Version: currentVersion,
			Height:  uint64(ctx.BlockHeight()),
			State:   types.IcaFail,
		}

		ctx.Logger().Error("HandleAckFail", "ZoneId", zoneInfo.ZoneId, "AutostakingCurrentVersion", currentVersion, "IcaWithdrawVersionState", versionInfo.Record[currentVersion].State)
		k.SetAutoStakingVersion(ctx, zoneInfo.ZoneId, versionInfo)
	case *transfertypes.MsgTransfer:
		msgData := packetData[0].(*transfertypes.MsgTransfer)

		//icawithdraw fail event
		event := types.EventTransferFail{
			MsgTypeUrl:       sdk.MsgTypeURL(&transfertypes.MsgTransfer{}),
			SourcePort:       msgData.SourcePort,
			SourceChannel:    msgData.SourceChannel,
			Token:            msgData.Token,
			Sender:           msgData.Sender,
			Receiver:         msgData.Receiver,
			TimeoutHeight:    msgData.TimeoutHeight.String(),
			TimeoutTimestamp: msgData.TimeoutTimestamp,
		}
		err = ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			ctx.Logger().Error("HandleAckFail", "MsgType", "MsgTransfer", "EventError", err)
		}

		ctx.Logger().Error("HandleAckFail", "MsgType", "MsgTransfer")
		k.AfterIcaWithdrawFail(ctx, *msgData)
	default:
		ctx.Logger().Error("HandleAckFail", "MsgType", types.ErrMsgNotFound)
		return nil
	}

	return nil
}
