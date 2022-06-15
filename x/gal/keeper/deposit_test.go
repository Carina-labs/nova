package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *KeeperTestSuite) TestRecordDepositAmt() {
	tcs := []struct {
		denom         string
		amt           int64
		userAddr      sdk.AccAddress
		expectedDenom string
		expectedAmt   int64
	}{
		{
			denom:         "osmo",
			amt:           10000,
			userAddr:      suite.GenRandomAddress(),
			expectedDenom: "osmo",
			expectedAmt:   10000,
		},
		{
			denom:         "atom",
			amt:           5555,
			userAddr:      suite.GenRandomAddress(),
			expectedDenom: "atom",
			expectedAmt:   5555,
		},
	}

	for _, tc := range tcs {
		amt := sdk.NewInt64Coin(tc.denom, tc.amt)
		depositMsg := types.DepositRecord{
			Address: tc.userAddr.String(),
			Amount:  &amt,
		}
		// Test RecordDepositAmt
		suite.App.GalKeeper.RecordDepositAmt(suite.Ctx, depositMsg)

		// Test GetRecordDepositAmt
		res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, tc.userAddr)
		suite.NoError(err)
		suite.Equal(tc.expectedAmt, res.Amount.Amount.Int64())
		suite.Equal(tc.expectedDenom, res.Amount.Denom)
		suite.Equal(tc.userAddr.String(), res.Address)
	}
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
