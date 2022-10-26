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
	// KeyDepositRecordInfo defines key to store deposit information of sn-assets.
	KeyDepositRecordInfo  = []byte{0x01}
	KeyDelegateRecordInfo = []byte{0x02}
	// KeyUndelegateRecordInfo defines key to store undelegate record information of wAsset.
	KeyUndelegateRecordInfo = []byte{0x03}
	// KeyWithdrawRecordInfo defines key to store withdraw record information of sn-assets.
	KeyWithdrawRecordInfo = []byte{0x04}

	KeyDelegateVersion   = []byte{0x05}
	KeyUndelegateVersion = []byte{0x06}
	KeyWithdrawVersion   = []byte{0x07}
)
