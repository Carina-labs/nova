package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) NewDelegateRecords(records []*DelegateRecord) {
	for _, record := range records {
		delegateRecord, ok := suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, record.zoneId, record.claimer)
		if !ok {
			delegateRecord = &types.DelegateRecord{
				ZoneId:  record.zoneId,
				Claimer: record.claimer.String(),
			}
			delegateRecord.Records = make(map[uint64]*types.DelegateRecordContent)
		}

		recordContent := types.DelegateRecordContent{
			Amount:        &record.amount,
			State:         record.state,
			OracleVersion: 1,
		}

		delegateRecord.Records[record.version] = &recordContent

		suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, delegateRecord)
	}
}

func (suite *KeeperTestSuite) TestSetDelegateRecord() {
	depositor := suite.GenRandomAddress()

	tcs := []struct {
		name           string
		zoneId         string
		depositor      sdk.AccAddress
		delegateRecord types.DelegateRecord
		result         types.DelegateRecord
		err            error
	}{
		{
			name:      "success delegate record",
			zoneId:    zoneId,
			depositor: depositor,
			delegateRecord: types.DelegateRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: map[uint64]*types.DelegateRecordContent{
					1: {
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseDenom,
						},
						State: types.DelegateRequest,
					},
					2: {
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  baseDenom,
						},
						State: types.DelegateRequest,
					},
				},
			},
			result: types.DelegateRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: map[uint64]*types.DelegateRecordContent{
					1: {
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseDenom,
						},
						State: types.DelegateRequest,
					},
					2: {
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  baseDenom,
						},
						State: types.DelegateRequest,
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, &tc.delegateRecord)

			result, ok := suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, zoneId, tc.depositor)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result, *result)
		})
	}
}

func (suite *KeeperTestSuite) TestSetDelegateRecords() {
	depositor := suite.GenRandomAddress()
	claimer := suite.GenRandomAddress()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	tcs := []struct {
		name          string
		zoneId        string
		depositor     sdk.AccAddress
		claimer       sdk.AccAddress
		depositRecord []*types.DepositRecord
		result        []*types.DelegateRecord
		err           bool
	}{
		{
			name:      "success",
			zoneId:    zoneId,
			depositor: depositor,
			claimer:   claimer,
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Claimer: claimer.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DepositSuccess,
						},
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DepositSuccess,
						},
						{
							Claimer: claimer.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DepositRequest,
						},
					},
				},
				{
					ZoneId:    zoneId,
					Depositor: claimer.String(),
					Records: []*types.DepositRecordContent{
						{
							Claimer: claimer.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DepositSuccess,
						},
						{
							Claimer: claimer.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State: types.DepositSuccess,
						},
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DepositRequest,
						},
					},
				},
			},
			result: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DelegateRequest,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: claimer.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(40000),
							},
							State: types.DelegateRequest,
						},
					},
				},
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, record := range tc.depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.SetDelegateRecords(suite.Ctx, zoneId)

			for _, res := range tc.result {
				claimerAddr, err := sdk.AccAddressFromBech32(res.Claimer)
				suite.Require().NoError(err)

				result, ok := suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, tc.zoneId, claimerAddr)
				suite.Require().True(ok)
				suite.Require().Equal(res, result)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDelegateOracleVersion() {
	depositor := suite.GenRandomAddress()

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name          string
		zoneId        string
		version       uint64
		oracleVersion uint64
	}{
		{
			name:          "oracle version is 1",
			zoneId:        zoneId,
			version:       1,
			oracleVersion: 1,
		},
		{
			name:          "oracle version is 2",
			zoneId:        zoneId,
			version:       2,
			oracleVersion: 2,
		},
		{
			name:          "oracle version is 3",
			zoneId:        zoneId,
			version:       1,
			oracleVersion: 3,
		},
		{
			name:          "oracle version is 4",
			zoneId:        zoneId,
			version:       2,
			oracleVersion: 4,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			delegateRecord := []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						1: {
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
						2: {
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
			}
			for _, record := range delegateRecord {
				suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.SetDelegateOracleVersion(suite.Ctx, tc.zoneId, tc.version, tc.oracleVersion)

			result, ok := suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, zoneId, depositor)
			suite.Require().True(ok)

			suite.Require().Equal(result.Records[tc.version].OracleVersion, tc.oracleVersion)

		})
	}
}

func (suite *KeeperTestSuite) TestChangeDelegateState() {
	claimer1 := suite.GenRandomAddress()
	claimer2 := suite.GenRandomAddress()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	tcs := []struct {
		name           string
		zoneId         string
		claimer        sdk.AccAddress
		delegateRecord []*types.DelegateRecord
		result         []*types.DelegateRecord
		version        uint64
		err            bool
	}{
		{
			name:   "success",
			zoneId: zoneId,
			delegateRecord: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: claimer1.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State: types.DelegateRequest,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: claimer2.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DelegateRequest,
						},
					},
				},
			},
			result: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: claimer1.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State: types.DelegateSuccess,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: claimer2.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DelegateSuccess,
						},
					},
				},
			},
			version: 0,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, record := range tc.delegateRecord {
				suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.ChangeDelegateState(suite.Ctx, tc.zoneId, tc.version)

			for i, res := range tc.result {
				claimer, err := sdk.AccAddressFromBech32(res.Claimer)
				suite.Require().NoError(err)
				suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, tc.zoneId, claimer)
				suite.Require().Equal(tc.result[i], res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetTotalDelegateAmtForZoneId() {
	claimer1 := suite.GenRandomAddress()
	claimer2 := suite.GenRandomAddress()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	tcs := []struct {
		name           string
		zoneId         string
		claimer        sdk.AccAddress
		delegateRecord []*types.DelegateRecord
		result         sdk.Coin
		version        uint64
		state          types.DelegateStatusType
		err            bool
	}{
		{
			name:   "success",
			zoneId: zoneId,
			delegateRecord: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: claimer1.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State: types.DelegateRequest,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: claimer2.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State: types.DelegateRequest,
						},
					},
				},
			},
			result:  sdk.NewCoin(ibcDenom, sdk.NewInt(30000)),
			state:   types.DelegateRequest,
			version: 0,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, record := range tc.delegateRecord {
				suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, record)
			}

			result := suite.App.GalKeeper.GetTotalDelegateAmtForZoneId(suite.Ctx, tc.zoneId, ibcDenom, tc.version, tc.state)
			suite.Require().Equal(tc.result, result)
		})
	}

}

func (suite *KeeperTestSuite) TestDeleteDelegateRecord() {
	claimer1 := suite.GenRandomAddress()
	claimer2 := suite.GenRandomAddress()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	tcs := []struct {
		name           string
		zoneId         string
		claimer        sdk.AccAddress
		delegateRecord []*types.DelegateRecord
		result         []*types.DelegateRecord
		oracleVersion  uint64
		version        uint64
		err            bool
	}{
		{
			name:   "success",
			zoneId: zoneId,
			delegateRecord: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: claimer1.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State:         types.DelegateRequest,
							OracleVersion: 1,
						},
						1: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(30000),
							},
							State:         types.DelegateSuccess,
							OracleVersion: 1,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: claimer2.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State:         types.DelegateSuccess,
							OracleVersion: 1,
						},
						1: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(10000),
							},
							State:         types.DelegateSuccess,
							OracleVersion: 1,
						},
					},
				},
			},
			result: []*types.DelegateRecord{
				{
					ZoneId:  zoneId,
					Claimer: claimer1.String(),
					Records: map[uint64]*types.DelegateRecordContent{
						0: {
							Amount: &sdk.Coin{
								Denom:  ibcDenom,
								Amount: sdk.NewInt(20000),
							},
							State:         types.DelegateRequest,
							OracleVersion: 1,
						},
					},
				},
				nil,
			},
			oracleVersion: 2,
			version:       0,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for i, record := range tc.delegateRecord {
				suite.App.GalKeeper.SetDelegateRecord(suite.Ctx, record)

				suite.App.GalKeeper.DeleteDelegateRecords(suite.Ctx, tc.delegateRecord[i], tc.oracleVersion)

				claimer, err := sdk.AccAddressFromBech32(tc.delegateRecord[i].Claimer)
				suite.Require().NoError(err)
				result, _ := suite.App.GalKeeper.GetUserDelegateRecord(suite.Ctx, tc.zoneId, claimer)
				suite.Require().Equal(tc.result[i], result)
			}
		})
	}
}
