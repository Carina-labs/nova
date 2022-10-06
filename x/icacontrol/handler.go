package icacontrol

import (
	"github.com/Carina-labs/nova/x/icacontrol/keeper"
	"github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func NewRegisterZoneProposalHandelr(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ZoneRegisterProposal:
			return k.HandleZoneRegisterProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized register zone proposal content type: %T", c)
		}
	}
}
