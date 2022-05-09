package cli

import (
	"fmt"
	"github.com/Carina-labs/novachain/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strings"
)

const (
	FlagDenom = "denom"
)

// GetQueryCmd creates and returns the intertx query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the gal module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetSharesCmd())

	return cmd
}

func GetSharesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shares [address]",
		Short: "Query for account shares by address",
		Long: strings.TrimSpace(fmt.Sprintf(`Query the total shares of an account or of a specific denomination.
Example:
  $ %s query %s shares [address]
  $ %s query %s shares [address] --denom=[denom]`,
			version.AppName, types.ModuleName, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(FlagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			params := types.NewQuerySharesRequest(addr, denom)
			res, err := queryClient.Share(ctx, params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
