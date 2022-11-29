package keeper

import (
	"github.com/Carina-labs/nova/v2/x/gal/types"
	icatypes "github.com/Carina-labs/nova/v2/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"time"
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
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	// construct the denomination trace from the full raw denomination
	denomTrace := transfertypes.ParseDenomTrace(data.Denom)
	voucherDenom := denomTrace.IBCDenom()

	ibcDenom := h.k.icaControlKeeper.GetIBCHashDenom(zoneInfo.TransferInfo.PortId, zoneInfo.TransferInfo.ChannelId, zoneInfo.BaseDenom)
	if voucherDenom != ibcDenom {
		return
	}

	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterTransferEnd")
	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, data.Sender)
	ctx.Logger().Info("AfterTransferEnd", "zone_id", zoneInfo.ZoneId, "sender", data.Sender, "receiver", data.Receiver, "amount", data.Amount)
}

func (h Hooks) AfterTransferFail(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, baseDenom string) {
	zoneInfo := h.k.icaControlKeeper.GetZoneForDenom(ctx, baseDenom)
	// if zoneInfo == nil, it may be a test situation.
	if zoneInfo == nil {
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return
	}
	amount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return
	}

	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterTransferFail")
	// remove deposit state
	err = h.k.DeleteRecordedDepositItem(ctx, zoneInfo.ZoneId, sender, types.DepositRequest, amount)
	if err != nil {
		ctx.Logger().Error("AfterTransferFail", "zone_id", zoneInfo.ZoneId, "sender", sender, "amount", amount, "err", err)
		return
	}
}

// AfterOnRecvPacket is executed IcaWithdraw request finished.
// 1. Increase the withdrawal version.
// 2. The withdrawal status registered as transferred changes
func (h Hooks) AfterOnRecvPacket(ctx sdk.Context, data transfertypes.FungibleTokenPacketData) {
	zone := h.k.icaControlKeeper.GetZoneForDenom(ctx, data.Denom)
	if zone == nil {
		return
	}

	// check receiveAddr == controllerAddr && receiver == hostAddr
	if data.Sender != zone.IcaAccount.HostAddress {
		return
	}

	if data.Receiver != zone.IcaAccount.ControllerAddress {
		return
	}

	asset, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return
	}

	if asset.IsZero() || asset.IsNil() {
		h.k.Logger(ctx).Error("AfterOnRecvPacket", "transfer amount is zero", data.Amount)
		return
	}

	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterOnRecvPacket")

	controllerAddr, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		h.k.Logger(ctx).Error("AfterOnRecvPacket", "receiver address is invalid", data.Receiver)
		return
	}

	denom := h.k.icaControlKeeper.GetIBCHashDenom(zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, zone.BaseDenom)
	err = h.k.bankKeeper.SendCoinsFromAccountToModule(ctx, controllerAddr, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, asset)))
	if err != nil {
		ctx.Logger().Error("AfterOnRecvPacket", "SendToModuleAccount", err)
		return
	}

	// get withdrawVersion
	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion
	ctx.Logger().Info("AfterOnRecvPacket", "zone_id", zone.ZoneId, "withdraw_current_version", currentVersion, "version_state", versionInfo.Record[currentVersion].State)

	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, currentVersion)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, types.WithdrawStatusTransferred)

	// change version state
	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaSuccess,
	}
	ctx.Logger().Info("AfterOnRecvPacket", "zone_id", zone.ZoneId, "withdraw_next_version", currentVersion, "version_state", versionInfo.Record[currentVersion].State)

	// set withdraw version
	nextVersion := versionInfo.CurrentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: nextVersion,
		State:   icatypes.IcaPending,
	}
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Info("AfterOnRecvPacket", "zone_id", zone.ZoneId, "withdraw_next_version", nextVersion, "version_state", versionInfo.Record[nextVersion].State)
}

func (h Hooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	// getZoneInfoForValidatorAddr
	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterDelegateEnd")
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)
	oracleVersion, _ := h.k.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	// get delegateVersion
	versionInfo := h.k.GetDelegateVersion(ctx, zoneInfo.ZoneId)
	if versionInfo.Size() == 0 {
		ctx.Logger().Error("AfterDelegateEnd", "version_info", "nil")
		return
	}
	currentVersion := versionInfo.CurrentVersion

	ctx.Logger().Info("AfterDelegateEnd", "zone_id", zoneInfo.ZoneId, "delegate_current_version", versionInfo.CurrentVersion, "version_state", versionInfo.Record[currentVersion].State)

	// change delegate state (DELEGATE_REQUEST -> DELEGATE_SUCCESS)
	h.k.ChangeDelegateState(ctx, zoneInfo.ZoneId, versionInfo.CurrentVersion)
	h.k.SetDelegateOracleVersion(ctx, zoneInfo.ZoneId, versionInfo.CurrentVersion, oracleVersion)

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaSuccess,
	}

	nextVersion := versionInfo.CurrentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: nextVersion,
		State:   types.IcaPending,
	}
	h.k.SetDelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
	ctx.Logger().Info("AfterDelegateEnd", "zone_id", zoneInfo.ZoneId, "delegate_next_version", nextVersion, "version_state", versionInfo.Record[nextVersion].State)
}

// AfterUndelegateEnd is executed when ICA undelegation request finished.
// 1. It removes undelegation history in store.
// 2. It saves undelegation finish time to store.
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, msg *stakingtypes.MsgUndelegateResponse) {
	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterUndelegateEnd")

	// get zone info from the validator address
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	if zoneInfo == nil {
		ctx.Logger().Error("Zone id is not found", "validator_address", undelegateMsg.ValidatorAddress, "hook", "AfterUndelegateEnd")
		return
	}

	versionInfo := h.k.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	if versionInfo.Size() == 0 {
		ctx.Logger().Error("Zone id is not found", "version_info", "nil")
		return
	}
	currentVersion := versionInfo.CurrentVersion
	h.k.SetUndelegateRecordVersion(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca, currentVersion)
	h.k.SetWithdrawRecords(ctx, zoneInfo.ZoneId, msg.CompletionTime)

	h.k.DeleteUndelegateRecords(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)
	ctx.Logger().Info("AfterUndelegateEnd", "zone_id", zoneInfo.ZoneId, "withdraw_current_version", currentVersion, "version_state", versionInfo.Record[currentVersion].State)

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaSuccess,
	}

	nextVersion := currentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: nextVersion,
		State:   types.IcaPending,
	}
	h.k.SetUndelegateVersion(ctx, zoneInfo.ZoneId, versionInfo)
	ctx.Logger().Info("AfterUndelegateEnd", "zone_id", zoneInfo.ZoneId, "undelegate_next_version", nextVersion, "version_state", versionInfo.Record[nextVersion].State)
}

func (h Hooks) AfterDelegateFail(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterDelegateFail")
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)
	ctx.Logger().Info("AfterDelegateFail", "zone", zone)

	versionInfo := h.k.GetDelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}
	h.k.SetDelegateVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Error("AfterDelegateFail", "zone_id", zone.ZoneId, "delegate_current_version", currentVersion, "delegate_version_state", versionInfo.Record[currentVersion].State)
}

func (h Hooks) AfterUndelegateFail(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate) {
	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterUndelegateFail")
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	ctx.Logger().Info("AfterUndelegateFail", "zone", zone)

	versionInfo := h.k.GetUndelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetUndelegateVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Error("AfterUndelegateFail", "zone_id", zone.ZoneId, "undelegate_current_version", currentVersion, "undelegate_version_state", versionInfo.Record[currentVersion].State)
}

func (h Hooks) AfterIcaWithdrawFail(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	defer telemetry.MeasureSince(time.Now(), "gal", "hook", "afterIcaWithdrawFail")
	zone, ok := h.k.icaControlKeeper.GetRegisterZoneForHostAddr(ctx, transferMsg.Sender)
	if !ok {
		ctx.Logger().Error("AfterIcaWithdrawFail", "err", "zone not found")
		return
	}

	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Error("AfterIcaWithdrawFail", "zone_id", zone.ZoneId, "icaWithdraw_current_version", currentVersion, "icaWithdraw_version_state", versionInfo.Record[currentVersion].State)
}
