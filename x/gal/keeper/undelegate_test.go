package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type UndelegateRecord struct {
	zoneId        string
	withdrawer    string
	delegator     string
	claimer       sdk.AccAddress
	snAsset       sdk.Coin
	withdrawAsset sdk.Coin
	state         int64
	result        types.UndelegateRecord
}

func (suite *KeeperTestSuite) NewUndelegateRecord(records []*UndelegateRecord) {

	for _, record := range records {
		undelegateRecord, ok := suite.App.GalKeeper.GetUndelegateRecord(suite.Ctx, record.zoneId, record.delegator)

		if !ok {
			undelegateRecord = &types.UndelegateRecord{
				ZoneId:    record.zoneId,
				Delegator: record.delegator,
			}
		}
		recordContent := &types.UndelegateRecordContent{
			Withdrawer:        record.withdrawer,
			SnAssetAmount:     &record.snAsset,
			WithdrawAmount:    record.withdrawAsset.Amount,
			State:             record.state,
			OracleVersion:     1,
			UndelegateVersion: 1,
		}

		undelegateRecord.Records = append(undelegateRecord.Records, recordContent)

		suite.App.GalKeeper.SetUndelegateRecord(suite.Ctx, undelegateRecord)

		if record.state == types.UndelegateRequestByUser {
			suite.App.BankKeeper.MintCoins(suite.Ctx, types.ModuleName, sdk.NewCoins(record.snAsset))
		}
	}
}

func (suite *KeeperTestSuite) TestGetAllUndelegateRecord() {
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
		name      string
		zoneId    string
		delegator sdk.AccAddress
		result    []*types.UndelegateRecord
	}{
		{
			name:      "success",
			zoneId:    zoneId,
			delegator: delegator,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(1500, 18),
							},
							WithdrawAmount:    sdk.NewInt(1500),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(4000, 18),
							},
							WithdrawAmount:    sdk.NewInt(4000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						}, {
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(5000, 18),
							},
							WithdrawAmount:    sdk.NewInt(5000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
					},
				},
			},
		},
		{
			name:      "fail - not registered zone",
			zoneId:    "test",
			delegator: delegator,
			result:    []*types.UndelegateRecord(nil),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			result := suite.App.GalKeeper.GetAllUndelegateRecord(suite.Ctx, tc.zoneId)
			suite.Require().Equal(tc.result, result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetUndelegateAmount() {
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

	suite.App.OracleKeeper.SetParams(suite.Ctx, oracletypes.Params{
		OracleOperators: []string{baseOwnerAcc.String()},
	})

	tcs := []struct {
		name          string
		zoneId        string
		snDenom       string
		oracleVersion uint64
		delegator     sdk.AccAddress
		snAsset       sdk.Coin
		wAsset        sdk.Int
	}{
		{
			name:          "success",
			zoneId:        zoneId,
			snDenom:       baseSnDenom,
			delegator:     delegator,
			oracleVersion: 2,
			snAsset:       sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5500, 18)),
			wAsset:        sdk.NewInt(5500),
		},
		{
			name:          "fail - oracle version",
			zoneId:        zoneId,
			snDenom:       baseSnDenom,
			delegator:     delegator,
			oracleVersion: 1,
			snAsset:       sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(0, 18)),
			wAsset:        sdk.NewInt(0),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			zone, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, tc.zoneId)
			suite.Require().True(ok)
			chainInfo := oracletypes.ChainInfo{
				ChainId:         tc.zoneId,
				OperatorAddress: baseOwnerAcc.String(),
				Coin:            sdk.NewCoin(baseDenom, tc.wAsset),
				OracleVersion:   1,
				AppHash:         []byte("test"),
				LastBlockHeight: 2,
			}

			err := suite.App.OracleKeeper.UpdateChainState(suite.Ctx, &chainInfo)
			suite.Require().NoError(err)
			suite.App.OracleKeeper.SetOracleVersion(suite.Ctx, tc.zoneId, tc.oracleVersion, uint64(suite.Ctx.BlockHeight()))

			snAsset, wAsset := suite.App.GalKeeper.GetUndelegateAmount(suite.Ctx, tc.snDenom, zone, tc.oracleVersion)
			suite.Require().Equal(tc.snAsset, snAsset)
			suite.Require().Equal(tc.wAsset, wAsset)
		})
	}
}

func (suite *KeeperTestSuite) TestChangeUndelegateState() {
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
		name   string
		zoneId string
		state  types.UndelegatedStatusType
		result []*types.UndelegateRecord
	}{
		{
			name:   "success",
			zoneId: zoneId,
			state:  types.UndelegateRequestByIca,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(1500, 18),
							},
							WithdrawAmount:    sdk.NewInt(1500),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(4000, 18),
							},
							WithdrawAmount:    sdk.NewInt(4000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						}, {
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(5000, 18),
							},
							WithdrawAmount:    sdk.NewInt(5000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.ChangeUndelegateState(suite.Ctx, tc.zoneId, tc.state)
			result := suite.App.GalKeeper.GetAllUndelegateRecord(suite.Ctx, tc.zoneId)
			suite.Require().Equal(result, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestUndelegateRecordVersion() {
	delegator := suite.GenRandomAddress()

	tcs := []struct {
		name              string
		zoneId            string
		records           []*UndelegateRecord
		undelegateVersion int64
		result            []*types.UndelegateRecord
	}{
		{
			name:              "success",
			zoneId:            zoneId,
			undelegateVersion: 2,
			records: []*UndelegateRecord{
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
			},
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(1500, 18),
							},
							WithdrawAmount:    sdk.NewInt(1500),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 2,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(4000, 18),
							},
							WithdrawAmount:    sdk.NewInt(4000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 2,
						}, {
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(5000, 18),
							},
							WithdrawAmount:    sdk.NewInt(5000),
							State:             types.UndelegateRequestByIca,
							OracleVersion:     1,
							UndelegateVersion: 2,
						},
					},
				},
			},
		},
		{
			name:   "Status does not change",
			zoneId: zoneId,
			records: []*UndelegateRecord{
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(1500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(1500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(2000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(2000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(3500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(3500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(4000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(4000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(5000)), state: types.UndelegateRequestByUser,
				},
			},
			undelegateVersion: 2,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(1500, 18),
							},
							WithdrawAmount:    sdk.NewInt(1500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(4000, 18),
							},
							WithdrawAmount:    sdk.NewInt(4000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						}, {
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(5000, 18),
							},
							WithdrawAmount:    sdk.NewInt(5000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
					},
				},
			},
		},
		{
			name:   "zone not found",
			zoneId: "test",
			records: []*UndelegateRecord{
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(1500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(1500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(2000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(2000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(3500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(3500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(4000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(4000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(5000)), state: types.UndelegateRequestByUser,
				},
			},
			undelegateVersion: 2,
			result:            []*types.UndelegateRecord(nil),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.NewUndelegateRecord(tc.records)

			ok := suite.App.GalKeeper.SetUndelegateRecordVersion(suite.Ctx, tc.zoneId, types.UndelegateRequestByIca, 2)
			suite.Require().True(ok)

			result := suite.App.GalKeeper.GetAllUndelegateRecord(suite.Ctx, tc.zoneId)
			suite.Require().Equal(tc.result, result)

			suite.App.GalKeeper.DeleteUndelegateRecords(suite.Ctx, tc.zoneId, types.UndelegateRequestByIca)
			suite.App.GalKeeper.DeleteUndelegateRecords(suite.Ctx, tc.zoneId, types.UndelegateRequestByUser)
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteUndelegateRecords() {
	delegator := suite.GenRandomAddress()

	tcs := []struct {
		name    string
		zoneId  string
		records []*UndelegateRecord
		state   types.UndelegatedStatusType
		result  []*types.UndelegateRecord
	}{
		{
			name:   "delete undelegateRecords case 1",
			zoneId: zoneId,
			records: []*UndelegateRecord{
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
			},
			state: types.UndelegateRequestByIca,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
					},
				},
			},
		},
		{
			name:   "delete undelegateRecords case 2",
			zoneId: zoneId,
			records: []*UndelegateRecord{
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(1500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(1500)), state: types.UndelegateRequestByIca,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(2000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(2000)), state: types.UndelegateRequestByIca,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(3500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(3500)), state: types.UndelegateRequestByIca,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(4000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(4000)), state: types.UndelegateRequestByIca,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(5000)), state: types.UndelegateRequestByIca,
				},
			},
			state: types.UndelegateRequestByIca,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records:   []*types.UndelegateRecordContent(nil),
				},
			},
		},
		{
			name:   "delete undelegateRecords case 3",
			zoneId: zoneId,
			records: []*UndelegateRecord{
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(1500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(1500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(2000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(2000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(3500, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(3500)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(4000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(4000)), state: types.UndelegateRequestByUser,
				},
				{
					zoneId: zoneId, withdrawer: delegator.String(), delegator: delegator.String(), claimer: delegator, snAsset: sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(5000, 18)), withdrawAsset: sdk.NewCoin(baseDenom, sdk.NewInt(5000)), state: types.UndelegateRequestByUser,
				},
			},
			state: types.UndelegateRequestByIca,
			result: []*types.UndelegateRecord{
				{
					ZoneId:    zoneId,
					Delegator: delegator.String(),
					Records: []*types.UndelegateRecordContent{
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(1500, 18),
							},
							WithdrawAmount:    sdk.NewInt(1500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(2000, 18),
							},
							WithdrawAmount:    sdk.NewInt(2000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(3500, 18),
							},
							WithdrawAmount:    sdk.NewInt(3500),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
						{
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(4000, 18),
							},
							WithdrawAmount:    sdk.NewInt(4000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						}, {
							Withdrawer: delegator.String(),
							SnAssetAmount: &sdk.Coin{
								baseSnDenom, sdk.NewIntWithDecimal(5000, 18),
							},
							WithdrawAmount:    sdk.NewInt(5000),
							State:             types.UndelegateRequestByUser,
							OracleVersion:     1,
							UndelegateVersion: 1,
						},
					},
				},
			},
		},
		{
			name:   "delete undelegateRecords case 4",
			zoneId: "test",
			records: []*UndelegateRecord{
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
			},
			state:  types.UndelegateRequestByIca,
			result: []*types.UndelegateRecord(nil),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.NewUndelegateRecord(tc.records)

			suite.App.GalKeeper.DeleteUndelegateRecords(suite.Ctx, tc.zoneId, types.UndelegateRequestByIca)

			result := suite.App.GalKeeper.GetAllUndelegateRecord(suite.Ctx, tc.zoneId)
			suite.Require().Equal(tc.result, result)

			suite.App.GalKeeper.DeleteUndelegateRecords(suite.Ctx, tc.zoneId, types.UndelegateRequestByUser)
		})
	}
}
