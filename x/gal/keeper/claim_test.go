package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func (suite *KeeperTestSuite) TestTotalClaimableAssets() {

	claimer1 := suite.GenRandomAddress()
	claimer2 := suite.GenRandomAddress()
	claimer3 := suite.GenRandomAddress()

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	depositRecords := []*DepositRecord{
		{zoneId: zoneId, claimer: claimer1, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(10000)), state: types.DelegateSuccess},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(20000)), state: types.DelegateSuccess},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(30000)), state: types.DepositSuccess},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(40000)), state: types.DelegateRequest},

		{zoneId: zoneId, claimer: claimer2, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(10000)), state: types.DelegateSuccess},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(20000)), state: types.DelegateRequest},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(30000)), state: types.DepositRequest},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(40000)), state: types.DelegateSuccess},
	}
	suite.NewDepositRecords(depositRecords)

	tcs := []struct {
		name          string
		claimer       sdk.AccAddress
		result        sdk.Coin
		oracleVersion uint64
		err           bool
	}{
		{
			name:          "claimer1 claimable assets",
			claimer:       claimer1,
			result:        sdk.NewCoin(ibcDenom, sdk.NewInt(30000)),
			oracleVersion: 2,
			err:           false,
		},
		{
			name:          "claimer2 claimable assets",
			claimer:       claimer2,
			result:        sdk.NewCoin(ibcDenom, sdk.NewInt(50000)),
			oracleVersion: 3,
			err:           false,
		},
		{
			name:          "claimer3 claimable assets",
			claimer:       claimer3,
			oracleVersion: 4,
			err:           true,
		},
		{
			name:          "oracleVersion issue",
			claimer:       claimer1,
			result:        sdk.NewCoin(ibcDenom, sdk.NewInt(0)),
			oracleVersion: 1,
			err:           false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			zone, ok := suite.App.IcaControlKeeper.GetRegisteredZone(suite.Ctx, zoneId)
			suite.Require().True(ok)

			trace := oracletypes.IBCTrace{
				Version: tc.oracleVersion,
				Height:  uint64(suite.Ctx.BlockHeight()),
			}
			suite.App.OracleKeeper.SetOracleVersion(suite.Ctx, zoneId, trace)
			result, err := suite.App.GalKeeper.TotalClaimableAssets(suite.Ctx, zone, tc.claimer)

			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.result, *result)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestCalculateDepositAlpha() {
	tcs := []struct {
		userDepositAmt        int64
		totalShareTokenSupply int64
		totalStakedAmount     int64
		expected              *big.Int
	}{
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              big.NewInt(1000_000000),
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              big.NewInt(1000_000000),
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 41000_000000,
			totalStakedAmount:     42000_000000,
			expected:              big.NewInt(976_190476),
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 41976190480,
			totalStakedAmount:     44500_000000,
			expected:              big.NewInt(943_285179),
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 42919_475660,
			totalStakedAmount:     47500_000000,
			expected:              big.NewInt(903_567908),
		},
	}

	for _, tc := range tcs {
		userDepositAmt := big.NewInt(tc.userDepositAmt)
		totalShareTokenSupply := big.NewInt(tc.totalShareTokenSupply)
		totalStakedAmount := big.NewInt(tc.totalStakedAmount)

		res := suite.App.GalKeeper.CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
		suite.Equal(tc.expected, res)
	}
}

func (suite *KeeperTestSuite) TestCalculateWithdrawAlpha() {
	tcs := []struct {
		userBurnStTokenAmt    int64
		totalShareTokenSupply int64
		totalStakedAmount     int64
		expected              *big.Int
	}{
		{
			userBurnStTokenAmt:    5000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              big.NewInt(5000_000000),
		},
		{
			userBurnStTokenAmt:    943_285180,
			totalShareTokenSupply: 41976_190480,
			totalStakedAmount:     44500_000000,
			expected:              big.NewInt(1000_000000),
		},
	}

	for _, tc := range tcs {
		burnedAmount := tc.userBurnStTokenAmt
		totalShareTokenSupply := tc.totalShareTokenSupply
		totalStakedAmount := tc.totalStakedAmount

		res := suite.App.GalKeeper.CalculateWithdrawAlpha(big.NewInt(burnedAmount), big.NewInt(totalShareTokenSupply), big.NewInt(totalStakedAmount))
		suite.Equal(tc.expected, res)
	}
}

func (suite *KeeperTestSuite) TestGetTotalStakedForLazyMinting() {
	type depositInfo struct {
		address string
		amount  sdk.Coin
	}

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom("transfer", "channel-0", "stake")
	tcs := []struct {
		name         string
		stakedAmount sdk.Coin
		depositInfo  []depositInfo
		expect       sdk.Coin
	}{
		{
			name:         "test case 1",
			stakedAmount: sdk.NewInt64Coin("stake", 1_000_000),
			depositInfo: []depositInfo{
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 100_000),
				},
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 100_000),
				},
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 100_000),
				},
			},
			expect: sdk.NewInt64Coin(ibcDenom, 700_000),
		},
		{
			name:         "test case 2",
			stakedAmount: sdk.NewInt64Coin("stake", 1_500_000),
			depositInfo: []depositInfo{
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 100_000),
				},
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 200_000),
				},
				{
					address: suite.GenRandomAddress().String(),
					amount:  sdk.NewInt64Coin("stake", 300_000),
				},
			},
			expect: sdk.NewInt64Coin(ibcDenom, 900_000),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// setup
			suite.Setup()
			operator := suite.GenRandomAddress().String()
			suite.App.OracleKeeper.InitGenesis(suite.Ctx, &oracletypes.GenesisState{
				Params: oracletypes.Params{
					OracleKeyManager: []string{
						operator,
					},
				},
				OracleAddressInfo: []oracletypes.OracleAddressInfo{
					{
						ZoneId: "stake-1",
						OracleAddress: []string{
							operator,
						},
					},
				},
				States: []oracletypes.ChainInfo{
					{
						Coin:            tc.stakedAmount,
						ZoneId:          "stake-1",
						OperatorAddress: operator,
					},
				},
			})

			baseTestZone := newBaseRegisteredZone()
			baseTestZone.ZoneId = "stake-1"
			baseTestZone.BaseDenom = "stake"
			baseTestZone.SnDenom = "snstake"
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, baseTestZone)

			for _, item := range tc.depositInfo {
				ibcAmount := sdk.NewInt64Coin(ibcDenom, item.amount.Amount.Int64())
				record := types.DepositRecord{
					ZoneId:  "stake-1",
					Claimer: item.address,
					Records: []*types.DepositRecordContent{
						{
							Amount: &ibcAmount,
							State:  types.DelegateSuccess,
						},
					},
				}
				suite.App.GalKeeper.SetDepositRecord(suite.Ctx, &record)
			}

			// execute
			res, err := suite.App.GalKeeper.GetTotalStakedForLazyMinting(suite.Ctx, "stake", "transfer", "channel-0")
			// verify
			suite.Require().NoError(err)
			suite.Require().True(tc.expect.IsEqual(res))
		})
	}
}

func (suite *KeeperTestSuite) TestConvertWAssetToSnAssetDecimal() {
	tcs := []struct {
		name    string
		amount  sdk.Coin
		decimal int64
		result  sdk.Coin
	}{
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
		},
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1300000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1300000000000000000)),
		},
		{
			name:    "success",
			decimal: 7,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(150000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(15000000000000000000)),
		},
		{
			name:    "success",
			decimal: 18,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			result, err := suite.App.GalKeeper.ConvertWAssetToSnAssetDecimal(tc.amount.Amount.BigInt(), tc.decimal, "snuatom")

			suite.Require().NoError(err)
			suite.Require().Equal(tc.result, *result)
		})
	}
}

func (suite *KeeperTestSuite) TestConvertSnAssetToWAssetDecimal() {
	tcs := []struct {
		name    string
		amount  sdk.Coin
		decimal int64
		result  sdk.Coin
	}{
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000)),
		},
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1300000000000000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1300000)),
		},
		{
			name:    "success",
			decimal: 7,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(15000000000000000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(150000000)),
		},
		{
			name:    "success",
			decimal: 18,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
			result:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000000000)),
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			result, err := suite.App.GalKeeper.ConvertSnAssetToWAssetDecimal(tc.amount.Amount.BigInt(), tc.decimal, "snuatom")

			suite.Require().NoError(err)
			suite.Require().Equal(tc.result, *result)
		})
	}
}
