package keeper

import (
	"fmt"

	icatypes "github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// Hooks wrapper struct for gal keeper
type Hooks struct {
	k Keeper
}

var _ transfertypes.TransferHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterTransferEnd(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, base_denom string) {

	depositor, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}

	amt, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		h.k.Logger(ctx).Error(fmt.Sprintf("type casting error, %s", data.Amount))
	}

	err = h.k.CacheDepositAmt(ctx, depositor, sdk.NewCoin(data.Denom, amt))
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}
}

// Hooks wrapper struct for gal keeper
type IHooks struct {
	k Keeper
}

var _ icatypes.ICAHooks = IHooks{}

func (k Keeper) IHooks() IHooks {
	return IHooks{k}
}

func (h IHooks) AfterDelegateEnd() {
}
func (h IHooks) AfterWithdrawEnd() {
}

func (h IHooks) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, txHash string) {
	// portId로 zone 정보 조회
	// zone := h.k.interTxKeeper.GetRegisteredZoneForPortId(ctx, packet.SourcePort)

	// getStatusForUndelegateRecords : status가 REQUEST_ICA인 UndelegateRecord 조회
	// undelegateRecords := h.k.GetUndelegateRecordsForZoneId(ctx, zone.ZoneName, UNDELEGATE_REQUEST_ICA)

	// setUndelegateReceipt : (Record의 주소, amount, zoneId)와 txHash 저장
	// h.k.SetUndelegateReceipt(ctx, txhash)
	// deleteUndelegateRecords : Receipt에 저장된 record삭제
	// h.k.DeleteUndelegateRecords(ctx, zone.ZoneName, UNDELEGATE_REQUEST_ICA)
}
func (h IHooks) AfterAutoStakingEnd() {
}
