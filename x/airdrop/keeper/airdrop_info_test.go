package keeper_test

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	"reflect"
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
