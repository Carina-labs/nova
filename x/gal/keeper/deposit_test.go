package keeper_test

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

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
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  "stake",
						},
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositRequest,
					},
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  "stake",
						},
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositRequest,
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(10000),
							Denom:  "stake",
						},
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositRequest,
					},
					{
						Depositor: depositor.String(),
						Amount: &sdk.Coin{
							Amount: sdk.NewInt(20000),
							Denom:  "stake",
						},
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositRequest,
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
	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
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
					ZoneId:  zoneId,
					Claimer: depositor,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
						},
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
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
					ZoneId:  zoneId,
					Claimer: depositor,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DepositRequest,
						},
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
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
					ZoneId:  zoneId,
					Claimer: depositor,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DepositRequest,
						},
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DepositRequest,
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
					ZoneId:  zoneId,
					Claimer: depositor,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
						},
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  ibcDenom,
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateSuccess,
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
					ZoneId:  zoneId,
					Claimer: depositor,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  "stake",
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
						},
						{
							Depositor: depositor,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(20000),
								Denom:  "stake",
							},
							OracleVersion:   1,
							DelegateVersion: 1,
							State:           types.DelegateRequest,
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

			result := suite.App.GalKeeper.GetTotalDepositAmtForZoneId(suite.Ctx, tc.zoneId, ibcDenom, types.DelegateRequest)
			suite.Require().Equal(tc.result, result)
		})
	}
}

func (suite *KeeperTestSuite) TestTotalDepositAmtForUserAddr() {
	depositor1 := suite.GenRandomAddress().String()
	depositor2 := suite.GenRandomAddress().String()
	testUser := suite.GenRandomAddress().String()

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name      string
		depositor string
		zoneId    string
		denom     string
		result    sdk.Coin
	}{
		{
			name:      "get depositor1 amount",
			depositor: depositor1,
			zoneId:    zoneId,
			denom:     ibcDenom,
			result:    sdk.NewCoin(ibcDenom, sdk.NewInt(40000)),
		},
		{
			name:      "get depositor2 amount",
			depositor: depositor2,
			zoneId:    zoneId,
			denom:     ibcDenom,
			result:    sdk.NewCoin(ibcDenom, sdk.NewInt(15000)),
		},
		{
			name:      "user address not found",
			depositor: testUser,
			zoneId:    zoneId,
			denom:     ibcDenom,
			result:    sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
		},
		{
			name:      "denom not found",
			depositor: depositor2,
			zoneId:    zoneId,
			denom:     "nil",
			result:    sdk.NewCoin("nil", sdk.NewInt(0)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor1,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor1,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State:           types.DepositSuccess,
							OracleVersion:   1,
							DelegateVersion: 1,
						},
						{
							Depositor: depositor1,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State:           types.DepositSuccess,
							OracleVersion:   1,
							DelegateVersion: 1,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: depositor2,
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor2,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State:           types.DepositSuccess,
							OracleVersion:   1,
							DelegateVersion: 1,
						},
						{
							Depositor: depositor2,
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State:           types.DelegateSuccess,
							OracleVersion:   1,
							DelegateVersion: 1,
						},
					},
				},
			}
			for _, record := range depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			result := suite.App.GalKeeper.GetTotalDepositAmtForUserAddr(suite.Ctx, tc.depositor, tc.denom)
			fmt.Println(result)
			fmt.Println(tc.result)
			suite.Require().Equal(tc.result, result)
		})
	}
}

func (suite *KeeperTestSuite) TestDepositOracleVersion() {
	depositor := suite.GenRandomAddress()

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name          string
		zoneId        string
		state         int64
		oracleVersion uint64
	}{
		{
			name:          "oracle version is 1",
			zoneId:        zoneId,
			state:         types.DepositSuccess,
			oracleVersion: 1,
		},
		{
			name:          "oracle version is 2",
			zoneId:        zoneId,
			state:         types.DepositSuccess,
			oracleVersion: 2,
		},
		{
			name:          "oracle version is 3",
			zoneId:        zoneId,
			state:         types.DepositSuccess,
			oracleVersion: 3,
		},
		{
			name:          "oracle version is 4",
			zoneId:        zoneId,
			state:         types.DepositSuccess,
			oracleVersion: 4,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
			}
			for _, record := range depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.SetDepositOracleVersion(suite.Ctx, tc.zoneId, tc.state, tc.oracleVersion)

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, depositor)
			suite.Require().True(ok)

			for _, record := range result.Records {
				if record.State == tc.state {
					suite.Require().Equal(record.OracleVersion, tc.oracleVersion)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestChangeDepositState() {
	depositor := suite.GenRandomAddress()

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name      string
		zoneId    string
		preState  int64
		postState int64
		result    []int64
	}{
		{
			name:      "state change : DepositRequest - DepositSuccess",
			zoneId:    zoneId,
			preState:  types.DelegateRequest,
			postState: types.DelegateSuccess,
			result: []int64{
				types.DepositRequest,
				types.DelegateSuccess,
				types.DepositRequest,
				types.DelegateSuccess,
			},
		},
		{
			name:      "state change : DelegateRequest - DelegateSuccess",
			zoneId:    zoneId,
			preState:  types.DepositRequest,
			postState: types.DepositSuccess,
			result: []int64{
				types.DepositSuccess,
				types.DelegateRequest,
				types.DepositSuccess,
				types.DelegateRequest,
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DelegateRequest,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Depositor: depositor.String(),
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

			suite.App.GalKeeper.ChangeDepositState(suite.Ctx, tc.zoneId, tc.preState, tc.postState)

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, depositor)
			suite.Require().True(ok)

			for i, record := range result.Records {
				suite.Require().Equal(record.State, tc.result[i])
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDelegateRecordVersion() {
	depositor := suite.GenRandomAddress()

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name            string
		zoneId          string
		state           int64
		delegateVersion uint64
	}{
		{
			name:            "delegate version 1",
			zoneId:          zoneId,
			state:           types.DelegateSuccess,
			delegateVersion: 1,
		},
		{
			name:            "delegate version 2",
			zoneId:          zoneId,
			state:           types.DelegateSuccess,
			delegateVersion: 2,
		},
		{
			name:            "delegate version 3",
			zoneId:          zoneId,
			state:           types.DelegateSuccess,
			delegateVersion: 3,
		},
		{
			name:            "delegate version 4",
			zoneId:          zoneId,
			state:           types.DelegateSuccess,
			delegateVersion: 4,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DepositSuccess,
						},
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: depositor.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
						{
							Depositor: depositor.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
			}
			for _, record := range depositRecord {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, record)
			}

			suite.App.GalKeeper.SetDelegateRecordVersion(suite.Ctx, tc.zoneId, tc.state, tc.delegateVersion)

			result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, depositor)
			suite.Require().True(ok)

			for _, record := range result.Records {
				if record.State == tc.state {
					suite.Require().Equal(record.DelegateVersion, tc.delegateVersion)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteRecordedDepositItem() {
	depositor := suite.GenRandomAddress().String()

	tcs := []struct {
		name          string
		depositor     string
		depositRecord types.DepositRecord
		result        types.DepositRecord
		wantErr       bool
	}{
		{
			name:      "success",
			depositor: depositor,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor,

				Records: []*types.DepositRecordContent{
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DelegateSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DelegateSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(2000),
						},
					},
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DelegateSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor,
				Records: []*types.DepositRecordContent(nil),
			},
			wantErr: false,
		},
		{
			name:      "no deleted item",
			depositor: depositor,
			depositRecord: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor,
				Records: []*types.DepositRecordContent{
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DelegateRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
			result: types.DepositRecord{
				ZoneId:  zoneId,
				Claimer: depositor,
				Records: []*types.DepositRecordContent{
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DepositSuccess,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(1000),
						},
					},
					{
						Depositor:       depositor,
						OracleVersion:   1,
						DelegateVersion: 1,
						State:           types.DelegateRequest,
						Amount: &sdk.Coin{
							Denom:  "stake",
							Amount: sdk.NewInt(3000),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name:          "no deposited history, error case",
			depositor:     depositor,
			depositRecord: types.DepositRecord{},
			result:        types.DepositRecord{},
			wantErr:       true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			if tc.depositRecord.Size() > 0 {
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, &tc.depositRecord)
			}

			testAcc, _ := sdk.AccAddressFromBech32(tc.depositor)
			err := suite.App.GalKeeper.DeleteRecordedDepositItem(suite.Ctx, zoneId, testAcc, types.DelegateSuccess)
			if !tc.wantErr {
				result, ok := suite.App.GalKeeper.GetUserDepositRecord(suite.Ctx, zoneId, testAcc)
				suite.Require().True(ok)
				suite.Require().Equal(*result, tc.result)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetAllAmountNotMintShareToken() {
	depositor1 := suite.GenRandomAddress()
	depositor2 := suite.GenRandomAddress()

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name   string
		zoneId string
		result sdk.Coin
	}{
		{
			name:   "success",
			zoneId: zoneId,
			result: sdk.NewCoin(ibcDenom, sdk.NewInt(50000)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			depositRecord := []*types.DepositRecord{
				{
					ZoneId:  zoneId,
					Claimer: depositor1.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor1.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
						{
							Depositor: depositor1.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(30000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
					},
				},
				{
					ZoneId:  zoneId,
					Claimer: depositor2.String(),
					Records: []*types.DepositRecordContent{
						{
							Depositor: depositor2.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(15000),
								Denom:  ibcDenom,
							},
							State: types.DepositRequest,
						},
						{
							Depositor: depositor2.String(),
							Amount: &sdk.Coin{
								Amount: sdk.NewInt(10000),
								Denom:  ibcDenom,
							},
							State: types.DelegateSuccess,
						},
						{
							Depositor: depositor2.String(),
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

			zone := ibcstakingtypes.RegisteredZone{
				ZoneId: zoneId,
				IcaAccount: &ibcstakingtypes.IcaAccount{
					ControllerAddress: baseOwnerAcc.String(),
					HostAddress:       baseHostAcc.String(),
				},
				IcaConnectionInfo: &ibcstakingtypes.IcaConnectionInfo{
					ConnectionId: icaConnection,
					PortId:       zoneId + "." + baseOwnerAcc.String(),
				},
				TransferInfo: &ibcstakingtypes.TransferConnectionInfo{
					ChannelId: transferChannel,
					PortId:    transferPort,
				},
				BaseDenom: baseDenom,
				SnDenom:   baseSnDenom,
			}

			result, err := suite.App.GalKeeper.GetAllAmountNotMintShareToken(suite.Ctx, &zone)
			suite.Require().NoError(err)

			fmt.Println(result)
			fmt.Println(tc.result)
			suite.Require().Equal(result, tc.result)
		})
	}
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
