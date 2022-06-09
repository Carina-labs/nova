package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *KeeperTestSuite) TestGetCachedDepositAmt() {
	tcs := []struct {
		expected sdk.Coin
	}{
		{
			expected: sdk.NewInt64Coin("uosmo", 2),
		},
		{
			expected: sdk.NewInt64Coin("uosmo", 2),
		},
		{
			expected: sdk.NewInt64Coin("uosmo", 2),
		},
		{
			expected: sdk.NewInt64Coin("uosmo", 2),
		},
		{
			expected: sdk.NewInt64Coin("uosmo", 2),
		},
	}

	for i, tc := range tcs {
		res, err := suite.App.GalKeeper.GetCachedDepositAmt(suite.Ctx, suite.TestAccs[i])
		suite.NoError(err)
		suite.Equal(tc.expected, res.Amount)
	}
}

func (suite *KeeperTestSuite) TestCacheDepositAmt() {
	err := suite.App.GalKeeper.CacheDepositAmt(suite.Ctx, suite.TestAccs[0], sdk.NewInt64Coin("atom", 1000))
	suite.NoError(err)

	res, err := suite.App.GalKeeper.GetCachedDepositAmt(suite.Ctx, suite.TestAccs[0])
	suite.Equal(suite.TestAccs[0].String(), res.Address)
	suite.Equal("atom", res.Amount.Denom)
	suite.Equal(sdk.NewInt(1000), res.Amount.Amount)
}

func (suite *KeeperTestSuite) TestClearCachedDepositAmt() {
	keeper := suite.App.GalKeeper
	err := keeper.CacheDepositAmt(suite.Ctx, suite.TestAccs[0], sdk.NewInt64Coin("atom", 1000))
	suite.NoError(err)

	err = keeper.ClearCachedDepositAmt(suite.Ctx, suite.TestAccs[0])
	suite.NoError(err)

	res, err := keeper.GetCachedDepositAmt(suite.Ctx, suite.TestAccs[0])
	suite.Nil(res)
	suite.Error(err)
}
