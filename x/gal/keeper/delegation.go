package keeper

import (
	"context"
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MintShareTokens mints st-token(share token) regard with deposited token.
func (k Keeper) MintShareTokens(ctx sdk.Context, depositor sdk.AccAddress, amt sdk.Coin) error {
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

// BurnShareTokens burns share token.
func (k Keeper) BurnShareTokens(ctx sdk.Context, burner sdk.Address, amt sdk.Coin) error {
	burnerAddr, err := sdk.AccAddressFromBech32(burner.String())
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, burnerAddr, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) Share(context context.Context, rq *types.QueryCacheDepositAmountRequest) (*types.QueryCachedDepositAmountResponse, error) {
	return nil, nil
}

func (k Keeper) DepositHistory(goCtx context.Context, rq *types.QueryDepositHistoryRequest) (*types.QueryDepositHistoryResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(sdkCtx, rq.Denom)
	if zoneInfo == nil {
		return nil, fmt.Errorf("can't find registered zone for denom : %s", rq.Denom)
	}

	acc, err := sdk.AccAddressFromBech32(rq.Address)
	if err != nil {
		return nil, err
	}

	dpInfo, err := k.GetRecordedDepositAmt(sdkCtx, acc)
	if err != nil {
		return nil, err
	}

	coins := sdk.Coins{}
	for _, r := range dpInfo.Records {
		coins.Add(*r.Amount)
	}

	return &types.QueryDepositHistoryResponse{
		Address: dpInfo.Address,
		Amount:  coins,
	}, nil
}
