package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Carina-labs/novachain/testutil/keeper"
	"github.com/Carina-labs/novachain/x/novachain/keeper"
	"github.com/Carina-labs/novachain/x/novachain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NovachainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
