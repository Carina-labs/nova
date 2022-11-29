package cli

import (
	"fmt"
	"github.com/Carina-labs/nova/v2/x/gal/types"
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
		txIcaWithdrawCmd(),
	)

	return cmd
}

// GetQueryCmd creates and returns the gal query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the gal module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		queryParams(),
		queryEstimatesnAsset(),
		queryClaimableAmount(),
		queryDepositAmount(),
		queryPendingWithdrawals(),
		queryActiveWithdrawal(),
		queryDepositRecords(),
		queryDelegateRecords(),
		queryUndelegateRecords(),
		queryWithdrawRecords(),
		queryDelegateVersion(),
		queryUndelegateVersion(),
		queryWithdrawVersion(),
		queryCurrentDelegateVersion(),
		queryCurrentUndelegateVersion(),
		queryCurrentWithdrawVersion(),
		queryTotalSnAssetSupply(),
	)

	return cmd
}
