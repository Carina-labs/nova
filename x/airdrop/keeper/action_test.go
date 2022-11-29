package keeper_test

import (
	"time"

	"github.com/Carina-labs/nova/v2/x/airdrop/keeper"
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestPostClaimedSnAsset() {
	validAirdrop := validAirdropInfo(suite.Ctx)
	defaultState := validUserState(suite.Ctx, validUser.String(), validAirdrop.MaximumTokenAllocPerUser)
	defaultQuest := defaultState.QuestStates[keeper.QuestTypeSnAssetClaim]

	tests := map[string]struct {
		airdropEndDate   time.Time
		userAddr         sdk.AccAddress
		targetState      types.QuestStateType
		targetArchivedAt time.Time
		shouldChangeDate bool
		isEligible       bool
	}{
		"invalid quest date, nothing happened": {
			airdropEndDate:   suite.Ctx.BlockTime().Add(-1 * time.Hour),
			userAddr:         validUser,
			targetState:      defaultQuest.State,
			shouldChangeDate: false,
			isEligible:       true,
		},
		"is not eligible user, nothing happened": {
			airdropEndDate:   validAirdrop.AirdropEndTimestamp,
			userAddr:         invalidUser,
			targetState:      defaultQuest.State,
			shouldChangeDate: false,
			isEligible:       false,
		},
		"eligible user quest should be marked as claimable": {
			airdropEndDate:   validAirdrop.AirdropEndTimestamp,
			userAddr:         validUser,
			targetState:      types.QuestStateType(keeper.QuestStateClaimable),
			shouldChangeDate: true,
			isEligible:       true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx

			airdrop := suite.App.AirdropKeeper
			airdropInfo := validAirdropInfo(ctx)
			airdropInfo.AirdropEndTimestamp = test.airdropEndDate
			airdrop.SetAirdropInfo(ctx, airdropInfo)
			err := airdrop.SetUserState(ctx, validUser, defaultState)
			suite.Require().NoError(err)

			// perform post claimed sn asset
			airdrop.PostClaimedSnAsset(ctx, test.userAddr)

			// check if user state is updated
			userState, err := airdrop.GetUserState(ctx, test.userAddr)

			if test.isEligible {
				suite.Require().NoError(err)

				quest := userState.QuestStates[keeper.QuestTypeSnAssetClaim]
				suite.Require().Equal(test.targetState, quest.State)
				if test.shouldChangeDate {
					suite.Require().Equal(ctx.BlockTime(), quest.AchievedAt)
				} else {
					suite.Require().Equal(defaultQuest.AchievedAt, quest.AchievedAt)
				}
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestPostProposalVote() {
	validAirdrop := validAirdropInfo(suite.Ctx)
	defaultState := validUserState(suite.Ctx, validUser.String(), validAirdrop.MaximumTokenAllocPerUser)
	defaultQuest := defaultState.QuestStates[keeper.QuestTypeVoteOnProposals]

	tests := map[string]struct {
		airdropEndDate   time.Time
		userAddr         sdk.AccAddress
		targetState      types.QuestStateType
		targetArchivedAt time.Time
		shouldChangeDate bool
		isEligible       bool
	}{
		"invalid quest date, nothing happened": {
			airdropEndDate:   suite.Ctx.BlockTime().Add(-1 * time.Hour),
			userAddr:         validUser,
			targetState:      defaultQuest.State,
			shouldChangeDate: false,
			isEligible:       true,
		},
		"is not eligible user, nothing happened": {
			airdropEndDate:   validAirdrop.AirdropEndTimestamp,
			userAddr:         invalidUser,
			targetState:      defaultQuest.State,
			shouldChangeDate: false,
			isEligible:       false,
		},
		"eligible user quest should be marked as claimable": {
			airdropEndDate:   validAirdrop.AirdropEndTimestamp,
			userAddr:         validUser,
			targetState:      types.QuestStateType(keeper.QuestStateClaimable),
			shouldChangeDate: true,
			isEligible:       true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx

			airdrop := suite.App.AirdropKeeper
			airdropInfo := validAirdropInfo(ctx)
			airdropInfo.AirdropEndTimestamp = test.airdropEndDate
			airdrop.SetAirdropInfo(ctx, airdropInfo)
			err := airdrop.SetUserState(ctx, validUser, defaultState)
			suite.Require().NoError(err)

			// perform post claimed sn asset
			airdrop.PostProposalVote(ctx, 0, test.userAddr)

			// check if user state is updated
			userState, err := airdrop.GetUserState(ctx, test.userAddr)

			if test.isEligible {
				suite.Require().NoError(err)

				quest := userState.QuestStates[keeper.QuestTypeVoteOnProposals]
				suite.Require().Equal(test.targetState, quest.State)
				if test.shouldChangeDate {
					suite.Require().Equal(ctx.BlockTime(), quest.AchievedAt)
				} else {
					suite.Require().Equal(defaultQuest.AchievedAt, quest.AchievedAt)
				}
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
