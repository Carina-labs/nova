package keeper_test

import (
	"reflect"
	"sort"

	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestSuite) TestGenesis() {
	airdropKeeper := s.App.AirdropKeeper
	genesis := types.DefaultGenesis()
	genesis.States = []*types.UserState{}

	// generate 10 states
	for i := 0; i < 10; i++ {
		state := types.UserState{Recipient: sdk.AccAddress{0x1 + byte(i)}.String(), TotalAmount: "100000", QuestStates: nil}
		genesis.States = append(genesis.States, &state)
	}

	// run init genesis and check the data
	airdropKeeper.InitGenesis(s.Ctx, genesis)

	exported := airdropKeeper.ExportGenesis(s.Ctx)
	s.Require().Equal(genesis.AirdropInfo.MaximumTokenAllocPerUser, exported.AirdropInfo.MaximumTokenAllocPerUser)
	s.Require().Equal(genesis.AirdropInfo.AirdropDenom, exported.AirdropInfo.AirdropDenom)
	s.Require().Equal(genesis.AirdropInfo.ControllerAddress, exported.AirdropInfo.ControllerAddress)
	s.Require().Equal(genesis.AirdropInfo.QuestsCount, exported.AirdropInfo.QuestsCount)

	// sort genesis state because exported state is iterated as sorted.
	sort.Slice(genesis.States, func(i, j int) bool {
		return genesis.States[i].Recipient < genesis.States[j].Recipient
	})

	for i := 0; i < len(genesis.States); i++ {
		genesis.States[i].QuestStates = types.EmptyQuestState(s.Ctx.BlockTime())
		if !reflect.DeepEqual(genesis.States[i], exported.States[i]) {
			s.T().Error("initial state and exported genesis are not equal", genesis.States[i], exported.States[i])
		}
	}
}
