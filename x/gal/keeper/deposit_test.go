package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// deposit content 생성

// func (suite *KeeperTestSuite) TestRecordDepositAmt() {
// 	depositor := suite.GenRandomAddress()
// 	claimer := suite.GenRandomAddress()
// 	type args struct {
// 		coin      sdk.Coin
// 		claimer   sdk.AccAddress
// 		depositor sdk.AccAddress
// 	}

// 	tcs := []struct {
// 		name    string
// 		args    []args
// 		expect  []args
// 		wantErr bool

// 		denom         string
// 		amt           int64
// 		userAddr      sdk.AccAddress
// 		expectedDenom string
// 		expectedAmt   int64
// 	}{
// 		{
// 			name: "should get recorded deposit amt",
// 			args: []args{
// 				{sdk.NewInt64Coin(baseDenom, 10000), depositor, claimer},
// 			},
// 			expect: []args{
// 				{sdk.NewInt64Coin(baseDenom, 10000), depositor, claimer},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "should not get deposit info",
// 			args: []args{},
// 			expect: []args{
// 				{sdk.NewInt64Coin(baseDenom, 10000), claimer, claimer},
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tc := range tcs {
// 		tc := tc

// 		suite.Run(tc.name, func() {
// 			suite.SetupTest()

// 			for _, arg := range tc.args {
// 				suite.App.GalKeeper.SetDepositAmt(
// 					suite.Ctx,
// 					&types.DepositRecord{
// 						ZoneId:  "test-zone-id",
// 						Claimer: string(arg.claimer),
// 						Records: []*types.DepositRecordContent{
// 							{
// 								Depositor: string(arg.depositor),
// 								Amount:    &arg.coin,
// 								State:     1,
// 							},
// 						},
// 					})
// 			}

// 			// totalAmt는 record.totalAmt에서 구할 수 있음
// 			for _, query := range tc.expect {
// 				res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, "test-zone-id", query.claimer)
// 				if tc.wantErr {
// 					suite.Require().NotNil(err, "error expected but no error found")
// 					suite.Require().Equal(err, types.ErrNoDepositRecord)
// 					continue
// 				}

// 				suite.Require().NoError(err)
// 				for _, record := range res.Records {
// 					suite.Require().Equal(record.Amount.Denom, query.coin.Denom)
// 					suite.Require().Equal(record.Amount.Amount, query.coin.Amount)
// 					suite.Require().Equal(res.Claimer, query.claimer.String())
// 				}
// 			}
// 		})
// 	}
// }

// func (suite *KeeperTestSuite) TestDeleteRecordedDepositItem() {
// 	type inputType struct {
// 		address string
// 		amount  sdk.Coin
// 	}
// 	testAddress := suite.GenRandomAddress().String()
// 	testAddress2 := suite.GenRandomAddress().String()
// 	tcs := []struct {
// 		name       string
// 		input      []inputType
// 		depositor  string
// 		deleteItem sdk.Coin
// 		result     []inputType
// 		wantErr    bool
// 	}{
// 		{
// 			name:      "valid case",
// 			depositor: testAddress,
// 			input: []inputType{
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 1000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 2000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 3000),
// 				},
// 			},
// 			deleteItem: sdk.NewInt64Coin("stake", 2000),
// 			result: []inputType{
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 1000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 3000),
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:      "no deleted item, but no error case",
// 			depositor: testAddress,
// 			input: []inputType{
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 1000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 3000),
// 				},
// 			},
// 			deleteItem: sdk.NewInt64Coin("stake", 4000),
// 			result: []inputType{
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 1000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 3000),
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:      "no deposited history, error case",
// 			depositor: testAddress2,
// 			input: []inputType{
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 1000),
// 				},
// 				{
// 					address: testAddress,
// 					amount:  sdk.NewInt64Coin("stake", 3000),
// 				},
// 			},
// 			deleteItem: sdk.NewInt64Coin("stake", 4000),
// 			result:     nil,
// 			wantErr:    true,
// 		},
// 	}

// 	for _, tc := range tcs {
// 		suite.Run(tc.name, func() {
// 			// setup
// 			record := types.DepositRecord{}
// 			record.Claimer = testAddress
// 			for _, item := range tc.input {
// 				coin := sdk.NewInt64Coin(item.amount.Denom, item.amount.Amount.Int64())
// 				record.Records = append(record.Records, &types.DepositRecordContent{
// 					State:  true,
// 					Amount: &coin,
// 				})
// 			}
// 			suite.App.GalKeeper.SetDepositAmt(suite.Ctx, &record)

// 			// execute
// 			testAcc, _ := sdk.AccAddressFromBech32(tc.depositor)
// 			err := suite.App.GalKeeper.DeleteRecordedDepositItem(suite.Ctx, testAcc, tc.deleteItem)
// 			if !tc.wantErr {
// 				suite.Require().NoError(err)
// 				result, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, testAcc)
// 				suite.Require().NoError(err)
// 				for i, resultItem := range result.Records {
// 					suite.Require().True(resultItem.Amount.IsEqual(tc.result[i].amount))
// 				}
// 			} else {
// 				suite.Require().Error(err)
// 			}
// 		})
// 	}
// }

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
