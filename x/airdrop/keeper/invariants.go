package keeper

import (
	"github.com/Carina-labs/nova/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterInvariants(registry sdk.InvariantRegistry, keeper Keeper, bankKeeper types.BankKeeper) {
}
