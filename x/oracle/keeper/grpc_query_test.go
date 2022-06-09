package keeper_test

import (
	"context"

	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGRPCState() {
	queryClient := suite.queryClient
	oracleKeeper := suite.App.OracleKeeper
	pk := ed25519.GenPrivKey().PubKey()
	operator := sdk.AccAddress(pk.Address())

	tests := []struct {
		name       string
		chainInfo  types.ChainInfo
		queryDenom string
		wantErr    bool
	}{
		{
			name: "should get state",
			chainInfo: types.ChainInfo{
				ChainDenom:         fooDenom,
				OperatorAddress:    operator.String(),
				LastBlockHeight:    10,
				TotalStakedBalance: 1000000000,
				Decimal:            6,
			},
			queryDenom: fooDenom,
			wantErr:    false,
		},
		{
			name: "should get error",
			chainInfo: types.ChainInfo{
				ChainDenom:         fooDenom,
				OperatorAddress:    operator.String(),
				LastBlockHeight:    10,
				TotalStakedBalance: 1000000000,
				Decimal:            6,
			},
			queryDenom: invalidDenom,
			wantErr:    true,
		},
	}

	oracleKeeper.SetParams(suite.Ctx, types.Params{
		OracleOperators: []string{operator.String()},
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
			suite.Require().Equal(val.TotalStakedBalance, tt.chainInfo.TotalStakedBalance)
			suite.Require().Equal(val.Decimal, tt.chainInfo.Decimal)
		})
	}
}
