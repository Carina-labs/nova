package types

func NewQueryChainStateRequest(chainDenom string) *QueryStateRequest {
	return &QueryStateRequest{
		ChainDenom: chainDenom,
	}
}

func NewQueryOracleVersionRequest(zoneId string) *QueryOracleVersionRequest {
	return &QueryOracleVersionRequest{
		ZoneId: zoneId,
	}
}
