package novatesting

import (
	"testing"
)

type NovaTestState struct {
	testZoneStates []*TestZoneState
}

type TestZoneState struct {
	denom      string
	accounts   TestAccounts
	validators TestValidators
}

type TestAccounts struct {
	numAccounts int
	amount      int64
}

type TestValidators struct {
	numValidators    int
	delegationAmount int64
}

func SetupTestZone(t *testing.T, novaTestState NovaTestState) *Coordinator {
	numZone := len(novaTestState.testZoneStates)
	return NewCoordinatorWithChainState(t, numZone, novaTestState)

}
