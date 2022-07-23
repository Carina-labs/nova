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
		NewDelegateCmd(),
		NewWithdrawCmd(),
		NewUndelegateRequestCmd(),
		NewUndelegateCmd(),
		NewClaimSnTokenCmd(),
		NewPendingWithdrawCmd(),
	)

	return cmd
}

func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [zone_id] [controller-address] [host-address] [amount] [transfer-port-id] [transfer-channel-id]",
		Short: "Deposit wrapped token to nova",
		Long: `Deposit wrapped token to nova.
Note, the '--from' flag is ignored as it is implied from [from_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[1]); err != nil {
				return err
			}

			zoneId := args[0]

			controllerAddr, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			portId := args[3]
			chanId := args[4]

			msg := types.NewMsgDeposit(zoneId, controllerAddr.GetFromAddress(), coin, portId, chanId)

			return tx.GenerateOrBroadcastTxCLI(controllerAddr, cmd.Flags(), msg)
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
			err := cmd.Flags().Set(flags.FlagFrom, args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			depositor := args[1]
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				panic(fmt.Sprintf("can't parse coin: %s", err.Error()))
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
		Use:  "undelegate [zone-id] [controller-address]",
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
			controllerAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgUndelegate(zoneId, controllerAddr)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [zone-id] [withdrawer] [receiver] [transfer-port-id] [transfer-channel-id]",
		Short: "Withdraw wrapped token to nova",
		Long: `Withdraw bonded token to wrapped-native token.
Note, the '--to' flag is ignored as it is implied from [to_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(5),
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

			withrawer, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			receiver := args[2]
			portId := args[3]
			chanId := args[4]

			msg := types.NewMsgWithdraw(zoneId, withrawer, receiver, portId, chanId)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewClaimSnTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claimsntoken [zone-id] [claimer-address] [transfer-port-id] [transfer-channel-id]",
		Short: "claim wrapped token to nova",
		Args:  cobra.ExactArgs(4),
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
			transferPortId := args[2]
			transferChanId := args[3]

			msg := types.NewMsgClaimSnAsset(zoneId, claimer, transferPortId, transferChanId)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewPendingWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pendingwithdraw [zone-id] [conroller-address] [ica-transfer-port-id] [ica-transfer-channel-id] [block-time]",
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[1]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			daomodifierAddr := clientCtx.GetFromAddress()
			portId := args[2]
			chanId := args[3]
			blockTime := args[4]
			t, err := time.Parse(time.RFC3339, blockTime)
			if err != nil {
				return err
			}

			msg := types.NewMsgPendingWithdraw(zoneId, daomodifierAddr, portId, chanId, t)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
func NewDelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate [zone-id] [controller-address] [transfer-port-id] [transfer-channel-id]",
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
			controllerAddr := clientCtx.GetFromAddress()
			transferPortId := args[2]
			transferChanId := args[3]

			msg := types.NewMsgDelegate(zoneId, controllerAddr, transferPortId, transferChanId)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
