package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	icatypes "github.com/Carina-labs/nova/x/icacontrol/types"
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
	zoneInfo := h.k.icaControlKeeper.GetZoneForDenom(ctx, baseDenom)
	// if zoneInfo == nil, it may be a test situation.
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("AfterTransferEnd", "err", "Zone id is not found", "Denom", data.Denom)
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, types.DepositRequest, types.DepositSuccess)
}

func (h Hooks) AfterOnRecvPacket(ctx sdk.Context, data transfertypes.FungibleTokenPacketData) {
	zone := h.k.icaControlKeeper.GetZoneForDenom(ctx, data.Denom)
	if zone == nil {
		h.k.Logger(ctx).Error("AfterOnRecvPacket", "err", "Zone id is not found", "Denom", data.Denom)
		return
	}

	// check receiveAddr == controllerAddr && receiver == hostAddr
	if data.Sender != zone.IcaAccount.HostAddress {
		h.k.Logger(ctx).Error("AfterOnRecvPacket", "err", "The sender address is different from the registered host address", "hostAddr", zone.IcaAccount.HostAddress, "senderAddr", data.Sender)
		return
	}

	if data.Receiver != zone.IcaAccount.ControllerAddress {
		h.k.Logger(ctx).Error("AfterOnRecvPacket", "err", "The sender address is different from the registered host address", "hostAddr", zone.IcaAccount.HostAddress, "senderAddr", data.Sender)
		return
	}

	asset, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		h.k.Logger(ctx).Error("")
		return
	}

	if asset.IsZero() || asset.IsNil() {
		h.k.Logger(ctx).Error("")
		return
	}

	// get withdrawVersion
	withdrawVersion, _ := h.k.GetWithdrawVersion(ctx, zone.ZoneId)

	trace := types.IBCTrace{
		Version: withdrawVersion + 1,
		Height:  uint64(ctx.BlockHeight()),
	}
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, trace)
	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusRegistered, withdrawVersion+1)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusRegistered, types.WithdrawStatusTransferred)
}

// delegateAddr(controllerAddr), validatorAddr, delegateAmt
func (h Hooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	// getZoneInfoForValidatorAddr
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)

	oracleVersion, _ := h.k.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)
	// get delegateVersion
	delegateVersion, _ := h.k.GetDelegateVersion(ctx, zoneInfo.ZoneId)

	trace := types.IBCTrace{
		Version: delegateVersion + 1,
		Height:  uint64(ctx.BlockHeight()),
	}
	h.k.SetDelegateVersion(ctx, zoneInfo.ZoneId, trace)

	// change deposit state (DELEGATE_REQUEST -> DELEGATE_SUCCESS)
	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, types.DelegateRequest, types.DelegateSuccess)
	h.k.SetDepositOracleVersion(ctx, zoneInfo.ZoneId, types.DelegateSuccess, oracleVersion)
	h.k.SetDelegateRecordVersion(ctx, zoneInfo.ZoneId, types.DelegateSuccess, delegateVersion+1)
}

// AfterWithdrawEnd is executed IcaWithdraw request finished.
// 1. Increase the withdrawal version.
// 2. The withdrawal status registered as transferred changes
func (h Hooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	asset := transferMsg.Token

	zoneInfo := h.k.icaControlKeeper.GetZoneForDenom(ctx, asset.Denom)

	if transferMsg.Receiver != zoneInfo.IcaAccount.ControllerAddress {
		h.k.Logger(ctx).Error("Receiver is not controller address", "receiver", transferMsg.Receiver, "Controller address", zoneInfo.IcaAccount.ControllerAddress, "hook", "AfterWithdrawEnd")
		return
	}

	// get withdrawVersion
	withdrawVersion, _ := h.k.GetWithdrawVersion(ctx, zoneInfo.ZoneId)

	trace := types.IBCTrace{
		Version: withdrawVersion + 1,
		Height:  uint64(ctx.BlockHeight()),
	}
	h.k.SetWithdrawVersion(ctx, zoneInfo.ZoneId, trace)
	h.k.SetWithdrawRecordVersion(ctx, zoneInfo.ZoneId, types.WithdrawStatusRegistered, withdrawVersion+1)
	h.k.ChangeWithdrawState(ctx, zoneInfo.ZoneId, types.WithdrawStatusRegistered, types.WithdrawStatusTransferred)

}

func (h Hooks) BeforeUndelegateStart(ctx sdk.Context, zoneId string) {
}

// AfterUndelegateEnd is executed when ICA undelegation request finished.
// 1. It removes undelegation history in store.
// 2. It saves undelegation finish time to store.
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, msg *stakingtypes.MsgUndelegateResponse) {
	// get zone info from the validator address
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("Zone id is not found", "validatorAddress", undelegateMsg.ValidatorAddress, "hook", "AfterUndelegateEnd")
		return
	}

	undelegateVersion, _ := h.k.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	h.k.SetUndelegateRecordVersion(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca, undelegateVersion+1)

	trace := types.IBCTrace{
		Version: undelegateVersion + 1,
		Height:  uint64(ctx.BlockHeight()),
	}
	h.k.SetWithdrawRecords(ctx, zoneInfo.ZoneId, msg.CompletionTime)
	h.k.SetUndelegateVersion(ctx, zoneInfo.ZoneId, trace)

	h.k.DeleteUndelegateRecords(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)
}

func (h Hooks) AfterAutoStakingEnd() {
}
