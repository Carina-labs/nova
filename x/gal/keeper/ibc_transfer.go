package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

type IBCTransferOption struct {
	SourcePort    string
	SourceChannel string
	Token         sdk.Coin
	Sender        string
	Receiver      string
}

// TransferToTargetZone transfers user's asset to target zone(Host chain)
// using IBC transfer.
func (k Keeper) TransferToTargetZone(ctx sdk.Context, option *IBCTransferOption) error {
	goCtx := sdk.WrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(option.Sender)
	if err != nil {
		return err
	}

	_, err = k.ibcTransferKeeper.Transfer(goCtx,
		&transfertypes.MsgTransfer{
			SourcePort:    option.SourcePort,
			SourceChannel: option.SourceChannel,
			Token:         option.Token,
			Sender:        sender.String(),
			Receiver:      option.Receiver,
			TimeoutHeight: ibcclienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		},
	)

	return err
}
