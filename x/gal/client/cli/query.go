package cli

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"
)

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "params",
		Long: "Query for parameter",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryClaimableAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "claimable [zone-id] [address]",
		Long: "Query for claimable snAssets",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.QueryClaimableAmountRequest{
				ZoneId:  args[0],
				Address: args[1],
			}
			res, err := queryClient.ClaimableAmount(ctx, query)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryIcaWithdrawal() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-withdrawal [zone-id] [address]",
		Long: "Query for pending withdrawals",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.QueryPendingWithdrawalsRequest{
				ZoneId:  args[0],
				Address: args[1],
			}
			res, err := queryClient.PendingWithdrawals(ctx, query)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryActiveWithdrawal() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "active-withdrawal [zone-id] [address]",
		Long: "Query for pending withdrawals",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.QueryActiveWithdrawalsRequest{
				ZoneId:  args[0],
				Address: args[1],
			}
			res, err := queryClient.ActiveWithdrawals(ctx, query)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryDepositRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-records [zone-id] [address]",
		Long: "Query for deposit records",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.DepositRecords(cmd.Context(), &types.QueryDepositRecordRequest{
				ZoneId:  args[0],
				Address: addr.String(),
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryDelegateRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate-records [zone-id] [address]",
		Long: "Query for delegate records",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.DelegateRecords(cmd.Context(), &types.QueryDelegateRecordRequest{
				ZoneId:  args[0],
				Address: addr.String(),
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryUndelegateRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate-records [zone-id] [address]",
		Long: "Query for undelegate records",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.UndelegateRecords(cmd.Context(), &types.QueryUndelegateRecordRequest{
				ZoneId:  args[0],
				Address: addr.String(),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryWithdrawRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-records [zone-id] [address]",
		Long: "Query for withdraw records",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.WithdrawRecords(cmd.Context(), &types.QueryWithdrawRecordRequest{
				ZoneId:  args[0],
				Address: addr.String(),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryCurrentDelegateVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate-current-version [zone-id]",
		Long: "Query for delegate current version",
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

			res, err := queryClient.DelegateCurrentVersion(cmd.Context(), &types.QueryCurrentDelegateVersion{
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

func queryCurrentUndelegateVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate-current-version [zone-id]",
		Long: "Query for undelegate current version",
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

			res, err := queryClient.UndelegateCurrentVersion(cmd.Context(), &types.QueryCurrentUndelegateVersion{
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

func queryCurrentWithdrawVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-current-version [zone-id]",
		Long: "Query for withdraw currnet version",
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

			res, err := queryClient.WithdrawCurrentVersion(cmd.Context(), &types.QueryCurrentWithdrawVersion{
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

func queryDelegateVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate-version [zone-id] [version]",
		Long: "Query for delegate version",
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

			res, err := queryClient.DelegateVersion(cmd.Context(), &types.QueryDelegateVersion{
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

func queryUndelegateVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "undelegate-version [zone-id] [version]",
		Long: "Query for undelegate version",
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

			res, err := queryClient.UndelegateVersion(cmd.Context(), &types.QueryUndelegateVersion{
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

func queryWithdrawVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-version [zone-id] [version]",
		Long: "Query for withdraw version",
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

			res, err := queryClient.WithdrawVersion(cmd.Context(), &types.QueryWithdrawVersion{
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

func queryTotalSnAssetSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "total-snasset-supply [zone-id]",
		Long: "Query for total snAsset supply",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalSnAssetSupply(cmd.Context(), &types.QueryTotalSnAssetSupply{
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
