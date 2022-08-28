package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	fooUser = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	barUser = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
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
				State:           types.DelegateSuccess,
				OracleVersion:   1,
				DelegateVersion: 1,
			},
		},
	})

	amt, err := queryClient.ClaimableAmount(ctx.Context(), &types.QueryClaimableAmountRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
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
	_, err := queryClient.PendingWithdrawals(ctx.Context(), &types.QueryPendingWithdrawalsRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	// query with invalid user
	invalidResult, err := queryClient.PendingWithdrawals(ctx.Context(), &types.QueryPendingWithdrawalsRequest{
		ZoneId:  zoneId,
		Address: "invalid_user",
	})
	suite.Require().NoError(err)
	suite.Require().Equal(denom, invalidResult.Amount.Denom)
	suite.Require().Equal(sdk.ZeroInt(), invalidResult.Amount.Amount)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	amount := sdk.NewInt64Coin(denom, 100)
	records[1] = &types.WithdrawRecordContent{
		Amount:          amount.Amount,
		State:           types.WithdrawStatusRegistered,
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	records[2] = &types.WithdrawRecordContent{
		Amount:          amount.Amount,
		State:           types.WithdrawStatusRegistered,
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
	result, err := queryClient.PendingWithdrawals(ctx.Context(), &types.QueryPendingWithdrawalsRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
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
	_, err := queryClient.ActiveWithdrawals(ctx.Context(), &types.QueryActiveWithdrawalsRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	pendingAmount := sdk.NewCoin(zoneBaseDenom, sdk.NewInt(150))
	activeAmount := sdk.NewCoin(zoneBaseDenom, sdk.NewInt(80))

	records[1] = &types.WithdrawRecordContent{
		Amount:          activeAmount.Amount,
		State:           types.WithdrawStatusTransferred,
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	records[2] = &types.WithdrawRecordContent{
		Amount:          pendingAmount.Amount,
		State:           types.WithdrawStatusRegistered,
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
	result, err := queryClient.ActiveWithdrawals(ctx.Context(), &types.QueryActiveWithdrawalsRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})

	// only transferred amount should be returned
	suite.Require().NoError(err)
	suite.Require().Equal(result.Amount.Denom, denom)
	suite.Require().Equal(result.Amount.Amount, activeAmount.Amount)
}

func (suite *KeeperTestSuite) TestQueryDepositRecord() {
	// Save the deposit record to the kv store
	// and query the result and check it is correct
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper

	// query with invalid zone
	_, err := queryClient.DepositRecords(ctx.Context(), &types.QueryDepositRecordRequest{
		ZoneId:  "invalid",
		Address: fooUser.String(),
	})
	suite.Require().Error(err)

	// Save the deposit record to the keeper
	token := sdk.NewInt64Coin(zoneBaseDenom, 100)
	fooRecords := &types.DepositRecord{
		ZoneId:  zoneId,
		Claimer: fooUser.String(),
		Records: []*types.DepositRecordContent{
			{
				Depositor:       fooUser.String(),
				Amount:          &token,
				State:           types.DepositRequest,
				OracleVersion:   0,
				DelegateVersion: 0,
			},
		},
	}

	galKeeper.SetDepositRecord(ctx, fooRecords)

	fooRecords.Claimer = barUser.String()
	galKeeper.SetDepositRecord(ctx, fooRecords)

	ret, err := queryClient.DepositRecords(ctx.Context(), &types.QueryDepositRecordRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})

	suite.Require().NoError(err)
	fooRecords.Claimer = fooUser.String()
	suite.Require().Equal(fooRecords, ret.DepositRecord)
}

func (suite *KeeperTestSuite) TestQueryWithdrawRecord() {
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper

	// query with invalid zone
	_, err := queryClient.WithdrawRecords(ctx.Context(), &types.QueryWithdrawRecordRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})
	suite.Require().Error(err)

	// Save the withdrawal record to the keeper
	records := make(map[uint64]*types.WithdrawRecordContent)
	token := sdk.NewCoin(zoneBaseDenom, sdk.NewInt(80))
	records[0] = &types.WithdrawRecordContent{
		Amount:          token.Amount,
		State:           types.WithdrawStatusTransferred,
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	withdrawRecords := &types.WithdrawRecord{
		ZoneId:     zoneId,
		Withdrawer: fooUser.String(),
		Records:    records,
	}

	galKeeper.SetWithdrawRecord(ctx, withdrawRecords)

	ret, err := queryClient.WithdrawRecords(ctx.Context(), &types.QueryWithdrawRecordRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})

	suite.Require().NoError(err)
	suite.Require().Equal(withdrawRecords, ret.WithdrawRecord)
}

func (suite *KeeperTestSuite) TestQueryUndelegateRecord() {
	// Save the undelegate record to the kv store
	// and query the result and check it is correct
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper

	// query with invalid zone
	_, err := queryClient.UndelegateRecords(ctx.Context(), &types.QueryUndelegateRecordRequest{
		ZoneId:  "invalid",
		Address: fooUser.String(),
	})
	suite.Require().Error(err)

	// save the undelegate recod to the keeper
	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(ctx, transferPort, transferChannel, zoneBaseDenom)
	snToken := sdk.NewInt64Coin(ibcDenom, 100)
	withdrawAmount := sdk.NewInt(150)

	records := &types.UndelegateRecord{
		ZoneId:    zoneId,
		Delegator: fooUser.String(),
		Records: []*types.UndelegateRecordContent{
			{
				Withdrawer:        fooUser.String(),
				SnAssetAmount:     &snToken,
				WithdrawAmount:    withdrawAmount,
				State:             0,
				OracleVersion:     0,
				UndelegateVersion: 0,
			},
		},
	}
	galKeeper.SetUndelegateRecord(ctx, records)
	ret, err := queryClient.UndelegateRecords(ctx.Context(), &types.QueryUndelegateRecordRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})
	suite.Require().NoError(err)
	suite.Require().Equal(records, ret.UndelegateRecord)
}
