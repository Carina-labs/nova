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
				ZoneId:          fooChainId,
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
				ZoneId:          fooChainId,
			},
			queryDenom: invalidDenom,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			oracleKeeper.SetOracleAddress(suite.Ctx, tt.chainInfo.ZoneId, []string{fooOperator.String()})
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
			suite.Require().Equal(val.ZoneId, tt.chainInfo.ZoneId)

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
				OracleKeyManager: tt.operators,
			})

			params, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			suite.Require().NoError(err)
			suite.Require().Equal(params.Params.OracleKeyManager, tt.expect)
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

	trace := types.IBCTrace{
		Version: 18,
		Height:  uint64(ctx.BlockHeight()),
	}
	//set delegate version
	suite.App.OracleKeeper.SetOracleVersion(ctx, "gaia", trace)
	res, err = queryClient.OracleVersion(ctx.Context(), &types.QueryOracleVersionRequest{
		ZoneId: "gaia",
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)
}
