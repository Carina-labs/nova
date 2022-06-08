package types

func NewQueryChainStateRequest(chainDenom string) *QueryStateRequest {
	return &QueryStateRequest{
		ChainDenom: chainDenom,
	}
}
