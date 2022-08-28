package types

const (
	WithdrawStatusRegistered WithdrawStatusType = iota + 1
	WithdrawStatusTransferred
)

const (
	DepositRequest DepositState = iota + 1
	DepositSuccess
	DelegateRequest
	DelegateSuccess
)

const (
	UndelegateRequestUser UndelegatedState = iota + 1
	UndelegateRequestIca
)

type (
	WithdrawStatusType = int64
	DepositState       = int64
	UndelegatedState   = int64
)
