package keeper_test

import (
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) SetWithdrawRecords(zoneId, withdrawer, recipient string, amount sdk.Coin, state int64, time time.Time) {
	record := types.WithdrawRecord{
		ZoneId:         zoneId,
		Withdrawer:     withdrawer,
		Recipient:      recipient,
		Amount:         &amount,
		State:          state,
		CompletionTime: time,
	}

	suite.App.GalKeeper.SetWithdrawRecord(suite.Ctx, record)
}

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

func (suite *KeeperTestSuite) TestGetTotalWithdrawAmountForZoneId() {
	suite.Setup()
	daoModifierAddr := suite.GenRandomAddress()
	zoneInfo := ibcstakingtypes.RegisteredZone{
		ZoneId: "gaia",
		IcaAccount: &ibcstakingtypes.IcaAccount{
			DaomodifierAddress: daoModifierAddr.String(),
		},
		BaseDenom: "stake",
		SnDenom:   "snuatom",
	}

	// 1 : WITHDRAW_REGISTER
	// 2 : TRANSFER_SUCCESS
	suite.App.IbcstakingKeeper.RegisterZone(suite.Ctx, &zoneInfo)
	suite.SetWithdrawRecords("gaia", "test1", "recipient1", sdk.NewCoin("stake", sdk.NewInt(1000)), int64(1),
		time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC))
	suite.SetWithdrawRecords("gaia", "test2", "recipient2", sdk.NewCoin("stake", sdk.NewInt(2000)), int64(2),
		time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour()-1, time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC))
	suite.SetWithdrawRecords("gaia", "test3", "recipient3", sdk.NewCoin("stake", sdk.NewInt(500)), int64(1),
		time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour()-1, time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC))
	suite.SetWithdrawRecords("gaia", "test4", "recipient4", sdk.NewCoin("stake", sdk.NewInt(700)), int64(2),
		time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC))

	tcs := []struct {
		name           string
		completionTime time.Time
		amount         int64
	}{
		{
			name:           "success",
			completionTime: time.Now().UTC(),
			amount:         1500,
		},
		{
			name:           "state",
			completionTime: time.Now().UTC().Add(10000000000),
			amount:         1500,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			res := suite.App.GalKeeper.GetTotalWithdrawAmountForZoneId(suite.Ctx, zoneInfo.ZoneId, tc.completionTime)
			suite.Require().Equal(tc.amount, res.Amount.Int64())
		})
	}
}
