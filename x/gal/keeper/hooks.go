package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"strconv"
)

// Hooks wrapper struct for gal keeper
type Hooks struct {
	k Keeper
}

var _ types.TransferHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterTransferEnd(ctx sdk.Context, data types.FungibleTokenPacketData, base_denom string) {
	stAsset := h.k.interTxKeeper.GetstDenomForBaseDenom(ctx, base_denom)

	senderAddr, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}

	// Mint share tokens
	totalSharedToken := h.k.bankKeeper.GetSupply(ctx, h.k.getPairSnToken(ctx, base_denom))
	userDepositToken, err := strconv.ParseInt(data.Amount, 10, 64)
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}

	// alpha = user_deposit_amount / total_staked_amount
	alpha, err := h.k.calculateAlpha(ctx, base_denom, int(userDepositToken))
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}

	// minted_amount = alpha * total_share_token_supply
	err = h.k.MintShareTokens(ctx, senderAddr, sdk.Coins{sdk.Coin{
		Denom:  stAsset,
		Amount: sdk.NewInt(int64(alpha * float64(totalSharedToken.Amount.Int64()))),
	}})
	if err != nil {
		h.k.Logger(ctx).Error(err.Error())
	}
}
