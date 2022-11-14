package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrNoDepositRecord      = errors.Register(ModuleName, 0, "no deposit history for this account")
	ErrNoUndelegateRecord   = errors.Register(ModuleName, 1, "no undelegate history for this account")
	ErrNoWithdrawRecord     = errors.Register(ModuleName, 2, "no withdraw history for this account")
	ErrNotFoundZoneInfo     = errors.Register(ModuleName, 3, "not found zone info")
	ErrCanNotWithdrawAsset  = errors.Register(ModuleName, 4, "cannot withdraw funds")
	ErrInvalidTime          = errors.Register(ModuleName, 5, "time is not zero")
	ErrCanNotChangeState    = errors.Register(ModuleName, 6, "cannot change state")
	ErrDelegateFail         = errors.Register(ModuleName, 7, "delegate fail")
	ErrInvalidDenom         = errors.Register(ModuleName, 8, "invalid denom")
	ErrInvalidAddress       = errors.Register(ModuleName, 9, "invalid address")
	ErrUnknown              = errors.Register(ModuleName, 10, "unknown error occurred")
	ErrTransferInfoNotFound = errors.Register(ModuleName, 11, "transfer info is not found")
	ErrConvertWAssetIsZero  = errors.Register(ModuleName, 12, "withdrawal amount is zero")
	ErrNoDeleteRecord       = errors.Register(ModuleName, 13, "fail delete deposit record")
	ErrMaxUndelegateEntries = errors.Register(ModuleName, 14, "too many undelegate request")
	ErrNegativeVersion      = errors.Register(ModuleName, 15, "every version must be positive")
	ErrMaxDepositEntries    = errors.Register(ModuleName, 16, "too many deposit request")
	ErrNoDelegateRecord     = errors.Register(ModuleName, 17, "no delegate history for this account")
	ErrInsufficientFunds    = errors.Register(ModuleName, 18, "insufficient funds")
	ErrInvalidAck           = errors.Register(ModuleName, 19, "ack is not receive")
)
