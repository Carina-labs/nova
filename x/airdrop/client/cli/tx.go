package cli

import (
	"fmt"
	"strconv"

	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func txClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [quest-type]",
		Short: "Claim airdrop token",
		Long: `Claim airdrop token.

Quest type is range from 0 to 4
Each quest is defined as follow:

QuestType_QUEST_NOTHING_TO_DO     QuestType = 0
QuestType_QUEST_SOCIAL            QuestType = 1
QuestType_QUEST_SN_ASSET_CLAIM    QuestType = 2
QuestType_QUEST_PROVIDE_LIQUIDITY QuestType = 3
QuestType_QUEST_VOTE_ON_PROPOSALS QuestType = 4

Note, the '--from' flag is ignored as it is implied from [from_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			context, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			questType, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			_, ok := types.QuestType_name[int32(questType)]
			if !ok {
				return fmt.Errorf("quest type %d is not supported", questType)
			}

			msg := &types.MsgClaimAirdropRequest{
				UserAddress: context.GetFromAddress().String(),
				QuestType:   types.QuestType(int32(questType)),
			}

			return tx.GenerateOrBroadcastTxCLI(context, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
