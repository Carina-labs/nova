package keeper

import (
	"fmt"
	"math/big"

	"github.com/Carina-labs/nova/x/gal/types"
	icatypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
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

// AfterTransferEnd coins user deposit information.
// It will be used in share token minting process.
func (h Hooks) AfterTransferEnd(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, baseDenom string) {
	packetAmount, ok := new(big.Int).SetString(data.Amount, 10)
	if !ok {
		panic(fmt.Sprintf("invalid ibc transfer packet: %v", data))
	}

	zoneInfo := h.k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)
	// if zoneInfo == nil, it may be a test situation.
	if zoneInfo == nil {
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	// change sender + zoneId
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return
	}

	depositRecord, err := h.k.GetRecordedDepositAmt(ctx, sender)
	if err != nil {
		return
	}
	for i, record := range depositRecord.Records {
		// assert amount of record equals to the amount of ibc transfer packet.
		if record.Amount.Amount.BigInt().Cmp(packetAmount) == 0 && !record.IsTransferred {
			depositRecord := &types.DepositRecord{
				Address: depositRecord.Address,
				Records: []*types.DepositRecordContent{
					{
						ZoneId:        record.ZoneId,
						Amount:        record.Amount,
						IsTransferred: true,
					},
				},
			}

			if err := h.k.MarkRecordTransfer(ctx, depositRecord.Address, i); err != nil {
				h.k.Logger(ctx).Error("error during replacing deposit information, %s", err.Error())
				panic(err)
			}

			if err != nil {
				h.k.Logger(ctx).Error("error during recording deposit information, %s", err.Error())
				panic(err)
			}

			// Delegate events
			if err := ctx.EventManager().EmitTypedEvent(zoneInfo); err != nil {
				panic(err)

			}
			if err := ctx.EventManager().EmitTypedEvent(depositRecord); err != nil {
				panic(err)
			}
		}
	}
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
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, msg *stakingtypes.MsgUndelegateResponse) {
	// get zone info from the validator address
	zone := h.k.ibcstakingKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	if zone == nil {
		return
	}

	h.k.DeleteUndelegateRecords(ctx, zone.ZoneId, UNDELEGATE_REQUEST_ICA)
	h.k.SetWithdrawTime(ctx, zone.ZoneId, WITHDRAW_REQUEST_USER, msg.CompletionTime)
}

func (h Hooks) AfterAutoStakingEnd() {
}
