package keeper

import (
	"context"
	"github.com/Carina-labs/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func (m msgServer) UpdateChainState(ctx context.Context, state *types.MsgUpdateChainState) (*types.MsgUpdateChainStateResponse, error) {
	goCtx := sdk.UnwrapSDKContext(ctx)

	err := m.keeper.UpdateChainState(goCtx, state)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateChainStateResponse{}, nil
}

func (m msgServer) GetChainState(ctx context.Context, request *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	goCtx := sdk.UnwrapSDKContext(ctx)

	result, err := m.keeper.GetChainState(goCtx, request.ChainDenom)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
