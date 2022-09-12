package types

const (
	WithdrawStatusRegistered WithdrawStatusType = iota + 1
	WithdrawStatusTransferred
)

const (
	DepositRequest DepositStatusType = iota + 1
	DepositSuccess
	DelegateRequest
	DelegateSuccess
)

const (
	UndelegateRequestByUser UndelegatedStatusType = iota + 1
	UndelegateRequestByIca
)

type (
	WithdrawStatusType    = int64
	DepositStatusType     = int64
	UndelegatedStatusType = int64

	UndelegateVersion = uint64
)
