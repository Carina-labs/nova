package types

import "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrTokenAllocCannotExceedMaxCap = errors.Register(ModuleName, 1, "token allocation cannot exceed max cap")
	ErrTimeConditionNotMet          = errors.Register(ModuleName, 2, "time condition not met")
	ErrAirdropWasOver               = errors.Register(ModuleName, 3, "airdrop was over")
	ErrUserNotEligible              = errors.Register(ModuleName, 4, "user not eligible for the airdrop")
)
