package cli

import (
	"github.com/Carina-labs/nova/x/poolincentive/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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
	cmd.AddCommand(NewUpdatePoolIncentivesProposalCmd())
	cmd.AddCommand(NewReplacePoolIncentivesProposalCmd())
	return cmd
}

func NewUpdatePoolIncentivesProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update-pool-incentives-proposal [pool-id] [contract-address] [weight]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(govcli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(govcli.FlagDescription)
			if err != nil {
				return err
			}

			poolId := strings.Split(args[0], ",")
			contractAddr := strings.Split(args[1], ",")
			weight, err := ParseUint64SliceFromString(args[2], ",")
			if err != nil {
				return err
			}

			var proposal []types.IncentivePool

			for i := 0; i < len(poolId); i++ {
				proposal = append(proposal, types.IncentivePool{
					PoolId:              poolId[i],
					PoolContractAddress: contractAddr[i],
					Weight:              weight[i],
				})
			}

			content := types.NewUpdatePoolIncentivesProposal(title, description, proposal)
			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func NewReplacePoolIncentivesProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace-pool-incentives-proposal [pool-id] [contract-address] [weight]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(govcli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(govcli.FlagDescription)
			if err != nil {
				return err
			}

			poolId := strings.Split(args[0], ",")
			contractAddr := strings.Split(args[1], ",")
			weight, err := ParseUint64SliceFromString(args[2], ",")
			if err != nil {
				return err
			}

			var proposal []types.IncentivePool

			for i := 0; i < len(poolId); i++ {
				proposal = append(proposal, types.IncentivePool{
					PoolId:              poolId[i],
					PoolContractAddress: contractAddr[i],
					Weight:              weight[i],
				})
			}

			content := types.NewReplacePoolIncentivesProposal(title, description, proposal)
			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func ParseUint64SliceFromString(s string, separator string) ([]uint64, error) {
	var parsedInts []uint64
	for _, s := range strings.Split(s, separator) {
		s = strings.TrimSpace(s)

		parsed, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return []uint64{}, err
		}
		parsedInts = append(parsedInts, parsed)
	}
	return parsedInts, nil
}
