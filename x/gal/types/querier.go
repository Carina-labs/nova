package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryShares      = "shares"
	QueryAllShares   = "all_shares"
	QueryTotalShares = "total_shares"
)

func NewQuerySharesRequest(addr sdk.AccAddress, denom string) *QuerySharesRequest {
	return &QuerySharesRequest{
		Address: addr.String(),
		Denom:   denom,
	}
}
