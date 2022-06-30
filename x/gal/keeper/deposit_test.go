package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *KeeperTestSuite) TestRecordDepositAmt() {
	randAddr := suite.GenRandomAddress()
	type args struct {
		coin sdk.Coin
		addr sdk.AccAddress
	}
	tcs := []struct {
		name    string
		args    []args
		expect  []args
		wantErr bool

		denom         string
		amt           int64
		userAddr      sdk.AccAddress
		expectedDenom string
		expectedAmt   int64
	}{
		{
			name: "should get recorded deposit amt",
			args: []args{
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			expect: []args{
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			wantErr: false,
		},
		{
			name: "should not get deposit info",
			args: []args{},
			expect: []args{
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest()

			for _, arg := range tc.args {
				err := suite.App.GalKeeper.SetDepositAmt(
					suite.Ctx,
					&types.DepositRecord{
						Address: arg.addr.String(),
						Records: []*types.DepositRecordContent{
							{
								ZoneId:        "test-zone-id",
								Amount:        &arg.coin,
								IsTransferred: false,
							},
						},
					})
				suite.Require().NoError(err)
			}

			for _, query := range tc.expect {
				res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, query.addr)
				if tc.wantErr {
					suite.Require().NotNil(err, "error expected but no error found")
					suite.Require().Equal(err, types.ErrNoDepositRecord)
					continue
				}

				suite.Require().NoError(err)
				for _, record := range res.Records {
					suite.Require().Equal(record.Amount.Denom, query.coin.Denom)
					suite.Require().Equal(record.Amount.Amount, query.coin.Amount)
					suite.Require().Equal(res.Address, query.addr.String())
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
