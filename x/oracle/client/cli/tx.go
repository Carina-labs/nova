package cli

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand()

	return cmd
}

func NewUpdateStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update_state [chain_denom] [balance] [decimal] [block_height]",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainDenom := args[0]
			balance, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}

			decimal, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			blockHeight, err := sdk.ParseUint(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateChainState(chainDenom, balance.Uint64(), decimal.Uint64(), blockHeight.Uint64())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
