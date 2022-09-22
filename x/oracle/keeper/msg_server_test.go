package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestServerUpdateChainState() {
	// store chain params
	ctx := suite.Ctx
	keeper := suite.App.OracleKeeper
	msgServer := suite.msgServer

	// set chain state
	chainInfo := types.ChainInfo{
		Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
		OperatorAddress: fooOperator.String(),
		LastBlockHeight: 1000,
		AppHash:         []byte(fooAppHash),
		ChainId:         fooChainId,
	}

	msg := types.MsgUpdateChainState{
		Coin:        chainInfo.Coin,
		Operator:    chainInfo.OperatorAddress,
		BlockHeight: chainInfo.LastBlockHeight,
		AppHash:     chainInfo.AppHash,
		ChainId:     chainInfo.ChainId,
	}

	// set operator
	keeper.SetParams(ctx, types.Params{
		OracleOperators: []string{fooOperator.String()},
	})

	// update chain state
	_, err := msgServer.UpdateChainState(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().NoError(err)

	// check updated chain state
	got, err := keeper.GetChainState(ctx, fooDenom)
	suite.Require().NoError(err)
	suite.Require().Equal(chainInfo.Coin, got.Coin)
	suite.Require().Equal(chainInfo.OperatorAddress, got.OperatorAddress)
	suite.Require().Equal(chainInfo.LastBlockHeight, got.LastBlockHeight)
	suite.Require().Equal(chainInfo.AppHash, got.AppHash)
	suite.Require().Equal(chainInfo.ChainId, got.ChainId)

	// oracle version should increase
	version, _ := keeper.GetOracleVersion(ctx, chainInfo.ChainId)
	suite.Require().Equal(uint64(1), version)

	// after updating chain state, version should increase
	_, err = msgServer.UpdateChainState(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().NoError(err)

	version, _ = keeper.GetOracleVersion(ctx, chainInfo.ChainId)
	suite.Require().Equal(uint64(2), version)

	// updating with invalid operator should fail
	msg.Operator = "invalid"
	_, err = msgServer.UpdateChainState(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().Error(err)
	suite.Require().Equal(err, types.ErrInvalidOperator)
}
