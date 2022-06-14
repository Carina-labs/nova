package keeper

import (
	"fmt"

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

func (h IHooks) AfterUndelegateEnd(ctx sdk.Context, packet channeltypes.Packet, response *stakingtypes.MsgUndelegateResponse) {

	h.k.Logger(ctx).Info("AfterUndelgateEnd", "msgData", response)

	// portId로 zone 정보 조회
	zone := h.k.interTxKeeper.GetRegisteredZoneForPortId(ctx, packet.SourcePort)

	// hook에서 withdraw record에 time 기록
	h.k.SetWithdrawTime(ctx, zone.ZoneName, WITHDRAW_REGISTER, response.CompletionTime)

	// undelegate record 삭제
	h.k.DeleteUndelegateRecords(ctx, zone.ZoneName, UNDELEGATE_REQUEST_ICA)
}
func (h IHooks) AfterAutoStakingEnd() {
}
