package cli

import (
	"fmt"
	"strconv"

	"github.com/Carina-labs/nova/x/icacontrol/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/x/authz"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagSpendLimit        = "spend-limit"
	FlagMsgType           = "msg-type"
	FlagExpiration        = "expiration"
	FlagAllowedValidators = "allowed-validators"
	FlagDenyValidators    = "deny-validators"
	FlagAllowList         = "allow-list"
	delegate              = "delegate"
	redelegate            = "redelegate"
	unbond                = "unbond"
)

// txRegisterZoneCmd is a transaction that registers new Zone information. This transaction can only be submitted by a given signatory.
func txRegisterZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register-zone [zone-id] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal] [max-entries]",
		Args: cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			icaConnId := args[1]
			transferPortId := args[2]
			transferChanId := args[3]
			validatorAddr := args[4]
			denom := args[5]
			decimal, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return err
			}

			maxEntries, err := strconv.ParseInt(args[7], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterZone(zoneId, icaConnId, clientCtx.GetFromAddress(), transferPortId, transferChanId, validatorAddr, denom, decimal, maxEntries)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txDelegateTxCmd is a transaction used for remote delegation using ICA. This transaction can only be submitted by a given signatory.
func txDelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-delegate [zone-id] [host-address] [amount]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			hostAddr := args[1]
			amount, err := sdk.ParseCoinNormalized(args[2])

			if err != nil {
				panic("coin error")
			}

			msg := types.NewMsgIcaDelegate(zoneId, clientCtx.GetFromAddress(), hostAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txUndelegateTxCmd is a transaction used for remote de-delegation using ICA. This transaction can only be submitted by a given signatory.
func txUndelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-undelegate [zone-id] [host-address] [amount]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			hostAddr := args[1]
			amount, _ := sdk.ParseCoinNormalized(args[2])

			msg := types.NewMsgIcaUnDelegate(zoneId, hostAddr, clientCtx.GetFromAddress(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txAutoStakingTxCmd is a transaction used for auto-compounding using ICA. This transaction can only be submitted by a given signatory.
func txAutoStakingTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-auto-staking [zone-id] [amount]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaAutoStaking(zoneId, clientCtx.GetFromAddress(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txTransferTxCmd is a transaction used to transfer assets between chains using ICA. This transaction can only be submitted by a given signatory.
func txTransferTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-transfer [zone-id] [host-address] [receiver] [ica-transfer-port-id] [ica-transfer-channel-id] [amount]",
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]
			hostAddr := args[1]
			receiver := args[2]
			portId := args[3]
			chanId := args[4]
			amount, err := sdk.ParseCoinNormalized(args[5])
			if err != nil {
				return err
			}

			msg := types.NewMsgIcaTransfer(zoneId, hostAddr, clientCtx.GetFromAddress(), receiver, portId, chanId, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txDeleteZoneTxCmd is a transaction that deletes the registered zone. This transaction can only be submitted by a given signatory.
func txDeleteZoneTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete-zone [zone-id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			zoneId := args[0]

			msg := types.NewMsgDeleteRegisteredZone(zoneId, clientCtx.GetFromAddress())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txChangeZoneInfoTxCmd is a transaction that modifies the registered zone. This transaction can only be submitted by a given signatory.
func txChangeZoneInfoTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "change-zone [zone-id] [host-address] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal] [max_entries]",
		Args: cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			hostAddr := args[1]
			icaConnId := args[2]
			transferPortId := args[3]
			transferChanId := args[4]
			validatorAddr := args[5]
			denom := args[6]
			decimal, err := strconv.ParseInt(args[7], 10, 64)
			if err != nil {
				return err
			}

			maxEntries, err := strconv.ParseInt(args[8], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgChangeZoneInfo(zoneId, hostAddr, clientCtx.GetFromAddress(), icaConnId, transferPortId, transferChanId, validatorAddr, denom, decimal, maxEntries)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// txAuthzGrantTxCmd is a transaction used to transfer permissions between accounts using ICA. This transaction can only be submitted by a given signatory.
func txAuthzGrantTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ica-authz-grant [zone-id] [grantee-address] [authorization-type] --from [granter]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			grantee := args[1]

			exp, err := cmd.Flags().GetInt64(FlagExpiration)
			if err != nil {
				return err
			}

			controllerAddr := clientCtx.GetFromAddress()

			var authorization authz.Authorization
			switch args[2] {
			case "send":
				limit, err := cmd.Flags().GetString(FlagSpendLimit)
				if err != nil {
					return err
				}

				spendLimit, err := sdk.ParseCoinsNormalized(limit)
				if err != nil {
					return err
				}

				if !spendLimit.IsAllPositive() {
					return fmt.Errorf("spend-limit should be greater than zero")
				}

				authorization = bank.NewSendAuthorization(spendLimit)

			case "generic":
				msgType, err := cmd.Flags().GetString(FlagMsgType)
				if err != nil {
					return err
				}

				authorization = authz.NewGenericAuthorization(msgType)
			case delegate, unbond, redelegate:
				limit, err := cmd.Flags().GetString(FlagSpendLimit)
				if err != nil {
					return err
				}

				allowValidators, err := cmd.Flags().GetStringSlice(FlagAllowedValidators)
				if err != nil {
					return err
				}

				denyValidators, err := cmd.Flags().GetStringSlice(FlagDenyValidators)
				if err != nil {
					return err
				}

				var delegateLimit *sdk.Coin
				if limit != "" {
					spendLimit, err := sdk.ParseCoinsNormalized(limit)
					if err != nil {
						return err
					}

					if !spendLimit.IsAllPositive() {
						return fmt.Errorf("spend-limit should be greater than zero")
					}
					delegateLimit = &spendLimit[0]
				}

				allowed, err := bech32toValidatorAddresses(allowValidators)
				if err != nil {
					return err
				}

				denied, err := bech32toValidatorAddresses(denyValidators)
				if err != nil {
					return err
				}

				switch args[1] {
				case delegate:
					authorization, err = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, delegateLimit)
				case unbond:
					authorization, err = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, delegateLimit)
				default:
					authorization, err = staking.NewStakeAuthorization(allowed, denied, staking.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, delegateLimit)
				}
				if err != nil {
					return err
				}

			default:
				return fmt.Errorf("invalid authorization type, %s", args[1])
			}

			msg, err := types.NewMsgAuthzGrant(zoneId, grantee, controllerAddr, authorization, time.Unix(exp, 0))
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagMsgType, "", "The Msg method name for which we are creating a GenericAuthorization")
	cmd.Flags().String(FlagSpendLimit, "", "SpendLimit for Send Authorization, an array of Coins allowed spend")
	cmd.Flags().StringSlice(FlagAllowedValidators, []string{}, "Allowed validators addresses separated by ,")
	cmd.Flags().StringSlice(FlagDenyValidators, []string{}, "Deny validators addresses separated by ,")
	cmd.Flags().StringSlice(FlagAllowList, []string{}, "Allowed addresses grantee is allowed to send funds separated by ,")
	cmd.Flags().Int64(FlagExpiration, time.Now().AddDate(2, 0, 0).Unix(), "The Unix timestamp. Default is one year.")
	return cmd
}

// txAuthzRevokeTxCmd is a transaction used to execute transferred permissions between accounts using ICA. This transaction can only be submitted by a given signatory.
func txAuthzRevokeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ica-authz-revoke [zone-id] [grantee-address] [msg_type]",

		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			zoneId := args[0]
			grantee := args[1]
			msgType := args[2]
			controllerAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgAuthzRevoke(zoneId, grantee, msgType, controllerAddr)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
func bech32toValidatorAddresses(validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
}
