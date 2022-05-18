package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// SubmitTx implements the Msg/SubmitTx interface
func (k Keeper) SendICATx(ctx sdk.Context, owner_address, connection_id string, msgs []sdk.Msg) error {

	portID, err := icatypes.NewControllerPortID(owner_address)
	if err != nil {
		return err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connection_id, portID)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}
	fmt.Println("msgs : ", msgs[0])
	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs)
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := time.Now().Add(time.Minute).UnixNano()
	_, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connection_id, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	return nil
}
