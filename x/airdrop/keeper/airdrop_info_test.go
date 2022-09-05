package keeper_test

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	"reflect"
	"time"
)

func (suite *KeeperTestSuite) TestSetAirdropInfo() {
	airdrop := suite.App.AirdropKeeper
	expected := types.DefaultGenesis()
	airdrop.SetAirdropInfo(suite.Ctx, expected.AirdropInfo)
	got := airdrop.GetAirdropInfo(suite.Ctx)

	if !reflect.DeepEqual(expected.AirdropInfo, got) {
		suite.T().Error("airdrop infos are not equal")
	}
}

func (suite *KeeperTestSuite) TestGetAirdropInfo() {
	airdrop := suite.App.AirdropKeeper
	expected := types.DefaultGenesis()

	// delete default airdrop info
	airdrop.DeleteAirdropInfo(suite.Ctx)

	defer func() {
		if r := recover(); r != nil {
			suite.Require().Equal("airdrop info is missing", r)
		}
	}()

	// check if panic is thrown when airdrop info is not set
	_ = airdrop.GetAirdropInfo(suite.Ctx)

	// set airdrop info and check if it is returned
	airdrop.SetAirdropInfo(suite.Ctx, expected.AirdropInfo)
	got := airdrop.GetAirdropInfo(suite.Ctx)
	if !reflect.DeepEqual(expected.AirdropInfo, got) {
		suite.T().Error("airdrop infos are not equal")
	}
}

func (suite *KeeperTestSuite) TestValidQuestDate() {
	airdrop := suite.App.AirdropKeeper

	// set valid quest date
	info := airdrop.GetAirdropInfo(suite.Ctx)
	info.AirdropEndTimestamp = suite.Ctx.BlockTime().Add(1 * time.Hour)
	airdrop.SetAirdropInfo(suite.Ctx, info)
	suite.Require().True(airdrop.ValidQuestDate(suite.Ctx))

	// set invalid quest date
	info.AirdropEndTimestamp = suite.Ctx.BlockTime().Add(-1 * time.Hour)
	airdrop.SetAirdropInfo(suite.Ctx, info)
	suite.Require().False(airdrop.ValidQuestDate(suite.Ctx))
}

func (suite *KeeperTestSuite) TestValidClaimableDate() {
	airdrop := suite.App.AirdropKeeper

	// set valid claimable date
	info := airdrop.GetAirdropInfo(suite.Ctx)
	info.AirdropStartTimestamp = suite.Ctx.BlockTime().Add(-1 * time.Hour)
	info.AirdropEndTimestamp = suite.Ctx.BlockTime().Add(1 * time.Hour)
	airdrop.SetAirdropInfo(suite.Ctx, info)
	suite.Require().True(airdrop.ValidClaimableDate(suite.Ctx))

	// set invalid claimable date
	info.AirdropStartTimestamp = suite.Ctx.BlockTime().Add(1 * time.Second)
	airdrop.SetAirdropInfo(suite.Ctx, info)
	suite.Require().False(airdrop.ValidClaimableDate(suite.Ctx))
}
