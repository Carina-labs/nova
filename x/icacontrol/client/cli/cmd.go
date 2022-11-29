package cli

import (
	"fmt"
	"github.com/Carina-labs/nova/v2/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetTxCmd creates and returns the icacontrol tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		txRegisterZoneCmd(),
		txDelegateTxCmd(),
		txUndelegateTxCmd(),
		txAutoStakingTxCmd(),
		txTransferTxCmd(),
		txDeleteZoneTxCmd(),
		txChangeZoneInfoTxCmd(),
		txAuthzGrantTxCmd(),
		txAuthzRevokeTxCmd(),
		txSetControllerAddrTxCmd(),
	)

	return cmd
}

// GetQueryCmd creates and returns the icacontrol query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the icacontrol module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		queryAllZones(),
		queryZone(),
		queryAutoStakingVersion(),
		queryAutoStakingCurrentVersion(),
	)

	return cmd
}
