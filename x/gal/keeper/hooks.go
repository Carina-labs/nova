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
	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, currentVersion)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, types.WithdrawStatusTransferred)

	// change version state
	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height:  uint64(ctx.BlockHeight()),
		Version: types.IcaSuccess,
	}

	// set withdraw version
	nextVersion := versionInfo.CurrentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: nextVersion,
		State:   icatypes.IcaPending,
	}
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
}

func (h Hooks) AfterDelegateEnd(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	// getZoneInfoForValidatorAddr
	zoneInfo := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)

	oracleVersion, _ := h.k.oracleKeeper.GetOracleVersion(ctx, zoneInfo.ZoneId)

	// get delegateVersion
	versionInfo := h.k.GetDelegateVersion(ctx, zoneInfo.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	// change deposit state (DELEGATE_REQUEST -> DELEGATE_SUCCESS)
	h.k.ChangeDepositState(ctx, zoneInfo.ZoneId, types.DelegateRequest, types.DelegateSuccess)
	h.k.SetDepositOracleVersion(ctx, zoneInfo.ZoneId, types.DelegateSuccess, oracleVersion)
	h.k.SetDelegateRecordVersion(ctx, zoneInfo.ZoneId, types.DelegateSuccess, currentVersion)

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
}

// AfterWithdrawEnd is executed IcaWithdraw request finished.
// 1. Increase the withdrawal version.
// 2. The withdrawal status registered as transferred changes
func (h Hooks) AfterWithdrawEnd(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	asset := transferMsg.Token

	zone := h.k.icaControlKeeper.GetZoneForDenom(ctx, asset.Denom)

	if transferMsg.Receiver != zone.IcaAccount.ControllerAddress {
		h.k.Logger(ctx).Error("Receiver is not controller address", "receiver", transferMsg.Receiver, "Controller address", zone.IcaAccount.ControllerAddress, "hook", "AfterWithdrawEnd")
		return
	}

	// get withdrawVersion
	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	h.k.SetWithdrawRecordVersion(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, currentVersion)
	h.k.ChangeWithdrawState(ctx, zone.ZoneId, types.WithdrawStatusTransferRequest, types.WithdrawStatusTransferred)

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaSuccess,
	}

	nextVersion := currentVersion + 1
	versionInfo.CurrentVersion = nextVersion
	versionInfo.Record[nextVersion] = &types.IBCTrace{
		Version: currentVersion + 1,
		State:   types.IcaPending,
	}
	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
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

	versionInfo := h.k.GetUndelegateVersion(ctx, zoneInfo.ZoneId)
	currentVersion := versionInfo.CurrentVersion
	h.k.SetUndelegateRecordVersion(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca, currentVersion)
	h.k.SetWithdrawRecords(ctx, zoneInfo.ZoneId, msg.CompletionTime)

	h.k.DeleteUndelegateRecords(ctx, zoneInfo.ZoneId, types.UndelegateRequestByIca)

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

}

func (h Hooks) AfterDelegateFail(ctx sdk.Context, delegateMsg stakingtypes.MsgDelegate) {
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, delegateMsg.ValidatorAddress)

	versionInfo := h.k.GetDelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetDelegateVersion(ctx, zone.ZoneId, versionInfo)
}

func (h Hooks) AfterUndelegateFail(ctx sdk.Context, undelegateMsg stakingtypes.MsgUndelegate) {
	zone := h.k.icaControlKeeper.GetRegisteredZoneForValidatorAddr(ctx, undelegateMsg.ValidatorAddress)

	versionInfo := h.k.GetUndelegateVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetUndelegateVersion(ctx, zone.ZoneId, versionInfo)
}

func (h Hooks) AfterTransferFail(ctx sdk.Context, transferMsg transfertypes.MsgTransfer) {
	zone, ok := h.k.icaControlKeeper.GetRegisterZoneForHostAddr(ctx, transferMsg.Sender)
	if !ok {
		return
	}

	versionInfo := h.k.GetWithdrawVersion(ctx, zone.ZoneId)
	currentVersion := versionInfo.CurrentVersion

	versionInfo.Record[currentVersion] = &types.IBCTrace{
		Height: uint64(ctx.BlockHeight()),
		State:  types.IcaFail,
	}

	h.k.SetWithdrawVersion(ctx, zone.ZoneId, versionInfo)
}
