package keeper

import (
	"context"
	"github.com/Carina-labs/novachain/x/gal/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func (k msgServer) Deposit(ctx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the intertx Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
