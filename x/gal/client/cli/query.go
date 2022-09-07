package cli

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "params",
		Long: "Query for parameter",
		Args: cobra.ExactArgs(4),
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
