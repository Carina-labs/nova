package types

const (
	ModuleName   = "airdrop"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	KeyAirdropInfo  = []byte{0x00}
	KeyAirdropState = []byte{0x01}
)

func GetKeyAirdropInfo() []byte {
	return KeyAirdropInfo
}

func GetKeyAirdropState(user string) []byte {
	return append(KeyAirdropState, []byte(user)...)
}
