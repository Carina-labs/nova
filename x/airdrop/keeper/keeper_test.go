package keeper_test

import (
	"testing"
	"time"

	"github.com/Carina-labs/nova/app/apptesting"
	"github.com/Carina-labs/nova/app/params"
	"github.com/Carina-labs/nova/x/airdrop/keeper"
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

// define common variables for tests
var (
	validUser            = sdk.AccAddress{0x1}
	invalidUser          = sdk.AccAddress{0x2}
	controllerUser       = sdk.AccAddress("controller_address__")
	maxTotalAllocPerUser = "100000000" // 100,000,000
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.AirdropKeeper)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func validUserState(ctx sdk.Context, recipient string, totalAsset string) *types.UserState {
	return &types.UserState{
		Recipient:   recipient,
		TotalAmount: totalAsset,
		QuestStates: types.EmptyQuestState(ctx.BlockTime()),
	}
}

func validAirdropInfo(ctx sdk.Context) *types.AirdropInfo {
	return &types.AirdropInfo{
		SnapshotTimestamp:        ctx.BlockTime(),
		AirdropStartTimestamp:    ctx.BlockTime().Add(-1 * time.Hour),
		AirdropEndTimestamp:      ctx.BlockTime().Add(time.Hour * 24 * 31),
		AirdropDenom:             params.BaseCoinUnit,
		QuestsCount:              5,
		ControllerAddress:        controllerUser.String(),
		MaximumTokenAllocPerUser: maxTotalAllocPerUser,
	}
}
