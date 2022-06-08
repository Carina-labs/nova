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

	if err := state.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdateChainState(goCtx, &types.ChainInfo{
		ChainDenom:         state.ChainDenom,
		OperatorAddress:    state.Operator,
		LastBlockHeight:    state.BlockHeight,
		TotalStakedBalance: state.TotalStakedBalance,
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
		TotalStakedBalance: result.TotalStakedBalance,
		Decimal: result.Decimal,
		LastBlockHeight: result.LastBlockHeight,
	}, nil
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
