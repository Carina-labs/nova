package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

func (server msgServer) UpdateChainState(goctx context.Context, state *types.MsgUpdateChainState) (*types.MsgUpdateChainStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if !server.keeper.IsValidOperator(ctx, state.Operator) {
		return nil, types.ErrInvalidOperator
	}

	newOracleState := &types.ChainInfo{
		Coin:            state.Coin,
		OperatorAddress: state.Operator,
		LastBlockHeight: state.BlockHeight,
		AppHash:         state.AppHash,
		ChainId:         state.ChainId,
	}

	oracleVersion, _ := server.keeper.GetOracleVersion(ctx, state.ChainId)
	server.keeper.SetOracleVersion(ctx, state.ChainId, oracleVersion+1, uint64(ctx.BlockHeight()))

	if err := server.keeper.UpdateChainState(ctx, newOracleState); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "err: %v", err)
	}

	if err := ctx.EventManager().EmitTypedEvent(newOracleState); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "err: %v", err)
	}

	return &types.MsgUpdateChainStateResponse{}, nil
}
