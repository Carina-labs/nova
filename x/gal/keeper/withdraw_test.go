package keeper_test

import (
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WithdrawRecords struct {
	zoneId           string
	withdrawAddr     string
	amount           sdk.Int
	state            types.WithdrawStatusType
	undelegateVerion types.UndelegateVersion
	time             time.Time
}

func (suite *KeeperTestSuite) NewWithdrawRecords(withdrawRecords []*WithdrawRecords) {
	records := make(map[types.UndelegateVersion]*types.WithdrawRecordContent)
	for _, record := range withdrawRecords {
		withdrawRecord, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, record.zoneId, record.withdrawAddr)
		if !ok {
			withdrawRecord = &types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: record.withdrawAddr,
				Records:    records,
			}
		}
		withdrawContents := types.WithdrawRecordContent{
			Amount:          record.amount,
			State:           record.state,
			OracleVersion:   1,
			WithdrawVersion: 1,
			CompletionTime:  record.time,
		}
		withdrawRecord.Records[record.undelegateVerion] = &withdrawContents
		suite.App.GalKeeper.SetWithdrawRecord(suite.Ctx, withdrawRecord)
	}
}

func (suite *KeeperTestSuite) TestDeleteWithdrawRecord() {
	withdrawAddr1 := suite.GenRandomAddress()
	withdrawAddr2 := suite.GenRandomAddress()

	records := []*WithdrawRecords{
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(1000), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(2000), state: types.WithdrawStatusTransferred, undelegateVerion: 2,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(500), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(700), state: types.WithdrawStatusRegistered, undelegateVerion: 2,
		},
	}

	suite.NewWithdrawRecords(records)

	tcs := []struct {
		name         string
		zoneId       string
		withdrawAddr string
		result       *types.WithdrawRecord
	}{
		{
			name: "success", zoneId: zoneId, withdrawAddr: withdrawAddr1.String(),
			result: &types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr1.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(1000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
				},
			},
		},
		{
			name: "success", zoneId: zoneId, withdrawAddr: withdrawAddr2.String(),
			result: &types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr2.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(500),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
					2: {
						Amount:          sdk.NewInt(700),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			withdrawRecords, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, tc.zoneId, tc.withdrawAddr)
			suite.Require().True(ok)
			suite.App.GalKeeper.DeleteWithdrawRecord(suite.Ctx, withdrawRecords)

			result, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, tc.zoneId, tc.withdrawAddr)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result.Records, result.Records)
		})
	}
}

func (suite *KeeperTestSuite) TestSetWithdrawRecordVersion() {
	suite.Setup()
	withdrawAddr1 := suite.GenRandomAddress()
	withdrawAddr2 := suite.GenRandomAddress()

	records := []*WithdrawRecords{
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(1000), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(2000), state: types.WithdrawStatusTransferred, undelegateVerion: 2,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(500), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(700), state: types.WithdrawStatusRegistered, undelegateVerion: 2,
		},
	}

	suite.NewWithdrawRecords(records)

	tcs := []struct {
		name            string
		zoneId          string
		state           types.WithdrawStatusType
		withdrawVersion uint64
		withdrawAddr    string
		result          types.WithdrawRecord
	}{
		{
			name:            "withdraw version test 1",
			zoneId:          zoneId,
			state:           types.WithdrawStatusRegistered,
			withdrawVersion: 2,
			withdrawAddr:    withdrawAddr1.String(),
			result: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr1.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(1000),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 2,
					},
					2: {
						Amount:          sdk.NewInt(2000),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
				},
			},
		},
		{
			name:            "withdraw version test 2",
			zoneId:          zoneId,
			withdrawAddr:    withdrawAddr2.String(),
			state:           types.WithdrawStatusRegistered,
			withdrawVersion: 3,
			result: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr2.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(500),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 3,
					},
					2: {
						Amount:          sdk.NewInt(700),
						State:           types.WithdrawStatusRegistered,
						OracleVersion:   1,
						WithdrawVersion: 3,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.SetWithdrawRecordVersion(suite.Ctx, tc.zoneId, tc.state, tc.withdrawVersion)
			result, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, tc.zoneId, tc.withdrawAddr)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result.Records, result.Records)
		})
	}
}

func (suite *KeeperTestSuite) TestSetWithdrawRecords() {
	delegator := suite.GenRandomAddress()

	undelegateRecord := []*UndelegateRecord{
		{
			zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(1500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(1500)), state: types.UndelegateRequestByIca,
		},
		{
			zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(2000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(2000)), state: types.UndelegateRequestByUser,
		},
		{
			zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(3500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(3500)), state: types.UndelegateRequestByUser,
		},
		{
			zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(4000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(4000)), state: types.UndelegateRequestByIca,
		},
		{
			zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(5000)), state: types.UndelegateRequestByIca,
		},
	}

	suite.NewUndelegateRecord(undelegateRecord)

	tcs := []struct {
		name         string
		zoneId       string
		withdrawAddr string
		result       types.WithdrawRecord
	}{
		{
			name:         "",
			zoneId:       zoneId,
			withdrawAddr: delegator.String(),
			result: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: delegator.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(10500),
						State:           types.WithdrawStatusRegistered,
						WithdrawVersion: 1,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.SetWithdrawRecords(suite.Ctx, tc.zoneId, time.Time{})
			result, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, tc.zoneId, tc.withdrawAddr)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result, *result)
		})
	}

}

func (suite *KeeperTestSuite) TestGetWithdrawAmountForUser() {
	withdrawAddr1 := suite.GenRandomAddress()
	withdrawAddr2 := suite.GenRandomAddress()

	records := []*WithdrawRecords{
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(1000), state: types.WithdrawStatusTransferred, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(2000), state: types.WithdrawStatusTransferred, undelegateVerion: 2,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(500), state: types.WithdrawStatusTransferred, undelegateVerion: 3,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(700), state: types.WithdrawStatusRegistered, undelegateVerion: 4,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(3000), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(5000), state: types.WithdrawStatusRegistered, undelegateVerion: 2,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(200), state: types.WithdrawStatusTransferred, undelegateVerion: 3,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(1200), state: types.WithdrawStatusTransferred, undelegateVerion: 4,
		},
	}

	suite.NewWithdrawRecords(records)

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name         string
		zoneId       string
		withdrawAddr string
		denom        string
		result       sdk.Coin
	}{
		{
			name:         "withdraw address 1",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr1.String(),
			denom:        ibcDenom,
			result:       sdk.NewCoin(ibcDenom, sdk.NewInt(3500)),
		},
		{
			name:         "withdraw address 2",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr2.String(),
			denom:        ibcDenom,
			result:       sdk.NewCoin(ibcDenom, sdk.NewInt(1400)),
		},
		{
			name:         "fail - not found withdraw address",
			zoneId:       zoneId,
			withdrawAddr: "test",
			denom:        ibcDenom,
			result:       sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			result := suite.App.GalKeeper.GetWithdrawAmountForUser(suite.Ctx, tc.zoneId, tc.denom, tc.withdrawAddr)
			suite.Require().Equal(tc.result, result)
		})
	}

}

func (suite *KeeperTestSuite) TestGetTotalWithdrawAmountForZoneId() {
	suite.Setup()
	withdrawAddr1 := suite.GenRandomAddress()
	withdrawAddr2 := suite.GenRandomAddress()
	withdrawAddr3 := suite.GenRandomAddress()
	withdrawAddr4 := suite.GenRandomAddress()

	records := []*WithdrawRecords{
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(1000), state: types.WithdrawStatusRegistered, undelegateVerion: 1, time: time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC),
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr2.String(), amount: sdk.NewInt(2000), state: types.WithdrawStatusTransferred, undelegateVerion: 1, time: time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC),
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr3.String(), amount: sdk.NewInt(500), state: types.WithdrawStatusRegistered, undelegateVerion: 1, time: time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC),
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr4.String(), amount: sdk.NewInt(700), state: types.WithdrawStatusTransferred, undelegateVerion: 1, time: time.Date(time.Now().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second(), time.Now().UTC().Nanosecond(), time.UTC),
		},
	}

	suite.NewWithdrawRecords(records)

	tcs := []struct {
		name           string
		zoneId         string
		denom          string
		completionTime time.Time
		amount         int64
	}{
		{
			name:           "success",
			zoneId:         zoneId,
			denom:          baseDenom,
			completionTime: time.Now().UTC(),
			amount:         1500,
		},
		{
			name:           "state",
			zoneId:         zoneId,
			denom:          baseDenom,
			completionTime: time.Now().UTC().Add(10000000000),
			amount:         1500,
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			res := suite.App.GalKeeper.GetTotalWithdrawAmountForZoneId(suite.Ctx, tc.zoneId, tc.denom, tc.completionTime)
			suite.Require().Equal(tc.amount, res.Amount.Int64())
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

func (suite *KeeperTestSuite) TestChangeWithdrawState() {
	withdrawAddr1 := suite.GenRandomAddress()

	records := []*WithdrawRecords{
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(1000), state: types.WithdrawStatusRegistered, undelegateVerion: 1,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(2000), state: types.WithdrawStatusTransferred, undelegateVerion: 2,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(500), state: types.WithdrawStatusRegistered, undelegateVerion: 3,
		},
		{
			zoneId: "gaia", withdrawAddr: withdrawAddr1.String(), amount: sdk.NewInt(700), state: types.WithdrawStatusRegistered, undelegateVerion: 4,
		},
	}

	suite.NewWithdrawRecords(records)

	tcs := []struct {
		name         string
		zoneId       string
		withdrawAddr string
		result       types.WithdrawRecord
		err          bool
	}{
		{
			name:         "change withdraw state test case 1",
			zoneId:       zoneId,
			withdrawAddr: withdrawAddr1.String(),
			result: types.WithdrawRecord{
				ZoneId:     zoneId,
				Withdrawer: withdrawAddr1.String(),
				Records: map[uint64]*types.WithdrawRecordContent{
					1: {
						Amount:          sdk.NewInt(1000),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
					2: {
						Amount:          sdk.NewInt(2000),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
					3: {
						Amount:          sdk.NewInt(500),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
					4: {
						Amount:          sdk.NewInt(700),
						State:           types.WithdrawStatusTransferred,
						OracleVersion:   1,
						WithdrawVersion: 1,
					},
				},
			},
			err: false,
		},
		{
			name:         "change withdraw state test case 2",
			zoneId:       zoneId,
			withdrawAddr: "test",
			err:          true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.ChangeWithdrawState(suite.Ctx, tc.zoneId, types.WithdrawStatusRegistered, types.WithdrawStatusTransferred)
			result, ok := suite.App.GalKeeper.GetWithdrawRecord(suite.Ctx, tc.zoneId, tc.withdrawAddr)
			if tc.err {
				suite.Require().False(ok)
			} else {
				suite.Require().True(ok)
				suite.Require().Equal(tc.result, *result)
			}
		})
	}
}
