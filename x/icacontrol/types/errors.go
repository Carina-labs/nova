package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 2, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 3, "interchain account not exist")
	ErrMsgNotFound            = sdkerrors.Register(ModuleName, 4, "message not found ")
	ErrMsgNotNil              = sdkerrors.Register(ModuleName, 5, "message is not nil")
	ErrNotFoundZoneInfo       = sdkerrors.Register(ModuleName, 6, "registered zone not found with given name")
	ErrZoneIdNotNil           = sdkerrors.Register(ModuleName, 7, "zone Id is not nil")
	ErrNotFoundHostAddr       = sdkerrors.Register(ModuleName, 8, "host address is not found")
)
