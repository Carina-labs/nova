package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"math/big"
)

func (suite *KeeperTestSuite) TestTotalClaimableAssets() {
	claimer1 := suite.GenRandomAddress()
	claimer2 := suite.GenRandomAddress()
	claimer3 := suite.GenRandomAddress()

	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	denomTrace := transfertypes.DenomTrace{
		Path:      transferPort + "/" + transferChannel,
		BaseDenom: baseDenom,
	}
	suite.App.TransferKeeper.SetDenomTrace(suite.Ctx, denomTrace)

	suite.App.OracleKeeper.InitGenesis(suite.Ctx, &oracletypes.GenesisState{
		Params: oracletypes.Params{
			OracleKeyManager: []string{
				baseOwnerAcc.String(),
			},
		},
		OracleAddressInfo: []oracletypes.OracleAddressInfo{
			{
				ZoneId:        zoneId,
				OracleAddress: []string{baseOwnerAcc.String()},
			},
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            sdk.NewCoin(baseDenom, sdk.NewInt(80000)),
				ZoneId:          zoneId,
				OperatorAddress: baseOwnerAcc.String(),
			},
		},
	})

	delegateRecords := []*DelegateRecord{
		{zoneId: zoneId, claimer: claimer1, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(10000)), state: types.DelegateSuccess, version: 1},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(20000)), state: types.DelegateSuccess, version: 2},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(30000)), state: types.DelegateRequest, version: 3},
		{zoneId: zoneId, claimer: claimer1, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(40000)), state: types.DelegateRequest, version: 4},

		{zoneId: zoneId, claimer: claimer2, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(10000)), state: types.DelegateSuccess, version: 1},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer2, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(20000)), state: types.DelegateRequest, version: 2},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(30000)), state: types.DelegateRequest, version: 3},
		{zoneId: zoneId, claimer: claimer2, depositor: claimer1, amount: sdk.NewCoin(ibcDenom, sdk.NewInt(40000)), state: types.DelegateSuccess, version: 4},
	}
	suite.App.GalKeeper.SetAssetInfo(suite.Ctx, &types.AssetInfo{
		ZoneId:         zoneId,
		UnMintedWAsset: sdk.NewInt(80000),
	})
	suite.NewDelegateRecords(delegateRecords)

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
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(30000, 18)),
			oracleVersion: 2,
			err:           false,
		},
		{
			name:          "claimer2 claimable assets",
			claimer:       claimer2,
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(50000, 18)),
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
			result:        sdk.NewCoin(baseSnDenom, sdk.NewIntWithDecimal(0, 18)),
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
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	tcs := []struct {
		name           string
		stakedAmount   sdk.Coin
		delegateRecord []*DelegateRecord
		expect         sdk.Int
	}{
		{
			name:         "test case 1",
			stakedAmount: sdk.NewInt64Coin(baseDenom, 1_000_000),
			delegateRecord: []*DelegateRecord{
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 100_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 100_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 100_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
			},
			expect: sdk.NewInt(700_000),
		},
		{
			name:         "test case 2",
			stakedAmount: sdk.NewInt64Coin(baseDenom, 1_500_000),
			delegateRecord: []*DelegateRecord{
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 100_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 200_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
				{
					zoneId:    zoneId,
					claimer:   suite.GenRandomAddress(),
					depositor: suite.GenRandomAddress(),
					amount:    sdk.NewInt64Coin(ibcDenom, 300_000),
					state:     types.DelegateSuccess,
					version:   1,
				},
			},
			expect: sdk.NewInt(900_000),
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
						ZoneId: zoneId,
						OracleAddress: []string{
							operator,
						},
					},
				},
				States: []oracletypes.ChainInfo{
					{
						Coin:            tc.stakedAmount,
						ZoneId:          zoneId,
						OperatorAddress: operator,
					},
				},
			})

			baseTestZone := newBaseRegisteredZone()
			suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, baseTestZone)

			suite.NewDelegateRecords(tc.delegateRecord)
			assetInfo := types.AssetInfo{
				ZoneId:         zoneId,
				UnMintedWAsset: sdk.NewInt(0),
			}

			for _, record := range tc.delegateRecord {
				assetInfo.UnMintedWAsset = assetInfo.UnMintedWAsset.Add(record.amount.Amount)
			}

			suite.App.GalKeeper.SetAssetInfo(suite.Ctx, &assetInfo)
			// execute
			res, err := suite.App.GalKeeper.GetTotalStakedForLazyMinting(suite.Ctx, baseDenom, transferPort, transferChannel, assetInfo)
			// verify
			suite.Require().NoError(err)
			suite.Require().True(tc.expect.Equal(res))
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

func (suite *KeeperTestSuite) TestCheckDecimal() {
	tcs := []struct {
		name    string
		amount  sdk.Coin
		decimal int64
		result  sdk.Coin
		err     bool
	}{
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000)),
			err:     true,
		},
		{
			name:    "fail",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000)),
			err:     false,
		},
		{
			name:    "success",
			decimal: 6,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(100000000000)),
			err:     true,
		},
		{
			name:    "success",
			decimal: 7,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1000000000000)),
			err:     false,
		},
		{
			name:    "success",
			decimal: 18,
			amount:  sdk.NewCoin("snuatom", sdk.NewIntFromUint64(1)),
			err:     false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			err := suite.App.GalKeeper.CheckDecimal(tc.amount, tc.decimal)
			if tc.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
