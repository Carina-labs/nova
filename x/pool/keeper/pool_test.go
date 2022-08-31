package keeper_test

import (
	"github.com/Carina-labs/nova/x/pool/types"
)

func (suite *KeeperTestSuite) TestCreateCandidatePool() {
	tcs := []struct {
		name string
		pool []types.CandidatePool
	}{
		{
			name: "valid case 1",
			pool: []types.CandidatePool{
				{
					PoolId:              "1",
					PoolContractAddress: "12345",
				},
			},
		},
		{
			name: "valid case 1",
			pool: []types.CandidatePool{
				{
					PoolId:              "1",
					PoolContractAddress: "12345",
				},
				{
					PoolId:              "2",
					PoolContractAddress: "abcde",
				},
				{
					PoolId:              "3",
					PoolContractAddress: "zxcv12",
				},
			},
		},
	}

	keeper := suite.App.PoolKeeper

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, p := range tc.pool {
				err := keeper.CreateCandidatePool(suite.Ctx, &p)
				suite.NoError(err)
			}

			keeper.IterateCandidatePools(suite.Ctx, func(i int64, pool *types.CandidatePool) bool {
				suite.Equal(tc.pool[i].PoolId, pool.PoolId)
				suite.Equal(tc.pool[i].PoolContractAddress, pool.PoolContractAddress)
				return false
			})

			keeper.ClearCandidatePools(suite.Ctx)
		})
	}
}

// TODO : implements tc for msg_server, testing SetPoolWeight is called with valid operator.
func (suite *KeeperTestSuite) TestSetPoolWeight() {
	keeper := suite.App.PoolKeeper

	tcs := []struct {
		name      string
		preset    []*types.IncentivePool
		targetId  string
		newWeight uint64
	}{
		{
			name: "valid case 1",
			preset: []*types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "12345",
					Weight:              0,
				},
			},
			targetId:  "pool-1",
			newWeight: 10,
		},
		{
			name: "valid case 2",
			preset: []*types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "12345",
					Weight:              5,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "abcde",
					Weight:              3,
				},
				{
					PoolId:              "pool-3",
					PoolContractAddress: "abcde12345",
					Weight:              2,
				},
			},
			targetId:  "pool-3",
			newWeight: 1,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// create pool with presets
			for _, p := range tc.preset {
				err := keeper.CreateIncentivePool(suite.Ctx, p)
				suite.NoError(err)
			}

			err := keeper.SetPoolWeight(suite.Ctx, tc.targetId, tc.newWeight)
			suite.NoError(err)

			pool, err := keeper.FindIncentivePoolById(suite.Ctx, tc.targetId)
			suite.NoError(err)
			suite.Equal(tc.newWeight, pool.Weight)

			keeper.ClearIncentivePools(suite.Ctx)
		})
	}
}

func (suite *KeeperTestSuite) TestGetTotalWeight() {
	keeper := suite.App.PoolKeeper
	tcs := []struct {
		name                string
		preset              []*types.IncentivePool
		expectedTotalWeight uint64
	}{
		{
			name: "valid case 1",
			preset: []*types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "12345",
					Weight:              5,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "12345abcde",
					Weight:              5,
				},
			},
			expectedTotalWeight: 10,
		},
		{
			name: "valid case 2",
			preset: []*types.IncentivePool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "12345",
					Weight:              5,
				},
				{
					PoolId:              "pool-2",
					PoolContractAddress: "abcde",
					Weight:              3,
				},
				{
					PoolId:              "pool-3",
					PoolContractAddress: "abcde12345",
					Weight:              2,
				},
			},
			expectedTotalWeight: 10,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// create pool with presets
			for _, p := range tc.preset {
				err := keeper.CreateIncentivePool(suite.Ctx, p)
				suite.NoError(err)
			}

			totalWeight := keeper.GetTotalWeight(suite.Ctx)
			suite.Equal(tc.expectedTotalWeight, totalWeight)

			keeper.ClearIncentivePools(suite.Ctx)
		})
	}
}
