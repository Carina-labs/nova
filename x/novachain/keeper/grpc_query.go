package keeper

import (
	"github.com/Carina-labs/novachain/x/novachain/types"
)

var _ types.QueryServer = Keeper{}
