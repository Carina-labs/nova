package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	"github.com/Carina-labs/nova/x/gal/types"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	baseDenom   = "nova"
	baseAcc     = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	baseHostAcc = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
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
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			expect: []args{
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			wantErr: false,
		},
		{
			name: "should not get deposit info",
			args: []args{},
			expect: []args{
				{sdk.NewInt64Coin(baseDenom, 10000), randAddr},
			},
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			for _, arg := range tc.args {
				err := suite.App.GalKeeper.SetDepositAmt(
					suite.Ctx,
					&types.DepositRecord{
						Address: arg.addr.String(),
						Records: []*types.DepositRecordContent{
							{
								ZoneId:        "test-zone-id",
								Amount:        &arg.coin,
								IsTransferred: false,
							},
						},
					})
				suite.Require().NoError(err)
			}

			for _, query := range tc.expect {
				res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, query.addr)
				if tc.wantErr {
					suite.Require().NotNil(err, "error expected but no error found")
					suite.Require().Equal(err, types.ErrNoDepositRecord)
					continue
				}

				suite.Require().NoError(err)
				for _, record := range res.Records {
					suite.Require().Equal(record.Amount.Denom, query.coin.Denom)
					suite.Require().Equal(record.Amount.Amount, query.coin.Amount)
					suite.Require().Equal(res.Address, query.addr.String())
				}
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
	type tcConfig struct {
		genesisStateMsgA banktypes.GenesisState
		registerZoneMsgs []*intertxtypes.RegisteredZone
	}
	type doArg struct {
		addr string
		amt  sdk.Coin
	}
	type verifyArg struct {
		addr string
	}
	tcs := []struct {
		name      string
		setting   tcConfig
		doArg     doArg
		do        func(doArg) error
		verifyArg verifyArg
		verify    func(verifyArg)
		wantErr   bool
	}{
		{
			name: "valid test case 1",
			setting: tcConfig{
				genesisStateMsgA: banktypes.GenesisState{
					Balances: []banktypes.Balance{
						{
							Address: "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
							Coins: sdk.Coins{
								sdk.NewInt64Coin(baseDenom, 1000),
							},
						},
					},
				},
			},
			doArg: doArg{
				addr: "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
				amt:  sdk.NewInt64Coin(baseDenom, 500),
			},
			do: func(arg doArg) error {
				return suite.chainA.App.GalKeeper.Deposit(
					suite.chainA.GetContext(),
					&types.MsgDeposit{
						Depositor: arg.addr,
						Amount:    sdk.Coins{sdk.NewInt64Coin(baseDenom, 500)},
						HostAddr:  arg.addr,
						ZoneId:    baseDenom,
					},
				)
			},
			verifyArg: verifyArg{
				addr: "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
			},
			verify: func(arg verifyArg) {
				acc, _ := sdk.AccAddressFromBech32(arg.addr)
				record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), acc)
				suite.Require().NoError(err)
				suite.Require().Equal(1, len(record.Records))
				suite.Require().True(record.Records[0].Amount.IsEqual(sdk.NewInt64Coin(baseDenom, 500)))
			},
		},
		{
			name: "valid multiple deposit case",
			setting: tcConfig{
				genesisStateMsgA: banktypes.GenesisState{
					Balances: []banktypes.Balance{
						{
							Address: "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
							Coins: sdk.Coins{
								sdk.NewInt64Coin(baseDenom, 1000),
							},
						},
					},
				},
			},
			doArg: doArg{
				addr: "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
				amt:  sdk.NewInt64Coin(baseDenom, 500),
			},
			do: func(arg doArg) error {
				err := suite.chainA.App.GalKeeper.Deposit(
					suite.chainA.GetContext(),
					&types.MsgDeposit{
						Depositor: arg.addr,
						Amount:    sdk.Coins{sdk.NewInt64Coin(baseDenom, 300)},
						HostAddr:  arg.addr,
						ZoneId:    baseDenom,
					},
				)
				if err != nil {
					return err
				}

				return suite.chainA.App.GalKeeper.Deposit(
					suite.chainA.GetContext(),
					&types.MsgDeposit{
						Depositor: arg.addr,
						Amount:    sdk.Coins{sdk.NewInt64Coin(baseDenom, 500)},
						HostAddr:  arg.addr,
						ZoneId:    baseDenom,
					},
				)
			},
			verifyArg: verifyArg{
				addr: "cosmos1a05qwsaeqgdp7pc3tsegw87w9c0j6xlhdk84f3",
			},
			verify: func(arg verifyArg) {
				acc, _ := sdk.AccAddressFromBech32(arg.addr)
				record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(suite.chainA.GetContext(), acc)
				suite.Require().NoError(err)
				suite.Require().Equal(2, len(record.Records))
				suite.Require().True(record.Records[0].Amount.IsEqual(sdk.NewInt64Coin(baseDenom, 300)))
				suite.Require().True(record.Records[1].Amount.IsEqual(sdk.NewInt64Coin(baseDenom, 500)))

				userBalance := suite.chainA.App.BankKeeper.GetBalance(suite.chainA.GetContext(), acc, baseDenom)
				suite.Require().True(userBalance.IsEqual(sdk.NewInt64Coin(baseDenom, 200)))
			},
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			// call SetupTest in each test case, for creating new keeper instance.
			suite.SetupTest()

			// setup state
			suite.chainA.App.BankKeeper.InitGenesis(suite.chainA.GetContext(), &tc.setting.genesisStateMsgA)
			suite.chainA.App.IntertxKeeper.RegisterZone(suite.chainA.GetContext(), newBaseRegisteredZone())

			// execute
			err := tc.do(tc.doArg)
			suite.Require().NoError(err)

			// verify
			tc.verify(tc.verifyArg)
		})
	}
}

// newBaseRegisteredZone returns a new zone info for testing purpose only
func newBaseRegisteredZone() *intertxtypes.RegisteredZone {
	icaControllerPort, _ := icatypes.NewControllerPortID(baseAcc.String())

	return &intertxtypes.RegisteredZone{
		ZoneId: baseDenom,
		IcaConnectionInfo: &intertxtypes.IcaConnectionInfo{
			ConnectionId: "connection-0",
			PortId:       icaControllerPort,
		},
		TransferConnectionInfo: &intertxtypes.TransferConnectionInfo{
			ConnectionId: "connection-0",
			PortId:       "transfer",
			ChannelId:    "channel-0",
		},
		IcaAccount: &intertxtypes.IcaAccount{
			OwnerAddress: baseAcc.String(),
			HostAddress:  baseHostAcc.String(),
		},
		ValidatorAddress: "",
		BaseDenom:        baseDenom,
		SnDenom:          "snOsmo",
	}
}

func (suite *KeeperTestSuite) TestMarkRecordTransfer() {
	// prepare
	type tcConfig struct {
		addr  string
		coins []sdk.Coin
	}
	type arg struct {
		addr      string
		markIndex []int
	}
	tcs := []struct {
		name    string
		config  tcConfig
		setup   func(tcConfig)
		arg     arg
		do      func(arg) error
		verify  func()
		wantErr bool
	}{
		{
			name: "single valid test case 1",
			config: tcConfig{
				addr:  "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
				coins: sdk.NewCoins(sdk.NewInt64Coin(baseDenom, 100)),
			},
			setup: func(arg tcConfig) {
				var records []*types.DepositRecordContent
				for _, coin := range arg.coins {
					records = append(records, &types.DepositRecordContent{
						ZoneId:        baseDenom,
						Amount:        &coin,
						IsTransferred: false,
					})
				}
				err := suite.App.GalKeeper.SetDepositAmt(suite.Ctx, &types.DepositRecord{
					Address: arg.addr,
					Records: records,
				})
				suite.Require().NoError(err)
			},
			arg: arg{
				addr:      "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
				markIndex: []int{0},
			},
			do: func(arg arg) error {
				return suite.App.GalKeeper.MarkRecordTransfer(suite.Ctx, arg.addr, arg.markIndex[0])
			},
			wantErr: false,
		},
		{
			name: "multiple execute, valid test case 2",
			config: tcConfig{
				addr: "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
				coins: []sdk.Coin{
					sdk.NewInt64Coin(baseDenom, 100),
					sdk.NewInt64Coin(baseDenom, 200),
					sdk.NewInt64Coin(baseDenom, 300),
				},
			},
			setup: func(arg tcConfig) {
				var records []*types.DepositRecordContent
				for _, coin := range arg.coins {
					records = append(records, &types.DepositRecordContent{
						ZoneId:        baseDenom,
						Amount:        &coin,
						IsTransferred: false,
					})
				}
				err := suite.App.GalKeeper.SetDepositAmt(suite.Ctx, &types.DepositRecord{
					Address: arg.addr,
					Records: records,
				})
				suite.Require().NoError(err)
			},
			arg: arg{
				addr:      "cosmos1l2pqgjx6qgavg8x984s5jgc6u2ehqkfq3azx7a",
				markIndex: []int{0, 1, 2},
			},
			do: func(arg arg) error {
				for _, index := range arg.markIndex {
					err := suite.App.GalKeeper.MarkRecordTransfer(suite.Ctx, arg.addr, index)
					if err != nil {
						return err
					}
				}

				return nil
			},
			wantErr: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup(tc.config)
			if !tc.wantErr {
				suite.Require().NoError(tc.do(tc.arg))
			} else {
				suite.Require().Error(tc.do(tc.arg))
				return
			}

			acc, _ := sdk.AccAddressFromBech32(tc.arg.addr)
			res, err := suite.App.GalKeeper.GetRecordedDepositAmt(suite.Ctx, acc)
			suite.Require().NoError(err)
			for _, index := range tc.arg.markIndex {
				suite.Require().True(res.Records[index].IsTransferred)
			}
		})
	}
}
