package poolincentive

import (
	"github.com/Carina-labs/nova/x/poolincentive/keeper"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func NewPoolIncentivesProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ReplacePoolIncentivesProposal:
			return k.HandleReplacePoolIncentivesProposal(ctx, c)
		case *types.UpdatePoolIncentivesProposal:
			return k.HandleUpdatePoolIncentivesProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized pool incentives proposal content type: %T", c)
		}
	}
}
