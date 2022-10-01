package cli

import (
	"github.com/Carina-labs/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetStateCmd())
	cmd.AddCommand(GetVersionCmd())

	return cmd
}

func GetStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "state [chain-denom]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chainDenom := args[0]
			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQueryChainStateRequest(chainDenom)
			res, err := queryClient.State(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func GetVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "oracle-version [zone-id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQueryOracleVersionRequest(zoneId)
			res, err := queryClient.OracleVersion(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
