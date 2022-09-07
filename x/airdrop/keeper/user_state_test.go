package keeper_test

import (
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetUserState() {
	state := validUserState(suite.Ctx, sdk.AccAddress{0x1}.String(), "100000")
	airdrop := suite.App.AirdropKeeper
	err := airdrop.SetUserState(suite.Ctx, sdk.AccAddress{0x1}, state)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestGetUserState() {
	state := validUserState(suite.Ctx, sdk.AccAddress{0x1}.String(), "100000")
	airdrop := suite.App.AirdropKeeper
	err := airdrop.SetUserState(suite.Ctx, sdk.AccAddress{0x1}, state)
	suite.Require().NoError(err)

	resp, err := airdrop.GetUserState(suite.Ctx, sdk.AccAddress{0x1})
	suite.Require().NoError(err)

	if !reflect.DeepEqual(state, resp) {
		suite.T().Errorf("we expected both states to be equal, but they are not, state: %v, resp: %v", state, resp)
	}

	// check non-exists user
	_, err = airdrop.GetUserState(suite.Ctx, sdk.AccAddress{0x2})
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestIsEligible() {
	state := validUserState(suite.Ctx, sdk.AccAddress{0x1}.String(), "100000")
	airdrop := suite.App.AirdropKeeper
	err := airdrop.SetUserState(suite.Ctx, sdk.AccAddress{0x1}, state)
	suite.Require().NoError(err)

	ok := airdrop.IsEligible(suite.Ctx, sdk.AccAddress{0x1})
	suite.Require().True(ok)

	// check non-exists user
	ok = airdrop.IsEligible(suite.Ctx, sdk.AccAddress{0x2})
	suite.Require().False(ok)
}
