package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (suite *KeeperTestSuite) TestIsAbleToWithdraw() {
	tcs := []struct {
		denom       string
		initialAmt  int64
		withdrawAmt int64
		expected    bool
	}{
		{
			denom:       "atom",
			initialAmt:  100000,
			withdrawAmt: 1000,
			expected:    true,
		},
		{
			denom:       "osmo",
			initialAmt:  5000,
			withdrawAmt: 6000,
			expected:    false,
		},
		{
			denom:       "juno",
			initialAmt:  9500,
			withdrawAmt: 9500,
			expected:    true,
		},
	}

	for _, tc := range tcs {
		err := suite.App.BankKeeper.MintCoins(suite.Ctx,
			types.ModuleName,
			sdk.NewCoins(sdk.NewInt64Coin(tc.denom, tc.initialAmt)))
		suite.NoError(err)

		res, err := suite.App.GalKeeper.IsAbleToWithdraw(suite.Ctx, sdk.NewInt64Coin(tc.denom, tc.withdrawAmt))
		suite.Equal(tc.expected, res)
	}
}

func (suite *KeeperTestSuite) TestClaimWithdrawAsset() {
	tcs := []struct {
		denom           string
		userKey         *secp256k1.PrivKey
		initialAmt      int64
		withdrawAmt     int64
		shouldError     bool
		postUserBalance int64
	}{
		{
			denom:           "atom",
			userKey:         secp256k1.GenPrivKey(),
			initialAmt:      100000,
			withdrawAmt:     1000,
			shouldError:     false,
			postUserBalance: 1000,
		},
		{
			denom:           "osmo",
			userKey:         secp256k1.GenPrivKey(),
			initialAmt:      5000,
			withdrawAmt:     6000,
			shouldError:     true,
			postUserBalance: 0,
		},
		{
			denom:           "juno",
			userKey:         secp256k1.GenPrivKey(),
			initialAmt:      9500,
			withdrawAmt:     9500,
			shouldError:     false,
			postUserBalance: 9500,
		},
	}

	for _, tc := range tcs {
		acc := authtypes.NewBaseAccount(tc.userKey.PubKey().Address().Bytes(), tc.userKey.PubKey(), 0, 0)
		err := suite.App.BankKeeper.MintCoins(suite.Ctx,
			types.ModuleName,
			sdk.NewCoins(sdk.NewInt64Coin(tc.denom, tc.initialAmt)))
		suite.NoError(err)

		err = suite.App.GalKeeper.ClaimWithdrawAsset(suite.Ctx,
			acc.Address, sdk.NewInt64Coin(tc.denom, tc.withdrawAmt))

		if tc.shouldError {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}

		goCtx := sdk.WrapSDKContext(suite.Ctx)
		balance, err := suite.App.BankKeeper.Balance(goCtx, &types2.QueryBalanceRequest{
			Address: acc.Address,
			Denom:   tc.denom,
		})
		suite.NoError(err)
		suite.Equal(tc.postUserBalance, balance.Balance.Amount.Int64())
	}
}
