# State

---

## CandidatePool
```protobuf
message CandidatePool {
  string pool_id = 1;
  string pool_contract_address = 2;
}
```

`CandidatePool` holds the pool id of the candidate pool and the contract address of the pool. Anyone can register this information.

## IncentivePool
```protobuf
message IncentivePool {
  string pool_id = 1;
  string pool_contract_address = 2;
  uint64 weight = 3;
}
```

`Incentive Pool` is a pool that receives incentives registered through governance among the Candidate pools.
Contains weight, which is information about how much incentives are allocated.