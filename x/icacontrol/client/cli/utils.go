package cli

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"strings"
)

func bech32toValidatorAddresses(prefix string, validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := ValAddressFromBech32(prefix, validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
}

func ValAddressFromBech32(prefix string, address string) (addr sdk.ValAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.ValAddress{}, errors.New("empty address string is not allowed")
	}

	bz, err := sdk.GetFromBech32(address, prefix)
	if err != nil {
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func NewStakeAuthorization(prefix string, allowed []sdk.ValAddress, denied []sdk.ValAddress, authzType staking.AuthorizationType, amount *sdk.Coin) (*staking.StakeAuthorization, error) {
	allowedValidators, deniedValidators, err := validateAndBech32fy(prefix, allowed, denied)
	if err != nil {
		return nil, err
	}

	a := staking.StakeAuthorization{}
	if allowedValidators != nil {
		a.Validators = &staking.StakeAuthorization_AllowList{AllowList: &staking.StakeAuthorization_Validators{Address: allowedValidators}}
	} else {
		a.Validators = &staking.StakeAuthorization_DenyList{DenyList: &staking.StakeAuthorization_Validators{Address: deniedValidators}}
	}

	if amount != nil {
		a.MaxTokens = amount
	}
	a.AuthorizationType = authzType

	return &a, nil
}

func validateAndBech32fy(prefix string, allowed []sdk.ValAddress, denied []sdk.ValAddress) ([]string, []string, error) {
	if len(allowed) == 0 && len(denied) == 0 {
		return nil, nil, sdkerrors.ErrInvalidRequest.Wrap("both allowed & deny list cannot be empty")
	}

	if len(allowed) > 0 && len(denied) > 0 {
		return nil, nil, sdkerrors.ErrInvalidRequest.Wrap("cannot set both allowed & deny list")
	}

	if prefix == "" {
		return nil, nil, sdkerrors.ErrInvalidRequest.Wrap("cannot set validator address prefix")
	}

	allowedValidators := make([]string, len(allowed))
	if len(allowed) > 0 {
		for i, validator := range allowed {
			validatorStr, err := bech32.ConvertAndEncode(prefix, validator)
			if err != nil {
				return nil, nil, err
			}
			allowedValidators[i] = validatorStr
		}
		return allowedValidators, nil, nil
	}

	deniedValidators := make([]string, len(denied))
	for i, validator := range denied {
		validatorStr, err := bech32.ConvertAndEncode(prefix, validator)
		if err != nil {
			return nil, nil, err
		}
		deniedValidators[i] = validatorStr
	}

	return nil, deniedValidators, nil
}
