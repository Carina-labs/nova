package keeper_test

import (
	"github.com/Carina-labs/nova/x/poolincentive/types"
	"strconv"
)

func (suite *KeeperTestSuite) setCandidatePools() types.CandidatePoolInfo {
	pools := types.CandidatePoolInfo{
		CandidatePools: []*types.CandidatePool{
			{
				PoolId:              "1",
				PoolContractAddress: "testAddr1",
			},
			{
				PoolId:              "2",
				PoolContractAddress: "testAddr2",
			},
			{
				PoolId:              "3",
				PoolContractAddress: "testAddr3",
			},
		},
	}
	suite.App.PoolKeeper.SetCandidatePoolInfo(suite.Ctx, pools)
	return pools
}

func (suite *KeeperTestSuite) setIncentivePools() types.IncentivePoolInfo {
	pools := types.IncentivePoolInfo{
		IncentivePools: []*types.IncentivePool{
			{
				PoolId:              "1",
				PoolContractAddress: "",
				Weight:              1,
			},
			{
				PoolId:              "2",
				PoolContractAddress: "",
				Weight:              1,
			},
			{
				PoolId:              "3",
				PoolContractAddress: "",
				Weight:              1,
			},
		},
	}
	suite.App.PoolKeeper.SetIncentivePoolInfo(suite.Ctx, pools)
	return pools
}
func (suite *KeeperTestSuite) TestSingleCandidatePool() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	poolId := "1"
	_, err := queryClient.SingleCandidatePool(ctx.Context(), &types.QuerySingleCandidatePoolRequest{
		PoolId: poolId,
	})
	suite.Require().Error(err)

	pools := suite.setCandidatePools()

	res, err := queryClient.SingleCandidatePool(ctx.Context(), &types.QuerySingleCandidatePoolRequest{
		PoolId: poolId,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(res.PoolId, pools.CandidatePools[0].PoolId)
	suite.Require().Equal(res.PoolAddress, pools.CandidatePools[0].PoolContractAddress)
}

func (suite *KeeperTestSuite) TestAllCandidatePool() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	res, err := queryClient.AllCandidatePool(ctx.Context(), &types.QueryAllCandidatePoolRequest{})
	suite.Require().Nil(res.CandidatePools)

	pools := suite.setCandidatePools()

	res, err = queryClient.AllCandidatePool(ctx.Context(), &types.QueryAllCandidatePoolRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(res.CandidatePools, pools.CandidatePools)
}

func (suite *KeeperTestSuite) TestSingleIncentivePool() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	poolId := "1"
	_, err := queryClient.SingleIncentivePool(ctx.Context(), &types.QuerySingleIncentivePoolRequest{
		PoolId: poolId,
	})
	suite.Require().Error(err)

	pools := suite.setIncentivePools()
	res, err := queryClient.SingleIncentivePool(ctx.Context(), &types.QuerySingleIncentivePoolRequest{
		PoolId: poolId,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(res.PoolId, pools.IncentivePools[0].PoolId)
	suite.Require().Equal(res.PoolAddress, pools.IncentivePools[0].PoolContractAddress)
	suite.Require().Equal(res.Weight, strconv.FormatUint(pools.IncentivePools[0].Weight, 10))
}

func (suite *KeeperTestSuite) TestAllIncentivePool() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	res, err := queryClient.AllIncentivePool(ctx.Context(), &types.QueryAllIncentivePoolRequest{})
	suite.Require().NoError(err)
	suite.Require().Nil(res.IncentivePools)

	pools := suite.setIncentivePools()

	res, err = queryClient.AllIncentivePool(ctx.Context(), &types.QueryAllIncentivePoolRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(res.IncentivePools, pools.IncentivePools)
}

func (suite *KeeperTestSuite) TestTotalWeight() {
	queryClient := suite.queryClient
	ctx := suite.Ctx

	res, err := queryClient.TotalWeight(ctx.Context(), &types.QueryTotalWeightRequest{})
	suite.Require().NoError(err)
	suite.Require().Zero(res.TotalWeight)

	suite.App.PoolKeeper.SetIncentivePoolInfo(ctx, types.IncentivePoolInfo{
		TotalWeight: 10,
		IncentivePools: []*types.IncentivePool{
			{
				PoolId:              "1",
				PoolContractAddress: "test1",
				Weight:              6,
			},
			{
				PoolId:              "1",
				PoolContractAddress: "test1",
				Weight:              4,
			},
			{
				PoolId:              "1",
				PoolContractAddress: "test1",
				Weight:              1,
			},
		},
	})

	res, err = queryClient.TotalWeight(ctx.Context(), &types.QueryTotalWeightRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(strconv.FormatUint(res.TotalWeight, 10), "10")
}
