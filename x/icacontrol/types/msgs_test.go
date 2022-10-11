package types_test

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"testing"
)

func TestValidateBasic(t *testing.T) {
	tcs := []struct {
		name    string
		msg     types.MsgRegisterZone
		wantErr bool
	}{
		{
			name: "valid case 1",
			msg: types.MsgRegisterZone{
				ZoneId: "test-zone",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: "connection-1",
					PortId:       "transfer",
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: "cosmos1zkarsurgym3hnm06qupyt96pu0k24k4f93tgjq",
					HostAddress:       "cosmos1zkarsurgym3hnm06qupyt96pu0k24k4f93tgjq",
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: "channel-1",
					PortId:    "transfer",
				},
				ValidatorAddress:     "cosmosvaloper1zkarsurgym3hnm06qupyt96pu0k24k4fq9la7n",
				BaseDenom:            "uatom",
				Decimal:              int64(6),
				DepositMaxEntries:    int64(100),
				UndelegateMaxEntries: int64(100),
			},
			wantErr: false,
		},
		{
			name: "error: invalid validator address",
			msg: types.MsgRegisterZone{
				ZoneId: "test-zone",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: "connection-1",
					PortId:       "transfer",
				},
				IcaAccount: &types.IcaAccount{
					ControllerAddress: "cosmos1zkarsurgym3hnm06qupyt96pu0k24k4f93tgjq",
					HostAddress:       "cosmos1zkarsurgym3hnm06qupyt96pu0k24k4f93tgjq",
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: "channel-1",
					PortId:    "transfer",
				},
				// ValidatorAddress must be format of "validator address"
				ValidatorAddress:     "cosmos1zkarsurgym3hnm06qupyt96pu0k24k4f93tgjq",
				BaseDenom:            "uatom",
				Decimal:              int64(6),
				DepositMaxEntries:    int64(100),
				UndelegateMaxEntries: int64(100),
			},
			wantErr: false,
		},
		{
			name: "valid case 1",
			msg: types.MsgRegisterZone{
				ZoneId: "test-zone",
				IcaInfo: &types.IcaConnectionInfo{
					ConnectionId: "connection-1",
					PortId:       "transfer",
				},
				IcaAccount: &types.IcaAccount{
					// Each addresses is must be format of bech32 address
					ControllerAddress: "abcd",
					HostAddress:       "abcd",
				},
				TransferInfo: &types.TransferConnectionInfo{
					ChannelId: "channel-1",
					PortId:    "transfer",
				},
				ValidatorAddress:     "cosmosvaloper1zkarsurgym3hnm06qupyt96pu0k24k4fq9la7n",
				BaseDenom:            "uatom",
				Decimal:              int64(6),
				DepositMaxEntries:    int64(100),
				UndelegateMaxEntries: int64(100),
			},
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				if err == nil {
					t.Errorf("should occur error")
				}
			} else {
				if err != nil {
					t.Errorf("should not error, err: %s", err.Error())
				}
			}
		})
	}
}
