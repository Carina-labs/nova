package cli

import (
	"fmt"

	"github.com/Carina-labs/novachain/x/inter-tx/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd creates and returns the intertx tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getRegisterZoneCmd(),
		getDelegateTxCmd(),
		getUndelegateTxCmd(),
	)

	return cmd
}

// TODO
func getRegisterZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register [zone-name] [chain-id] [owner-address] [connection-id] [validator_address] [denom]",
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zone_name := args[0]
			chain_id := args[1]
			owner_address := clientCtx.GetFromAddress().String()
			connection_id := args[3]
			validator_address := args[4]
			denom := args[5]

			msg := types.NewMsgRegisterZone(zone_name, chain_id, connection_id, owner_address, validator_address, denom)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getDelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate [zone-name] [sender(host-address)] [owner-address] [amount]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zone_name := args[0]
			sender := args[1]
			owner_address := clientCtx.GetFromAddress().String()
			amount, err := sdk.ParseCoinNormalized(args[3])

			if err != nil {
				panic("coin error")
			}

			msg := types.NewMsgIcaDelegate(zone_name, sender, owner_address, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getUndelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate [zone_name] [sender(host-address)] [owner-address] [amount]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[2])
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zone_name := args[0]
			sender := args[1]
			owner_address := clientCtx.GetFromAddress().String()
			amount, _ := sdk.ParseCoinNormalized(args[3])

			msg := types.NewMsgIcaUnDelegate(zone_name, sender, owner_address, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
