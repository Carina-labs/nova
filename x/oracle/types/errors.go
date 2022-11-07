package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNoSupportChain      = sdkerrors.Register(ModuleName, 0, "this chain is not supported")
	ErrInvalidOperator     = sdkerrors.Register(ModuleName, 1, "invalid operator address")
	ErrUnknown             = sdkerrors.Register(ModuleName, 2, "unknown error")
	ErrNotFoundZoneInfo    = sdkerrors.Register(ModuleName, 3, "not found zone info")
	ErrInvalidKeyManager   = sdkerrors.Register(ModuleName, 4, "invalid key manager address")
	ErrNegativeBlockHeight = sdkerrors.Register(ModuleName, 5, "blockHeight must be positive")
	ErrInvalidBlockHeight  = sdkerrors.Register(ModuleName, 6, "current block height must be higher than the previous block height.")
)
