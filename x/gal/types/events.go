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
)

func CreateDepositEvent(ctx sdk.Context, sender sdk.AccAddress, zoneId, ownerAddr, hostAddr string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		TypeEvtDeposit,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyZoneId, zoneId),
		sdk.NewAttribute(AttributeKeyOwnerAddr, ownerAddr),
		sdk.NewAttribute(AttributeKeyHostAddr, hostAddr),
		sdk.NewAttribute(AttributeKeyAmount, amount.String()),
	)
}
