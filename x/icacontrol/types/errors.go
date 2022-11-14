package types

import (
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = errors.Register(ModuleName, 2, "interchain account already registered")
	ErrIBCAccountNotExist     = errors.Register(ModuleName, 3, "interchain account not exist")
	ErrMsgNotFound            = errors.Register(ModuleName, 4, "message not found ")
	ErrMsgNotNil              = errors.Register(ModuleName, 5, "message is not nil")
	ErrNotFoundZoneInfo       = errors.Register(ModuleName, 6, "registered zone not found with given name")
	ErrZoneIdNotNil           = errors.Register(ModuleName, 7, "zone Id is not nil")
	ErrNotFoundHostAddr       = errors.Register(ModuleName, 8, "host address is not found")
	ErrDenomDuplicates        = errors.Register(ModuleName, 9, "denom is already exists")
	ErrInvalidPortId          = errors.Register(ModuleName, 10, "invalid port id")
	ErrInvalidConnId          = errors.Register(ModuleName, 11, "invalid connection id")
	ErrInvalidAck             = errors.Register(ModuleName, 12, "ack is not receive")
)
