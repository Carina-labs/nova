package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	"github.com/Carina-labs/nova/x/gal/types"
	intertxtypes "github.com/Carina-labs/nova/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
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

	userAcc := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	moduleAcc := suite.App.AccountKeeper.GetModuleAddress(types.ModuleName)

	ctxA := suite.chainA.GetContext()
	ctxB := suite.chainB.GetContext()

	// initialize chainA
	suite.chainA.App.BankKeeper.InitGenesis(ctxA, &banktypes.GenesisState{
		Balances: []banktypes.Balance{
			{userAcc.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, 6000)}},
			{moduleAcc.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, 1000)}},
		},
	})

	// initialize chainB
	suite.chainB.App.BankKeeper.InitGenesis(ctxB, &banktypes.GenesisState{
		Balances: []banktypes.Balance{
			{userAcc.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, 0)}},
			{moduleAcc.String(), sdk.Coins{sdk.NewInt64Coin(baseDenom, 1000)}},
		},
	})

	// register zone
	suite.chainA.App.IntertxKeeper.RegisterZone(ctxA, newBaseRegisteredZone())

	// deposit to gal keeper
	err := suite.chainA.App.GalKeeper.Deposit(ctxA, &types.MsgDeposit{
		Depositor: userAcc.String(),
		Amount:    sdk.Coins{sdk.NewInt64Coin(baseDenom, 500)},
		HostAddr:  userAcc.String(),
		ZoneId:    baseDenom,
	})

	suite.Require().NoError(err)

	// reveal chain A block
	suite.chainA.NextBlock()

	// relay packet to the chain B
	p, err := ibctesting.ParsePacketFromEvents(ctxA.EventManager().Events())
	suite.Require().NoError(err)
	err = suite.path.RelayPacket(p)
	suite.Require().NoError(err)

	// reveal chain B block
	suite.chainB.NextBlock()

	// user of chain B should get 500 osmo
	balance := suite.chainB.App.BankKeeper.GetAllBalances(ctxB, userAcc)
	fmt.Println(balance[0].Denom)
	fmt.Println(balance[0].Amount)
	suite.Require().Equal(balance[0].Amount.Int64(), int64(500))

	// check deposit record
	record, err := suite.chainA.App.GalKeeper.GetRecordedDepositAmt(ctxA, userAcc)
	suite.Require().NoError(err)
	suite.Require().Equal(record.ZoneId, baseDenom)
	suite.Require().Equal(record.Address, userAcc.String())
	suite.Require().Equal(record.Amount.Amount.Int64(), int64(500))
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
