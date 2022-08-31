package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidCoinAmount = errors.Register(ModuleName, 0, "invalid coin amount put into string")
)
