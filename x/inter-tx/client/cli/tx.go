package cli

import (
	"fmt"

	"github.com/Carina-labs/nova/x/inter-tx/types"
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
		getAutoStakingTxCmd(),
		getWithdrawTxCmd(),
	)

	return cmd
}

// TODO
func getRegisterZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register [zone-name] [chain-id] [owner-address] [connection-id] [transfer-channel-id] [transfer-connection-id] [transfer-port-id] [validator_address] [base-denom]",
		Args: cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zone_name := args[0]
			chain_id := args[1]
			ica_owner_address := clientCtx.GetFromAddress().String()
			ica_connection_id := args[3]
			transfer_channel_id := args[4]
			transfer_connection_id := args[5]
			transfer_port_id := args[6]
			validator_address := args[7]
			denom := args[8]

			msg := types.NewMsgRegisterZone(zone_name, chain_id, ica_connection_id, ica_owner_address, transfer_channel_id, transfer_connection_id, transfer_port_id, validator_address, denom)

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
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}

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
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneName := args[0]
			sender := args[1]
			ownerAddress := clientCtx.GetFromAddress().String()
			amount, _ := sdk.ParseCoinNormalized(args[3])

			msg := types.NewMsgIcaUnDelegate(zoneName, sender, ownerAddress, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getAutoStakingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "autostaking [zone-name] [sender(host-address)] [owner-address] [amount]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneName := args[0]
			sender := args[1]
			owner := clientCtx.GetFromAddress().String()
			amount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaAutoStaking(zoneName, sender, owner, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getWithdrawTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw [zone-name] [sender-address(host-address)] [owner-address] [reveiver] [amount]",
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneName := args[0]
			sender := args[1]
			owner := clientCtx.GetFromAddress().String()
			receiver := args[3]
			amount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaWithdraw(zoneName, sender, owner, receiver, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
