package cli

import (
	"fmt"
	"github.com/Carina-labs/nova/x/poolincentive/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Tx commands for the pool module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CreateCandidatePoolCmd())
	cmd.AddCommand(CreateIncentivePoolCmd())
	cmd.AddCommand(SetPoolWeightCmd())
	cmd.AddCommand(SetMultiplePoolWeightCmd())

	return cmd
}

func CreateCandidatePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create-candidate-pool [pool_id] [pool_contract_address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId := args[0]
			poolContractAddress := args[1]

			msg := types.NewMsgCreateCandidatePool(poolId, poolContractAddress, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func CreateIncentivePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create-incentive-pool [pool_id] [pool_contract_address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId := args[0]
			poolContractAddress := args[1]

			msg := types.NewMsgCreateIncentivePool(poolId, poolContractAddress, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func SetPoolWeightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-pool-weight [pool_id] [new_weight]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId := args[0]
			newWeight, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPoolWeight(poolId, newWeight, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func SetMultiplePoolWeightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-multiple-pool-weight [pool_ids] [weights]",
		Args: cobra.MatchAll(),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolIds := strings.Split(args[0], ",")
			weights := strings.Split(args[1], ",")
			if len(poolIds) != len(weights) {
				return fmt.Errorf("the number of pool id set and weight set must be same")
			}

			var newPoolData []types.NewPoolWeight
			for i := range weights {
				newWeight, err := strconv.ParseUint(weights[i], 10, 64)
				if err != nil {
					return err
				}
				newPoolData = append(newPoolData, types.NewPoolWeight{
					NewWeight: newWeight,
					PoolId:    poolIds[i],
				})
			}

			msg := types.NewMsgSetMultipleWeight(newPoolData, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
