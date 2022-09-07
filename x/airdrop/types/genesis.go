package types

import (
	"fmt"
	"time"

	"github.com/Carina-labs/nova/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewGenesisState(state []*UserState, airdropInfo *AirdropInfo) *GenesisState {
	return &GenesisState{
		States:      state,
		AirdropInfo: airdropInfo,
	}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		States: nil,
		AirdropInfo: &AirdropInfo{
			SnapshotTimestamp:        time.Now().UTC(),
			AirdropStartTimestamp:    time.Now().UTC().Add(time.Hour * 24 * 7),
			AirdropEndTimestamp:      time.Now().UTC().Add(time.Hour * 24 * 31),
			AirdropDenom:             params.BaseCoinUnit,
			QuestsCount:              5,
			ControllerAddress:        sdk.AccAddress([]byte{0x1}).String(),
			MaximumTokenAllocPerUser: sdk.NewInt(10000_00000).String(),
		},
	}
}

func (gs GenesisState) Validate() error {
	if gs.AirdropInfo == nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "airdrop info is required")
	}

	maxTokenAlloc, ok := sdk.NewIntFromString(gs.AirdropInfo.MaximumTokenAllocPerUser)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("maximum token alloc %s is invalid", gs.AirdropInfo.MaximumTokenAllocPerUser))
	}

	if maxTokenAlloc.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("maximum token alloc should be posivie: %v", gs.AirdropInfo.MaximumTokenAllocPerUser))
	}

	for _, state := range gs.States {
		tokenAlloc, ok := sdk.NewIntFromString(state.TotalAmount)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("token amount %s is invalid", state.TotalAmount))
		}

		if tokenAlloc.IsNegative() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("token amount should be posivie: %v", state.TotalAmount))
		}

		if tokenAlloc.GT(maxTokenAlloc) {
			return sdkerrors.Wrap(ErrTokenAllocCannotExceedMaxCap, "airdrop token allocation on each user must be less than or equal to maxTokenAllocPerUser")
		}

		// check recipient address is valid
		_, err := sdk.AccAddressFromBech32(state.Recipient)
		if err != nil {
			return err
		}
	}

	if !gs.AirdropInfo.SnapshotTimestamp.Before(gs.AirdropInfo.AirdropStartTimestamp) {
		return sdkerrors.Wrap(ErrTimeConditionNotMet, "snapshot date must be before airdrop start date")
	}

	if !gs.AirdropInfo.AirdropStartTimestamp.Before(gs.AirdropInfo.AirdropEndTimestamp) {
		return sdkerrors.Wrap(ErrTimeConditionNotMet, "airdrop start date must be before airdrop end date")
	}

	if _, err := sdk.AccAddressFromBech32(gs.AirdropInfo.ControllerAddress); err != nil {
		return err
	}

	return nil
}

// EmptyQuestState returns initial state of quest.
func EmptyQuestState(blockTime time.Time) map[int32]*QuestState {
	defaultQuestState := func() *QuestState {
		return &QuestState{
			State:         QuestStateType_QUEST_STATE_NOT_STARTED,
			AchievedAt:    time.Time{},
			ClaimedAt:     time.Time{},
			ClaimedAmount: "", // claimed amount should be empty string
		}
	}

	result := map[int32]*QuestState{
		int32(QuestType_QUEST_NOTHING_TO_DO):     defaultQuestState(),
		int32(QuestType_QUEST_SOCIAL):            defaultQuestState(),
		int32(QuestType_QUEST_PROVIDE_LIQUIDITY): defaultQuestState(),
		int32(QuestType_QUEST_SN_ASSET_CLAIM):    defaultQuestState(),
		int32(QuestType_QUEST_VOTE_ON_PROPOSALS): defaultQuestState(),
	}

	// All user starts with the empty quests except for one quest.
	// This quest is the one that allows the user to claim 20% of the airdrop.
	result[int32(QuestType_QUEST_NOTHING_TO_DO)].State = QuestStateType_QUEST_STATE_CLAIMABLE
	result[int32(QuestType_QUEST_NOTHING_TO_DO)].AchievedAt = blockTime

	return result
}
