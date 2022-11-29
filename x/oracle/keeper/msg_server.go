package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Carina-labs/nova/v2/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

func (server msgServer) UpdateChainState(goctx context.Context, state *types.MsgUpdateChainState) (*types.MsgUpdateChainStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if !server.keeper.IsValidOracleAddress(ctx, state.ZoneId, state.Operator) {
		return nil, types.ErrInvalidOperator
	}

	ok := server.keeper.ValidZoneAddress(ctx, state.ZoneId)
	if !ok {
		return nil, types.ErrNotFoundZoneInfo
	}

	oracleState, _ := server.keeper.GetChainState(ctx, state.Coin.Denom)
	if oracleState != nil && oracleState.LastBlockHeight >= state.BlockHeight {
		return nil, types.ErrInvalidBlockHeight
	}

	newOracleState := &types.ChainInfo{
		Coin:            state.Coin,
		OperatorAddress: state.Operator,
		LastBlockHeight: state.BlockHeight,
		AppHash:         state.AppHash,
		ZoneId:          state.ZoneId,
	}

	oracleVersion, _ := server.keeper.GetOracleVersion(ctx, state.ZoneId)

	trace := types.IBCTrace{
		Version: oracleVersion + 1,
		Height:  uint64(ctx.BlockHeight()),
	}
	server.keeper.SetOracleVersion(ctx, state.ZoneId, trace)

	if err := server.keeper.UpdateChainState(ctx, newOracleState); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "err: %v", err)
	}

	if err := ctx.EventManager().EmitTypedEvent(newOracleState); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknown, "err: %v", err)
	}

	return &types.MsgUpdateChainStateResponse{}, nil
}

func (server msgServer) RegisterOracleAddress(goctx context.Context, msg *types.MsgRegisterOracleAddr) (*types.MsgRegisterOracleAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	if !server.keeper.IsValidOracleKeyManager(ctx, msg.FromAddress) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidKeyManager, msg.OracleAddress)
	}

	oracleAddrInfo := server.keeper.GetOracleAddress(ctx, msg.ZoneId)
	oracleAddrInfo.OracleAddress = append(oracleAddrInfo.OracleAddress, msg.OracleAddress)

	server.keeper.SetOracleAddress(ctx, msg.ZoneId, oracleAddrInfo.OracleAddress)

	return &types.MsgRegisterOracleAddrResponse{}, nil
}
