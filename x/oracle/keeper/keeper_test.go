package keeper_test

import (
	"github.com/Carina-labs/nova/x/oracle/keeper"
	"testing"

	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

var (
	fooDenom               = "uatom"
	invalidDenom           = "invalid_denom"
	fooChainId             = "testchain"
	fooAppHash             = "apphash"
	fooBlockProposer       = "cosmos-block-proposer"
	fooOperator            = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	fooBalance       int64 = 1000000000
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.OracleKeeper)
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
		AppHash:         []byte(fooAppHash),
		ZoneId:          fooChainId,
	}

	tests := []struct {
		name       string
		chainInfo  types.ChainInfo
		queryDenom string
		operator   *sdk.AccAddress
		wantErr    bool
	}{
		{
			"no operator",
			chainInfo, fooDenom, nil, true,
		},
		{
			"no data with incorrect query",
			chainInfo, invalidDenom, &fooOperator, true,
		},
		{
			"should success",
			chainInfo, fooDenom, &fooOperator, false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if tt.operator != nil {
				oracleKeeper.SetOracleAddress(suite.Ctx, chainInfo.ZoneId, []string{tt.operator.String()})
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

func (suite *KeeperTestSuite) TestSetOracleVersion() {
	testZoneId := "cosmos"
	version := uint64(1)

	// set & get oracle version
	oracleKeeper := suite.App.OracleKeeper
	trace := types.IBCTrace{
		Version: version,
		Height:  uint64(suite.Ctx.BlockHeight()),
	}
	oracleKeeper.SetOracleVersion(suite.Ctx, testZoneId, trace)
	got, _ := oracleKeeper.GetOracleVersion(suite.Ctx, testZoneId)
	suite.Require().Equal(version, got)

	// update to new version
	newVersion := uint64(2)
	trace = types.IBCTrace{
		Version: newVersion,
		Height:  uint64(suite.Ctx.BlockHeight()),
	}
	oracleKeeper.SetOracleVersion(suite.Ctx, testZoneId, trace)
	got, _ = oracleKeeper.GetOracleVersion(suite.Ctx, testZoneId)
	suite.Require().Equal(newVersion, got)

	// if there's no version info, return 0
	got, _ = oracleKeeper.GetOracleVersion(suite.Ctx, "unknown")
	suite.Require().Equal(uint64(0), got)
}
