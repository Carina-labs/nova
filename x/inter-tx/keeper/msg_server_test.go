package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (suite *KeeperTestSuite) TestRegisterZone() {

	testOwnerAddress := suite.GenRandomAddress().String()

	var (
		path             *ibctesting.Path
		packetData       icatypes.InterchainAccountPacketData
		chanCap          *capabilitytypes.Capability
		timeoutTimestamp uint64
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				interchainAccountAddr, found := suite.chainA.GetApp().ICAControllerKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), ibctesting.FirstConnectionID, path.EndpointA.ChannelConfig.PortID)
				suite.Require().True(found)

				msg := &banktypes.MsgSend{
					FromAddress: interchainAccountAddr,
					ToAddress:   suite.chainB.SenderAccount.GetAddress().String(),
					Amount:      sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))),
				}

				data, err := icatypes.SerializeCosmosTx(suite.chainB.GetApp().AppCodec(), []sdk.Msg{msg})
				suite.Require().NoError(err)

				packetData = icatypes.InterchainAccountPacketData{
					Type: icatypes.EXECUTE_TX,
					Data: data,
				}
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.msg, func() {
			suite.SetupTest()             // reset
			timeoutTimestamp = ^uint64(0) // default

			path := NewICAPAth(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			err := suite.SetupICAPath(path, testOwnerAddress)
			fmt.Println(err)
			suite.Require().NoError(err)

			var ok bool
			chanCap, ok = suite.chainA.GetApp().ScopedICAMockKeeper.GetCapability(path.EndpointA.Chain.GetContext(), host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
			suite.Require().True(ok)

			tc.malleate() // malleate mutates test data

			_, err = suite.chainA.GetApp().ICAControllerKeeper.SendTx(suite.chainA.GetContext(), chanCap, ibctesting.FirstConnectionID, path.EndpointA.ChannelConfig.PortID, packetData, timeoutTimestamp)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
