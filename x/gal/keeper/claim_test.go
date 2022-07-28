package keeper_test

import (
	"math/big"

	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestCalculateMintAmount() {
	tcs := []struct {
		userDepositAmt        int64
		totalShareTokenSupply int64
		totalStakedAmount     int64
		expected              int64
	}{
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              1000_000000,
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              1000_000000,
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 41000_000000,
			totalStakedAmount:     42000_000000,
			expected:              976_190476,
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 41976190480,
			totalStakedAmount:     44500_000000,
			expected:              943_285179,
		},
		{
			userDepositAmt:        1000_000000,
			totalShareTokenSupply: 42919_475660,
			totalStakedAmount:     47500_000000,
			expected:              903_567908,
		},
	}

	for _, tc := range tcs {
		userDepositAmt := big.NewInt(tc.userDepositAmt)
		totalShareTokenSupply := big.NewInt(tc.totalShareTokenSupply)
		totalStakedAmount := big.NewInt(tc.totalStakedAmount)

		res := suite.App.GalKeeper.CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
		suite.Equal(tc.expected, res.Int64())
	}
}

func (suite *KeeperTestSuite) TestCalculateBurnAmount() {
	tcs := []struct {
		userBurnStTokenAmt    int64
		totalShareTokenSupply int64
		totalStakedAmount     int64
		expected              int64
	}{
		{
			userBurnStTokenAmt:    5000_000000,
			totalShareTokenSupply: 40000_000000,
			totalStakedAmount:     40000_000000,
			expected:              5000_000000,
		},
		{
			userBurnStTokenAmt:    943_285180,
			totalShareTokenSupply: 41976_190480,
			totalStakedAmount:     44500_000000,
			expected:              1000_000000,
		},
	}

	for _, tc := range tcs {
		burnedAmount := big.NewInt(tc.userBurnStTokenAmt)
		totalShareTokenSupply := big.NewInt(tc.totalShareTokenSupply)
		totalStakedAmount := big.NewInt(tc.totalStakedAmount)

		res := suite.App.GalKeeper.CalculateWithdrawAlpha(burnedAmount, totalShareTokenSupply, totalStakedAmount)
		suite.Equal(tc.expected, res.Int64())
	}
}

func (suite *KeeperTestSuite) TestGetTotalStakedForLazyMinting() {
	type depositInfo struct {
		address string
		amount  sdk.Coin
	}

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.Ctx, "transfer", "channel-0", "stake")
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
					OracleOperators: []string{
						operator,
					},
				},
				States: []oracletypes.ChainInfo{
					{
						Coin:            tc.stakedAmount,
						ChainId:         "stake-1",
						Decimal:         8,
						OperatorAddress: operator,
					},
				},
			})
			suite.App.IbcstakingKeeper.RegisterZone(suite.Ctx, &ibcstakingtypes.RegisteredZone{
				ZoneId:            "stake-1",
				IcaConnectionInfo: nil,
				IcaAccount:        nil,
				ValidatorAddress:  "",
				BaseDenom:         "stake",
				SnDenom:           "snstake",
			})
			for _, item := range tc.depositInfo {
				ibcAmount := sdk.NewInt64Coin(ibcDenom, item.amount.Amount.Int64())
				record := types.DepositRecord{
					ZoneId:  "stake-1",
					Address: item.address,
					Records: []*types.DepositRecordContent{
						{
							Amount: &ibcAmount,
							State:  4,
						},
					},
				}
				suite.App.GalKeeper.SetDepositAmt(suite.Ctx, &record)
			}

			// execute
			res, err := suite.App.GalKeeper.GetTotalStakedForLazyMinting(suite.Ctx, "stake", "transfer", "channel-0")

			// verify
			suite.Require().NoError(err)
			suite.Require().True(tc.expect.IsEqual(res))
		})
	}
}
