package keeper

import (
	"github.com/Carina-labs/nova/v2/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the auth module's parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams gets the auth module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}
