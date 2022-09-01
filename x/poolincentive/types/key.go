package types

const (
	ModuleName = "poolincentive"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

// KVStore keys
var (
	KeyCandidatePool = []byte{0x00}
	KeyIncentivePool = []byte{0x01}
)
