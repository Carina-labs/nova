package keeper_test

import (
	ibcstakingkeeper "github.com/Carina-labs/nova/x/icacontrol/keeper"
	ibcstakingtypes "github.com/Carina-labs/nova/x/icacontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"time"
)

func bech32toValidatorAddresses(validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
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

func (suite *KeeperTestSuite) getGrantMsg(msg, zoneId, grantee string, controllerAddr sdk.AccAddress) ibcstakingtypes.MsgIcaAuthzGrant {
	var authorization authz.Authorization
	var allowed []sdk.ValAddress
	var denied []sdk.ValAddress
	var allowValidators []string
	var denyValidators []string
	var delegateLimit *sdk.Coin

	switch msg {
	case "send":
		spendLimit, _ := sdk.ParseCoinsNormalized("10000")
		authorization = bank.NewSendAuthorization(spendLimit)
	case "generic":
		msgType := ""
		authorization = authz.NewGenericAuthorization(msgType)
	case "delegate", "unbond", "redelegate":
		allowValidators = append(allowValidators, "")
		denyValidators = append(denyValidators, "")

		allowed, _ = bech32toValidatorAddresses(allowValidators)
		denied, _ = bech32toValidatorAddresses(denyValidators)
	}

	switch msg {
	case "delegate":
		authorization, _ = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, delegateLimit)
	case "undelegate":
		authorization, _ = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, delegateLimit)
	case "redelegate":
		authorization, _ = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, delegateLimit)
	}

	time := time.Now().AddDate(2, 0, 0).UTC()
	grantMsg, _ := ibcstakingtypes.NewMsgAuthzGrant(zoneId, grantee, controllerAddr, authorization, time)

	return *grantMsg
}

func (suite *KeeperTestSuite) TestAuthzGrant() {
	suite.setControllerAddr(suite.GetControllerAddr())
	granteeAddr := suite.GenRandomAddress()
	controllerAddr, _ := sdk.AccAddressFromBech32(suite.GetControllerAddr())
	randAddr := suite.GenRandomAddress()

	tcs := []struct {
		name      string
		grantMsg  ibcstakingtypes.MsgIcaAuthzGrant
		shouldErr bool
	}{
		{
			name:      "success",
			grantMsg:  suite.getGrantMsg("send", "gaia0", granteeAddr.String(), controllerAddr),
			shouldErr: false,
		},
		{
			name:      "success",
			grantMsg:  suite.getGrantMsg("send", "gaia0", granteeAddr.String(), controllerAddr),
			shouldErr: false,
		},
		{
			name:      "fail - not found zone name",
			grantMsg:  suite.getGrantMsg("send", "osmo", granteeAddr.String(), controllerAddr),
			shouldErr: true,
		},
		{
			name:      "fail - controller address is not found",
			grantMsg:  suite.getGrantMsg("send", "gaia", granteeAddr.String(), randAddr),
			shouldErr: true,
		},
	}

	for _, tc := range tcs {
		suite.Run(tc.name, func() {
			msgServer := ibcstakingkeeper.NewMsgServerImpl(*suite.chainA.App.IbcstakingKeeper)
			_, err := msgServer.IcaAuthzGrant(sdk.WrapSDKContext(suite.chainA.GetContext()), &tc.grantMsg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
