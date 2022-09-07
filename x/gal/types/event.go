package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewEventDeposit(depositor, claimer string, depositedAmount *sdk.Coin) *EventDeposit {
	return &EventDeposit{
		Depositor:     depositor,
		Claimer:       claimer,
		DepositAmount: depositedAmount,
	}
}

func NewEventDelegate(
	hostAddress,
	validatorAddress string,
	delegatedAmount *sdk.Coin,
	channelId,
	portId string) *EventDelegate {
	return &EventDelegate{
		HostAddress:      hostAddress,
		ValidatorAddress: validatorAddress,
		DelegatedAmount:  delegatedAmount,
		ChannelId:        channelId,
		PortId:           portId,
	}
}

func NewEventPendingUndelegate(
	zoneId, delegator, withdrawer string,
	burnedAmount, undelegatedAmount *sdk.Coin,
) *EventPendingUndelegate {
	return &EventPendingUndelegate{
		ZoneId:            zoneId,
		Delegator:         delegator,
		Withdrawer:        withdrawer,
		BurnedAmount:      burnedAmount,
		UndelegatedAmount: undelegatedAmount,
	}
}

func NewEventUndelegate(
	zoneId string, burnedAsset, undelegatedAsset *sdk.Coin) *EventUndelegate {
	return &EventUndelegate{
		ZoneId:            zoneId,
		BurnedAmount:      burnedAsset,
		UndelegatedAmount: undelegatedAsset,
	}
}

func NewEventWithdraw(zoneId, withdrawer string, withdrawnAmount *sdk.Coin) *EventWithdraw {
	return &EventWithdraw{
		ZoneId:          zoneId,
		Withdrawer:      withdrawer,
		WithdrawnAmount: withdrawnAmount,
	}
}

func NewEventClaimSnToken(claimer string,
	claimedToken *sdk.Coin,
	oracleVersion uint64) *EventClaimSnToken {
	return &EventClaimSnToken{
		Claimer:       claimer,
		ClaimedToken:  claimedToken,
		OracleVersion: oracleVersion,
	}
}

func NewEventIcaWithdraw(hostAddress,
	controllerAddress string,
	withdrawalToken *sdk.Coin,
	connectionId, channelId, portId string) *EventIcaWithdraw {
	return &EventIcaWithdraw{
		HostAddress:       hostAddress,
		ControllerAddress: controllerAddress,
		WithdrawnToken:    withdrawalToken,
		ConnectionId:      connectionId,
		ChannelId:         channelId,
		PortId:            portId,
	}
}
