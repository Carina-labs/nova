package keeper

import (
	"context"

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

	newOracleState := &types.ChainInfo{
		Coin:            state.Coin,
		OperatorAddress: state.Operator,
		LastBlockHeight: state.BlockHeight,
		Decimal:         state.Decimal,
		AppHash:         state.AppHash,
		ChainId:         state.ChainId,
	}

	if err := server.keeper.UpdateChainState(ctx, newOracleState); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(newOracleState); err != nil {
		return nil, err
	}

	return &types.MsgUpdateChainStateResponse{}, nil
}

func (server msgServer) GetChainState(goCtx context.Context, request *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result, err := server.keeper.GetChainState(ctx, request.ChainDenom)
	if err != nil {
		return nil, err
	}

	return &types.QueryStateResponse{
		Coin:            result.Coin,
		Decimal:         result.Decimal,
		LastBlockHeight: result.LastBlockHeight,
		AppHash:         result.AppHash,
		ChainId:         result.ChainId,
	}, nil
}
