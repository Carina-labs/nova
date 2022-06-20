package keeper_test

// func (suite *KeeperTestSuite) TestDeposit() {
//tcs := []struct {
//	userPrivKey          *secp256k1.PrivKey
//	denom                string
//	preUserAmt           int64
//	preModuleAccountAmt  int64
//	depositAmt           int64
//	postUserAmt          int64
//	postModuleAccountAmt int64
//	shouldErr            bool
//}{
//	{
//		userPrivKey:          secp256k1.GenPrivKey(),
//		denom:                "atom",
//		preUserAmt:           1000,
//		preModuleAccountAmt:  1000,
//		depositAmt:           500,
//		postUserAmt:          500,
//		postModuleAccountAmt: 1000,
//		shouldErr:            false,
//	},
//	{
//		userPrivKey:          secp256k1.GenPrivKey(),
//		denom:                "atom",
//		preUserAmt:           4000,
//		preModuleAccountAmt:  5000,
//		depositAmt:           3000,
//		postUserAmt:          1000,
//		postModuleAccountAmt: 5000,
//		shouldErr:            false,
//	},
//	{
//		// ERROR CASE
//		userPrivKey:          secp256k1.GenPrivKey(),
//		denom:                "juno",
//		preUserAmt:           5000,
//		preModuleAccountAmt:  1000,
//		depositAmt:           6000,
//		postUserAmt:          5000,
//		postModuleAccountAmt: 1000,
//		shouldErr:            true,
//	},
//}
//
//msgServer := keeper.NewMsgServerImpl(*suite.App.GalKeeper)
//
//for _, tc := range tcs {
//	acc := authtypes.NewBaseAccount(tc.userPrivKey.PubKey().Address().Bytes(), tc.userPrivKey.PubKey(), 0, 0)
//	galAddr := suite.App.AccountKeeper.GetModuleAddress(types.ModuleName)
//	accAddr, err := sdk.AccAddressFromBech32(acc.Address)
//	suite.NoError(err)
//	// setup
//	suite.App.BankKeeper.InitGenesis(suite.Ctx, &banktypes.GenesisState{
//		Balances: []banktypes.Balance{
//			{
//				Address: accAddr.String(),
//				Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.preUserAmt)},
//			},
//			{
//				Address: galAddr.String(),
//				Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.preModuleAccountAmt)},
//			},
//		},
//	})
//	coin := sdk.NewInt64Coin(tc.denom, tc.depositAmt)
//	msg := types.NewMsgDeposit(accAddr, sdk.NewCoins(coin), "channel-1")
//	goCtx := sdk.WrapSDKContext(suite.Ctx)
//
//	// execute
//	_, err = msgServer.Deposit(goCtx, msg)
//
//	// assert
//	if tc.shouldErr {
//		suite.Error(err)
//	} else {
//		suite.NoError(err)
//	}
//
//	postUserBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, accAddr, tc.denom)
//	postGalBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, galAddr, tc.denom)
//	suite.Equal(tc.postUserAmt, postUserBalance.Amount.Int64())
//	suite.Equal(tc.postModuleAccountAmt, postGalBalance.Amount.Int64())
//}
// }
