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
		return
	}

	if data.Receiver != zoneInfo.IcaAccount.HostAddress {
		return
	}

	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId)
	ctx.Logger().Info("AfterTransferEnd", "zone id", zoneInfo.ZoneId, "sender", data.Sender, "receiver", data.Receiver, "amount", data.Amount)
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

	// remove deposit state
	err = h.k.DeleteRecordedDepositItem(ctx, zoneInfo.ZoneId, sender, types.DepositRequest, amount)
	if err != nil {
		ctx.Logger().Error("AfterTransferFail", "zoneId", zoneInfo.ZoneId, "sender", sender, "amount", amount, "err", err)
		return
	}
}

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

	controllerAddr, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		h.k.Logger(ctx).Error(",AfterOnRecvPacket", "receiver address is invalid", data.Receiver)
		return
	}

	denom := h.k.icaControlKeeper.GetIBCHashDenom(zone.TransferInfo.PortId, zone.TransferInfo.ChannelId, zone.BaseDenom)
	err = h.k.bankKeeper.SendCoinsFromAccountToModule(ctx, controllerAddr, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, asset)))
	if err != nil {
		return
	}

	// get withdrawVersion
	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion
	ctx.Logger().Info("AfterOnRecvPacket", "ZoneId", zone.ZoneId, "WithdrawCurrentVersion", currentVersion, "VersionState", versionInfo.Record[currentVersion].State)

	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, currentVersion)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, types.WithdrawStatusTransferred)

	// change version state
	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height:  uint64(ctx.BlockHeight()),
		Version: types.IcaSuccess,
	}
	ctx.Logger().Info("AfterOnRecvPacket", "ZoneId", zone.ZoneId, "WithdrawNextVersion", currentVersion, "VersionState", versionInfo.Record[currentVersion].State)

	// set withdraw version
	nextVersion := versionInfo.CurrentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: nextVersion,
		State:   icatypes.IcaPending,
	}
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Info("AfterOnRecvPacket", "ZoneId", zone.ZoneId, "WithdrawNextVersion", nextVersion, "VersionState", versionInfo.Record[nextVersion].State)
}

func (h Hooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	// getZoneInfoForValidatorAddr
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)
	oracleVersion, _ := h.k.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	// get delegateVersion
	versionInfo := h.k.GetDelegateVersion(ctx, zoneInfo.ZoneId)
	if versionInfo.Size() == 0 {
		return
	}
	currentVersion := versionInfo.CurrentVersion

	ctx.Logger().Info("AfterDelegateEnd", "ZoneId", zoneInfo.ZoneId, "DelegateCurrentVersion", versionInfo.CurrentVersion, "VersionState", versionInfo.Record[currentVersion].State)

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
	ctx.Logger().Info("AfterDelegateEnd", "ZoneId", zoneInfo.ZoneId, "DelegateNextVersion", nextVersion, "VersionState", versionInfo.Record[nextVersion].State)
}

// AfterWithdrawEnd is executed IcaWithdraw request finished.
// 1. Increase the withdrawal version.
// 2. The withdrawal status registered as transferred changes
func (h Hooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	asset := transferMsg.Token

	zone := h.k.icaControlKeeper.GetZoneForDenom(ctx, asset.Denom)

	if transferMsg.Receiver != zone.IcaAccount.ControllerAddress {
		h.k.Logger(ctx).Error("Receiver is not controller address", "Receiver", transferMsg.Receiver, "ControllerAddress", zone.IcaAccount.ControllerAddress, "Hook", "AfterWithdrawEnd")
		return
	}

	// get withdrawVersion
	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	if versionInfo.Size() == 0 {
		return
	}
	currentVersion := versionInfo.CurrentVersion
	ctx.Logger().Info("AfterWithdrawEnd", "ZoneId", zone.ZoneId, "WithdrawCurrentVersion", currentVersion, "VersionState", versionInfo.Record[currentVersion].State)

	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, currentVersion)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, types.WithdrawStatusTransferred)

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
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Info("AfterWithdrawEnd", "ZoneId", zone.ZoneId, "WithdrawNextVersion", nextVersion, "VersionState", versionInfo.Record[nextVersion].State)
}

// AfterUndelegateEnd is executed when ICA undelegation request finished.
// 1. It removes undelegation history in store.
// 2. It saves undelegation finish time to store.
func (h Hooks) AfterUndelegateEnd(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate, msg *stakingtypes.MsgUndelegateResponse) {
	// get zone info from the validator address
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	if zoneInfo == nil {
		h.k.Logger(ctx).Error("Zone id is not found", "ValidatorAddress", undelegateMsg.ValidatorAddress, "hook", "AfterUndelegateEnd")
		return
	}

	versionInfo := h.k.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	if versionInfo.Size() == 0 {
		return
	}
	currentVersion := versionInfo.CurrentVersion
	h.k.SetUndelegateRecordVersion(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca, currentVersion)
	h.k.SetWithdrawRecords(ctx, zoneInfo.ZoneId, msg.CompletionTime)

	h.k.DeleteUndelegateRecords(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)
	ctx.Logger().Info("AfterUndelegateEnd", "ZoneId", zoneInfo.ZoneId, "WithdrawCurrentVersion", currentVersion, "VersionState", versionInfo.Record[currentVersion].State)

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
	ctx.Logger().Info("AfterUndelegateEnd", "ZoneId", zoneInfo.ZoneId, "WithdrawNextVersion", nextVersion, "VersionState", versionInfo.Record[nextVersion].State)
}

func (h Hooks) AfterDelegateFail(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)
	ctx.Logger().Info("AfterDelegateFail", "Zone", zone)

	versionInfo := h.k.GetDelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}
	h.k.SetDelegateVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Error("AfterDelegateFail", "ZoneId", zone.ZoneId, "WithdrawCurrentVersion", currentVersion, "WithdrawVersionState", versionInfo.Record[currentVersion].State)
}

func (h Hooks) AfterUndelegateFail(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate) {
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)
	ctx.Logger().Info("AfterUndelegateFail", "Zone", zone)

	versionInfo := h.k.GetUndelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetUndelegateVersion(ctx, zone.ZoneId, versionInfo)
	ctx.Logger().Error("AfterUndelegateFail", "ZoneId", zone.ZoneId, "UndelegateCurrentVersion", currentVersion, "UndelegateVersionState", versionInfo.Record[currentVersion].State)
}

func (h Hooks) AfterIcaWithdrawFail(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
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
	ctx.Logger().Error("AfterIcaWithdrawFail", "ZoneId", zone.ZoneId, "IcaWithdrawCurrentVersion", currentVersion, "IcaWithdrawVersionState", versionInfo.Record[currentVersion].State)
}
