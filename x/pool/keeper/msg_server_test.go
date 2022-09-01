package keeper_test

import (
	"github.com/Carina-labs/nova/x/pool/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	TestOperator = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
)

func (suite *KeeperTestSuite) TestCreateIncentivePool_withOperator() {
	poolKeeper := suite.App.PoolKeeper
	// set operator
	poolKeeper.SetParams(suite.Ctx, types.NewParams([]string{TestOperator.String()}))
	msgServer := suite.msgServer

	msg := &types.MsgCreateIncentivePool{
		PoolId:              "test-pool",
		PoolContractAddress: "contract",
		Operator:            TestOperator.String(),
	}
	res, err := msgServer.CreateIncentivePool(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.NoError(err)
	suite.NotNil(res)
}

func (suite *KeeperTestSuite) TestCreateIncentivePool_withInvalidOperator() {
	poolKeeper := suite.App.PoolKeeper
	// set operator
	poolKeeper.SetParams(suite.Ctx, types.NewParams([]string{TestOperator.String()}))
	msgServer := suite.msgServer

	fakeOperator := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	msg := &types.MsgCreateIncentivePool{
		PoolId:              "test-pool",
		PoolContractAddress: "contract",
		Operator:            fakeOperator.String(),
	}

	res, err := msgServer.CreateIncentivePool(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.Error(err)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestSetMultiplePoolWeight() {
	poolKeeper := suite.App.PoolKeeper
	// set operator
	poolKeeper.SetParams(suite.Ctx, types.NewParams([]string{TestOperator.String()}))
	msgServer := suite.msgServer

	// preset
	err := poolKeeper.CreateIncentivePool(suite.Ctx, &types.IncentivePool{
		PoolId:              "pool-1",
		PoolContractAddress: "pool-contract-1",
		Weight:              uint64(1),
	})
	suite.NoError(err)
	err = poolKeeper.CreateIncentivePool(suite.Ctx, &types.IncentivePool{
		PoolId:              "pool-2",
		PoolContractAddress: "pool-contract-2",
		Weight:              uint64(2),
	})
	suite.NoError(err)
	err = poolKeeper.CreateIncentivePool(suite.Ctx, &types.IncentivePool{
		PoolId:              "pool-3",
		PoolContractAddress: "pool-contract-3",
		Weight:              uint64(3),
	})
	suite.NoError(err)

	msg := &types.MsgSetMultiplePoolWeight{
		NewPoolData: []types.NewPoolWeight{
			{
				PoolId:    "pool-1",
				NewWeight: uint64(3),
			},
			{
				PoolId:    "pool-2",
				NewWeight: uint64(2),
			},
			{
				PoolId:    "pool-3",
				NewWeight: uint64(5),
			},
		},
		Operator: TestOperator.String(),
	}
	res, err := msgServer.SetMultiplePoolWeight(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.NoError(err)
	suite.NotNil(res)

	pool, err := poolKeeper.FindIncentivePoolById(suite.Ctx, "pool-1")
	suite.Equal(uint64(3), pool.Weight)
	pool, err = poolKeeper.FindIncentivePoolById(suite.Ctx, "pool-2")
	suite.Equal(uint64(2), pool.Weight)
	pool, err = poolKeeper.FindIncentivePoolById(suite.Ctx, "pool-3")
	suite.Equal(uint64(5), pool.Weight)
}
