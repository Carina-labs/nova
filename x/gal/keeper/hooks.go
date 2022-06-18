package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	icatypes "github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// Hooks wrapper struct for gal keeper
type Hooks struct {
	k Keeper
}

var _ transfertypes.TransferHooks = Hooks{}
var _ icatypes.ICAHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// AfterTransferEnd records user deposit information.
// It will be used in share token minting process.
func (h Hooks) AfterTransferEnd(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, base_denom string) {
	amt, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		h.k.Logger(ctx).Error("can't cast int to string, str: %s", data.Amount)
		return
	}

	zoneInfo := h.k.interTxKeeper.GetZoneForDenom(ctx, base_denom)
	// zoneID
	coin := sdk.NewInt64Coin(data.Denom, amt.Int64())

	record := &types.DepositRecord{
		ZoneId:  zoneInfo.ZoneName,
		Address: data.Sender,
		Amount:  &coin,
	}
	h.k.RecordDepositAmt(ctx, *record)

	// Delegate events
	ctx.EventManager().EmitTypedEvent(zoneInfo)
	ctx.EventManager().EmitTypedEvent(record)

}

func (h Hooks) AfterDelegateEnd() {
}

func (h Hooks) AfterWithdrawEnd() {
}

func (h Hooks) BeforeUndelegateStart(ctx sdk.Context, zoneId string) {
}

// AfterUndelegateEnd is executed when ICA undelegation request finished.
// 1. It removes undelegation history in store.
// 2. It saves undelegation finish time to store.
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, msg *stakingtypes.MsgUndelegateResponse) {
	zone := h.k.interTxKeeper.GetRegisteredZoneForPortId(ctx, packet.SourcePort)
	h.k.DeleteUndelegateRecords(ctx, zone.ZoneName, UNDELEGATE_REQUEST_ICA)
	h.k.SetWithdrawTime(ctx, zone.ZoneName, WITHDRAW_REQUEST_USER, msg.CompletionTime)
}

func (h Hooks) AfterAutoStakingEnd() {
}
