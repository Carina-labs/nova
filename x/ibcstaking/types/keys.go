package types

const (
	ModuleName = "ibcstaking"

	StoreKey = "storeibcstaking"

	RouterKey = ModuleName

	QuerierRoute = ModuleName

	PortID = ModuleName

	Version = "ics27-1"
)

// prefix bytes for the epoch persistent store
const (
	prefixZone           = iota + 1
	prefixConnectionInfo = iota + 1
)

var (
	KeyPrefixZone           = []byte{prefixZone}
	KeyPrefixConnectionInfo = []byte{prefixConnectionInfo}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
