package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/Carina-labs/nova/x/icacontrol/keeper"
)

// WeightedOperations returns all the operations from the module with their respective sights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, k keeper.Keeper,
) simulation.WeightedOperations {
	// TODO : implements this!
	return nil
}
