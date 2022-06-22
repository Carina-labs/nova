package cli

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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

	cmd.AddCommand(NewUpdateStateCmd())

	return cmd
}

func NewUpdateStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update_state [from_key_or_address] [chain_denom] [balance] [decimal] [block_height] [app_hash] [chain_id] [block_proposer]",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainDenom := args[1]
			balance, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			decimal, err := sdk.ParseUint(args[3])
			if err != nil {
				return err
			}

			blockHeight, err := sdk.ParseUint(args[4])
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &types.MsgUpdateChainState{
				Coin:          sdk.NewCoin(chainDenom, sdk.NewIntFromBigInt(balance.BigInt())),
				Operator:      clientCtx.GetFromAddress().String(),
				Decimal:       decimal.Uint64(),
				BlockHeight:   blockHeight.Uint64(),
				AppHash:       args[5],
				ChainId:       args[6],
				BlockProposer: args[7],
			})
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
