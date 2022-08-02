package keeper

import (
	"errors"

	proto "github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
