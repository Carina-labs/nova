package keeper_test

import (
	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestExportGenesis() {
	depositor := suite.GenRandomAddress()
	claimer := suite.GenRandomAddress()
	withdrawer := suite.GenRandomAddress()

	genesis := &types.GenesisState{
		Params: types.Params{},
		DepositRecord: []*types.DepositRecord{
			{
				ZoneId:    zoneId,
				Depositor: depositor.String(),
				Records: []*types.DepositRecordContent{
					{
						Claimer: claimer.String(),
						Amount:  &sdk.Coin{Amount: sdk.NewInt(1000000), Denom: baseDenom},
						State:   types.DepositSuccess,
					},
					{
						Claimer: claimer.String(),
						Amount:  &sdk.Coin{Amount: sdk.NewInt(2000000), Denom: baseDenom},
						State:   types.DepositSuccess,
					},
				},
			},
		},
		DelegateRecord: []*types.DelegateRecord{
			{
				ZoneId:  zoneId,
				Claimer: claimer.String(),
				Records: map[uint64]*types.DelegateRecordContent{
					0: {
						Amount:        &sdk.Coin{Amount: sdk.NewInt(1000000), Denom: baseDenom},
						State:         types.DelegateSuccess,
						OracleVersion: 1,
					},
				},
			},
		},
		UndelegateRecord: []*types.UndelegateRecord{
			{
				ZoneId:    zoneId,
				Delegator: depositor.String(),
				Records: []*types.UndelegateRecordContent{
					{
						Withdrawer:        withdrawer.String(),
						SnAssetAmount:     &sdk.Coin{Amount: sdk.NewInt(1000000), Denom: baseSnDenom},
						WithdrawAmount:    sdk.NewInt(1000005),
						State:             types.WithdrawStatusTransferred,
						OracleVersion:     1,
						UndelegateVersion: 1,
					},
				},
			},
		},
		WithdrawRecord: []*types.WithdrawRecord{},
		RecordInfo:     []*types.RecordInfo{},
	}

	suite.App.GalKeeper.InitGenesis(suite.Ctx, genesis)

	exported := suite.App.GalKeeper.ExportGenesis(suite.Ctx)
	suite.Require().Equal(genesis, exported)
}
