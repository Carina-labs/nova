package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/suite"

	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/cosmos/cosmos-sdk/x/distribution"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestMintCoinsToFeeCollectorAndGetProportions() {
	mintKeeper := suite.App.MintKeeper

	// When coin is minted to the fee collector
	fee := sdk.NewCoin("nova", sdk.NewInt(0))
	coin := mintKeeper.GetProportions(suite.Ctx, fee, sdk.NewDecWithPrec(2, 1))
	suite.Equal("0nova", coin.String())

	// When mint the 100K stake coin to the fee collector
	fee = sdk.NewCoin("nova", sdk.NewInt(100000))
	fees := sdk.NewCoins(fee)

	err := simapp.FundModuleAccount(suite.App.BankKeeper,
		suite.Ctx,
		authtypes.FeeCollectorName,
		fees)
	suite.NoError(err)

	// check proportion for 20%
	coin = mintKeeper.GetProportions(suite.Ctx, fee, sdk.NewDecWithPrec(2, 1))
	suite.Equal(fees[0].Amount.Quo(sdk.NewInt(5)), coin.Amount)
}

func (suite *KeeperTestSuite) TestMintIncentives() {
	mintKeeper := suite.App.MintKeeper

	params := suite.App.MintKeeper.GetParams(suite.Ctx)
	// At this time, there is no distr record, so the asset should be allocated to the community pool.
	mintCoin := sdk.NewCoin("nova", sdk.NewInt(100000))
	mintCoins := sdk.Coins{mintCoin}
	err := mintKeeper.MintCoins(suite.Ctx, mintCoins)
	suite.NoError(err)
	err = mintKeeper.DistributeMintedCoin(suite.Ctx, mintCoin)
	suite.NoError(err)

	distribution.BeginBlocker(suite.Ctx, abci.RequestBeginBlock{}, *suite.App.DistrKeeper)

	feePool := suite.App.DistrKeeper.GetFeePool(suite.Ctx)
	feeCollector := suite.App.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)

	suite.Equal(
		mintCoins[0].Amount.ToDec().Mul(params.DistributionProportions.Staking).TruncateInt(),
		suite.App.BankKeeper.GetBalance(suite.Ctx, feeCollector, "nova").Amount)
	suite.Equal(
		mintCoins[0].Amount.ToDec().Mul(params.DistributionProportions.CommunityPool),
		feePool.CommunityPool.AmountOf("nova"))
}
