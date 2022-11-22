package keeper_test

import (
	"github.com/Carina-labs/nova/x/poolincentive/types"
)

func (suite *KeeperTestSuite) TestHandleReplacePoolIncentivesProposal() {
	tcs := []struct {
		name        string
		preset      []types.IncentivePool
		replacePool []types.IncentivePool
		shouldErr   bool
	}{
		{
			name: "valid case",
			preset: []types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "contract-1",
					Weight:              5,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "contract-2",
					Weight:              3,
				},
				{
					PoolId:              "pool-3",
					PoolContractAddress: "contract-3",
					Weight:              2,
				},
			},
			replacePool: []types.IncentivePool{
				{
					PoolId:              "pool-4",
					PoolContractAddress: "contract-4",
					Weight:              10,
				},
				{
					PoolId:              "pool-5",
					PoolContractAddress: "contract-5",
					Weight:              11,
				},
				{
					PoolId:              "pool-6",
					PoolContractAddress: "contract-6",
					Weight:              12,
				},
			},
			shouldErr: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// setup
			for i := range tc.preset {
				err := suite.App.PoolKeeper.CreateIncentivePool(suite.Ctx, &tc.preset[i])
				suite.NoError(err)
			}

			msg := types.ReplacePoolIncentivesProposal{
				Title:         "title",
				Description:   "description",
				NewIncentives: tc.replacePool,
			}

			err := suite.App.PoolKeeper.HandleReplacePoolIncentivesProposal(suite.Ctx, &msg)
			suite.NoError(err)

			if !tc.shouldErr {
				for _, pool := range tc.replacePool {
					res, err := suite.App.PoolKeeper.FindIncentivePoolById(suite.Ctx, pool.PoolId)
					suite.NoError(err)
					suite.Equal(pool, *res)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleUpdatePoolIncentivesProposal() {
	tcs := []struct {
		name                string
		preset              []types.IncentivePool
		updatePool          []types.IncentivePool
		expected            []types.IncentivePool
		expectedTotalWeight uint64
		shouldErr           bool
	}{
		{
			name: "valid case",
			preset: []types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "contract-1",
					Weight:              5,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "contract-2",
					Weight:              3,
				},
				{
					PoolId:              "pool-3",
					PoolContractAddress: "contract-3",
					Weight:              2,
				},
			},
			updatePool: []types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "contract-1",
					Weight:              10,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "contract-2",
					Weight:              11,
				},
				{
					PoolId:              "pool-6",
					PoolContractAddress: "contract-6",
					Weight:              12,
				},
			},
			expected: []types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "contract-1",
					Weight:              10,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "contract-2",
					Weight:              11,
				},
				{
					PoolId:              "pool-3",
					PoolContractAddress: "contract-3",
					Weight:              2,
				},
				{
					PoolId:              "pool-6",
					PoolContractAddress: "contract-6",
					Weight:              12,
				},
			},
			expectedTotalWeight: 35,
			shouldErr:           false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// setup
			for i := range tc.preset {
				err := suite.App.PoolKeeper.CreateIncentivePool(suite.Ctx, &tc.preset[i])
				suite.NoError(err)
			}

			msg := types.UpdatePoolIncentivesProposal{
				Title:             "title",
				Description:       "description",
				UpdatedIncentives: tc.updatePool,
			}

			err := suite.App.PoolKeeper.HandleUpdatePoolIncentivesProposal(suite.Ctx, &msg)
			suite.NoError(err)

			if !tc.shouldErr {
				for _, item := range tc.expected {
					res, err := suite.App.PoolKeeper.FindIncentivePoolById(suite.Ctx, item.PoolId)
					suite.NoError(err)
					suite.Equal(item, *res)
				}
				tw := suite.App.PoolKeeper.GetTotalWeight(suite.Ctx)
				suite.Require().Equal(tc.expectedTotalWeight, tw)
			}
		})
	}
}
