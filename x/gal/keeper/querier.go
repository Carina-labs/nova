package keeper

import (
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier returns a new sdk.Keeper instance.
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryDepositHistory:
			return queryDepositHistory(ctx, req, k, legacyQuerierCdc)
		case types.QueryUndelegateHistory:
			return queryUndelegateHistory(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryDepositHistory(ctx sdk.Context,
	req abci.RequestQuery,
	k Keeper,
	legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDepositHistoryRequest

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	address, err := sdk.AccAddressFromBech32(params.Address)
	if err != nil {
		return nil, err
	}

	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, params.Denom)

	depositHistory, err := k.GetRecordedDepositAmt(ctx, zoneInfo.ZoneId, address)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, depositHistory)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryUndelegateHistory(ctx sdk.Context,
	req abci.RequestQuery,
	k Keeper,
	legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryUndelegateHistoryRequest

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	address, err := sdk.AccAddressFromBech32(params.Address)
	if err != nil {
		return nil, err
	}

	zoneInfo := k.ibcstakingKeeper.GetZoneForDenom(ctx, params.Denom)

	withdrawHistory, ok := k.GetUndelegateRecord(ctx, zoneInfo.ZoneId+address.String())
	if !ok {
		return nil, fmt.Errorf("can't find error for denom: %s", params.Denom)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, withdrawHistory)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
