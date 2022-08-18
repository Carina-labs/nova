package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNoSupportChain  = sdkerrors.Register(ModuleName, 0, "this chain is not supported")
	ErrInvalidOperator = sdkerrors.Register(ModuleName, 1, "invalid operator address")
	ErrUnknown         = sdkerrors.Register(ModuleName, 2, "unknown error")
)
