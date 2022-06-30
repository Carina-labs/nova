package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrNoDepositRecord  = errors.Register(ModuleName, 0, "no deposit history for this account.")
	ErrNoWithdrawRecord = errors.Register(ModuleName, 1, "no withdraw history for this account.")
)
