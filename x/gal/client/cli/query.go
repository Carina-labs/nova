package cli

import (
	"fmt"

	"strings"

	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

const (
	FlagDenom = "denom"
)

func queryClaimableAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "claimable_asset [zone_id] [address] [transfer_port_id] [transfer_channel_id]",
		Long: "Query for claimable snAssets",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.ClaimableAmountRequest{
				ZoneId:            args[0],
				Address:           args[1],
				TransferPortId:    args[2],
				TransferChannelId: args[3],
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

func queryPendingWithdrawals() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pending_withdrawals [zone_id] [address] [transfer_port_id] [transfer_channel_id]",
		Long: "Query for pending withdrawals",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.PendingWithdrawalsRequest{
				ZoneId:            args[0],
				Address:           args[1],
				TransferPortId:    args[2],
				TransferChannelId: args[3],
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

func queryActiveWithdrawals() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "active_withdrawals [zone_id] [address] [transfer_port_id] [transfer_channel_id]",
		Long: "Query for pending withdrawals",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			query := &types.ActiveWithdrawalsRequest{
				ZoneId:            args[0],
				Address:           args[1],
				TransferPortId:    args[2],
				TransferChannelId: args[3],
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
		Use:  "deposit_records [zone_id] [address]",
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
		Use:  "undelegate_records [zone_id] [address]",
		Long: "Query for undelegate records",
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
		Use: "withdraw_records [zone_id] [address]",
		Long: strings.TrimSpace(fmt.Sprintf(`Query withdraw history of an account or of a specific denomination.
Example:
	$ %s query %s withdraw [address]
	$ %s query %s withdraw [address] --denom=[denom]`, version.AppName, types.ModuleName, version.AppName, types.ModuleName)),
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
