package novatesting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestChainSetup(t *testing.T) {
	valState := TestValidators{
		numValidators:    5,
		delegationAmount: 111111,
	}

	accountState := TestAccounts{
		numAccounts: 5,
		amount:      1000000000,
	}

	testZoneState1 := TestZoneState{
		denom:      "uosmo",
		accounts:   accountState,
		validators: valState,
	}

	testZoneState2 := TestZoneState{
		denom:      "unova",
		accounts:   accountState,
		validators: valState,
	}
	testZoneState3 := TestZoneState{
		denom:      "uatom",
		accounts:   accountState,
		validators: valState,
	}
	testZoneStates := []*TestZoneState{&testZoneState1, &testZoneState2, &testZoneState3}
	novaTestState := NovaTestState{testZoneStates: testZoneStates}

	coor := SetupTestZone(t, novaTestState)
	for _, chain := range coor.Chains {
		println("=============================")
		chainId := chain.ChainID
		println(chainId)
		app := chain.App
		ctx := chain.GetContext()

		println("account info ==")
		chainAccounts := chain.SenderAccounts

		for _, acc := range chainAccounts {
			println("address : ", acc.SenderAccount.GetAddress().String())
			balance := app.BankKeeper.GetAllBalances(ctx, acc.SenderAccount.GetAddress())
			println("balance : ", balance.String())

			delegation := app.StakingKeeper.GetAllDelegatorDelegations(ctx, acc.SenderAccount.GetAddress())
			for _, del := range delegation {
				delegatorTokens, _ := app.StakingKeeper.GetValidator(ctx, del.GetValidatorAddr())
				println("delegations : ", sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), delegatorTokens.TokensFromShares(del.GetShares()).TruncateInt()).String())

			}
		}

		println("validator info ==")
		validators := app.StakingKeeper.GetAllValidators(ctx)
		for _, val := range validators {
			println("operator address : ", val.OperatorAddress)
			balance := app.StakingKeeper.GetValidatorDelegations(ctx, val.GetOperator())
			println("balance : ", balance[0].String())
		}
	}

}
