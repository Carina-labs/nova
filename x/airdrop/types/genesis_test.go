package types_test

import (
	"testing"
	"time"

	"github.com/Carina-labs/nova/v2/app/params"
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	maxCap = sdk.NewInt(10000_00000)
)

func validAirdrop() *types.AirdropInfo {
	return &types.AirdropInfo{
		SnapshotTimestamp:        time.Now(),
		AirdropStartTimestamp:    time.Now().Add(time.Hour * 24 * 7),
		AirdropEndTimestamp:      time.Now().Add(time.Hour * 24 * 31),
		AirdropDenom:             params.BaseCoinUnit,
		QuestsCount:              5,
		ControllerAddress:        sdk.AccAddress([]byte{0x1}).String(),
		MaximumTokenAllocPerUser: maxCap.String(),
	}
}

func TestGenesis_ValidateAirdropInfo(t *testing.T) {
	invalidAddr := "foobar"
	negativeCoin := sdk.NewInt(-1000)

	// check airdrop cannot be nil
	gs := types.NewGenesisState(nil, nil)
	require.Error(t, gs.Validate())

	// check airdrop time condition didn't met
	gs = types.NewGenesisState(nil, validAirdrop())
	gs.AirdropInfo.AirdropStartTimestamp = gs.AirdropInfo.AirdropEndTimestamp.Add(time.Second)
	require.Error(t, gs.Validate())

	gs = types.NewGenesisState(nil, validAirdrop())
	gs.AirdropInfo.SnapshotTimestamp = gs.AirdropInfo.AirdropStartTimestamp.Add(time.Second)
	require.Error(t, gs.Validate())

	// check invalid controller address
	gs = types.NewGenesisState(nil, validAirdrop())
	gs.AirdropInfo.ControllerAddress = invalidAddr
	require.Error(t, gs.Validate())

	// check invalid max cap
	gs = types.NewGenesisState(nil, validAirdrop())
	gs.AirdropInfo.MaximumTokenAllocPerUser = negativeCoin.String()
	require.Error(t, gs.Validate())

	gs = types.NewGenesisState(nil, validAirdrop())
	require.NoError(t, gs.Validate())
}

func TestGenesis_ValidateState(t *testing.T) {
	invalidAddr := "foobar"

	validState := func() *types.UserState {
		return &types.UserState{
			Recipient:   sdk.AccAddress{0x1}.String(),
			TotalAmount: maxCap.String(),
			QuestStates: types.EmptyQuestState(time.Time{}),
		}
	}

	gs := types.NewGenesisState([]*types.UserState{validState()}, validAirdrop())
	require.NoError(t, gs.Validate())

	// check recipient is invalid address
	state := validState()
	state.Recipient = invalidAddr
	gs = types.NewGenesisState([]*types.UserState{state}, validAirdrop())
	require.Error(t, gs.Validate())

	// check total amount is negative
	state = validState()
	state.TotalAmount = sdk.NewInt(-1000).String()
	gs = types.NewGenesisState([]*types.UserState{state}, validAirdrop())
	require.Error(t, gs.Validate())

	// check total amount is greater than max cap
	state = validState()
	state.TotalAmount = maxCap.Add(sdk.NewInt(1)).String()
	gs = types.NewGenesisState([]*types.UserState{state}, validAirdrop())
	require.Error(t, gs.Validate())
}
