package keeper_test

import (
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func ParseAddressToIbcAddress(destPort string, destChannel string, denom string) string {
	sourcePrefix := transfertypes.GetDenomPrefix(destPort, destChannel)
	prefixedDenom := sourcePrefix + denom
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	voucherDenom := denomTrace.IBCDenom()
	return voucherDenom
}

