package keeper_test

import (
	"strconv"
	"testing"

	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/stretchr/testify/suite"

	"github.com/Carina-labs/nova/app/apptesting"
	novatesting "github.com/Carina-labs/nova/testing"
)

var (
	key1 = secp256k1.GenPrivKey()
	acc1 = authtypes.NewBaseAccount(key1.PubKey().Address().Bytes(), key1.PubKey(), 0, 0)
	key2 = secp256k1.GenPrivKey()
	acc2 = authtypes.NewBaseAccount(key2.PubKey().Address().Bytes(), key2.PubKey(), 0, 0)
	key3 = secp256k1.GenPrivKey()
	acc3 = authtypes.NewBaseAccount(key3.PubKey().Address().Bytes(), key3.PubKey(), 0, 0)
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient

	coordinator  *novatesting.Coordinator
	chainA       *novatesting.TestChain
	chainB       *novatesting.TestChain
	ctxA         sdk.Context
	ctxB         sdk.Context
	path         *novatesting.Path
	icaPath      *novatesting.Path
	transferPath *novatesting.Path
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) GenRandomAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(key.PubKey().Address().Bytes(), key.PubKey(), 0, 0)
	return acc.GetAddress()
}

func (suite *KeeperTestSuite) setZone(num int) []types.RegisteredZone {

	addr := make([]sdk.AccAddress, 0)
	zones := make([]types.RegisteredZone, 0)

	for i := 0; i < num; i++ {
		addr = append(addr, suite.GenRandomAddress())
		zones = append(zones, types.RegisteredZone{
			ZoneId: "gaia" + strconv.Itoa(i),
			IcaConnectionInfo: &types.IcaConnectionInfo{
				ConnectionId: "connection-" + strconv.Itoa(i),
				PortId:       "gaia" + strconv.Itoa(i) + "." + addr[i].String(),
			},
			IcaAccount: &types.IcaAccount{
				ControllerAddress: addr[i].String(),
			},
			TransferInfo: &types.TransferConnectionInfo{
				PortId:    "transfer",
				ChannelId: "channel-" + strconv.Itoa(i),
			},
			ValidatorAddress: sdk.ValAddress(addr[i]).String(),
			BaseDenom:        "atom" + strconv.Itoa(i),
			SnDenom:          "snatom" + strconv.Itoa(i),
		})
	}

	return zones
}

// RegisterInterchainAccount is a helper function for starting the channel handshake
func (suite *KeeperTestSuite) RegisterInterchainAccount(endpoint *novatesting.Endpoint, owner string) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelSequence := endpoint.Chain.App.IBCKeeper.ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := endpoint.Chain.App.ICAControllerKeeper.RegisterInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, owner); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = ibcchanneltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID
	return nil
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)

	suite.coordinator = novatesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(novatesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(novatesting.GetChainID(2))

	suite.transferPath = NewIbcTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.transferPath)
	suite.icaPath = newIcaPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.icaPath)

	suite.chainA.GetApp().IcaControlKeeper.RegisterZone(suite.chainA.GetContext(), newBaseRegisteredZone())

	err := setupIcaPath(suite.icaPath, zoneId+"."+baseOwnerAcc.String())
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) SetupTestIBCZone(zoneMsgs []types.RegisteredZone) {
	for _, msg := range zoneMsgs {
		suite.App.IcaControlKeeper.RegisterZone(suite.Ctx, &msg)
	}
}

func (suite *KeeperTestSuite) TestIsValidControllerAddr() {
	var keyManager []string

	addr1 := suite.GenRandomAddress().String()
	addr2 := suite.GenRandomAddress().String()
	addr3 := suite.GenRandomAddress().String()

	keyManager = append(keyManager, addr1)
	keyManager = append(keyManager, addr2)

	params := types.Params{
		ControllerKeyManager: keyManager,
	}

	suite.App.IcaControlKeeper.SetParams(suite.Ctx, params)

	suite.App.IcaControlKeeper.SetControllerAddr(suite.Ctx, zoneId, addr3)

	tcs := []struct {
		name   string
		addr   string
		expect bool
	}{
		{
			name:   "success",
			addr:   addr3,
			expect: true,
		},
		{
			name:   "nil address",
			addr:   "",
			expect: false,
		},
		{
			name:   "random address",
			addr:   addr1,
			expect: false,
		},
		{
			name:   "invalid address",
			addr:   "test",
			expect: false,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			ok := suite.App.IcaControlKeeper.IsValidControllerAddr(suite.Ctx, zoneId, tc.addr)

			suite.Require().Equal(ok, tc.expect)
		})
	}
}

func (suite *KeeperTestSuite) GetControllerAddr() string {
	return acc1.Address
}
