package apptesting

import (
	"time"

	"github.com/Carina-labs/nova/v2/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestHelper struct {
	suite.Suite

	App         *app.NovaApp
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
	nTestAccs   int
}

func (s *KeeperTestHelper) Setup() {
	s.App = Setup(false)
	s.Ctx = s.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "nova-1", Time: time.Now().UTC()})
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}
	s.nTestAccs = 5
	s.TestAccs = CreateRandomAccounts(s.nTestAccs)
}

func CreateRandomAccounts(n int) []sdk.AccAddress {
	testAddr := make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddr[i] = sdk.AccAddress(pk.Address())
	}

	return testAddr
}

// FundAcc add coins to the address
func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	err := simapp.FundAccount(s.App.BankKeeper, s.Ctx, acc, amounts)
	s.Require().NoError(err)
}

// EndBlock run end block process
func (s *KeeperTestHelper) EndBlock() {
	reqEndBlock := abci.RequestEndBlock{Height: s.Ctx.BlockHeight()}
	s.App.EndBlocker(s.Ctx, reqEndBlock)
}

// BuildTx create transaction
func (s *KeeperTestHelper) BuildTx(
	txBuilder client.TxBuilder,
	msgs []sdk.Msg,
	sigV2 signing.SignatureV2,
	memo string, txFee sdk.Coins,
	gasLimit uint64,
) authsigning.Tx {
	err := txBuilder.SetMsgs(msgs[0])
	s.Require().NoError(err)

	err = txBuilder.SetSignatures(sigV2)
	s.Require().NoError(err)

	txBuilder.SetMemo(memo)
	txBuilder.SetFeeAmount(txFee)
	txBuilder.SetGasLimit(gasLimit)

	return txBuilder.GetTx()
}
