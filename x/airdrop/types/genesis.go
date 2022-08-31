package types

import (
	"github.com/Carina-labs/nova/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return nil
}
