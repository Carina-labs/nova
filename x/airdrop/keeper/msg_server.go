package keeper

import (
	context "context"
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (m msgServer) ClaimAirdrop(goCtx context.Context, request *types.MsgClaimAirdropRequest) (*types.MsgClaimAirdropResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) MarkSocialQuestPerformed(goCtx context.Context, request *types.MsgMarkSocialQuestPerformedRequest) (*types.MsgMarkSocialQuestPerformedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := request.GetSigners()[0]

	if !m.keeper.ValidAirdropDate(ctx) {
		ctx.Logger().Debug("airdrop was over")
		return nil, types.ErrAirdropWasOver
	}

	if !m.keeper.IsValidControllerAddr(ctx, signer) {
		ctx.Logger().Debug("invalid controller address", "addr", signer)
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, userAddr := range request.UserAddresses {
		addr, err := sdk.AccAddressFromBech32(userAddr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "addr: %v", addr)
		}

		err = m.keeper.MarkUserPerformedQuest(ctx, addr, types.QuestType_QUEST_SOCIAL)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "addr: %v, err: %v", addr, err)
		}
	}

	return &types.MsgMarkSocialQuestPerformedResponse{
		Succeeded: true, // TODO: delete this field
	}, nil
}

func (m msgServer) MarkUserProvidedLiquidity(goCtx context.Context, request *types.MsgMarkUserProvidedLiquidityRequest) (*types.MsgMarkUserProvidedLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := request.GetSigners()[0]

	if !m.keeper.ValidAirdropDate(ctx) {
		ctx.Logger().Debug("airdrop was over")
		return nil, types.ErrAirdropWasOver
	}

	if !m.keeper.IsValidControllerAddr(ctx, signer) {
		ctx.Logger().Debug("invalid controller address", "addr", signer)
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, userAddr := range request.UserAddresses {
		addr, err := sdk.AccAddressFromBech32(userAddr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "addr: %v", addr)
		}

		err = m.keeper.MarkUserPerformedQuest(ctx, addr, types.QuestType_QUEST_PROVIDE_LIQUIDITY)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "addr: %v, err: %v", addr, err)
		}
	}

	return &types.MsgMarkUserProvidedLiquidityResponse{
		Succeeded: true, // TODO: delete this field
	}, nil
}
