package keeper

import (
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

	zoneInfo := h.k.ibcstakingKeeper.GetZoneForDenom(ctx, baseDenom)
	// if zoneInfo == nil, it may be a test situation.
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("Zone id is not found", "Denom", data.Denom, "hook", "AfterTransferEnd")
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, DEPOSIT_REQUEST, DEPOSIT_SUCCESS)
}

// delegateAddr(controllerAddr), validatorAddr, delegateAmt
func (h Hooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	// getZoneInfoForValidatorAddr
	zoneInfo := h.k.ibcstakingKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)

	oracleVersion := h.k.oracleKeeper.GetOracleVersion(ctx, zoneInfo.BaseDenom)

	// get delegateVersion
	delegateVersion := h.k.GetDelegateVersion(ctx, zoneInfo.ZoneId)
	h.k.SetDelegateVersion(ctx, zoneInfo.ZoneId, delegateVersion+1)

	// change deposit state (DELEGATE_REQUEST -> DELEGATE_SUCCESS)
	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, DELEGATE_REQUEST, DELEGATE_SUCCESS)
	h.k.SetDepositOracleVersion(ctx, zoneInfo.ZoneId, DELEGATE_SUCCESS, oracleVersion)
	h.k.SetDelegateRecordVersion(ctx, zoneInfo.ZoneId, DELEGATE_SUCCESS, delegateVersion+1)
}

// ica transfer
func (h Hooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	asset := transferMsg.Token

	zoneInfo := h.k.ibcstakingKeeper.GetZoneForDenom(ctx, asset.Denom)
	if transferMsg.Receiver != zoneInfo.IcaAccount.DaomodifierAddress {
		h.k.Logger(ctx).Error("Receiver is not daomodifier address", "receiver", transferMsg.Receiver, "daomodifier address", zoneInfo.IcaAccount.DaomodifierAddress, "hook", "AfterWithdrawEnd")
		return
	}

	if asset.Amount.IsZero() || asset.Amount.IsNil() {
		// TODO: withdraw fail event
		return
	}

	// get withdrawVersion
	withdrawVersion := h.k.GetWithdrawVersion(ctx, zoneInfo.ZoneId)

	h.k.SetWithdrawVersion(ctx, zoneInfo.ZoneId, withdrawVersion+1)
	h.k.ChangeWithdrawState(ctx, zoneInfo.ZoneId, WITHDRAW_REGISTER, TRANSFER_SUCCESS)
}

func (h Hooks) BeforeUndelegateStart(ctx sdk.Context, zoneId string) {
}

// AfterUndelegateEnd is executed when ICA undelegation request finished.
// 1. It removes undelegation history in store.
// 2. It saves undelegation finish time to store.
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, msg *stakingtypes.MsgUndelegateResponse) {
	// get zone info from the validator address
	zoneInfo := h.k.ibcstakingKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("Zone id is not found", "validatorAddress", undelegateMsg.ValidatorAddress, "hook", "AfterUndelegateEnd")
		return
	}

	undelegateVersion := h.k.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	h.k.SetUndelegateRecordVersion(ctx, zoneInfo.ZoneId, UNDELEGATE_REQUEST_ICA, undelegateVersion+1)

	h.k.SetWithdrawRecords(ctx, zoneInfo.ZoneId, msg.CompletionTime)
	h.k.SetUndelegateVersion(ctx, zoneInfo.ZoneId, undelegateVersion+1)

	h.k.DeleteUndelegateRecords(ctx, zoneInfo.ZoneId, UNDELEGATE_REQUEST_ICA)
}

func (h Hooks) AfterAutoStakingEnd() {
}
