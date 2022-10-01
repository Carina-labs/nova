package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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
		ZoneId:          fooChainId,
	}

	msg := types.MsgUpdateChainState{
		Coin:        chainInfo.Coin,
		Operator:    chainInfo.OperatorAddress,
		BlockHeight: chainInfo.LastBlockHeight,
		AppHash:     chainInfo.AppHash,
		ZoneId:      chainInfo.ZoneId,
	}

	// set oracle address
	keeper.SetOracleAddress(ctx, fooChainId, []string{fooOperator.String()})

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
	suite.Require().Equal(chainInfo.ZoneId, got.ZoneId)

	// oracle version should increase
	version, _ := keeper.GetOracleVersion(ctx, chainInfo.ZoneId)
	suite.Require().Equal(uint64(1), version)

	// after updating chain state, version should increase
	_, err = msgServer.UpdateChainState(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().NoError(err)

	version, _ = keeper.GetOracleVersion(ctx, chainInfo.ZoneId)
	suite.Require().Equal(uint64(2), version)

	// updating with invalid operator should fail
	msg.Operator = "invalid"
	_, err = msgServer.UpdateChainState(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().Error(err)
	suite.Require().Equal(err, types.ErrInvalidOperator)
}

func (suite *KeeperTestSuite) TestRegisterOracleAddress() {
	ctx := suite.Ctx
	keeper := suite.App.OracleKeeper
	msgServer := suite.msgServer
	oracleAddr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())

	//set param
	managerAddr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	keeper.SetParams(ctx, types.Params{OracleKeyManager: []string{managerAddr.String()}})

	msg := types.MsgRegisterOracleAddr{
		ZoneId:        fooChainId,
		OracleAddress: oracleAddr.String(),
		FromAddress:   managerAddr.String(),
	}

	_, err := msgServer.RegisterOracleAddress(sdk.WrapSDKContext(ctx), &msg)
	suite.Require().NoError(err)

	oracleInfo := keeper.GetOracleAddress(ctx, fooChainId)

	suite.Require().Equal(oracleInfo.ZoneId, fooChainId)
	suite.Require().Equal(oracleInfo.OracleAddress[0], oracleAddr.String())

	msg = types.MsgRegisterOracleAddr{
		ZoneId:        fooChainId,
		OracleAddress: oracleAddr.String(),
		FromAddress:   oracleAddr.String(),
	}

	_, err = msgServer.RegisterOracleAddress(sdk.WrapSDKContext(ctx), &msg)
	suite.Require().Error(err)
}
