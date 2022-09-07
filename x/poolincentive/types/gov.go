package types

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeUpdatePoolIncentives  = "UpdatePoolIncentives"
	ProposalTypeReplacePoolIncentives = "ReplacePoolIncentives"
)

// Init registers proposals to update and replace pool incentives.
func init() {
	govtypes.RegisterProposalType(ProposalTypeUpdatePoolIncentives)
	govtypes.RegisterProposalTypeCodec(&UpdatePoolIncentivesProposal{}, "nova/UpdatePoolIncentivesProposal")
	govtypes.RegisterProposalType(ProposalTypeReplacePoolIncentives)
	govtypes.RegisterProposalTypeCodec(&ReplacePoolIncentivesProposal{}, "nova/ReplacePoolIncentivesProposal")
}

var (
	_ govtypes.Content = &UpdatePoolIncentivesProposal{}
	_ govtypes.Content = &ReplacePoolIncentivesProposal{}
)

// NewReplacePoolIncentivesProposal returns a new instance of a replace pool incentives proposal struct.
func NewReplacePoolIncentivesProposal(title, description string, pools []IncentivePool) govtypes.Content {
	return &ReplacePoolIncentivesProposal{
		Title:         title,
		Description:   description,
		NewIncentives: pools,
	}
}

// GetTitle gets the title of the proposal
func (p *ReplacePoolIncentivesProposal) GetTitle() string { return p.Title }

// GetDescription gets the description of the proposal
func (p *ReplacePoolIncentivesProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the router key for the proposal
func (p *ReplacePoolIncentivesProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the proposal
func (p *ReplacePoolIncentivesProposal) ProposalType() string {
	return ProposalTypeReplacePoolIncentives
}

// ValidateBasic validates a governance proposal's abstract and basic contents
func (p *ReplacePoolIncentivesProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	if len(p.NewIncentives) == 0 {
		return fmt.Errorf("there is no incetive pool information")
	}

	for _, pool := range p.NewIncentives {
		if err := pool.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// String returns a string containing the pool incentives proposal.
func (p ReplacePoolIncentivesProposal) String() string {
	return ""
}

// NewUpdatePoolIncentivesProposal returns a new instance of a replace pool incentives proposal struct.
func NewUpdatePoolIncentivesProposal(title, description string, pools []IncentivePool) govtypes.Content {
	return &UpdatePoolIncentivesProposal{
		Title:             title,
		Description:       description,
		UpdatedIncentives: pools,
	}
}

// GetTitle gets the title of the proposal
func (p *UpdatePoolIncentivesProposal) GetTitle() string { return p.Title }

// GetDescription gets the description of the proposal
func (p *UpdatePoolIncentivesProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the router key for the proposal
func (p *UpdatePoolIncentivesProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the proposal
func (p *UpdatePoolIncentivesProposal) ProposalType() string { return ProposalTypeUpdatePoolIncentives }

// ValidateBasic validates a governance proposal's abstract and basic contents.
func (p *UpdatePoolIncentivesProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	if len(p.UpdatedIncentives) == 0 {
		return fmt.Errorf("there is no incetive pool information")
	}

	for _, incentive := range p.UpdatedIncentives {
		if err := incentive.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// String returns a string containing the pool incentives proposal.
func (p UpdatePoolIncentivesProposal) String() string {
	return ""
}
