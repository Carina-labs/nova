# Client

---
## Query
### Candidate Pool
```shell
novad query poolincentive candidate-pool [pool_id]
```

Example:
```shell
novad query poolincentive candidate-pool pool-1
```

Example Output:
```json
{
  "poolId": "pool-1",
  "poolContractAddress": "0xabcd..."
}
```

### All Candidate Pool
```shell
novad query poolincentive all-candidate-pool
```

Example:
```shell
novad query poolincentive all-candidate-pool
```

Example output:
```json
[
  {
    "poolId": "pool-1",
    "poolContractAddress": "0xabcd..."
  },
  {
    "poolId": "pool-2",
    "poolContractAddress": "0x1234..."
  }
]
```

### Incentive Pool
```shell
novad query poolincentive incentive-pool [pool_id]
```

Example:
```shell
novad query poolincentive incentive-pool pool-1
```

Example Output:
```json
{
  "poolId": "pool-1",
  "poolContractAddress": "0xabcd...",
  "weight": 3
}
```

### All Incentive Pool
```shell
novad query poolincentive all-incentive-pool
```

Example:
```shell
novad query poolincentive all-incentive-pool
```

Example output:
```json
[
  {
    "poolId": "pool-1",
    "poolContractAddress": "0xabcd...",
    "weight": 3
  },
  {
    "poolId": "pool-2",
    "poolContractAddress": "0x1234...",
    "weight": 5
  }
]
```

## Tx
### Create Candidate Pool
```shell
novad tx poolincentive create-candidate-pool [pool_id] [pool_contract_address]
```

### Create Incentive Pool
```shell
novad tx poolincentive create-incentive-pool [pool_id] [pool_contract_address]
```

### Set Pool Weight
```shell
novad tx poolincentive set-pool-weight [pool_id] [new_weight]
```

### Set Multiple Pool Weight
```shell
novad tx poolincentive set-multiple-pool-weight [pool_ids] [weights]
```