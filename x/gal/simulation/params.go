package simulation

import (
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

// ParamChanges defines the parameters that can be modified by param change proposals on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return nil
}
