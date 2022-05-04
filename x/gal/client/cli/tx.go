package cli

import (
	"fmt"
	"strings"

	"github.com/Carina-labs/novachain/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		getRegisterAccountCmd(),
		getSubmitTxCmd(),
	)

	return cmd
}

func getRegisterAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "register",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterAccount(
				clientCtx.GetFromAddress().String(),
				viper.GetString(FlagConnectionID),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionID)
	_ = cmd.MarkFlagRequired(FlagConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getSubmitTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "submit [path/to/sdk_msg.json]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			var msgs = strings.Split(args[0], "&&")
			txMsgs := make([]sdk.Msg, len(msgs))

			for i, msg := range msgs {
				var txMsg sdk.Msg
				if err := cdc.UnmarshalInterfaceJSON([]byte(msg), &txMsg); err != nil {
					// check for file path if JSON input is not provided
					// contents, err := ioutil.ReadFile(msg)
					if err != nil {
						return errors.Wrap(err, "neither JSON input nor path to .json file for sdk msg were provided")
					}
					// if err := cdc.UnmarshalInterfaceJSON(contents, txMsg); err != nil {
					// 	return errors.Wrap(err, "error unmarshalling sdk msg file")
					// }
				}

				txMsgs[i] = txMsg
			}

			msg, err := types.NewMsgSubmitTx(txMsgs, viper.GetString(FlagConnectionID), clientCtx.GetFromAddress().String())
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionID)
	_ = cmd.MarkFlagRequired(FlagConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
