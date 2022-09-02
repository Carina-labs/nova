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

// ClaimAirdrop mint airdrop tokens to the user
// Users must have performed the airdrop quest if not, return error
func (m msgServer) ClaimAirdrop(goCtx context.Context, request *types.MsgClaimAirdropRequest) (*types.MsgClaimAirdropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddr, err := sdk.AccAddressFromBech32(request.UserAddress)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	if !m.keeper.ValidClaimableDate(ctx) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "this is not claimable date")
	}

	if !m.keeper.IsEligible(ctx, userAddr) {
		return nil, types.ErrUserNotEligible
	}

	userState, err := m.keeper.GetAirdropState(ctx, userAddr)
	if err != nil {
		return nil, err
	}

	quest, ok := userState.QuestStates[int32(request.QuestType)]
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotSupported, "this quest is not supported: %v", request.QuestType)
	}

	// check user performed given quest
	if quest.State != types.QuestStateType_QUEST_STATE_CLAIMABLE {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "this quest is not claimable: %v, please perform this quest before claim request", request.QuestType)
	}

	// check user already claimed airdrop tokens for given quest
	if !quest.ClaimedAt.IsZero() || quest.ClaimedAmount != "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "you already claimed for this airdrop quest: %v", types.QuestType_name[int32(request.QuestType)])
	}

	// calculate claimable amount
	// TODO: Refactoring this codes
	info := m.keeper.GetAirdropInfo(ctx)
	getTotalClaimed := func() (claimedCnt int32, claimed sdk.Int) {
		claimedCnt = 0
		claimed = sdk.ZeroInt()

		for _, v := range userState.QuestStates {
			if !v.ClaimedAt.IsZero() {
				amt, ok := sdk.NewIntFromString(v.ClaimedAmount)
				if !ok {
					m.keeper.Logger(ctx).Error("failed to parse claimed amount, this is not intended..", "amount", v.ClaimedAmount)
					panic("invalid claimed amount")
				}

				claimedCnt += 1
				claimed = claimed.Add(amt)
			}
		}

		return claimedCnt, claimed
	}
	getClaimAmount := func() sdk.Int {
		cnt, amt := getTotalClaimed()
		total, ok := sdk.NewIntFromString(userState.TotalAmount)
		if !ok {
			panic("invalid total amount")
		}

		onlyOneRemained := cnt == info.QuestsCount-1
		if onlyOneRemained {
			// if only one quest remained, claim all remaining amount
			return total.Sub(amt)
		}

		// if not, claim fixed rate amount
		return total.QuoRaw(int64(info.QuestsCount))
	}

	// mint and send airdrop tokens to the user
	amount := getClaimAmount()
	airdropToken := sdk.NewCoin(info.AirdropDenom, amount)
	err = m.keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(airdropToken))
	if err != nil {
		m.keeper.Logger(ctx).Error("failed to mint token", "module", types.ModuleName, "token", airdropToken)
		return nil, err
	}

	err = m.keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddr, sdk.NewCoins(airdropToken))
	if err != nil {
		m.keeper.Logger(ctx).Error("failed to send token", "module", types.ModuleName, "token", airdropToken)
		return nil, err
	}

	quest.State = types.QuestStateType_QUEST_STATE_CLAIMED
	quest.ClaimedAt = ctx.BlockTime()
	if err = m.keeper.SetAirdropState(ctx, userAddr, userState); err != nil {
		m.keeper.Logger(ctx).Error("failed to mark user claimed airdrop", "user", userAddr, "quest", request.QuestType)
		return nil, err
	}

	return &types.MsgClaimAirdropResponse{}, nil
}

func (m msgServer) MarkSocialQuestPerformed(goCtx context.Context, request *types.MsgMarkSocialQuestPerformedRequest) (*types.MsgMarkSocialQuestPerformedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := request.GetSigners()[0]

	if !m.keeper.ValidQuestDate(ctx) {
		ctx.Logger().Debug("user cannot perform the quest anymore")
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

	if !m.keeper.ValidQuestDate(ctx) {
		ctx.Logger().Debug("user cannot perform the quest anymore")
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
