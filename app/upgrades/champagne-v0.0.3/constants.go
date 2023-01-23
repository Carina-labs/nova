package champagne_v0_0_3

import (
	"github.com/Carina-labs/nova/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "champagne-v0.0.3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateChampagne_v0_0_3_handler,
	StoreUpgrades:        store.StoreUpgrades{},
}
