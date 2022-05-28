package keeper

import (
	"context"
	"github.com/Carina-labs/novachain/x/oracle/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func (m msgServer) UpdateChainState(ctx context.Context, state *types.MsgUpdateChainState) (*types.MsgUpdateChainStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
