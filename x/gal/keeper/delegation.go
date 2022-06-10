package keeper

import (
	"context"
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

func (k Keeper) SetPairToken(ctx sdk.Context, denom string, shareTokenDenom string) {
	data := make(map[string]string)
	data[denom] = shareTokenDenom
	k.paramSpace.Set(ctx, types.KeyWhiteListedTokenDenoms, data)
}

func (k Keeper) Share(context context.Context, rq *types.QueryCacheDepositAmountRequest) (*types.QueryCachedDepositAmountResponse, error) {
	return nil, nil
}
