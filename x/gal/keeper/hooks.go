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
	zoneInfo := h.k.interTxKeeper.GetZoneForDenom(ctx, base_denom)
	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}
  
	// change sender + zoneId
	depositRecord, err := h.k.GetRecordedDepositAmt(ctx, sdk.AccAddress(data.Sender))
	if err != nil {
		return
	}

	record := &types.DepositRecord{
		ZoneId:        depositRecord.ZoneId,
		Address:       depositRecord.Address,
		Amount:        depositRecord.Amount,
		IsTransferred: true,
	}

	if err := h.k.RecordDepositAmt(ctx, record); err != nil {
		panic(err)
	}

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
	h.k.DeleteUndelegateRecords(ctx, zone.ZoneId, UNDELEGATE_REQUEST_ICA)
	h.k.SetWithdrawTime(ctx, zone.ZoneId, WITHDRAW_REQUEST_USER, msg.CompletionTime)
}

func (h Hooks) AfterAutoStakingEnd() {
}
