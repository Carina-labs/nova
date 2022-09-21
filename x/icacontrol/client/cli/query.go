package cli

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func queryAllZones() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "all-zone",
		Long: "Query for all zone",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllZones(cmd.Context(), &types.QueryAllZonesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "zone",
		Long: "Query for zone",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Zone(cmd.Context(), &types.QueryZoneRequest{
				ZoneId:  args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
