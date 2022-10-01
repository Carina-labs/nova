package types

const (
	IcaPending IcaStatus = iota + 1
	IcaRequest
	IcaSuccess
)

type IcaStatus = uint64
