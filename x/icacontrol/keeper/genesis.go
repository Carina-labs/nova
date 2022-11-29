package keeper

import (
	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, controllerInfo := range genState.ControllerAddressInfo {
		k.SetControllerAddr(ctx, controllerInfo.ZoneId, controllerInfo.ControllerAddress)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params:                k.GetParams(ctx),
		ControllerAddressInfo: []*types.ControllerAddressInfo{},
	}

	var controllerAddr []*types.ControllerAddressInfo
	k.IterateRegisteredZones(ctx, func(index int64, zoneInfo types.RegisteredZone) (stop bool) {
		res := k.GetControllerAddr(ctx, zoneInfo.ZoneId)
		controllerAddr = append(controllerAddr, &res)
		return false
	})

	result.ControllerAddressInfo = controllerAddr

	return result
}
