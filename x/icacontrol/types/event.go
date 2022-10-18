package types

func NewEventRegisterZone(zoneInfo *RegisteredZone) *EventRegisterZone {
	return &EventRegisterZone{
		ZoneId:               zoneInfo.ZoneId,
		IcaAccount:           zoneInfo.IcaAccount,
		TransferInfo:         zoneInfo.TransferInfo,
		IcaConnectionInfo:    zoneInfo.IcaConnectionInfo,
		ValidatorAddress:     zoneInfo.ValidatorAddress,
		Decimal:              zoneInfo.Decimal,
		SnDenom:              zoneInfo.SnDenom,
		DepositMaxEntries:    zoneInfo.DepositMaxEntries,
		UndelegateMaxEntries: zoneInfo.UndelegateMaxEntries,
	}
}

func NewEventDeleteZone(zoneInfo *MsgDeleteRegisteredZone) *EventDeleteZone {
	return &EventDeleteZone{
		ZoneId:            zoneInfo.ZoneId,
		ControllerAddress: zoneInfo.ControllerAddress,
	}
}

func NewEventChangeRegisterZone(zoneInfo *RegisteredZone) *EventChangeRegisteredZone {
	return &EventChangeRegisteredZone{
		ZoneId:               zoneInfo.ZoneId,
		IcaInfo:              zoneInfo.IcaConnectionInfo,
		IcaAccount:           zoneInfo.IcaAccount,
		TransferInfo:         zoneInfo.TransferInfo,
		ValidatorAddress:     zoneInfo.ValidatorAddress,
		BaseDenom:            zoneInfo.BaseDenom,
		Decimal:              zoneInfo.Decimal,
		SnDenom:              zoneInfo.SnDenom,
		UndelegateMaxEntries: zoneInfo.UndelegateMaxEntries,
		DepositMaxEntries:    zoneInfo.DepositMaxEntries,
	}
}

func NewEventIcaAutoStaking(icaAutoStaking *MsgIcaAutoStaking) *EventIcaAutoStaking {
	return &EventIcaAutoStaking{
		ZoneId:            icaAutoStaking.ZoneId,
		ControllerAddress: icaAutoStaking.ControllerAddress,
		Amount:            icaAutoStaking.Amount,
	}
}

func NewEventIcaAuthzGrant(icaAuthzGrant *MsgIcaAuthzGrant) *EventIcaAuthzGrant {
	return &EventIcaAuthzGrant{
		ZoneId:            icaAuthzGrant.ZoneId,
		ControllerAddress: icaAuthzGrant.ControllerAddress,
		Grant:             icaAuthzGrant.Grant,
	}
}

func NewEventIcaAuthzRevoke(icaAuthzRevoke *MsgIcaAuthzRevoke) *EventIcaAuthzRevoke {
	return &EventIcaAuthzRevoke{
		ZoneId:            icaAuthzRevoke.ZoneId,
		Grantee:           icaAuthzRevoke.Grantee,
		ControllerAddress: icaAuthzRevoke.ControllerAddress,
		MsgTypeUrl:        icaAuthzRevoke.MsgTypeUrl,
	}
}

func NewEventRegisterControllerAddress(controllerAddressInfo *MsgRegisterControllerAddr) *EventRegisterControllerAddress {
	return &EventRegisterControllerAddress{
		ZoneId:            controllerAddressInfo.ZoneId,
		ControllerAddress: controllerAddressInfo.ControllerAddress,
		FromAddress:       controllerAddressInfo.FromAddress,
	}
}

func NewEventIcaDelegate(delegateInfo *MsgIcaDelegate) *EventIcaDelegate {
	return &EventIcaDelegate{
		ZoneId:            delegateInfo.ZoneId,
		ControllerAddress: delegateInfo.ControllerAddress,
		Amount:            delegateInfo.Amount,
	}
}

func NewEventIcaUndelegate(undelegateInfo *MsgIcaUndelegate) *EventIcaUndelegate {
	return &EventIcaUndelegate{
		ZoneId:            undelegateInfo.ZoneId,
		ControllerAddress: undelegateInfo.ControllerAddress,
		Amount:            undelegateInfo.Amount,
	}
}

func NewEventIcaTransfer(transferInfo *MsgIcaTransfer) *EventIcaTransfer {
	return &EventIcaTransfer{
		ZoneId:               transferInfo.ZoneId,
		ControllerAddress:    transferInfo.ControllerAddress,
		ReceiverAddress:      transferInfo.ReceiverAddress,
		IcaTransferPortId:    transferInfo.IcaTransferPortId,
		IcaTransferChannelId: transferInfo.IcaTransferChannelId,
		Amount:               transferInfo.Amount,
	}
}
