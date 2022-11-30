package v2

import (
	"github.com/Carina-labs/nova/v2/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the nova v2 upgrade.
const UpgradeName = "v2"
const AirdropModuleName = "airdrop"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{AirdropModuleName},
	},
}
