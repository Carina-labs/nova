package keeper_test

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	oracletypes "github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"time"
)

var (
	fooUser = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	barUser = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
)

func (suite *KeeperTestSuite) TestClaimableAssetQuery() {
	denomTrace := transfertypes.DenomTrace{
		Path:      transferPort + "/" + transferChannel,
		BaseDenom: baseDenom,
	}
	suite.App.TransferKeeper.SetDenomTrace(suite.Ctx, denomTrace)

	suite.App.OracleKeeper.InitGenesis(suite.Ctx, &oracletypes.GenesisState{
		Params: oracletypes.Params{
			OracleKeyManager: []string{
				baseOwnerAcc.String(),
			},
		},
		OracleAddressInfo: []oracletypes.OracleAddressInfo{
			{
				ZoneId:        zoneId,
				OracleAddress: []string{baseOwnerAcc.String()},
			},
		},
		States: []oracletypes.ChainInfo{
			{
				Coin:            sdk.NewCoin(baseDenom, sdk.NewInt(1000_000000)),
				ZoneId:          zoneId,
				OperatorAddress: baseOwnerAcc.String(),
			},
		},
	})

	queryClient := suite.queryClient
	keeper := suite.App.GalKeeper
	icaControlKeeper := suite.App.IcaControlKeeper
	oracleKeeper := suite.App.OracleKeeper
	ctx := suite.Ctx

	denom := icaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	amount := sdk.NewInt(1000_000000)
	coin := sdk.NewCoin(denom, amount)

	trace := oracletypes.IBCTrace{
		Version: 2,
		Height:  uint64(ctx.BlockHeight()),
	}

	oracleKeeper.SetOracleVersion(ctx, zoneId, trace)
	keeper.SetDelegateRecord(ctx, &types.DelegateRecord{
		ZoneId:  zoneId,
		Claimer: fooUser.String(),
		Records: map[uint64]*types.DelegateRecordContent{
			1: {
				Amount: &coin,
				State:  types.DelegateSuccess,
			},
		},
	})

	amt, err := queryClient.ClaimableAmount(ctx.Context(), &types.QueryClaimableAmountRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})

	suite.Require().NoError(err)
	suite.Require().Equal(amt.Amount.Denom, baseSnDenom)
	suite.Require().Equal(amt.Amount.Amount, sdk.NewIntWithDecimal(amount.Int64(), 18))
}

func (suite *KeeperTestSuite) TestDepositAmountQuery() {
	queryClient := suite.queryClient
	keeper := suite.App.GalKeeper
	ibcstaking := suite.App.IcaControlKeeper
	oracleKeeper := suite.App.OracleKeeper
	ctx := suite.Ctx

	denom := ibcstaking.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	amount := sdk.NewInt(1000_000000)
	coin := sdk.NewCoin(denom, amount)

	trace := oracletypes.IBCTrace{
		Version: 2,
		Height:  uint64(ctx.BlockHeight()),
	}

	oracleKeeper.SetOracleVersion(ctx, zoneId, trace)
	keeper.SetDepositRecord(ctx, &types.DepositRecord{
		ZoneId:    zoneId,
		Depositor: fooUser.String(),
		Records: []*types.DepositRecordContent{
			{
				Claimer: fooUser.String(),
				Amount:  &coin,
				State:   types.DepositSuccess,
			},
		},
	})

	amt, err := queryClient.DepositAmount(ctx.Context(), &types.QueryDepositAmountRequest{
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
	suite.Require().Equal(baseSnDenom, invalidResult.Amount.Denom)
	suite.Require().Equal(sdk.ZeroInt(), invalidResult.Amount.Amount)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	amount := sdk.NewInt64Coin(baseSnDenom, 100)
	snAmount := sdk.NewCoin(baseSnDenom, sdk.NewInt(100))
	records[1] = &types.WithdrawRecordContent{
		Amount:          amount.Amount,
		UnstakingAmount: &snAmount,
		State:           types.WithdrawStatusRegistered,
		OracleVersion:   1,
		WithdrawVersion: 1,
		CompletionTime:  time.Time{},
	}

	records[2] = &types.WithdrawRecordContent{
		Amount:          amount.Amount,
		UnstakingAmount: &snAmount,
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
	suite.Require().Equal(result.Amount.Denom, baseSnDenom)
	suite.Require().Equal(result.Amount.Amount, sdk.NewInt(200))
}

func (suite *KeeperTestSuite) TestQueryActiveWithdrawals() {
	queryClient := suite.queryClient
	ctx := suite.Ctx
	galKeeper := suite.App.GalKeeper
	denom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)

	// query with invalid zone
	_, err := queryClient.ActiveWithdrawals(ctx.Context(), &types.QueryActiveWithdrawalsRequest{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	// fill the fake content
	records := make(map[uint64]*types.WithdrawRecordContent)

	pendingAmount := sdk.NewCoin(baseDenom, sdk.NewInt(150))
	activeAmount := sdk.NewCoin(baseDenom, sdk.NewInt(80))

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
	token := sdk.NewInt64Coin(baseDenom, 100)
	fooRecords := &types.DepositRecord{
		ZoneId:    zoneId,
		Depositor: fooUser.String(),
		Records: []*types.DepositRecordContent{
			{
				Claimer: fooUser.String(),
				Amount:  &token,
				State:   types.DepositRequest,
			},
		},
	}

	galKeeper.SetDepositRecord(ctx, fooRecords)

	fooRecords.Depositor = barUser.String()
	galKeeper.SetDepositRecord(ctx, fooRecords)

	ret, err := queryClient.DepositRecords(ctx.Context(), &types.QueryDepositRecordRequest{
		ZoneId:  zoneId,
		Address: fooUser.String(),
	})

	suite.Require().NoError(err)
	fooRecords.Depositor = fooUser.String()
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
	token := sdk.NewCoin(baseDenom, sdk.NewInt(80))
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
	ibcDenom := suite.App.IcaControlKeeper.GetIBCHashDenom(transferPort, transferChannel, baseDenom)
	snAsset := sdk.NewInt64Coin(ibcDenom, 100)
	withdrawAmount := sdk.NewInt(150)

	records := &types.UndelegateRecord{
		ZoneId:    zoneId,
		Delegator: fooUser.String(),
		Records: []*types.UndelegateRecordContent{
			{
				Withdrawer:        fooUser.String(),
				SnAssetAmount:     &snAsset,
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

func (suite *KeeperTestSuite) TestQueryDelegateVersion() {

	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	// query with invalid zone
	_, err := queryClient.DelegateVersion(ctx.Context(), &types.QueryDelegateVersion{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//version info is nil
	exp := types.QueryDelegateVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 0,
		},
	}

	res, err := queryClient.DelegateVersion(ctx.Context(), &types.QueryDelegateVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	//sequence is 8
	exp = types.QueryDelegateVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 8,
			Height:  100,
			State:   types.IcaPending,
		},
	}

	//set delegate version
	versionInfo := types.VersionState{
		ZoneId:         zoneId,
		CurrentVersion: 8,
		Record: map[uint64]*types.IBCTrace{
			8: {
				Version: 8,
				Height:  100,
				State:   types.IcaPending,
			},
		},
	}
	suite.App.GalKeeper.SetDelegateVersion(ctx, zoneId, versionInfo)
	res, err = queryClient.DelegateVersion(ctx.Context(), &types.QueryDelegateVersion{
		ZoneId:  zoneId,
		Version: 8,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	// get current sequence
	currentVersion, err := queryClient.DelegateCurrentVersion(ctx.Context(), &types.QueryCurrentDelegateVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(currentVersion.Version, versionInfo.CurrentVersion)
}

func (suite *KeeperTestSuite) TestQueryUndelegateVersion() {
	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	// query with invalid zone
	_, err := queryClient.UndelegateVersion(ctx.Context(), &types.QueryUndelegateVersion{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//version info is nil
	exp := types.QueryUndelegateVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 0,
		},
	}

	res, err := queryClient.UndelegateVersion(ctx.Context(), &types.QueryUndelegateVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	//sequence is 20
	exp = types.QueryUndelegateVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 20,
			Height:  1,
			State:   types.IcaPending,
		},
	}

	//set delegate version
	versionInfo := types.VersionState{
		ZoneId:         zoneId,
		CurrentVersion: 20,
		Record: map[uint64]*types.IBCTrace{
			20: {
				Version: 20,
				Height:  uint64(ctx.BlockHeight()),
				State:   types.IcaPending,
			},
		},
	}

	suite.App.GalKeeper.SetUndelegateVersion(ctx, zoneId, versionInfo)
	res, err = queryClient.UndelegateVersion(ctx.Context(), &types.QueryUndelegateVersion{
		ZoneId:  zoneId,
		Version: 20,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	currentVersion, err := queryClient.UndelegateCurrentVersion(ctx.Context(), &types.QueryCurrentUndelegateVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(currentVersion.Version, versionInfo.CurrentVersion)
}

func (suite *KeeperTestSuite) TestQueryWithdrawVersion() {
	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	// query with invalid zone
	_, err := queryClient.WithdrawVersion(ctx.Context(), &types.QueryWithdrawVersion{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//version info is nil
	exp := types.QueryWithdrawVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 0,
		},
	}

	res, err := queryClient.WithdrawVersion(ctx.Context(), &types.QueryWithdrawVersion{
		ZoneId:  zoneId,
		Version: 0,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	//sequence is 20
	exp = types.QueryWithdrawVersionResponse{
		VersionInfo: &types.IBCTrace{
			Version: 20,
			Height:  1,
			State:   types.IcaPending,
		},
	}

	//set delegate version
	versionInfo := types.VersionState{
		ZoneId:         zoneId,
		CurrentVersion: 20,
		Record: map[uint64]*types.IBCTrace{
			20: {
				Version: 20,
				Height:  uint64(ctx.BlockHeight()),
				State:   types.IcaPending,
			},
		},
	}

	suite.App.GalKeeper.SetWithdrawVersion(ctx, zoneId, versionInfo)
	res, err = queryClient.WithdrawVersion(ctx.Context(), &types.QueryWithdrawVersion{
		ZoneId:  zoneId,
		Version: 20,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(res, &exp)

	// get current sequence
	currentVersion, err := queryClient.WithdrawCurrentVersion(ctx.Context(), &types.QueryCurrentWithdrawVersion{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(currentVersion.Version, versionInfo.CurrentVersion)
}

func (suite *KeeperTestSuite) TestQueryTotalSnAssetSupply() {
	//invalid zone
	queryClient := suite.queryClient
	ctx := suite.Ctx

	// query with invalid zone
	_, err := queryClient.TotalSnAssetSupply(ctx.Context(), &types.QueryTotalSnAssetSupply{
		ZoneId: "invalid",
	})
	suite.Require().Error(err)

	//total snAsset supply is zero
	res, err := queryClient.TotalSnAssetSupply(ctx.Context(), &types.QueryTotalSnAssetSupply{
		ZoneId: zoneId,
	})

	suite.Require().NoError(err)
	suite.Require().True(res.Amount.IsZero())

	suite.App.BankKeeper.MintCoins(suite.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(baseSnDenom, sdk.NewInt(10000))))
	res, err = queryClient.TotalSnAssetSupply(ctx.Context(), &types.QueryTotalSnAssetSupply{
		ZoneId: zoneId,
	})
	suite.Require().NoError(err)
	fmt.Println(res)
	suite.Require().Equal(res.Amount, sdk.NewCoin(baseSnDenom, sdk.NewInt(10000)))

}
