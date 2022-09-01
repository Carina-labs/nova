package keeper_test

import (
	"github.com/Carina-labs/nova/x/pool/types"
)

func (suite *KeeperTestSuite) TestCreateCandidatePool() {
	tcs := []struct {
		name      string
		preset    []types.CandidatePool
		pool      types.CandidatePool
		shouldErr bool
	}{
		{
			name:   "valid case 1",
			preset: []types.CandidatePool{},
			pool: types.CandidatePool{
				PoolId:              "1",
				PoolContractAddress: "12345",
			},
			shouldErr: false,
		},
		{
			name: "valid case 1",
			preset: []types.CandidatePool{
				{
					PoolId:              "1",
					PoolContractAddress: "12345",
				},
				{
					PoolId:              "2",
					PoolContractAddress: "abcde",
				},
			},
			pool: types.CandidatePool{
				PoolId:              "3",
				PoolContractAddress: "zxcv12",
			},
			shouldErr: false,
		},
		{
			name: "error case : duplicated pool id",
			preset: []types.CandidatePool{
				{
					PoolId:              "1",
					PoolContractAddress: "12345",
				},
				{
					PoolId:              "2",
					PoolContractAddress: "abcde",
				},
			},
			pool: types.CandidatePool{
				PoolId:              "1",
				PoolContractAddress: "abcde",
			},
			shouldErr: true,
		},
	}

	keeper := suite.App.PoolKeeper

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// setup
			for _, p := range tc.preset {
				err := keeper.CreateCandidatePool(suite.Ctx, &p)
				suite.NoError(err)
			}

			err := keeper.CreateCandidatePool(suite.Ctx, &tc.pool)
			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				candidatePool, err := keeper.FindCandidatePoolById(suite.Ctx, tc.pool.PoolId)
				suite.NoError(err)
				suite.Equal(tc.pool.PoolId, candidatePool.PoolId)
				suite.Equal(tc.pool.PoolContractAddress, candidatePool.PoolContractAddress)
			}

			keeper.ClearCandidatePools(suite.Ctx)
		})
	}
}

func (suite *KeeperTestSuite) TestCreateIncentivePool() {
	tcs := []struct {
		name      string
		preset    []types.IncentivePool
		pool      types.IncentivePool
		shouldErr bool
	}{
		{
			name:   "valid case",
			preset: []types.IncentivePool{},
			pool: types.IncentivePool{
				PoolId:              "pool-1",
				PoolContractAddress: "12345",
				Weight:              0,
			},
			shouldErr: false,
		},
	}
	keeper := suite.App.PoolKeeper

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// setup
			for i := range tc.preset {
				err := keeper.CreateIncentivePool(suite.Ctx, &tc.preset[i])
				suite.NoError(err)
			}

			err := keeper.CreateIncentivePool(suite.Ctx, &tc.pool)
			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.NoError(err)

				incentivePool, err := keeper.FindIncentivePoolById(suite.Ctx, tc.pool.PoolId)
				suite.NoError(err)
				suite.Equal(tc.pool.PoolId, incentivePool.PoolId)
				suite.Equal(tc.pool.PoolContractAddress, incentivePool.PoolContractAddress)
				suite.Equal(tc.pool.Weight, incentivePool.Weight)
			}
		})
	}
}

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
