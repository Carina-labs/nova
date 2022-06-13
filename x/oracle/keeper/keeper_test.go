package keeper_test

import (
	"testing"

	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

var (
	fooDenom           = "uatom"
	invalidDenom       = "invalid_denom"
	fooOperator        = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	fooBalance   int64 = 1000000000
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestUpdateChainState() {
	oracleKeeper := suite.App.OracleKeeper

	chainInfo := types.ChainInfo{
		Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
		OperatorAddress: fooOperator.String(),
		LastBlockHeight: 10,
		Decimal:         6,
	}

	tests := []struct {
		name       string
		chainInfo  types.ChainInfo
		queryDenom string
		operator   *sdk.AccAddress
		wantErr    bool
	}{
		{
			name: "no operator",
			chainInfo: types.ChainInfo{
				Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
				OperatorAddress: fooOperator.String(),
				LastBlockHeight: 10,
				Decimal:         6,
			},
			queryDenom: fooDenom,
			operator:   nil,
			wantErr:    true,
		},
		{
			name: "no data with incorrect query",
			chainInfo: types.ChainInfo{
				Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
				OperatorAddress: fooOperator.String(),
				LastBlockHeight: 10,
				Decimal:         6,
			},
			queryDenom: invalidDenom,
			operator:   &fooOperator,
			wantErr:    true,
		},
		{
			name: "should success",
			chainInfo: types.ChainInfo{
				Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
				OperatorAddress: fooOperator.String(),
				LastBlockHeight: 10,
				Decimal:         6,
			},
			queryDenom: fooDenom,
			operator:   &fooOperator,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if tt.operator != nil {
				oracleKeeper.SetParams(suite.Ctx, types.Params{
					OracleOperators: []string{tt.operator.String()},
				})
			}

			err := oracleKeeper.UpdateChainState(suite.Ctx, &chainInfo)
			if tt.operator == nil && tt.wantErr {
				suite.Require().NotNil(err, "[%s] error expected but no error found", tt.name)
				return
			}
			suite.Require().NoError(err)

			got, err := oracleKeeper.GetChainState(suite.Ctx, tt.queryDenom)
			if tt.wantErr {
				suite.Require().NotNil(err, "[%s] error expected but no error found", tt.name)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(&chainInfo, got)
			}
		})
	}
}
