package keeper_test

import "math/big"

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

		res := suite.App.GalKeeper.CalculateAlpha(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
		println(res.Int64())
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

		res := suite.App.GalKeeper.CalculateWithdrawAmount(burnedAmount, totalShareTokenSupply, totalStakedAmount)
		print(res.Int64())
		suite.Equal(tc.expected, res.Int64())
	}
}
