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

	if err := m.keeper.UpdateChainState(goCtx, &types.ChainInfo{
		Coin:               state.Coin,
		OperatorAddress:    state.Operator,
		LastBlockHeight:    state.BlockHeight,
		Decimal:            state.Decimal,
	}); err != nil {
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

	return &types.QueryStateResponse{
		Coin:               result.Coin,
		Decimal:            result.Decimal,
		LastBlockHeight:    result.LastBlockHeight,
	}, nil
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
