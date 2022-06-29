package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestIsAbleToWithdraw() {
	tcs := []struct {
		name           string
		initialCoin    sdk.Coin
		withdrawAmount sdk.Coin
		expected       bool
	}{
		{
			name:           "valid test case (equal)",
			initialCoin:    sdk.NewInt64Coin("stake", 1_000_000),
			withdrawAmount: sdk.NewInt64Coin("stake", 1_000_000),
			expected:       true,
		},
		{
			name:           "valid test case (sub-withdraw)",
			initialCoin:    sdk.NewInt64Coin("stake", 1_000_000),
			withdrawAmount: sdk.NewInt64Coin("stake", 100_000),
			expected:       true,
		},
		{
			name:           "error test case (over-withdraw)",
			initialCoin:    sdk.NewInt64Coin("stake", 1_000_000),
			withdrawAmount: sdk.NewInt64Coin("stake", 1_000_000_000),
			expected:       false,
		},
	}

	for _, tc := range tcs {
		testOwnerAddress := "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3"
		testOwnerAcc, _ := sdk.AccAddressFromBech32(testOwnerAddress)
		suite.Run(tc.name, func() {
			// prepare
			err := suite.App.BankKeeper.MintCoins(suite.Ctx, types.ModuleName, sdk.Coins{tc.initialCoin})
			suite.Require().NoError(err)
			err = suite.App.BankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, types.ModuleName, testOwnerAcc, sdk.Coins{tc.initialCoin})
			suite.Require().NoError(err)

			// execute
			ok := suite.App.GalKeeper.IsAbleToWithdraw(suite.Ctx, testOwnerAcc, tc.withdrawAmount)

			// verify
			suite.Require().Equal(tc.expected, ok)
		})
	}
}

func (suite *KeeperTestSuite) TestClaimWithdrawAsset() {
	tcs := []struct {
		name        string
		initAmount  sdk.Coin
		userAddress string
		wantToClaim sdk.Coin
		shouldErr   bool
	}{
		{
			name:        "valid case",
			initAmount:  sdk.NewInt64Coin("stake", 1_000_000),
			userAddress: "cosmos1f4kwn0jqmten8kydsa6c9x3zsu4ctcexqpd64f",
			wantToClaim: sdk.NewInt64Coin("stake", 100_000),
			shouldErr:   false,
		},
		{
			name:        "error case",
			initAmount:  sdk.NewInt64Coin("stake", 1_000_000),
			userAddress: "cosmos1f4kwn0jqmten8kydsa6c9x3zsu4ctcexqpd64f",
			wantToClaim: sdk.NewInt64Coin("stake", 100_000_000),
			shouldErr:   true,
		},
	}

	for _, tc := range tcs {
		testOwnerAddress := "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3"
		testOwnerAcc, _ := sdk.AccAddressFromBech32(testOwnerAddress)
		suite.Run(tc.name, func() {
			// setup
			suite.Setup()
			err := suite.App.BankKeeper.MintCoins(suite.Ctx, types.ModuleName, sdk.Coins{tc.initAmount})
			suite.Require().NoError(err)
			err = suite.App.BankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, types.ModuleName, testOwnerAcc, sdk.Coins{tc.initAmount})
			suite.Require().NoError(err)
			userAcc, _ := sdk.AccAddressFromBech32(tc.userAddress)

			// execute
			err = suite.App.GalKeeper.ClaimWithdrawAsset(suite.Ctx, testOwnerAcc, userAcc, tc.wantToClaim)

			// verify
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
