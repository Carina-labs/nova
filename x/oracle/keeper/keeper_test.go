package keeper_test

import (
	"testing"

	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestUpdateChainState() {
	oracleKeeper := suite.App.OracleKeeper
	pk := ed25519.GenPrivKey().PubKey()
	operator := sdk.AccAddress(pk.Address())

	chainInfo := types.ChainInfo{
		ChainDenom:         "atom",
		OperatorAddress:    operator.String(),
		LastBlockHeight:    10,
		TotalStakedBalance: 1000000000,
		Decimal:            6,
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
				ChainDenom:         "atom",
				OperatorAddress:    operator.String(),
				LastBlockHeight:    10,
				TotalStakedBalance: 1000000000,
				Decimal:            6,
			},
			queryDenom: "atom",
			operator:   nil,
			wantErr:    true,
		},
		{
			name: "no data with incorrect query",
			chainInfo: types.ChainInfo{
				ChainDenom:         "atom",
				OperatorAddress:    operator.String(),
				LastBlockHeight:    10,
				TotalStakedBalance: 1000000000,
				Decimal:            6,
			},
			queryDenom: "nova",
			operator:   &operator,
			wantErr:    true,
		},
		{
			name: "should success",
			chainInfo: types.ChainInfo{
				ChainDenom:         "atom",
				OperatorAddress:    operator.String(),
				LastBlockHeight:    10,
				TotalStakedBalance: 1000000000,
				Decimal:            6,
			},
			queryDenom: "atom",
			operator:   &operator,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		if tt.operator != nil {
			oracleKeeper.SetParams(suite.Ctx, types.Params{
				OracleOperators: []string{tt.operator.String()},
			})
		}

		err := oracleKeeper.UpdateChainState(suite.Ctx, &chainInfo)
		if tt.operator == nil && tt.wantErr {
			suite.Require().NotNil(err, "[%s] error expected but no error found", tt.name)
			continue
		}
		suite.Require().NoError(err)

		got, err := oracleKeeper.GetChainState(suite.Ctx, tt.queryDenom)
		if tt.wantErr {
			suite.Require().NotNil(err, "[%s] error expected but no error found", tt.name)
		} else {
			suite.Require().NoError(err)
			suite.Require().Equal(&chainInfo, got)
		}
	}
}
