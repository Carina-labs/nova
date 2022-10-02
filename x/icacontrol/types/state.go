package types

const (
	IcaPending IcaStatus = iota + 1
	IcaRequest
	IcaSuccess
	IcaFail
)

type IcaStatus = uint64
