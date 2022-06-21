package keeper_test

import (
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (suite *KeeperTestSuite) TestRecordDepositAmt() {
	randAddr := suite.GenRandomAddress()
	type args struct {
		coin sdk.Coin
		addr sdk.AccAddress
	}
	tcs := []struct {
		name    string
		args    []args
		expect  []args
		wantErr bool

		denom         string
		amt           int64
		userAddr      sdk.AccAddress
		expectedDenom string
		expectedAmt   int64
	}{
		{
			name: "should get recorded deposit amt",
			args: []args{
				{sdk.NewInt64Coin("osmo", 10000), randAddr},
			},
			expect: []args{
				{sdk.NewInt64Coin("osmo", 10000), randAddr},
			},
			wantErr: false,
		},
		{
			name: "should not get deposit info",
			args: []args{},
			expect: []args{
				{sdk.NewInt64Coin("osmo", 10000), randAddr},
			},
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, arg := range tc.args {
				err := suite.App.GalKeeper.RecordDepositAmt(suite.Ctx, &types.DepositRecord{
					Address: arg.addr.String(),
					Amount:  &arg.coin,
				})
				suite.Require().NoError(err)
			}

			for _, query := range tc.expect {
				res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, query.addr)
				if tc.wantErr {
					suite.Require().NotNil(err, "error expected but no error found")
					continue
				}

				suite.Require().NoError(err)
				suite.Require().Equal(res.Amount.Denom, query.coin.Denom)
				suite.Require().Equal(res.Amount.Amount, query.coin.Amount)
				suite.Require().Equal(res.Address, query.addr.String())
			}

			for _, arg := range tc.args {
				err := suite.App.GalKeeper.ClearRecordedDepositAmt(suite.Ctx, arg.addr)
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}

func (suite *KeeperTestSuite) TestDeposit() {
	suite.SetupTest()
	tcs := []struct {
		name                 string
		userPrivKey          *secp256k1.PrivKey
		denom                string
		preUserAmt           int64
		preModuleAccountAmt  int64
		depositAmt           int64
		postUserAmt          int64
		postModuleAccountAmt int64
		shouldErr            bool
	}{
		{
			name:                 "valid test case 1",
			userPrivKey:          secp256k1.GenPrivKey(),
			denom:                "osmo",
			preUserAmt:           1000,
			preModuleAccountAmt:  1000,
			depositAmt:           500,
			postUserAmt:          500,
			postModuleAccountAmt: 1000,
			shouldErr:            false,
		},
		{
			name:                 "valid test case 2",
			userPrivKey:          secp256k1.GenPrivKey(),
			denom:                "osmo",
			preUserAmt:           4000,
			preModuleAccountAmt:  5000,
			depositAmt:           3000,
			postUserAmt:          1000,
			postModuleAccountAmt: 5000,
			shouldErr:            false,
		},
		{
			// ERROR CASE
			name:                 "error test case 1",
			userPrivKey:          secp256k1.GenPrivKey(),
			denom:                "osmo",
			preUserAmt:           5000,
			preModuleAccountAmt:  1000,
			depositAmt:           6000,
			postUserAmt:          5000,
			postModuleAccountAmt: 1000,
			shouldErr:            true,
		},
	}

	ctxA := suite.chainA.GetContext()
	ctxB := suite.chainB.GetContext()

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.chainB.App.BankKeeper.IterateAllBalances(ctxB, func(address sdk.AccAddress, coin sdk.Coin) bool {
				fmt.Printf("addr: %s, balance: %s\n", address.String(), coin.String())
				return false
			})

			acc := authtypes.NewBaseAccount(tc.userPrivKey.PubKey().Address().Bytes(), tc.userPrivKey.PubKey(), 0, 0)
			accAddr, err := sdk.AccAddressFromBech32(acc.Address)
			suite.Require().NoError(err)

			galAddr := suite.App.AccountKeeper.GetModuleAddress(types.ModuleName)
			suite.chainA.App.BankKeeper.InitGenesis(ctxA, &banktypes.GenesisState{
				Balances: []banktypes.Balance{
					{
						Address: accAddr.String(),
						Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.preUserAmt)},
					},
					{
						Address: galAddr.String(),
						Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.preModuleAccountAmt)},
					},
				},
			})

			suite.chainB.App.BankKeeper.InitGenesis(ctxB, &banktypes.GenesisState{
				Balances: []banktypes.Balance{
					{
						Address: accAddr.String(),
						Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, 0)},
					},
					{
						Address: galAddr.String(),
						Coins:   sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.preModuleAccountAmt)},
					},
				},
			})

			suite.chainA.App.IntertxKeeper.RegisterZone(ctxA, &intertxtypes.RegisteredZone{
				ZoneId: "osmo",
				IcaConnectionInfo: &intertxtypes.IcaConnectionInfo{
					ConnectionId: "connection-0",
					PortId:       "icacontroller-" + accAddr.String(),
				},
				TransferConnectionInfo: &intertxtypes.TransferConnectionInfo{
					ConnectionId: "connection-0",
					PortId:       "transfer",
					ChannelId:    "channel-0",
				},
				IcaAccount: &intertxtypes.IcaAccount{
					OwnerAddress: accAddr.String(),
					HostAddress:  "",
				},
				ValidatorAddress: "",
				BaseDenom:        "osmo",
				SnDenom:          "snOsmo",
				StDenom:          "stOsmo",
			})

			err = suite.chainA.App.GalKeeper.Deposit(ctxA, &types.MsgDeposit{
				Depositor: accAddr.String(),
				Amount:    sdk.Coins{sdk.NewInt64Coin(tc.denom, tc.depositAmt)},
				HostAddr:  accAddr.String(),
				ZoneId:    "osmo",
			})

			if tc.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.chainA.NextBlock()

			p, err := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events())
			suite.Require().NoError(err)

			err = suite.path.RelayPacket(p)
			suite.Require().NoError(err)
			suite.chainB.NextBlock()

			res := suite.chainB.App.BankKeeper.GetAllBalances(ctxB, accAddr)
			suite.Require().Equal(res[0].Amount.Int64(), tc.depositAmt)

			// Check record
			record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(ctxA, accAddr)
			fmt.Printf("record: %s\n", record.String())
			suite.Require().Equal(record.Address, accAddr.String())
			suite.Require().Equal(record.Amount.Denom, tc.denom)
			suite.Require().Equal(record.Amount.Amount.Int64(), tc.depositAmt)
		})

	}
}

/*
Comment : TEST EVENT DATA
event type: coin_spent attr: key:"spender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp" attr: key:"amount" value:"500osmo"
event type: coin_received attr: key:"receiver" value:"cosmos1qx63cevacrwd4wrqlfvmdy03vttgynz3gyd9yp" attr: key:"amount" value:"500osmo"
event type: transfer attr: key:"recipient" value:"cosmos1qx63cevacrwd4wrqlfvmdy03vttgynz3gyd9yp" attr: key:"sender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp" attr: key:"amount" value:"500osmo"
event type: message attr: key:"sender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp"
event type: coin_spent attr: key:"spender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp" attr: key:"amount" value:"500osmo"
event type: coin_received attr: key:"receiver" value:"cosmos1a53udazy8ayufvy0s434pfwjcedzqv34kvz9tw" attr: key:"amount" value:"500osmo"
event type: transfer attr: key:"recipient" value:"cosmos1a53udazy8ayufvy0s434pfwjcedzqv34kvz9tw" attr: key:"sender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp" attr: key:"amount" value:"500osmo"
event type: message attr: key:"sender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp"
event type: send_packet attr: key:"packet_data" value:"{\"amount\":\"500\",\"denom\":\"osmo\",\"receiver\":\"\",\"sender\":\"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp\"}" attr: key:"packet_data_hex" value:"7b22616d6f756e74223a22353030222c2264656e6f6d223a226f736d6f222c227265636569766572223a22222c2273656e646572223a22636f736d6f7331306566613273726a33723572663665797a78647875796a346374326138666166366539666470227d" attr: key:"packet_timeout_height" value:"0-0" attr: key:"packet_timeout_timestamp" value:"1655715387013916000" attr: key:"packet_sequence" value:"1" attr: key:"packet_src_port" value:"transfer" attr: key:"packet_src_channel" value:"channel-0" attr: key:"packet_dst_port" value:"transfer" attr: key:"packet_dst_channel" value:"channel-0" attr: key:"packet_channel_ordering" value:"ORDER_UNORDERED" attr: key:"packet_connection" value:"connection-0"
event type: message attr: key:"module" value:"ibc_channel"
event type: ibc_transfer attr: key:"sender" value:"cosmos10efa2srj3r5rf6eyzxdxuyj4ct2a8faf6e9fdp" attr: key:"receiver"
event type: message attr: key:"module" value:"transfer"
*/

func ContainEvent(em *sdk.EventManager, eventType, key, value string) bool {
	for _, event := range em.Events() {
		if event.Type == eventType {
			for _, attr := range event.Attributes {
				k := string(attr.Key[:])
				v := string(attr.Value[:])
				if key == k && value == v {
					return true
				}
			}
		}
	}
	return false
}

func PrintEvents(em *sdk.EventManager) {
	for _, event := range em.Events() {
		fmt.Printf("type: %s, ", event.Type)
		for _, attr := range event.Attributes {
			fmt.Printf("(key: %s, value: %s)", attr.Key, attr.Value)
		}
		fmt.Println()
	}
}
