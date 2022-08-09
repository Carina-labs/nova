package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestCalculateMintAmount() {
	tcs := []struct {
		userDepositAmt        sdk.Dec
		totalShareTokenSupply sdk.Dec
		totalStakedAmount     sdk.Dec
		expected              sdk.Int
	}{
		{
			userDepositAmt:        sdk.NewDec(1000_000000),
			totalShareTokenSupply: sdk.NewDec(40000_000000),
			totalStakedAmount:     sdk.NewDec(40000_000000),
			expected:              sdk.NewInt(1000_000000),
		},
		{
			userDepositAmt:        sdk.NewDec(1000_000000),
			totalShareTokenSupply: sdk.NewDec(40000_000000),
			totalStakedAmount:     sdk.NewDec(40000_000000),
			expected:              sdk.NewInt(1000_000000),
		},
		{
			userDepositAmt:        sdk.NewDec(1000_000000),
			totalShareTokenSupply: sdk.NewDec(41000_000000),
			totalStakedAmount:     sdk.NewDec(42000_000000),
			expected:              sdk.NewInt(976_190476),
		},
		{
			userDepositAmt:        sdk.NewDec(1000_000000),
			totalShareTokenSupply: sdk.NewDec(41976190480),
			totalStakedAmount:     sdk.NewDec(44500_000000),
			expected:              sdk.NewInt(943_285179),
		},
		{
			userDepositAmt:        sdk.NewDec(1000_000000),
			totalShareTokenSupply: sdk.NewDec(42919_475660),
			totalStakedAmount:     sdk.NewDec(47500_000000),
			expected:              sdk.NewInt(903_567908),
		},
	}

	for _, tc := range tcs {
		userDepositAmt := tc.userDepositAmt
		totalShareTokenSupply := tc.totalShareTokenSupply
		totalStakedAmount := tc.totalStakedAmount

		res := suite.App.GalKeeper.CalculateDepositAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
		suite.Equal(tc.expected, res)
	}
}

func (suite *KeeperTestSuite) TestCalculateBurnAmount() {
	tcs := []struct {
		userBurnStTokenAmt    sdk.Dec
		totalShareTokenSupply sdk.Dec
		totalStakedAmount     sdk.Dec
		expected              sdk.Int
	}{
		{
			userBurnStTokenAmt:    sdk.NewDec(5000_000000),
			totalShareTokenSupply: sdk.NewDec(40000_000000),
			totalStakedAmount:     sdk.NewDec(40000_000000),
			expected:              sdk.NewInt(5000_000000),
		},
		{
			userBurnStTokenAmt:    sdk.NewDec(943_285180),
			totalShareTokenSupply: sdk.NewDec(41976_190480),
			totalStakedAmount:     sdk.NewDec(44500_000000),
			expected:              sdk.NewInt(1000_000000),
		},
	}

	for _, tc := range tcs {
		burnedAmount := tc.userBurnStTokenAmt
		totalShareTokenSupply := tc.totalShareTokenSupply
		totalStakedAmount := tc.totalStakedAmount

		res := suite.App.GalKeeper.CalculateWithdrawAlpha(burnedAmount, totalShareTokenSupply, totalStakedAmount)
		suite.Equal(tc.expected, res)
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
					Claimer: item.address,
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
