package types

import (
	"fmt"
	"github.com/Carina-labs/nova/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func NewGenesisState(state []*AirdropState, airdropInfo *AirdropInfo) *GenesisState {
	return &GenesisState{
		States:      state,
		AirdropInfo: airdropInfo,
	}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		States: nil,
		AirdropInfo: &AirdropInfo{
			SnapshotTimestamp:        time.Now(),
			AirdropStartTimestamp:    time.Now().Add(time.Hour * 24 * 7),
			AirdropEndTimestamp:      time.Now().Add(time.Hour * 24 * 31),
			AirdropDenom:             params.BaseCoinUnit,
			QuestsCount:              5,
			ControllerAddress:        sdk.AccAddress([]byte{0x1}).String(),
			MaximumTokenAllocPerUser: sdk.NewInt(10000_00000).String(),
		},
	}
}

func (gs GenesisState) Validate() error {
	maxTokenAlloc, ok := sdk.NewIntFromString(gs.AirdropInfo.MaximumTokenAllocPerUser)
	if !ok {
		return sdkerrors.Wrap(ErrInvalidCoinAmount, fmt.Sprintf("maximum token alloc %s is invalid", gs.AirdropInfo.MaximumTokenAllocPerUser))
	}

	for _, state := range gs.States {
		tokenAlloc, ok := sdk.NewIntFromString(state.TotalAmount)
		if !ok {
			return sdkerrors.Wrap(ErrInvalidCoinAmount, fmt.Sprintf("token amount %s is invalid", state.TotalAmount))
		}

		if tokenAlloc.IsNegative() {
			return fmt.Errorf("token amount should be posivie: %v", state.TotalAmount)
		}

		if tokenAlloc.GT(maxTokenAlloc) {
			return fmt.Errorf("airdrop token allocation on each user must be less than or equal to maxTokenAllocPerUser")
		}

		// check recipient address is valid
		_, err := sdk.AccAddressFromBech32(state.Recipient)
		if err != nil {
			return err
		}
	}

	if !gs.AirdropInfo.SnapshotTimestamp.Before(gs.AirdropInfo.AirdropStartTimestamp) {
		return fmt.Errorf("snpashot date must be before airdrop start date")
	}

	if !gs.AirdropInfo.AirdropStartTimestamp.Before(gs.AirdropInfo.AirdropEndTimestamp) {
		return fmt.Errorf("airdrop start date must be before airdrop end date")
	}

	if _, err := sdk.AccAddressFromBech32(gs.AirdropInfo.ControllerAddress); err != nil {
		return err
	}

	return nil
}
