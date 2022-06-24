package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrNoDepositRecord = errors.Register(ModuleName, 0, "no deposit history for this address.")
)
