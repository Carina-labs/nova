package keeper_test

import (
	galkeeper "github.com/Carina-labs/nova/x/gal/keeper"
	"github.com/Carina-labs/nova/x/gal/types"
	ibcstakingtypes "github.com/Carina-labs/nova/x/ibcstaking/types"
	minttypes "github.com/Carina-labs/nova/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"time"
)

func ParseAddressToIbcAddress(destPort string, destChannel string, denom string) string {
	sourcePrefix := transfertypes.GetDenomPrefix(destPort, destChannel)
	prefixedDenom := sourcePrefix + denom
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	voucherDenom := denomTrace.IBCDenom()
	return voucherDenom
}

func (suite *KeeperTestSuite) setControllerAddr(address string) {
	var addresses []string
	addr1 := address
	addresses = append(addresses, addr1)
	params := ibcstakingtypes.Params{
		DaoModifiers: addresses,
	}
	suite.chainA.App.IbcstakingKeeper.SetParams(suite.chainA.GetContext(), params)
}

func (suite *KeeperTestSuite) setMintWAsset(amt sdk.Int, toAddr sdk.AccAddress) sdk.Coin {
	var wAssets sdk.Coins

	ibcDenom := suite.App.IbcstakingKeeper.GetIBCHashDenom(suite.chainA.GetContext(), transferPort, transferChannel, baseDenom)

	wAsset := sdk.NewCoin(ibcDenom, amt)
	wAssets = append(wAssets, wAsset)

	//snAsset mint
	suite.chainA.App.BankKeeper.MintCoins(suite.chainA.GetContext(), minttypes.ModuleName, wAssets)
	suite.chainA.App.BankKeeper.SendCoinsFromModuleToAccount(suite.chainA.GetContext(), minttypes.ModuleName, toAddr, wAssets)

	return wAsset
}

func setWithdrawRecordContents(cnt int) map[uint64]*types.WithdrawRecordContent {
	records := make(map[uint64]*types.WithdrawRecordContent)

	for i := 1; i <= cnt; i++ {
		amt := int64(i * 100)
		records[uint64(i)] = &types.WithdrawRecordContent{
			Amount:          sdk.NewInt(amt),
			CompletionTime:  time.Now(),
			WithdrawVersion: uint64(0),
			State:           int64(2),
			OracleVersion:   int64(0),
		}
	}

	return records
}

func setWithdrawRecords(zoneId, withdrawer string, recordContentsCnt int) *types.WithdrawRecord {
	var record *types.WithdrawRecord

	recordContents := setWithdrawRecordContents(recordContentsCnt)
	record = &types.WithdrawRecord{
		ZoneId:     zoneId,
		Withdrawer: withdrawer,
		Records:    recordContents,
	}
	return record
}

func (suite *KeeperTestSuite) TestWithdraw() {
	suite.setControllerAddr(suite.GetControllerAddr())

	withdrawer := suite.GenRandomAddress().String()
	controllerAddr, _ := sdk.AccAddressFromBech32(suite.GetControllerAddr())

	suite.chainA.App.IbcstakingKeeper.RegisterZone(suite.chainA.GetContext(), &ibcstakingtypes.RegisteredZone{
		ZoneId: zoneId,
		IcaConnectionInfo: &ibcstakingtypes.IcaConnectionInfo{
			ConnectionId: suite.icaPath.EndpointA.ConnectionID,
			PortId:       suite.icaPath.EndpointA.ChannelConfig.PortID,
		},
		IcaAccount: &ibcstakingtypes.IcaAccount{
			DaomodifierAddress: controllerAddr.String(),
		},
		ValidatorAddress: "",
		BaseDenom:        baseDenom,
		SnDenom:          baseSnDenom,
	})

	record := setWithdrawRecords(zoneId, withdrawer, 3)

	suite.chainA.App.GalKeeper.SetWithdrawRecord(suite.chainA.GetContext(), record)
	suite.setMintWAsset(sdk.NewInt(10000), controllerAddr)

	tcs := []struct {
		name        string
		withdrawMsg types.MsgWithdraw
		shouldErr   bool
	}{
		{
			name: "success",
			withdrawMsg: types.MsgWithdraw{
				ZoneId:            zoneId,
				Withdrawer:        withdrawer,
				TransferChannelId: transferChannel,
				TransferPortId:    transferPort,
			},
			shouldErr: false,
		},
		{
			name: "transfer channel id is not found",
			withdrawMsg: types.MsgWithdraw{
				ZoneId:            zoneId,
				Withdrawer:        withdrawer,
				TransferChannelId: "channel-111",
				TransferPortId:    transferPort,
			},
			shouldErr: true,
		},
		{
			name: "transfer port id is not found",
			withdrawMsg: types.MsgWithdraw{
				ZoneId:            zoneId,
				Withdrawer:        withdrawer,
				TransferChannelId: transferChannel,
				TransferPortId:    "testport",
			},
			shouldErr: true,
		},
		{
			name: "wAsset is zero amount",
			withdrawMsg: types.MsgWithdraw{
				ZoneId:            zoneId,
				Withdrawer:        withdrawer,
				TransferChannelId: transferChannel,
				TransferPortId:    transferPort,
			},
			shouldErr: true,
		},
		{
			name: "zone not found",
			withdrawMsg: types.MsgWithdraw{
				ZoneId:            "test",
				Withdrawer:        withdrawer,
				TransferChannelId: transferChannel,
				TransferPortId:    transferPort,
			},
			shouldErr: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			msgServer := galkeeper.NewMsgServerImpl(suite.chainA.App.GalKeeper)
			_, err := msgServer.Withdraw(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.withdrawMsg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
