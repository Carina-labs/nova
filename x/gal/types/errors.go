package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrNoDepositRecord     = errors.Register(ModuleName, 0, "no deposit history for this account")
	ErrNoWithdrawRecord    = errors.Register(ModuleName, 1, "no withdraw history for this account")
	ErrCanNotReplaceRecord = errors.Register(ModuleName, 2, "cannot replace record")
	ErrInsufficientFunds   = errors.Register(ModuleName, 3, "cannot withdraw funds : insufficient fund")
	ErrNotFoundZoneInfo    = errors.Register(ModuleName, 4, "not found zone info")
	ErrCanNotWithdrawAsset = errors.Register(ModuleName, 5, "cannot withdraw funds")
	ErrInvalidTime         = errors.Register(ModuleName, 6, "time is not zero")
	ErrCanNotChangeState   = errors.Register(ModuleName, 7, "cannot change state")
	ErrDelegateFail        = errors.Register(ModuleName, 8, "delegate fail")
	ErrInvalidDenom        = errors.Register(ModuleName, 9, "invalid denom")
)
