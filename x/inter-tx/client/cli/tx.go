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
		getHostAddressTxCmd(),
	)

	return cmd
}

func getRegisterZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register [zone-name] [chain-id] [controller-address] [connection-id] [transfer-channel-id] [transfer-connection-id] [transfer-port-id] [validator_address] [base-denom]",
		Args: cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneName := args[0]
			chainId := args[1]
			icaControllerAddr := clientCtx.GetFromAddress().String()
			icaConnId := args[3]
			transferChanId := args[4]
			transferConnId := args[5]
			transferPortId := args[6]
			validatorAddr := args[7]
			denom := args[8]

			msg := types.NewMsgRegisterZone(zoneName, chainId, icaConnId, icaControllerAddr, transferChanId, transferConnId, transferPortId, validatorAddr, denom)

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
		Use:  "delegate [zone-name] [sender(host-address)] [controller-address] [amount]",
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
			controllerAddr := clientCtx.GetFromAddress().String()
			amount, err := sdk.ParseCoinNormalized(args[3])

			if err != nil {
				panic("coin error")
			}

			msg := types.NewMsgIcaDelegate(zoneName, sender, controllerAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getUndelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate [zone-name] [sender(host-address)] [controller-address] [amount]",
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
			controllerAddr := clientCtx.GetFromAddress().String()
			amount, _ := sdk.ParseCoinNormalized(args[3])

			msg := types.NewMsgIcaUnDelegate(zoneName, sender, controllerAddr, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getAutoStakingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "autostaking [zone-name] [sender(host-address)] [controller-address] [amount]",
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
			controllerAddr := clientCtx.GetFromAddress().String()
			amount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaAutoStaking(zoneName, sender, controllerAddr, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getWithdrawTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw [zone-name] [sender-address(host-address)] [controller-address] [reveiver] [amount]",
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
			controllerAddr := clientCtx.GetFromAddress().String()
			receiver := args[3]
			amount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaWithdraw(zoneName, sender, controllerAddr, receiver, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getHostAddressTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "registerhostaddress [controller-address] [host-address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			controllerAddr := args[0]
			hostAddr := args[1]

			msg := types.NewMsgRegisterHostAccount(controllerAddr, hostAddr)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
