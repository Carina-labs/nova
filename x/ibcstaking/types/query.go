package types

// NewQueryInterchainAccountRequest creates and returns a new QueryInterchainAccountFromZoneRequest
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
