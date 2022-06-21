package types

// NewQueryInterchainAccountRequest creates and returns a new QueryInterchainAccountFromZoneRequest
// Zone_name으로 조회하도록 변경(zone name)
func NewQueryInterchainAccountRequest(portID, connectionID string) *QueryInterchainAccountFromZoneRequest {
	return &QueryInterchainAccountFromZoneRequest{
		PortId:       portID,
		ConnectionId: connectionID,
	}
}

// NewQueryInterchainAccountResponse creates and returns a new QueryInterchainAccountFromZoneResponse
func NewQueryInterchainAccountResponse(interchainAccAddr string) *QueryInterchainAccountFromZoneResponse {
	return &QueryInterchainAccountFromZoneResponse{
		InterchainAccountAddress: interchainAccAddr,
	}
}
