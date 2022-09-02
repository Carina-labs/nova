package keeper

import (
	context "context"

	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.QueryServer = &Querier{}

type Querier struct {
	keeper Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{keeper: k}
}

// AirdropInfo returns the airdrop info
func (q Querier) AirdropInfo(goCtx context.Context, _ *types.QueryAirdropInfoRequest) (*types.QueryAirdropInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	info := q.keeper.GetAirdropInfo(ctx)
	return &types.QueryAirdropInfoResponse{AirdropInfo: &info}, nil
}

// TODO: rename this method to TotalAllocatedAirropTokens
func (q Querier) TotalAssetForAirdrop(goCtx context.Context, request *types.QueryTotalAssetForAirdropRequest) (*types.QueryTotalAssetForAirdropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddr, err := sdk.AccAddressFromBech32(request.Address)
	info := q.keeper.GetAirdropInfo(ctx)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	if !q.keeper.IsEligible(ctx, userAddr) {
		return &types.QueryTotalAssetForAirdropResponse{TotalAssets: sdk.NewCoin(info.AirdropDenom, sdk.ZeroInt())}, nil
	}

	state, err := q.keeper.GetAirdropState(ctx, userAddr)
	if err != nil {
		q.keeper.Logger(ctx).Error("cannot get airdrop state, this error must never happen", "err", err)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "err: %v", err)
	}

	amt, ok := sdk.NewIntFromString(state.TotalAmount)
	if !ok {
		q.keeper.Logger(ctx).Error("cannot parse total amount, this error must never happen", "err", err)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "err: %v", err)
	}
	asset := sdk.NewCoin(info.AirdropDenom, amt)

	return &types.QueryTotalAssetForAirdropResponse{TotalAssets: asset}, nil
}

// QuestState returns state of the quest for the given user
func (q Querier) QuestState(goCtx context.Context, request *types.QueryQuestStateRequest) (*types.QueryQuestStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	if !q.keeper.IsEligible(ctx, userAddr) {
		return nil, sdkerrors.ErrUnauthorized
	}

	state, err := q.keeper.GetAirdropState(ctx, userAddr)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "err: %v", err)
	}

	return &types.QueryQuestStateResponse{QuestStates: state.QuestStates}, nil
}
