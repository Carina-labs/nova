package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryDepositHistory    = "depositHistory"
	QueryUndelegateHistory = "undelegateHistory"
)

func NewQuerySharesRequest(addr sdk.AccAddress, zoneId string) *QueryMyShareRequest {
	return &QueryMyShareRequest{
		Address: addr.String(),
		ZoneId:  zoneId,
	}
}

func NewDepositHistoryRequest(addr sdk.AccAddress, zoneId string) *QueryDepositHistoryRequest {
	return &QueryDepositHistoryRequest{
		Address: addr.String(),
		ZoneId:  zoneId,
	}
}

func NewUndelegateHistoryRequest(addr sdk.AccAddress, zoneId string) *QueryUndelegateHistoryRequest {
	return &QueryUndelegateHistoryRequest{
		Address: addr.String(),
		ZoneId:  zoneId,
	}
}

func NewWithdrawHistoryRequest(addr sdk.AccAddress, zoneId string) *QueryWithdrawHistoryRequest {
	return &QueryWithdrawHistoryRequest{
		Address: addr.String(),
		ZoneId:  zoneId,
	}
}
