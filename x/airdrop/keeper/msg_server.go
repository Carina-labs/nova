package keeper

import (
	context "context"
	"github.com/Carina-labs/nova/x/airdrop/types"
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

func (m msgServer) ClaimAirdrop(ctx context.Context, request *types.MsgClaimAirdropRequest) (*types.MsgClaimAirdropResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) MarkSocialQuestPerformed(ctx context.Context, request *types.MsgMarkSocialQuestPerformedRequest) (*types.MsgMarkSocialQuestPerformedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) MarkUserProvidedLiquidity(ctx context.Context, request *types.MsgMarkUserProvidedLiquidityRequest) (*types.MsgMarkUserProvidedLiquidityResponse, error) {
	//TODO implement me
	panic("implement me")
}
