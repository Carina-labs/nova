package oracle

import (
	"github.com/Carina-labs/nova/v2/x/oracle/keeper"
	"github.com/Carina-labs/nova/v2/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "oracle" type messages.
func NewHandler(k *keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUpdateChainState:
			res, err := msgServer.UpdateChainState(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRegisterOracleAddr:
			res, err := msgServer.RegisterOracleAddress(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
