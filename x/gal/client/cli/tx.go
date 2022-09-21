package cli

import (
	"fmt"
	"time"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func txDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [zone-id] [depositor] [claimer] [amount]",
		Short: "Deposit wrapped token to nova",
		Long: `Deposit wrapped token to nova.
Note, the '--from' flag is ignored as it is implied from [from_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[1]); err != nil {
				return err
			}

			zoneId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			claimer, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(zoneId, clientCtx.GetFromAddress(), claimer, coin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUndelegateRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pending-undelegate [zone-id] [delegator] [withdrawer] [amount]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			delegator, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			withdrawer, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				panic(fmt.Sprintf("can't parse coin: %s", err.Error()))
			}

			msg := types.NewMsgPendingUndelegate(zoneId, delegator, withdrawer, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUndelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate [zone-id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]

			msg := types.NewMsgUndelegate(zoneId, clientCtx.GetFromAddress())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [zone-id] [withdrawer]",
		Short: "Withdraw wrapped token to nova",
		Long: `Withdraw bonded token to wrapped-native token.
Note, the '--to' flag is ignored as it is implied from [to_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]

			withdrawer, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(zoneId, withdrawer)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txClaimSnAssetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [zone-id] [claimer-address]",
		Short: "claim wrapped coin to nova",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			claimer := clientCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimSnAsset(zoneId, claimer)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func txPendingWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pending-withdraw [zone-id] [ica-transfer-port-id] [ica-transfer-channel-id] [block-time]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[1]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			portId := args[1]
			chanId := args[2]
			blockTime := args[3]
			t, err := time.Parse(time.RFC3339, blockTime)
			if err != nil {
				return err
			}

			msg := types.NewIcaWithdraw(zoneId, clientCtx.GetFromAddress(), portId, chanId, t)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
func txDelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate [zone-id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]

			msg := types.NewMsgDelegate(zoneId, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
