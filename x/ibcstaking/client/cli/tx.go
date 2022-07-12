package cli

import (
	"fmt"

	"github.com/Carina-labs/nova/x/ibcstaking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd creates and returns the ibcstaking tx command
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
		getTransferTxCmd(),
		getHostAddressTxCmd(),
		getDeleteZoneTxCmd(),
		getChangeZoneInfoTxCmd(),
	)

	return cmd
}

func getRegisterZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register [zone-id] [daomodifier-address] [connection-id] [validator_address] [base-denom]",
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
			icaConnId := args[2]
			validatorAddr := args[3]
			denom := args[4]

			msg := types.NewMsgRegisterZone(zoneId, icaConnId, daomodifierAddr, validatorAddr, denom)

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
		Use:  "delegate [zone-id] [daomodifier-address] [host-address] [amount]",
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
			daomodifierAddr := clientCtx.GetFromAddress().String()
			hostAddr := args[2]
			amount, err := sdk.ParseCoinNormalized(args[3])

			if err != nil {
				panic("coin error")
			}

			msg := types.NewMsgIcaDelegate(zoneId, daomodifierAddr, hostAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getUndelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate [zone-id] [daomodifier-address] [host-address] [amount]",
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
			daomodifierAddr := clientCtx.GetFromAddress()
			hostAddr := args[2]
			amount, _ := sdk.ParseCoinNormalized(args[3])

			msg := types.NewMsgIcaUnDelegate(zoneId, hostAddr, daomodifierAddr, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getAutoStakingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "autostaking [zone-id] [daomodifier-address] [host-address] [amount]",
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
			daomodifierAddr := clientCtx.GetFromAddress().String()
			hostAddr := args[2]
			amount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaAutoStaking(zoneId, hostAddr, daomodifierAddr, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getTransferTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "icatransfer [zone-id] [daomodifier-address] [host-address] [receiver] [ica-transfer-port-id] [ica-transfer-channel-id] [amount]",
		Args: cobra.ExactArgs(7),
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
			hostAddr := args[2]
			receiver := args[3]
			portId := args[4]
			chanId := args[5]
			amount, err := sdk.ParseCoinNormalized(args[6])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaTransfer(zoneId, hostAddr, daomodifierAddr, receiver, portId, chanId, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getHostAddressTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "registerhostaddress [zone-id] [host-address] [daomodifier-address]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[2]); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			hostAddr := args[1]
			daomodifierAddr := args[2]

			msg := types.NewMsgRegisterHostAccount(zoneId, hostAddr, daomodifierAddr)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getDeleteZoneTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deletezone [zone-id] [daomodifier-address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[1]); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			daomodifier := args[1]

			msg := types.NewMsgDeleteRegisteredZone(zoneId, daomodifier)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//TODO: add host address
func getChangeZoneInfoTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "changezoneinfo [zone-id] [daomodifier-address] [connection-id] [validator_address] [base-denom]",
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
			icaConnId := args[2]
			validatorAddr := args[3]
			denom := args[4]

			msg := types.NewMsgChangeZoneInfo(zoneId, icaConnId, daomodifierAddr, validatorAddr, denom)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
