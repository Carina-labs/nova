package keeper_test

import (
	"context"

	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGRPCState() {
	queryClient := suite.queryClient
	oracleKeeper := suite.App.OracleKeeper

	tests := []struct {
		name       string
		chainInfo  types.ChainInfo
		queryDenom string
		wantErr    bool
	}{
		{
			name: "should get state",
			chainInfo: types.ChainInfo{
				Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
				OperatorAddress: fooOperator.String(),
				LastBlockHeight: 10,
			},
			queryDenom: fooDenom,
			wantErr:    false,
		},
		{
			name: "should get error",
			chainInfo: types.ChainInfo{
				Coin:            sdk.NewCoin(fooDenom, sdk.NewInt(fooBalance)),
				OperatorAddress: fooOperator.String(),
				LastBlockHeight: 10,
			},
			queryDenom: invalidDenom,
			wantErr:    true,
		},
	}

	oracleKeeper.SetParams(suite.Ctx, types.Params{
		OracleOperators: []string{fooOperator.String()},
	})

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := oracleKeeper.UpdateChainState(suite.Ctx, &tt.chainInfo)
			suite.Require().NoError(err)

			val, err := queryClient.State(context.Background(), &types.QueryStateRequest{
				ChainDenom: tt.queryDenom,
			})

			if tt.wantErr {
				suite.Require().NotNil(err, "[%s] error expected but no found", tt.name)
				return
			}

			suite.Require().NoError(err)
			suite.Require().Equal(val.LastBlockHeight, tt.chainInfo.LastBlockHeight)
			suite.Require().Equal(val.Coin, tt.chainInfo.Coin)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryParam() {
	queryClient := suite.queryClient
	tests := []struct {
		name      string
		operators []string
		expect    []string
	}{
		// grpc query replace empty value to nil
		{"empty operator", []string{}, nil},
		{"got operator", []string{fooOperator.String()}, []string{fooOperator.String()}},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.App.OracleKeeper.SetParams(suite.Ctx, types.Params{
				OracleOperators: tt.operators,
			})

			params, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			suite.Require().NoError(err)
			suite.Require().Equal(params.Params.OracleOperators, tt.expect)
		})
	}
}
