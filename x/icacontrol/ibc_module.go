package icacontrol

import (
	distributiontype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	proto "github.com/gogo/protobuf/proto"

	"github.com/Carina-labs/nova/x/icacontrol/keeper"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	return im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID))
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID string,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	connId, err := im.keeper.GetConnectionId(ctx, portID)
	if err != nil {
		return err
	}

	hostAddr, ok := im.keeper.IcaControllerKeeper.GetInterchainAccountAddress(ctx, connId, portID)
	if !ok {
		return types.ErrNotFoundHostAddr
	}

	zoneInfo, ok := im.keeper.GetRegisterZoneForPortId(ctx, portID)
	if !ok {
		return types.ErrNotFoundZoneInfo
	}

	zoneInfo.IcaAccount.HostAddress = hostAddr
	im.keeper.RegisterZone(ctx, zoneInfo)
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is successfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement("cannot receive packet via interchain accounts authentication module")
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 packet acknowledgement: %v", err)
	}
	txMsgData := &sdk.TxMsgData{}
	if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 tx message data: %v", err)
	}

	if !ack.Success() {
		if err := im.keeper.HandleAckFail(ctx, packet); err != nil {
			return err
		}
		im.keeper.Logger(ctx).Error("ICA ack result is fail", "receive packet", packet)
	}

	switch len(txMsgData.Data) {
	case 1: // Delegate, Undelegate, IcaWithdraw
		response, err := im.keeper.HandleAckMsgData(ctx, packet, txMsgData.Data[0])
		if err != nil {
			return err
		}
		im.keeper.Logger(ctx).Info("message response in ICS-27 packet response", "response", response)

		return nil
	case 2: // AutoStaking
		if txMsgData.Data[0].MsgType == sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{}) &&
			txMsgData.Data[1].MsgType == sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}) {
			response, err := im.keeper.HandleAckMsgData(ctx, packet, txMsgData.Data[0])
			if err != nil {
				return err
			}
			im.keeper.Logger(ctx).Info("message response in ICS-27 packet response", "response", response)
		}

		return nil
	default:
		ctx.Logger().Debug("Unknown ICA msg", "receive packet", packet)
		return nil
	}
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.keeper.HandleAckFail(ctx, packet)
}

// NegotiateAppVersion implements the IBCModule interface
func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (string, error) {
	return "", nil
}