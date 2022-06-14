package keeper

import (
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

// AfterTransferEnd mints st token the result of IBC transfer.
func (h Hooks) AfterTransferEnd(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, base_denom string) {

	////depositor를 받아오고
	//depositor, err := sdk.AccAddressFromBech32(data.Sender)
	//if err != nil {
	//	h.k.Logger(ctx).Error(err.Error())
	//}
	////보낸 금액을 받아와서
	//amt, ok := sdk.NewIntFromString(data.Amount)
	//if !ok {
	//	h.k.Logger(ctx).Error(fmt.Sprintf("type casting error, %s", data.Amount))
	//}
	//
	////cacheDepositAmt하는데 denom이 hash다?
	////cacheDepositAmt는 현재 depositor와 amt를 캐시화해서 저장
	//err = h.k.RecordDepositAmt(ctx, sdk.NewCoin(data.Denom, amt))
	//if err != nil {
	//	h.k.Logger(ctx).Error(err.Error())
	//}
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
func (h IHooks) BeforeUndelegateStart(ctx sdk.Context, zoneId string) {
	// getAllUndelegateRecords
	undelegateRecoreds := h.k.GetAllUndelegateRecord(ctx, zoneId)

	// changeUndelegateStatus : Request
	h.k.ChangeUndelegateStatus(ctx, undelegateRecoreds)
}
func (h IHooks) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, txHash string) {
	// portId로 zone 정보 조회
	// zone := h.k.interTxKeeper.GetRegisteredZoneForPortId(ctx, packet.SourcePort)

	// getStatusForUndelegateRecords : status가 Request인 UndelegateRecord 조회
	// undelegateRecords := h.k.GetUndelegateRecordsForZoneId(ctx, zone.ZoneName, UNDELEGATE_REQUEST_ICA)

	// setUndelegateReceipt : (Record의 주소, amount, zoneId)와 txHash 저장
	// h.k.SetUndelegateReceipt(ctx, txhash)
	// deleteUndelegateRecords : Receipt에 저장된 record삭제
	// h.k.DeleteUndelegateRecords(ctx, undelegateRecords)
}
func (h IHooks) AfterAutoStakingEnd() {
}
