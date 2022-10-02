package types

const (
	WithdrawStatusRegistered WithdrawStatusType = iota + 1
	WithdrawStatusTransferRequest
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

const (
	IcaPending IcaStatus = iota + 1
	IcaRequest
	IcaSuccess
	IcaFail
)

type (
	WithdrawStatusType    = int64
	DepositStatusType     = int64
	UndelegatedStatusType = int64
	IcaStatus             = uint64

	UndelegateVersion = uint64
)
