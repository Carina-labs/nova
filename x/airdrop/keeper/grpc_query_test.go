package keeper_test

import (
	"reflect"
	"time"

	"github.com/Carina-labs/nova/x/airdrop/types"
)

func (suite *KeeperTestSuite) TestQueryAirdropInfo() {
	client := suite.queryClient
	ctx := suite.Ctx
	expected := validAirdropInfo(ctx)
	suite.App.AirdropKeeper.SetAirdropInfo(ctx, expected)

	got, err := client.AirdropInfo(ctx.Context(), &types.QueryAirdropInfoRequest{})
	suite.Require().NoError(err)

	if !reflect.DeepEqual(expected, got.AirdropInfo) {
		suite.T().Errorf("airdrop infos are not equal: got %v, expected %v", got, expected)
	}
}

func (suite *KeeperTestSuite) TestQueryTotalAllocatedAirdropToken() {
	airdrop := validAirdropInfo(suite.Ctx)
	totalAsset := "10000"

	tests := map[string]struct {
		user         string
		expectAmount string
		expectDenom  string
		shouldErr    bool
	}{
		"invalid user address format": {
			user:      "invalid address format",
			shouldErr: true,
		},
		"user is not eligible for airdrop": {
			user:         invalidUser.String(),
			expectAmount: "0",
			expectDenom:  airdrop.AirdropDenom,
			shouldErr:    false,
		},
		"valid user for airdrop": {
			user:         validUser.String(),
			expectAmount: totalAsset,
			expectDenom:  airdrop.AirdropDenom,
			shouldErr:    false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx

			// set initial airdrop info
			airdropKeeper := suite.App.AirdropKeeper
			airdropKeeper.SetAirdropInfo(ctx, airdrop)

			// set valid user state
			userState := validUserState(ctx, validUser.String(), totalAsset)
			err := airdropKeeper.SetUserState(ctx, validUser, userState)
			suite.Require().NoError(err)

			// check test cases
			client := suite.queryClient
			resp, err := client.TotalAllocatedAirdropToken(ctx.Context(), &types.QueryTotalAllocatedAirdropTokenRequest{Address: test.user})
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expectAmount, resp.TotalAssets.Amount.String())
				suite.Require().Equal(test.expectDenom, resp.TotalAssets.Denom)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryQuestState() {
	client := suite.queryClient
	ctx := suite.Ctx

	airdropKeeper := suite.App.AirdropKeeper
	airdropKeeper.SetAirdropInfo(suite.Ctx, validAirdropInfo(suite.Ctx))

	userState := validUserState(ctx, validUser.String(), "10000")
	err := airdropKeeper.SetUserState(ctx, validUser, userState)
	suite.Require().NoError(err)

	// should return err if user address is invalid format
	_, err = client.QuestState(ctx.Context(), &types.QueryQuestStateRequest{Address: "invalid address format"})
	suite.Require().Error(err)

	// should return unauthorized error if user is not eligible for airdrop
	_, err = client.QuestState(ctx.Context(), &types.QueryQuestStateRequest{Address: invalidUser.String()})
	suite.Require().Error(err)

	// should return user state if user is eligible for airdrop
	resp, err := client.QuestState(ctx.Context(), &types.QueryQuestStateRequest{Address: validUser.String()})
	suite.Require().NoError(err)
	suite.Require().Equal(userState.QuestStates, resp.QuestStates)

	// quest state 0 starts as completed state from the beginning
	// so the `achieved at` field must be filled with time of genesis block.
	// therefore, at this time, we use quest state 1 to write a unit test
	suite.Require().True(userState.QuestStates[1].AchievedAt.IsZero())
	suite.Require().True(userState.QuestStates[1].ClaimedAt.IsZero())

	// double check user state after changing it
	targetDate := ctx.BlockTime()
	userState.QuestStates[1].AchievedAt = targetDate
	userState.QuestStates[1].ClaimedAt = targetDate.Add(1 * time.Hour)
	err = airdropKeeper.SetUserState(ctx, validUser, userState)
	suite.Require().NoError(err)

	resp, err = client.QuestState(ctx.Context(), &types.QueryQuestStateRequest{Address: validUser.String()})
	suite.Require().NoError(err)
	suite.Require().Equal(targetDate, resp.QuestStates[1].AchievedAt)
	suite.Require().Equal(targetDate.Add(1*time.Hour), resp.QuestStates[1].ClaimedAt)
}
