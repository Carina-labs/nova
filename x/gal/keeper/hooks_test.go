package keeper_test

import (
	"fmt"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (suite *KeeperTestSuite) TestAfterTransferEnd() {
	tcs := []struct {
		packet        types.FungibleTokenPacketData
		denom         string
		expectedDenom string
		expectedAmt   int64
		shouldErr     bool
	}{
		{
			packet: types.FungibleTokenPacketData{
				Denom:    "osmo",
				Amount:   "100000",
				Sender:   suite.GenRandomAddress().String(),
				Receiver: suite.GenRandomAddress().String(),
			},
			denom:         "osmo",
			expectedDenom: "osmo",
			expectedAmt:   100000,
			shouldErr:     false,
		},
		{
			packet: types.FungibleTokenPacketData{
				Denom:    "atom",
				Amount:   "55555",
				Sender:   suite.GenRandomAddress().String(),
				Receiver: suite.GenRandomAddress().String(),
			},
			denom:         "atom",
			expectedDenom: "atom",
			expectedAmt:   55555,
			shouldErr:     false,
		},
	}

	hooks := suite.App.GalKeeper.Hooks()
	for _, tc := range tcs {
		// register zone
		zoneInfo := &intertxtypes.RegisteredZone{
			ZoneName:               tc.denom,
			IcaConnectionInfo:      &intertxtypes.IcaConnectionInfo{},
			TransferConnectionInfo: &intertxtypes.TransferConnectionInfo{},
			ValidatorAddress:       "",
			BaseDenom:              tc.denom,
			StDenom:                fmt.Sprintf("st%s", tc.denom),
			SnDenom:                fmt.Sprintf("sn%s", tc.denom),
		}

		suite.App.IntertxKeeper.RegisterZone(suite.Ctx, zoneInfo)

		hooks.AfterTransferEnd(suite.Ctx, tc.packet, tc.denom)

		senderAddr, err := sdk.AccAddressFromBech32(tc.packet.Sender)
		suite.NoError(err)

		record, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, senderAddr)
		suite.Equal(tc.expectedDenom, record.Amount.Denom)
		suite.Equal(tc.expectedAmt, record.Amount.Amount.Int64())
		suite.Equal(tc.packet.Sender, record.Address)
	}
}
