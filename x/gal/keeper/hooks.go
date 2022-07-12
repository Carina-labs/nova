package keeper

import (
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
		h.k.Logger(ctx).Error("Invalid ibc transfer packet", "data", data, "hook", "AfterTransferEnd")
		return
	}

	zoneInfo := h.k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)
	// if zoneInfo == nil, it may be a test situation.
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("Zone id is not found", "Denom", data.Denom, "hook", "AfterTransferEnd")
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return
	}
	// change sender + zoneId
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
				h.k.Logger(ctx).Error("Error during replacing deposit information", "err", err.Error(), "hook", "AfterTransferEnd")
				return
			}

			if err != nil {
				h.k.Logger(ctx).Error("Error during recording deposit information", "err,", err.Error(), "hook", "AfterTransferEnd")
				return
			}

			// TODO: Delegate events
		}
	}
}

func (h Hooks) AfterDelegateEnd() {
}

// ica transfer
func (h Hooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	asset := transferMsg.Token

	zoneInfo := h.k.ibcstakingKeeper.GetZoneForDenom(ctx, asset.Denom)
	if transferMsg.Receiver != zoneInfo.IcaAccount.DaomodifierAddress {
		h.k.Logger(ctx).Error("Receiver is not found", "receiver", transferMsg.Receiver, "hook", "AfterWithdrawEnd")
		return
	}

	if asset.Amount.IsZero() || asset.Amount.IsNil() {
		// TODO: withdraw fail event
		return
	}

	h.k.ChangeWithdrawState(ctx, zoneInfo.ZoneId, int64(ICA_WITHDRAW_REQUEST), int64(TRANSFER_SUCCESS))
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
		h.k.Logger(ctx).Error("Zone id is not found", "validatorAddress", undelegateMsg.ValidatorAddress, "hook", "AfterUndelegateEnd")
		return
	}

	h.k.DeleteUndelegateRecords(ctx, zone.ZoneId, UNDELEGATE_REQUEST_ICA)
	h.k.SetWithdrawTime(ctx, zone.ZoneId, WITHDRAW_REGISTER, msg.CompletionTime)
}

func (h Hooks) AfterAutoStakingEnd() {
}
