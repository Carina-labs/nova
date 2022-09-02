package types

const (
	ModuleName   = "airdrop"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	KeyAirdropInfo = []byte{0x00}
	KeyUserState   = []byte{0x01}
)

func GetKeyAirdropInfo() []byte {
	return KeyAirdropInfo
}

func GetKeyUserState(user string) []byte {
	return append(KeyUserState, []byte(user)...)
}
