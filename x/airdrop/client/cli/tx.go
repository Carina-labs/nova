package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Carina-labs/nova/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func txAirdropData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "airdrop-data [user_state_json_file]",
		Short: "Enter user state data",
		Long:  "Enter the user's airdrop information. Data can only be added by the controller account.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			userState, err := ReadUserStateFromFile(args[0])
			if err != nil {
				return err
			}

			controllerAddr := context.GetFromAddress().String()

			msg := &types.MsgAirdropDataRequest{
				States:            userState,
				ControllerAddress: controllerAddr,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(context, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

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
			context, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			questType, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid integer for quest type")
			}

			// check questType is smaller than maximum int32
			if questType.Int64() < 0 || questType.Int64() > 4 {
				return fmt.Errorf("invalid quest type")
			}

			msg := &types.MsgClaimAirdropRequest{
				UserAddress: context.GetFromAddress().String(),
				QuestType:   types.QuestType(int32(questType.Int64())),
			}

			return tx.GenerateOrBroadcastTxCLI(context, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txMarkSocialQuest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-social-quest [user-address]",
		Short: "Mark that the user has completed social quest",
		Long: `Mark that the user has completed social quest.
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

			_, err = sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgMarkSocialQuestPerformedRequest{
				ControllerAddress: context.GetFromAddress().String(),
				UserAddresses:     []string{args[0]},
			}

			return tx.GenerateOrBroadcastTxCLI(context, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txMarkProvideLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-provide-liquidity [user-address]",
		Short: "Mark that the user has completed providing liquidity quest",
		Long: `Mark that the user has completed providing liquidity quest.
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

			_, err = sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgMarkUserProvidedLiquidityRequest{
				ControllerAddress: context.GetFromAddress().String(),
				UserAddresses:     []string{args[0]},
			}

			return tx.GenerateOrBroadcastTxCLI(context, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func ReadUserStateFromFile(filename string) ([]*types.UserState, error) {
	var userState []*types.UserState
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &userState)
	if err != nil {
		return nil, err
	}

	return userState, nil
}
