package cli

import (
	"github.com/Carina-labs/nova/x/poolincentive/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the pool module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetSingleCandidatePool())
	cmd.AddCommand(GetAllCandidatePool())
	cmd.AddCommand(GetSingleIncentivePool())
	cmd.AddCommand(GetAllIncentivePool())

	return cmd
}

func GetSingleCandidatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "candidate-pool [pool_id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId := args[0]
			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQuerySingleCandidatePool(poolId)
			res, err := queryClient.SingleCandidatePool(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func GetAllCandidatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "all-candidate-pool",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQueryAllCandidatePool()
			res, err := queryClient.AllCandidatePool(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func GetSingleIncentivePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "incentive-pool [pool_id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId := args[0]
			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQuerySingleIncentivePool(poolId)
			res, err := queryClient.SingleIncentivePool(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func GetAllIncentivePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "all-incentive-pool",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			ctx := cmd.Context()

			msg := types.NewQueryAllIncentivePool()
			res, err := queryClient.AllIncentivePool(ctx, msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}