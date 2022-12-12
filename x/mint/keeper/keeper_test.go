package keeper_test

import (
	pooltypes "github.com/Carina-labs/nova/x/poolincentive/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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
	coin := mintKeeper.GetProportions(fee, sdk.NewDecWithPrec(2, 1))
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
	coin = mintKeeper.GetProportions(fee, sdk.NewDecWithPrec(2, 1))
	suite.Equal(fees[0].Amount.Quo(sdk.NewInt(5)), coin.Amount)
}

func (suite *KeeperTestSuite) TestMintIncentives() {
	mintKeeper := suite.App.MintKeeper

	gaiaAddr := suite.GenRandomAddress()
	osmoAddr := suite.GenRandomAddress()
	junoAddr := suite.GenRandomAddress()

	params := suite.App.MintKeeper.GetParams(suite.Ctx)
	// At this time, there is no distr record, so the asset should be allocated to the community pool.
	mintCoin := sdk.NewCoin("nova", sdk.NewInt(100000))
	mintCoins := sdk.Coins{mintCoin}
	err := mintKeeper.MintCoins(suite.Ctx, mintCoins)
	suite.NoError(err)
	err = mintKeeper.DistributeMintedCoin(suite.Ctx, mintCoin)
	suite.NoError(err)

	distribution.BeginBlocker(suite.Ctx, abci.RequestBeginBlock{}, *suite.App.DistrKeeper)

	// pool does not exist
	feePool := suite.App.DistrKeeper.GetFeePool(suite.Ctx)
	feeCollector := suite.App.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)

	suite.Equal(
		mintCoins[0].Amount.ToDec().Mul(params.DistributionProportions.Staking).TruncateInt(),
		suite.App.BankKeeper.GetBalance(suite.Ctx, feeCollector, "nova").Amount)
	suite.Equal(
		mintCoins[0].Amount.ToDec().Mul(params.DistributionProportions.CommunityPool),
		feePool.CommunityPool.AmountOf("nova"))

	// set pool
	pools := []*pooltypes.IncentivePool{
		{
			PoolId:              "gaia",
			Weight:              5,
			PoolContractAddress: gaiaAddr.String(),
		},
		{
			PoolId:              "osmo",
			Weight:              3,
			PoolContractAddress: osmoAddr.String(),
		},
		{
			PoolId:              "juno",
			Weight:              2,
			PoolContractAddress: junoAddr.String(),
		},
	}
	for _, pool := range pools {
		suite.App.PoolKeeper.CreateIncentivePool(suite.Ctx, pool)
	}

	err = mintKeeper.MintCoins(suite.Ctx, mintCoins)
	suite.NoError(err)
	err = mintKeeper.DistributeMintedCoin(suite.Ctx, mintCoin)
	suite.NoError(err)

	gaiaBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, gaiaAddr, "nova")
	suite.Require().Equal(gaiaBalance.Amount.String(), "20000")
	osmoBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, osmoAddr, "nova")
	suite.Require().Equal(osmoBalance.Amount.String(), "12000")
	junoBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, junoAddr, "nova")
	suite.Require().Equal(junoBalance.Amount.String(), "8000")
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}
