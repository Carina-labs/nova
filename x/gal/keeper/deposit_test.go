package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type DelegateRecord struct {
	zoneId    string
	depositor sdk.AccAddress
	claimer   sdk.AccAddress
	amount    sdk.Coin
	state     int64
	version   uint64
}

func (suite *KeeperTestSuite) TestDepositRecord() {
	depositor := suite.GenRandomAddress()

	tcs := []struct {
		name          string
		depositor     sdk.AccAddress
		depositRecord types.DepositRecord
		result        types.DepositRecord
		err           error
	}{
		{
			name:      "success deposit records",
			depositor: depositor,
			depositRecord: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseDenom,
						},
						State: types.DepositRequest,
					},
					{
						Claimer: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  baseDenom,
						},
						State: types.DepositRequest,
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  baseDenom,
						},
						State: types.DepositRequest,
					},
					{
						Claimer: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  baseDenom,
						},
						State: types.DepositRequest,
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.App.GalKeeper.SetDepositRecord(suite.Ctx, &tc.depositRecord)

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, tc.depositor)
			suite.Require().True(ok)
			suite.Require().Equal(tc.result, *result)
		})
	}
}

func (suite *KeeperTestSuite) TestTotalDepositAmtForZoneId() {
	depositor := suite.GenRandomAddress().String()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name      string
		depositor string
		zoneId    string

		depositRecord []*types.DepositRecord
		result        sdk.Coin
	}{
		{
			name:      "get all deposit amount",
			depositor: depositor,
			zoneId:    zoneId,
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor,
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
					},
				},
			},
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(30000)),
		},
		{
			name:      "check deposit state1",
			depositor: depositor,
			zoneId:    zoneId,
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor,
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
					},
				},
			},
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(20000)),
		},
		{
			name:      "check deposit state2",
			depositor: depositor,
			zoneId:    zoneId,
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor,
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
					},
				},
			},
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
		},
		{
			name:      "zoneId is not found",
			depositor: depositor,
			zoneId:    "nil",
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor,
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateRequest,
						},
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
			},
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
		},
		{
			name:      "denom is not found",
			depositor: depositor,
			zoneId:    zoneId,
			depositRecord: []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor,
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  "stake",
							},
							State: types.DepositRequest,
						},
						{
							Claimer: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  "stake",
							},
							State: types.DepositRequest,
						},
					},
				},
			},
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, record := range tc.depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			result := suite.App.GalKeeper.GetTotalDepositAmtForZoneId(suite.Ctx, tc.zoneId, ibcDenom, types.DepositSuccess)
			suite.Require().Equal(tc.result, result)
		})
	}
}

func (suite *KeeperTestSuite) TestChangeDepositState() {
	depositor := suite.GenRandomAddress()

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name      string
		zoneId    string
		preState  int64
		postState int64
		result    []int64
	}{
		{
			name:      "state change",
			zoneId:    zoneId,
			preState:  types.DepositRequest,
			postState: types.DepositSuccess,
			result: []int64{
				types.DepositSuccess,
				types.DepositSuccess,
				types.DepositSuccess,
				types.DepositSuccess,
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:    zoneId,
					Depositor: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DelegateRequest,
						},
					},
				},
				{
					ZoneId:    zoneId,
					Depositor: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Claimer: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateRequest,
						},
					},
				},
			}
			for _, record := range depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.ChangeDepositState(suite.Ctx, tc.zoneId, depositor.String())

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, depositor)
			suite.Require().True(ok)

			for i, record := range result.Records {
				suite.Require().Equal(record.State, tc.result[i])
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteDepositRecords() {
	depositor := suite.GenRandomAddress().String()

	tcs := []struct {
		name          string
		depositor     string
		depositRecord types.DepositRecord
		result        types.DepositRecord
	}{
		{
			name:      "success",
			depositor: depositor,
			depositRecord: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor,
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor,
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
					{
						Claimer: depositor,
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(2000),
						},
					},
					{
						Claimer: depositor,
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor,
				Records:   []*types.DepositRecordContent(nil),
			},
		},
		{
			name:      "no deleted item",
			depositor: depositor,
			depositRecord: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor,
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor,
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Claimer: depositor,
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor,
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor,
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Claimer: depositor,
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			if tc.depositRecord.Size() > 0 {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, &tc.depositRecord)
			}

			testAcc, _ := sdk.AccAddressFromBech32(tc.depositor)
			suite.App.GalKeeper.DeleteDepositRecords(suite.Ctx, zoneId, types.DepositSuccess)

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, testAcc)
			suite.Require().True(ok)
			suite.Require().Equal(*result, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteRecordedDepositItem() {
	depositor := suite.GenRandomAddress()
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	tcs := []struct {
		name          string
		zoneId        string
		depositor     sdk.AccAddress
		amount        sdk.Int
		state         types.DepositStatusType
		depositRecord types.DepositRecord
		result        types.DepositRecord
		wantErr       bool
	}{
		{
			name:      "success",
			zoneId:    zoneId,
			amount:    sdk.NewInt(3000),
			depositor: depositor,
			state:     types.DepositRequest,
			depositRecord: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor.String(),
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3500),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DelegateSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(2000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Claimer: depositor.String(),
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositRequest,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3500),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DelegateSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(2000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Claimer: depositor.String(),
						State:   types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  ibcDenom,
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tcs {
		suite.App.GalKeeper.SetDepositRecord(suite.Ctx, &tc.depositRecord)

		err := suite.App.GalKeeper.DeleteRecordedDepositItem(suite.Ctx, tc.zoneId, tc.depositor, tc.state, tc.amount)
		suite.NoError(err)

		res, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, tc.zoneId, tc.depositor)
		suite.Require().True(ok)

		suite.Require().Equal(tc.result, *res)
	}
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
