package cli

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewUpdateStateCmd())

	return cmd
}

func NewUpdateStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update_state [amount] [block-height] [app-hash] [chain-id]",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			blockHeight, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			appHash := []byte(args[2])
			chainId := args[3]

			msg := types.NewMsgUpdateChainState(clientCtx.GetFromAddress(), chainId, amount, blockHeight, appHash)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
