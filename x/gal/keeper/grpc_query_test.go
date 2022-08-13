package keeper_test

import (
	galkeeper "github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	fooUser = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
)

func (suite *KeeperTestSuite) TestClaimableAssetQuery() {
	queryClient := suite.queryClient
	keeper := suite.App.GalKeeper
	ibcstaking := suite.App.IbcstakingKeeper
	oracleKeeper := suite.App.OracleKeeper
	ctx := suite.Ctx

	denom := ibcstaking.GetIBCHashDenom(ctx, transferPort, transferChannel, zoneBaseDenom)
	amount := sdk.NewInt(1000_000000)
	coin := sdk.NewCoin(denom, amount)

	oracleKeeper.SetOracleVersion(ctx, zoneId, 2)
	keeper.SetDepositRecord(ctx, &types.DepositRecord{
		ZoneId:  zoneId,
		Claimer: fooUser.String(),
		Records: []*types.DepositRecordContent{
			{
				Depositor:       fooUser.String(),
				Amount:          &coin,
				State:           int64(galkeeper.DELEGATE_SUCCESS),
				OracleVersion:   1,
				DelegateVersion: 1,
			},
		},
	})

	amt, err := queryClient.ClaimableAmount(ctx.Context(), &types.ClaimableAmountRequest{
		ZoneId:            zoneId,
		Address:           fooUser.String(),
		TransferPortId:    transferPort,
		TransferChannelId: transferChannel,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(amt.Amount.Denom, denom)
	suite.Require().Equal(amt.Amount.Amount.Int64(), amount.Int64())
}

func (suite *KeeperTestSuite) TestQueryPendingWithdrawals() {
	// add pending withdrawals to the kv store
	// query pending withdrawals and check if the result is correct
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper
	denom := suite.App.IbcstakingKeeper.GetIBCHashDenom(ctx, transferPort, transferChannel, zoneBaseDenom)

	// query with invalid zone
	_, err := queryClient.PendingWithdrawals(ctx.Context(), &types.PendingWithdrawalsRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	// query with invalid user
	invalidResult, err := queryClient.PendingWithdrawals(ctx.Context(), &types.PendingWithdrawalsRequest{
		ZoneId:            zoneId,
		Address:           "invalid_user",
		TransferPortId:    transferPort,
		TransferChannelId: transferChannel,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(denom, invalidResult.Amount.Denom)
	suite.Require().Equal(sdk.ZeroInt(), invalidResult.Amount.Amount)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	amount := sdk.NewInt64Coin(denom, 100)
	records[1] = &types.WithdrawRecordContent{
		Amount:          &amount,
		State:           int64(galkeeper.WithdrawStatus_Registered),
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	records[2] = &types.WithdrawRecordContent{
		Amount:          &amount,
		State:           int64(galkeeper.WithdrawStatus_Registered),
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	galKeeper.SetWithdrawRecord(ctx, &types.WithdrawRecord{
		ZoneId:     zoneId,
		Withdrawer: fooUser.String(),
		Records:    records,
	})

	// query the pending withdrawal amount and check if the amount is correct
	result, err := queryClient.PendingWithdrawals(ctx.Context(), &types.PendingWithdrawalsRequest{
		ZoneId:            zoneId,
		Address:           fooUser.String(),
		TransferPortId:    transferPort,
		TransferChannelId: transferChannel,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(result.Amount.Denom, denom)
	suite.Require().Equal(result.Amount.Amount, sdk.NewInt(200))
}

func (suite *KeeperTestSuite) TestQueryActiveWithdrawals() {
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper
	denom := suite.App.IbcstakingKeeper.GetIBCHashDenom(ctx, transferPort, transferChannel, zoneBaseDenom)

	// query with invalid zone
	_, err := queryClient.ActiveWithdrawals(ctx.Context(), &types.ActiveWithdrawalsRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	pendingAmount := sdk.NewCoin(zoneBaseDenom, sdk.NewInt(150))
	activeAmount := sdk.NewCoin(zoneBaseDenom, sdk.NewInt(80))

	records[1] = &types.WithdrawRecordContent{
		Amount:          &activeAmount,
		State:           int64(galkeeper.WithdrawStatus_Transferred),
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	records[2] = &types.WithdrawRecordContent{
		Amount:          &pendingAmount,
		State:           int64(galkeeper.WithdrawStatus_Registered),
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	galKeeper.SetWithdrawRecord(ctx, &types.WithdrawRecord{
		ZoneId:     zoneId,
		Withdrawer: fooUser.String(),
		Records:    records,
	})

	// query the pending withdrawal amount and check if the amount is correct
	result, err := queryClient.ActiveWithdrawals(ctx.Context(), &types.ActiveWithdrawalsRequest{
		ZoneId:            zoneId,
		Address:           fooUser.String(),
		TransferPortId:    transferPort,
		TransferChannelId: transferChannel,
	})

	// only transferred amount should be returned
	suite.Require().NoError(err)
	suite.Require().Equal(result.Amount.Denom, denom)
	suite.Require().Equal(result.Amount.Amount, activeAmount.Amount)
}

func (suite *KeeperTestSuite) TestQueryDepositRecord() {

}

func (suite *KeeperTestSuite) TestQueryWithdrawRecord() {

}

func (suite *KeeperTestSuite) TestQueryUndelegateRecord() {

}
