package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgUpdateChainStateValidation(t *testing.T) {
	addr := sdk.AccAddress([]byte("addr________________"))

	msg := MsgUpdateChainState{
		Coin:        sdk.NewCoin("atom", sdk.NewInt(1000)),
		Operator:    addr.String(),
		BlockHeight: 10,
		AppHash:     []byte("apphash"),
		ZoneId:      "cosmos",
	}
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// check invalid address
	msg.Operator = "invalid"
	err = msg.ValidateBasic()
	require.Error(t, err)
}
