package airdrop

import (
	"github.com/Carina-labs/nova/x/airdrop/keeper"
	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"time"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	airdropInfo := k.GetAirdropInfo(ctx)
	blockTime := ctx.BlockTime()

	if airdropInfo.AirdropEndTimestamp.Before(blockTime) {
		k.BurnToken(ctx)
	}
}
