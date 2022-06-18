package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeEvtDeposit    = "deposit"
	TypeEvtUndelegate = "undelegate"
	TypeEvtWithdraw   = "withdraw"

	AttributeValueCategory = ModuleName
	AttributeKeyZoneId     = "zone_id"
	AttributeKeyOwnerAddr  = "owner_addr"
	AttributeKeyHostAddr   = "host_addr"
	AttributeKeyAmount     = "amount"
	AttributeKeySenderAddr = "send_addr"
	AttributeKeyReceivAddr = "receive_addr"
)

func CreateDepositEvent(ctx sdk.Context, sender sdk.AccAddress, zoneId, ownerAddr, hostAddr string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		TypeEvtDeposit,
		sdk.NewAttribute(AttributeKeyZoneId, zoneId),
		sdk.NewAttribute(AttributeKeyOwnerAddr, ownerAddr),
		sdk.NewAttribute(AttributeKeyHostAddr, hostAddr),
		sdk.NewAttribute(AttributeKeyAmount, amount.String()),
	)

}

func CreatUndelegateEvent(ctx sdk.Context, sender sdk.AccAddress, zoneId, ownerAddr, hostAddr string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		TypeEvtUndelegate,
		sdk.NewAttribute(AttributeKeyZoneId, zoneId),
		sdk.NewAttribute(AttributeKeyOwnerAddr, ownerAddr),
		sdk.NewAttribute(AttributeKeyHostAddr, hostAddr),
		sdk.NewAttribute(AttributeKeyAmount, amount.String()),
	)
}

func CreatWithdrawEvent(ctx sdk.Context, sender, zoneId, ownerAddr, receive_addr string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		TypeEvtWithdraw,
		sdk.NewAttribute(AttributeKeyZoneId, zoneId),
		sdk.NewAttribute(AttributeKeyOwnerAddr, ownerAddr),
		sdk.NewAttribute(AttributeKeySenderAddr, sender),
		sdk.NewAttribute(AttributeKeyReceivAddr, receive_addr),
		sdk.NewAttribute(AttributeKeyAmount, amount.String()),
	)
}
