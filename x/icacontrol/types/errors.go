package types

import (
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = errors.Register(ModuleName, 2, "interchain account already registered")
	ErrIBCAccountNotExist     = errors.Register(ModuleName, 3, "interchain account not exist")
	ErrMsgNotFound            = errors.Register(ModuleName, 4, "message not found")
	ErrInvalidMsg             = errors.Register(ModuleName, 5, "invalid message")
	ErrNotFoundZone           = errors.Register(ModuleName, 6, "registered zone not found with given name")
	ErrInvalidZoneId          = errors.Register(ModuleName, 7, "zone Id cannot be nil")
	ErrNotFoundHostAddr       = errors.Register(ModuleName, 8, "host address is not found")
	ErrDenomDuplicates        = errors.Register(ModuleName, 9, "denom is already exists")
	ErrInvalidPortId          = errors.Register(ModuleName, 10, "invalid port id")
	ErrInvalidConnId          = errors.Register(ModuleName, 11, "invalid connection id")
	ErrInvalidAck             = errors.Register(ModuleName, 12, "ack is not receive")
	ErrInvalidIcaVersion      = errors.Register(ModuleName, 13, "invalid ica version")
	ErrAlreadyExistZone       = errors.Register(ModuleName, 14, "zone is already exist")
	ErrInvalidAddress         = errors.Register(ModuleName, 15, "invalid address")
	ErrInvalidMaxEntries      = errors.Register(ModuleName, 16, "invalid max entries")
	ErrInvalidDecimal         = errors.Register(ModuleName, 17, "invalid decimal")
	ErrInvalidChanId          = errors.Register(ModuleName, 19, "invalid channel id")
)
