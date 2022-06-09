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
			userDepositAmt:        950 * 1000000,
			totalShareTokenSupply: 900000000 * 1000000,
			totalStakedAmount:     2500000 * 1000000,
			expected:              342000 * 1000000,
		},
		{
			userDepositAmt:        123 * 1000000,
			totalShareTokenSupply: 871312859 * 1000000,
			totalStakedAmount:     2159872 * 1000000,
			expected:              49619367099,
		},
	}

	for _, tc := range tcs {
		userDepositAmt := big.NewInt(tc.userDepositAmt)
		totalShareTokenSupply := big.NewInt(tc.totalShareTokenSupply)
		totalStakedAmount := big.NewInt(tc.totalStakedAmount)

		res := suite.App.GalKeeper.CalculateMintAmount(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
		println(res.Int64() / 1000000)
		suite.Equal(tc.expected, res.Int64())
	}
}
