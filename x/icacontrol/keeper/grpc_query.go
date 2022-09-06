package keeper

import (
	context "context"
	"github.com/Carina-labs/nova/x/icacontrol/types"
)

type QueryServer struct {
	types.QueryServer
	keeper Keeper
}

func NewQueryServer(keeper Keeper) *QueryServer {
	return &QueryServer{
		keeper: keeper,
	}
}

func (q *QueryServer) AllZones(ctx context.Context, request *types.QueryAllZonesRequest) (*types.QueryAllZonesResponse, error) {
	//TODO implement me
	panic("implement me")
}
