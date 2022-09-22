package keeper_test

import (
	"context"
	icacontroltypes "github.com/Carina-labs/nova/x/icacontrol/types"

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
				AppHash:         []byte("apphash"),
				ChainId:         fooChainId,
				OracleVersion:   0,
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
				AppHash:         []byte("apphash"),
				ChainId:         fooChainId,
				OracleVersion:   0,
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
			suite.Require().Equal(val.AppHash, tt.chainInfo.AppHash)
			suite.Require().Equal(val.ChainId, tt.chainInfo.ChainId)

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

func (suite *KeeperTestSuite) TestQueryOracleVersion() {
	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	zone := &icacontroltypes.RegisteredZone{
		ZoneId: "gaia",
		IcaConnectionInfo: &icacontroltypes.IcaConnectionInfo{
			ConnectionId: "connection-0",
			PortId:       "test",
		},
		TransferInfo: &icacontroltypes.TransferConnectionInfo{
			PortId:    "transfer",
			ChannelId: "channel-0",
		},
		ValidatorAddress: "valAddr",
		BaseDenom:        "uatom",
		SnDenom:          "snuatom",
		Decimal:          6,
	}

	//set chain info
	suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, zone)

	// query with invalid zone
	_, err := queryClient.OracleVersion(ctx.Context(), &types.QueryOracleVersionRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//sequence is zero
	exp := types.QueryOracleVersionResponse{Version: 0, Height: 0}

	res, err := queryClient.OracleVersion(ctx.Context(), &types.QueryOracleVersionRequest{
		ZoneId: "gaia",
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	//sequence is 18
	exp = types.QueryOracleVersionResponse{Version: 18, Height: 1}

	//set delegate version
	suite.App.OracleKeeper.SetOracleVersion(ctx, "gaia", 18, uint64(ctx.BlockHeight()))
	res, err = queryClient.OracleVersion(ctx.Context(), &types.QueryOracleVersionRequest{
		ZoneId: "gaia",
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)
}
