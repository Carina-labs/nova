package keeper

import (
	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterZoneProposal stores metadata for the new zone proposal.
func (k Keeper) RegisterZoneProposal(ctx sdk.Context, proposal *types.ZoneRegisterProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyZoneProposal)
	bz := k.cdc.MustMarshal(proposal)
	store.Set([]byte(proposal.Zone.ZoneId), bz)
}

// GetRegisterZoneProposal gets information about the stored zone that fits the zoneId.
func (k Keeper) GetRegisterZoneProposal(ctx sdk.Context, zoneId string) (types.ZoneRegisterProposal, bool) {
	zone := types.ZoneRegisterProposal{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyZoneProposal)
	bz := store.Get([]byte(zoneId))

	if len(bz) == 0 {
		return zone, false
	}

	k.cdc.MustUnmarshal(bz, &zone)
	return zone, true
}

//set prefix store
func (k Keeper) HandleZoneRegisterProposal(ctx sdk.Context, proposal *types.ZoneRegisterProposal) error {
	k.RegisterZoneProposal(ctx, proposal)
	return nil
}
