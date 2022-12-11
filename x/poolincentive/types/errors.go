package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrPoolNotFound = errors.Register(ModuleName, 1, "pool does not exist")
)
