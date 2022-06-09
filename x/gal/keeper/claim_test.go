package keeper_test

import "math/big"

func (suite *KeeperTestSuite) TestCalculateMintAmount() {
	userDepositAmt := big.NewInt(1000)
	totalShareTokenSupply := big.NewInt(10000)
	totalStakedAmount := big.NewInt(1000000)

	res := suite.App.GalKeeper.CalculateMintAmount(userDepositAmt, totalShareTokenSupply, totalStakedAmount)
	suite.Equal(int64(10), res.Int64())
}
