package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewMsgUpdateChainState(t *testing.T) {
	operator := sdk.AccAddress("_______alice________")
	msg := NewMsgUpdateChainState(sdk.NewInt64Coin("nova", 123), operator, 6, 10)
	signers := msg.GetSigners()

	require.Equal(t, 1, len(signers))
	require.Equal(t, signers[0], operator)
}
