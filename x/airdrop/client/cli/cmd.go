package cli

import (
	"github.com/Carina-labs/nova/v2/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Airdrop transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(txClaim())
	cmd.AddCommand(txMarkSocialQuest())
	cmd.AddCommand(txMarkProvideLiquidity())

	return cmd
}

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Airdrop query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(queryAirdropInfo())
	cmd.AddCommand(queryTotalAirdropToken())
	cmd.AddCommand(queryQuestState())

	return cmd
}
