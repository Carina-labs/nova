package types

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeRegisterZone = "RegisterZone"
)

// Init registers proposals to update and replace pool incentives.
func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterZone)
	govtypes.RegisterProposalTypeCodec(&ZoneRegisterProposal{}, "icacontrol/ZoneRegisterProposal")
}

var (
	_ govtypes.Content = &ZoneRegisterProposal{}
)

// NewRegisterZoneProposal returns a new instance of a replace pool incentives proposal struct.
func NewRegisterZoneProposal(title, description string, zoneInfo ZoneProposalInfo) govtypes.Content {
	return &ZoneRegisterProposal{
		Title:       title,
		Description: description,
		Zone:        zoneInfo,
	}
}

// GetTitle gets the title of the proposal
func (p *ZoneRegisterProposal) GetTitle() string { return p.Title }

// GetDescription gets the description of the proposal
func (p *ZoneRegisterProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the router key for the proposal
func (p *ZoneRegisterProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the proposal
func (p *ZoneRegisterProposal) ProposalType() string {
	return ProposalTypeRegisterZone
}

// ValidateBasic validates a governance proposal's abstract and basic contents
func (p *ZoneRegisterProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if p.Zone.ZoneId == "" {
		return fmt.Errorf("zone id is not nil")
	}

	if p.Zone.BaseDenom == "" {
		return fmt.Errorf("base denom is not nil")
	}

	return nil
}

// String returns a string containing the pool incentives proposal.
func (p ZoneRegisterProposal) String() string {
	return ""
}
