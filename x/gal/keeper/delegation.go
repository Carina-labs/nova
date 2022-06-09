package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

//delegate wAsset
func (k Keeper) DepositCoin(ctx sdk.Context,
	depositor sdk.AccAddress,
	receiver sdk.AccAddress,
	sourcePort string,
	sourceChannel string,
	amt sdk.Coin) error {
	goCtx := sdk.WrapSDKContext(ctx)

	// 1. IBC transfer
	_, err := k.ibcTransferKeeper.Transfer(goCtx,
		&transfertypes.MsgTransfer{
			SourcePort:    sourcePort,
			SourceChannel: sourceChannel,
			Token:         amt,
			Sender:        depositor.String(),
			Receiver:      receiver.String(),
			TimeoutHeight: ibcclienttypes.Height{
				RevisionHeight: 500000,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: 0,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

//stAsset mint
func (k Keeper) MintShareTokens(ctx sdk.Context,
	depositor sdk.Address,
	amt sdk.Coin) error {
	depositorAddr, err := sdk.AccAddressFromBech32(depositor.String())
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) CalculateShares(ctx sdk.Context,
	targetDenom string,
	coin sdk.Coin) (float64, error) {
	targetTotalSupply, err := k.oracleKeeper.GetChainState(ctx, targetDenom)
	if err != nil {
		return 0, err
	}

	amt := coin.Amount.Uint64()
	shares := amt/amt + targetTotalSupply.TotalStakedBalance

	return float64(shares * 100), nil
}

func (k Keeper) getPairSnToken(ctx sdk.Context, denom string) (string, error) {
	data := make(map[string]string)
	k.paramSpace.Get(ctx, types.KeyWhiteListedTokenDenoms, &data)
	if _, ok := data[denom]; !ok {
		return "", errors.New(fmt.Sprintf("%s is not in white list", denom))
	}

	return data[denom], nil
}

func (k Keeper) SetPairToken(ctx sdk.Context, denom string, shareTokenDenom string) {
	data := make(map[string]string)
	data[denom] = shareTokenDenom
	k.paramSpace.Set(ctx, types.KeyWhiteListedTokenDenoms, data)
}

func (k Keeper) Share(context context.Context, rq *types.QueryCacheDepositAmountRequest) (*types.QueryCachedDepositAmountResponse, error) {
	return nil, nil
}
