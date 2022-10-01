package types

const (
	ModuleName   = "oracle"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	KeyOracleVersion    = []byte{0x00}
	KeyOracleAddr       = []byte{0x01}
	KeyOracleChainState = []byte{0x02}
)
