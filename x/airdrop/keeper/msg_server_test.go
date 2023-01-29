package keeper_test

import (
	"time"

	"github.com/Carina-labs/nova/x/airdrop/keeper"
	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgClaimAirdrop() {
	tests := map[string]struct {
		user               sdk.AccAddress
		questType          types.QuestType
		before             func(ctx sdk.Context, k *keeper.Keeper)
		shouldErr          bool
		performDoubleClaim bool
	}{
		"invalid claimable date": {
			user:      validUser,
			questType: types.QuestType(keeper.QuestTypeNothingToDo),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				// set airdrop date before current time
				airdrop := validAirdropInfo(ctx)
				airdrop.AirdropStartTimestamp = ctx.BlockTime().Add(1 * time.Hour)
				k.SetAirdropInfo(ctx, airdrop)
			},
			shouldErr: true,
		},
		"not eligible user": {
			user:      invalidUser,
			questType: types.QuestType(keeper.QuestTypeNothingToDo),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				err := k.SetUserState(ctx, validUser, validUserState(ctx, validUser.String(), maxTotalAllocPerUser))
				suite.Require().NoError(err)
			},
			shouldErr: true,
		},
		"not claimable state": {
			user:      validUser,
			questType: types.QuestType(keeper.QuestTypeSnAssetClaim),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				err := k.SetUserState(ctx, validUser, validUserState(ctx, validUser.String(), maxTotalAllocPerUser))
				suite.Require().NoError(err)
			},
			shouldErr: true,
		},
		"already claimed": {
			user:      validUser,
			questType: types.QuestType(keeper.QuestTypeSnAssetClaim),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				userState := validUserState(ctx, validUser.String(), maxTotalAllocPerUser)
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].AchievedAt = ctx.BlockTime()
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].State = types.QuestStateType(keeper.QuestStateClaimed)
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].ClaimedAt = ctx.BlockTime()
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].ClaimedAmount = "2000"
				err := k.SetUserState(ctx, validUser, userState)
				suite.Require().NoError(err)
			},
			shouldErr: true,
		},
		"user can claim asset": {
			user:      validUser,
			questType: types.QuestType(keeper.QuestTypeSnAssetClaim),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				userState := validUserState(ctx, validUser.String(), maxTotalAllocPerUser)
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].AchievedAt = ctx.BlockTime()
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].State = types.QuestStateType(keeper.QuestStateClaimable)
				err := k.SetUserState(ctx, validUser, userState)
				suite.Require().NoError(err)
			},
			shouldErr: false,
		},
		"user cannot claim more than once": {
			user:      validUser,
			questType: types.QuestType(keeper.QuestTypeSnAssetClaim),
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				userState := validUserState(ctx, validUser.String(), maxTotalAllocPerUser)
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].AchievedAt = ctx.BlockTime()
				userState.QuestStates[keeper.QuestTypeSnAssetClaim].State = types.QuestStateType(keeper.QuestStateClaimable)
				err := k.SetUserState(ctx, validUser, userState)
				suite.Require().NoError(err)
			},
			shouldErr:          false,
			performDoubleClaim: true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			airdropInfo := validAirdropInfo(suite.Ctx)
			bankKeeper := suite.App.BankKeeper
			airdropKeeper := suite.App.AirdropKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx

			test.before(ctx, airdropKeeper)

			// check user balance is zero before sending claim transaction
			balance := bankKeeper.GetBalance(ctx, test.user, airdropInfo.AirdropDenom)
			expectClaimed, _, _ := airdropKeeper.CalcClaimableAmount(ctx, test.user)
			suite.Require().Equal(sdk.ZeroInt(), balance.Amount)

			// send claim msg
			msg := types.MsgClaimAirdropRequest{
				UserAddress: test.user.String(),
				QuestType:   test.questType,
			}
			_, err := msgServer.ClaimAirdrop(sdk.WrapSDKContext(suite.Ctx), &msg)
			if test.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)

			// if claim tx is successfully executed,
			// user balance should be increased
			coin := bankKeeper.GetBalance(ctx, test.user, airdropInfo.AirdropDenom)
			suite.Require().Equal(coin.Amount, expectClaimed)

			// if performDoubleClaim is true, send claim msg again
			// error expected
			if test.performDoubleClaim {
				_, err := msgServer.ClaimAirdrop(sdk.WrapSDKContext(suite.Ctx), &msg)
				suite.Require().Error(err)
				return
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMarkSocialQuestPerformed() {
	var users []string
	for i := 0; i < 10; i++ {
		users = append(users, sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String())
	}

	tests := map[string]struct {
		signer       sdk.AccAddress
		addresses    []string
		before       func(ctx sdk.Context, k *keeper.Keeper)
		shouldErr    bool
		executeTwice bool
	}{
		"only controller user can execute the mark tx": {
			signer: validUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
			},
			shouldErr: true,
		},
		"controller user must execute the tx within the airdrop period": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				airdropInfo := validAirdropInfo(ctx)

				// airdrop is already ended
				airdropInfo.AirdropEndTimestamp = ctx.BlockTime().Add(-1 * time.Second)
				k.SetAirdropInfo(ctx, airdropInfo)
			},
			shouldErr: true,
		},
		"the quest state of users is marked as claimable": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))

				for _, user := range users {
					userState := validUserState(ctx, user, maxTotalAllocPerUser)
					addr, err := sdk.AccAddressFromBech32(user)
					suite.Require().NoError(err)

					err = k.SetUserState(ctx, addr, userState)
					suite.Require().NoError(err)
				}
			},
			addresses: users,
			shouldErr: false,
		},
		"return error if controller user try to mark the same user twice": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				err := k.SetUserState(ctx, validUser, validUserState(ctx, validUser.String(), maxTotalAllocPerUser))
				suite.Require().NoError(err)
			},
			addresses:    []string{validUser.String()},
			shouldErr:    false,
			executeTwice: true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			airdropKeeper := suite.App.AirdropKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx

			test.before(ctx, airdropKeeper)

			// send mark tx
			msg := types.MsgMarkSocialQuestPerformedRequest{
				ControllerAddress: test.signer.String(),
				UserAddresses:     test.addresses,
			}
			_, err := msgServer.MarkSocialQuestPerformed(sdk.WrapSDKContext(suite.Ctx), &msg)
			if test.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)

			// check quest state of users
			for _, user := range test.addresses {
				addr, err := sdk.AccAddressFromBech32(user)
				suite.Require().NoError(err)

				userState, err := airdropKeeper.GetUserState(ctx, addr)
				suite.Require().NoError(err)

				// only social quest is claimable
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeSocial].State, types.QuestStateType(keeper.QuestStateClaimable))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeVoteOnProposals].State, types.QuestStateType(keeper.QuestStateNotStarted))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeProvideLiquidity].State, types.QuestStateType(keeper.QuestStateNotStarted))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeSnAssetClaim].State, types.QuestStateType(keeper.QuestStateNotStarted))
			}

			// error returned if controller user try to mark the same user twice
			if test.executeTwice {
				_, err := msgServer.MarkSocialQuestPerformed(sdk.WrapSDKContext(suite.Ctx), &msg)
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMarkUserProvidedLiquidity() {
	var users []string
	for i := 0; i < 100; i++ {
		users = append(users, sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String())
	}

	tests := map[string]struct {
		signer       sdk.AccAddress
		addresses    []string
		before       func(ctx sdk.Context, k *keeper.Keeper)
		shouldErr    bool
		executeTwice bool
	}{
		"only controller user can execute the mark tx": {
			signer: validUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
			},
			shouldErr: true,
		},
		"controller user must execute the tx within the airdrop period": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				airdropInfo := validAirdropInfo(ctx)

				// airdrop is already ended
				airdropInfo.AirdropEndTimestamp = ctx.BlockTime().Add(-1 * time.Second)
				k.SetAirdropInfo(ctx, airdropInfo)
			},
			shouldErr: true,
		},
		"the quest state of users is marked as claimable": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))

				for _, user := range users {
					userState := validUserState(ctx, user, maxTotalAllocPerUser)
					addr, err := sdk.AccAddressFromBech32(user)
					suite.Require().NoError(err)

					err = k.SetUserState(ctx, addr, userState)
					suite.Require().NoError(err)
				}
			},
			addresses: users,
			shouldErr: false,
		},
		"return error if controller user try to mark the same user twice": {
			signer: controllerUser,
			before: func(ctx sdk.Context, k *keeper.Keeper) {
				k.SetAirdropInfo(ctx, validAirdropInfo(ctx))
				err := k.SetUserState(ctx, validUser, validUserState(ctx, validUser.String(), maxTotalAllocPerUser))
				suite.Require().NoError(err)
			},
			addresses:    []string{validUser.String()},
			shouldErr:    false,
			executeTwice: true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			airdropKeeper := suite.App.AirdropKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx

			test.before(ctx, airdropKeeper)

			// send mark tx
			msg := types.MsgMarkUserProvidedLiquidityRequest{
				ControllerAddress: test.signer.String(),
				UserAddresses:     test.addresses,
			}
			_, err := msgServer.MarkUserProvidedLiquidity(sdk.WrapSDKContext(suite.Ctx), &msg)
			if test.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)

			// check quest state of users
			for _, user := range test.addresses {
				addr, err := sdk.AccAddressFromBech32(user)
				suite.Require().NoError(err)

				userState, err := airdropKeeper.GetUserState(ctx, addr)
				suite.Require().NoError(err)

				// only social quest is claimable
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeSocial].State, types.QuestStateType(keeper.QuestStateNotStarted))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeVoteOnProposals].State, types.QuestStateType(keeper.QuestStateNotStarted))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeProvideLiquidity].State, types.QuestStateType(keeper.QuestStateClaimable))
				suite.Require().Equal(userState.QuestStates[keeper.QuestTypeSnAssetClaim].State, types.QuestStateType(keeper.QuestStateNotStarted))
			}

			// error returned if controller user try to mark the same user twice
			if test.executeTwice {
				_, err := msgServer.MarkUserProvidedLiquidity(sdk.WrapSDKContext(suite.Ctx), &msg)
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAirdropData() {

	tests := map[string]struct {
		signer     sdk.AccAddress
		userAddr   []sdk.AccAddress
		userStates []*types.UserState
		shouldErr  bool
	}{
		"only controller user can execute the airdrop data tx": {
			signer: invalidUser,
			userAddr: []sdk.AccAddress{
				[]byte{0x01}, []byte{0x02},
			},
			userStates: []*types.UserState{
				{
					Recipient:   sdk.AccAddress{0x01}.String(),
					TotalAmount: "100000000",
				},
				{
					Recipient:   sdk.AccAddress{0x02}.String(),
					TotalAmount: "200000000",
				},
			},
			shouldErr: true,
		},
		"enter user states data": {
			signer: controllerUser,
			userAddr: []sdk.AccAddress{
				[]byte{0x01}, []byte{0x02},
			},
			userStates: []*types.UserState{
				{
					Recipient:   sdk.AccAddress{0x01}.String(),
					TotalAmount: "300000000",
				},
				{
					Recipient:   sdk.AccAddress{0x02}.String(),
					TotalAmount: "1200000000",
				},
			},
			shouldErr: false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			airdropKeeper := suite.App.AirdropKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx

			airdropKeeper.SetAirdropInfo(ctx, validAirdropInfo(ctx))

			// send mark tx
			msg := types.MsgAirdropDataRequest{
				ControllerAddress: test.signer.String(),
				States:            test.userStates,
			}
			_, err := msgServer.AirdropData(sdk.WrapSDKContext(suite.Ctx), &msg)
			if test.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)

			for i, userAddr := range test.userAddr {
				userState, err := airdropKeeper.GetUserState(ctx, userAddr)
				suite.NoError(err)
				suite.Require().EqualValues(userState, test.userStates[i])
			}
		})
	}
}
