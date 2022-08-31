package keeper_test

import (
	"github.com/Carina-labs/nova/x/pool/types"
)

func (suite *KeeperTestSuite) TestCreatePool() {
	tcs := []struct {
		name string
		pool []types.Pool
	}{
		{
			name: "valid case 1",
			pool: []types.Pool{
				{
					PoolId:              "1",
					PoolContractAddress: "12345",
				},
			},
		},
		{
			name: "valid case 1",
			pool: []types.Pool{
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
				err := keeper.CreatePool(suite.Ctx, &p)
				suite.NoError(err)
			}

			keeper.IteratePools(suite.Ctx, func(i int64, pool *types.Pool) bool {
				suite.Equal(tc.pool[i].PoolId, pool.PoolId)
				suite.Equal(tc.pool[i].PoolContractAddress, pool.PoolContractAddress)
				suite.Equal(tc.pool[i].Weight, pool.Weight)
				return false
			})

			keeper.ClearPools(suite.Ctx)
		})
	}
}

// TODO : implements tc for msg_server, testing SetPoolWeight is called with valid operator.
func (suite *KeeperTestSuite) TestSetPoolWeight() {
	keeper := suite.App.PoolKeeper

	tcs := []struct {
		name      string
		preset    []*types.Pool
		targetId  string
		newWeight uint64
	}{
		{
			name: "valid case 1",
			preset: []*types.Pool{
				{
					PoolId:              "pool-1",
					PoolContractAddress: "12345",
					Weight:              0,
				},
			},
			targetId:  "pool-1",
			newWeight: 10,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// create pool with presets
			for _, p := range tc.preset {
				err := keeper.CreatePool(suite.Ctx, p)
				suite.NoError(err)
			}

			err := keeper.SetPoolWeight(suite.Ctx, tc.targetId, tc.newWeight)
			suite.NoError(err)

			pool, ok := keeper.FindPoolById(suite.Ctx, tc.targetId)
			suite.True(ok)
			suite.Equal(tc.newWeight, pool.Weight)
		})
	}
}