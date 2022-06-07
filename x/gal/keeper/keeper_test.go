package keeper_test

import (
	"github.com/Carina-labs/nova/app"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	app.KeeperTestHelper
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.setRandomState()
}

func (suite *KeeperTestSuite) setRandomState() {
	for _, acc := range suite.TestAccs {
		err := suite.App.GalKeeper.SetShare(suite.Ctx, acc, 0.2)
		if err != nil {
			panic(err)
		}
	}
}

func (suite *KeeperTestSuite) TestGetShares() {
	tcs := []struct {
		expected float64
	}{
		{
			expected: 0.2,
		},
		{
			expected: 0.2,
		},
		{
			expected: 0.2,
		},
		{
			expected: 0.2,
		},
		{
			expected: 0.2,
		},
	}

	for i, tc := range tcs {
		shares, err := suite.App.GalKeeper.GetShare(suite.Ctx, suite.TestAccs[i])
		suite.NoError(err)
		suite.Same(tc.expected, shares.Shares)
	}
}
