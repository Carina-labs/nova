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
	// KeySupply defines key to store the total supply of sn-assets.
	KeySupply = []byte{0x00}
	// KeyDepositInfo defines key to store deposit information of sn-assets.
	KeyDepositInfo = []byte{0x01}
	// KeyWithdrawRecordInfo defines key to store withdraw record information of sn-assets.
	KeyWithdrawRecordInfo = []byte{0x02}
	// KeyUndelegateRecordInfo defines key to store undelegate record information of wAsset.
	KeyUndelegateRecordInfo = []byte{0x03}
	// KeyShare defines key to store deposit information of sn-assets.
	KeyShare = []byte{0x06}
	// KeyWithdrawInfo defines key to store withdraw information of sn-assets.
	KeyWithdrawInfo = []byte{0x07}

	KeyDelegateVersion   = []byte{0x08}
	KeyUndelegateVersion = []byte{0x09}
	KeyWithdrawVersion   = []byte{0x1}
)
