package keeper_test

import (
	galkeeper "github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	keeper.SetDepositAmt(ctx, &types.DepositRecord{
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

	suite.NoError(err)
	suite.Equal(amt.Amount.Denom, denom)
	suite.Equal(amt.Amount.Amount.Int64(), amount.Int64())
}
