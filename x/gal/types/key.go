package types

const (
	// ModuleName defines the module name
	ModuleName = "gal"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// KVStore keys
var (
	// KeySupply defines key to store the total supply of snTokens.
	KeySupply = []byte{0x00}
	// KeyDepositInfo defines key to store deposit information of snTokens.
	KeyDepositInfo = []byte{0x01}
	// KeyWithdrawInfo defines key to store withdraw information of snTokens.
	KeyWithdrawInfo = []byte{0x02}
)
