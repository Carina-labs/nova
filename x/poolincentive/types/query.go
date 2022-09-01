package types

func NewQuerySingleCandidatePool(poolId string) *QuerySingleCandidatePool {
	return &QuerySingleCandidatePool{
		PoolId: poolId,
	}
}

func NewQueryAllCandidatePool() *QueryAllCandidatePool {
	return &QueryAllCandidatePool{}
}

func NewQuerySingleIncentivePool(poolId string) *QuerySingleIncentivePool {
	return &QuerySingleIncentivePool{
		PoolId: poolId,
	}
}

func NewQueryAllIncentivePool() *QueryAllIncentivePool {
	return &QueryAllIncentivePool{}
}
