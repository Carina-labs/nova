package keeper

import (
	"context"

	"github.com/Carina-labs/novachain/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

//delegate wAsset
func (k Keeper) DepositCoin(ctx sdk.Context,
	depositor string,
	receiver string,
	sourcePort string,
	sourceChannel string,
	amt sdk.Coins) error {
	// wAtom -> [ GAL ] -> snAtom
	for _, coin := range amt {
		goCtx := sdk.WrapSDKContext(ctx)

		_, err := k.ibcTransferKeeper.Transfer(goCtx,
			&transfertypes.MsgTransfer{
				SourcePort:    sourcePort,
				SourceChannel: sourceChannel,
				Token:         coin,
				Sender:        depositor,
				Receiver:      receiver,
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
	}

	return nil
}

//Redelegate wAsset

//stAsset mint
func (k Keeper) MintStTokenAndDistribute(ctx sdk.Context, depositor string, amt sdk.Coins) error {
	depositorAddr, err := sdk.AccAddressFromBech32(depositor)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, amt); err != nil {
		return err
	}

	return nil
}

func (k Keeper) CalculateShares(ctx sdk.Context, depositor string, coin sdk.Coin) (float64, error) {
	totalSupply := k.bankKeeper.GetSupply(ctx, coin.Denom)
	depositorBalance, err := k.bankKeeper.Balance(ctx.Context(), &banktypes.QueryBalanceRequest{
		Address: depositor,
		Denom:   coin.Denom,
	})

	if err != nil {
		return 0, err
	}

	shares := float64(depositorBalance.Balance.Amount.Int64()/totalSupply.Amount.Int64()) * 100
	return shares, nil
}

func (k Keeper) getPairSnToken(ctx sdk.Context, denom string) (stTokenDenom string) {
	k.paramSpace.Get(ctx, types.KeyWhiteListedTokenDenoms, &stTokenDenom)
	return
}

func (k Keeper) Share(context context.Context, rq *types.QuerySharesRequest) (*types.QuerySharesResponse, error) {
	return nil, nil
}
