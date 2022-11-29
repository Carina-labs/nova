package keeper

import (
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Carina-labs/nova/v2/x/mint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc                 codec.BinaryCodec
	storeKey            sdk.StoreKey
	paramSpace          paramtypes.Subspace
	distrKeeper         types.DistrKeeper
	stakingKeeper       types.StakingKeeper
	accountKeeper       types.AccountKeeper
	bankKeeper          types.BankKeeper
	PoolIncentiveKeeper types.PoolKeeper
	feeCollectorName    string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper, dk types.DistrKeeper, pk types.PoolKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                 cdc,
		storeKey:            key,
		paramSpace:          paramSpace,
		stakingKeeper:       sk,
		accountKeeper:       ak,
		bankKeeper:          bk,
		distrKeeper:         dk,
		PoolIncentiveKeeper: pk,
		feeCollectorName:    feeCollectorName,
	}
}

// CreateLpIncentiveModuleAccount creates the module account for developer vesting.
// Should only be called in initial genesis creation, never again.
func (k Keeper) CreateLpIncentiveModuleAccount(ctx sdk.Context, amount sdk.Coin) {
	moduleAcc := authtypes.NewEmptyModuleAccount(
		types.LpIncentiveModuleAccName, authtypes.Minter)

	k.accountKeeper.SetModuleAccount(ctx, moduleAcc)

	err := k.bankKeeper.MintCoins(ctx, types.LpIncentiveModuleAccName, sdk.NewCoins(amount))
	if err != nil {
		panic(err)
	}
}

// _____________________________________________________________________

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetMinter returns current minter.
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// SetMinter sets new minter.
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}

// GetProportions gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetProportions(mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}

// DistributeMintedCoin implements distribution of minted coins from mint to external modules.
func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {
	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	// allocate staking incentives into fee collector account to be moved to on next begin blocker by staking module
	stakingIncentivesCoins := sdk.NewCoins(k.GetProportions(mintedCoin, proportions.Staking))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingIncentivesCoins)
	if err != nil {
		return err
	}
	ctx.Logger().Info("Mint", "StakingIncentives", stakingIncentivesCoins)

	lpIncentivesCoin := k.GetProportions(mintedCoin, proportions.LpIncentives)
	ctx.Logger().Info("Mint", "LpIncentives", lpIncentivesCoin)
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.LpIncentiveModuleAccName, sdk.NewCoins(lpIncentivesCoin))
	if err != nil {
		return err
	}

	err = k.distributeLPIncentivePools(ctx, lpIncentivesCoin.Denom)
	if err != nil {
		return err
	}

	communityPoolCoins := sdk.NewCoins(mintedCoin).Sub(stakingIncentivesCoins).Sub(sdk.NewCoins(lpIncentivesCoin))
	ctx.Logger().Info("Mint", "CommunityPoolCoins", communityPoolCoins)
	err = k.distrKeeper.FundCommunityPool(ctx, communityPoolCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	return err
}

func (k Keeper) distributeLPIncentivePools(ctx sdk.Context, denom string) error {
	pools := k.PoolIncentiveKeeper.GetAllIncentivePool(ctx)
	if len(pools) == 0 {
		return nil
	}

	totalWeight := k.PoolIncentiveKeeper.GetTotalWeight(ctx)
	moduleAddr := k.accountKeeper.GetModuleAddress(types.LpIncentiveModuleAccName)
	lpIncentiveCoin := k.bankKeeper.GetBalance(ctx, moduleAddr, denom)

	for _, pool := range pools {
		poolWeight := sdk.NewIntFromUint64(pool.Weight).ToDec().Quo(sdk.NewIntFromUint64(totalWeight).ToDec())
		incentive := sdk.NewDecFromInt(lpIncentiveCoin.Amount).Mul(poolWeight)
		incentivesCoins := sdk.NewCoins(sdk.NewCoin(lpIncentiveCoin.Denom, incentive.TruncateInt()))
		poolAddr, err := sdk.AccAddressFromBech32(pool.PoolContractAddress)
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.LpIncentiveModuleAccName, poolAddr, incentivesCoins)
		if err != nil {
			return err
		}
	}
	return nil
}
