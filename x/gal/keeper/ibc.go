package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// TransferToTargetZone transfers user's asset to target zone(Host chain)
// using IBC transfer.
func (k Keeper) TransferToTargetZone(ctx sdk.Context,
	sourcePort, sourceChannel, depositor, icaController string, amt sdk.Coin) error {
	goCtx := sdk.WrapSDKContext(ctx)
	depositorAddr, err := sdk.AccAddressFromBech32(depositor)
	if err != nil {
		return err
	}

	_, err = k.ibcTransferKeeper.Transfer(goCtx,
		&transfertypes.MsgTransfer{
			SourcePort:    sourcePort,
			SourceChannel: sourceChannel,
			Token:         amt,
			Sender:        depositorAddr.String(),
			Receiver:      icaController,
			TimeoutHeight: ibcclienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		},
	)

	return err
}
