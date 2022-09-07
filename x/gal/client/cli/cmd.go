package cli

import (
	"fmt"
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetTxCmd creates and returns the gal tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		txDepositCmd(),
		txDelegateCmd(),
		txWithdrawCmd(),
		txUndelegateRequestCmd(),
		txUndelegateCmd(),
		txClaimSnAssetCmd(),
		txPendingWithdrawCmd(),
	)

	return cmd
}

// GetQueryCmd creates and returns the ibcstaking query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the gal module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(queryParams())
	cmd.AddCommand(queryClaimableAsset())
	cmd.AddCommand(queryIcaWithdrawal())
	cmd.AddCommand(queryActiveWithdrawal())
	cmd.AddCommand(queryDepositRecords())
	cmd.AddCommand(queryUndelegateRecords())
	cmd.AddCommand(queryWithdrawRecords())

	return cmd
}
