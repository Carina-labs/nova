package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryDepositHistory    = "depositHistory"
	QueryUndelegateHistory = "undelegateHistory"
)

func NewQuerySharesRequest(addr sdk.AccAddress, denom string) *QueryCacheDepositAmountRequest {
	return &QueryCacheDepositAmountRequest{
		Address: addr.String(),
		Denom:   denom,
	}
}

func NewDepositHistoryRequest(addr sdk.AccAddress, denom string) *QueryDepositHistoryRequest {
	return &QueryDepositHistoryRequest{
		Address: addr.String(),
		Denom:   denom,
	}
}

func NewUndelegateHistoryRequest(addr sdk.AccAddress, denom string) *QueryUndelegateHistoryRequest {
	return &QueryUndelegateHistoryRequest{
		Address: addr.String(),
		Denom:   denom,
	}
}

func NewWithdrawHistoryRequest(addr sdk.AccAddress, denom string) *QueryWithdrawHistoryRequest {
	return &QueryWithdrawHistoryRequest{
		Address: addr.String(),
		Denom:   denom,
	}
}
