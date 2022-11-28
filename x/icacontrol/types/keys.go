package types

const (
	ModuleName   = "icacontrol"
	StoreKey     = "storeIcaControl"
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	PortID       = ModuleName
	Version      = "ics27-1"
)

const (
	PrefixSnAsset = "sn"
)

var (
	KeyPrefixZone         = []byte{0x01}
	KeyAutoStakingVersion = []byte{0x02}

	KeyControllerAddress = []byte{0x03}
	KeyZoneProposal      = []byte{0x04}
	KeyAuthzGrantInfo    = []byte{0x05}
)
