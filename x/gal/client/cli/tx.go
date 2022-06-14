package cli

import (
	"fmt"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetTxCmd creates and returns the gal tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewDepositCmd(),
		// NewWithdrawCmd(),
		NewUndelegateRequestCmd(),
		NewUndelegateCmd(),
	)

	return cmd
}

func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [from_key_or_address] [to_address] [amount] [channel]",
		Short: "Deposit wrapped token to nova",
		Long: `Deposit wrapped token to nova.
Note, the '--from' flag is ignored as it is implied from [from_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			channel := args[2]

			msg := types.NewMsgDeposit(clientCtx.GetFromAddress(), coins, channel)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewUndelegateRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegaterequest [zone-id] [depositor] [amount]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[1])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			depositor := args[1]
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				panic("coin error")
			}

			msg := types.NewMsgUndelegateRecord(zoneId, depositor, amount)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewUndelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate [zone-id] [controller-address] [host-address]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[1])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			controllerAddr := args[1]
			hostAddr := args[2]

			msg := types.NewMsgUndelegate(zoneId, controllerAddr, hostAddr)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// func NewWithdrawCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "withdraw [to_key_or_address] [amount]",
// 		Short: "Withdraw wrapped token to nova",
// 		Long: `Withdraw bonded token to wrapped-native token.
// Note, the '--to' flag is ignored as it is implied from [to_key_or_address].
// When using '--dry-run' a key name cannot be used, only a bech32 address.`,
// 		Args: cobra.ExactArgs(0),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cmd.Flags().Set("to", args[0])
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			toAddr, err := sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			coins, err := sdk.ParseCoinsNormalized(args[1])
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgWithdraw(toAddr, coins)

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	return cmd
// }
