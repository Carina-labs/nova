package cli

import (
	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"strconv"
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
		Use:  "zone [zone-id]",
		Long: "Query for zone id",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Zone(cmd.Context(), &types.QueryZoneRequest{
				ZoneId: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryAutoStakingVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "autostaking-version [zone-id] [version]",
		Long: "Query for autostaking version",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			version, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.AutoStakingVersion(cmd.Context(), &types.QueryAutoStakingVersion{
				ZoneId:  args[0],
				Version: version,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryAutoStakingCurrentVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "autostaking-current-version [zone-id]",
		Long: "Query for autostaking current veresion",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			if err != nil {
				return err
			}

			res, err := queryClient.AutoStakingCurrentVersion(cmd.Context(), &types.QueryCurrentAutoStakingVersion{
				ZoneId: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
