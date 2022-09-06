# Client

---

## Query
### State
```shell
novad query oracle query state [chain_name]
```

Example:
```shell
novad query oracle query state cosmoshub-1
```

Example Output:
```json
{
  "coin": {
    "denom": "uatom",
    "amount": 100000
  },
  "operator": "nova1a2b...",
  "decimal": 8,
  "lastBlockHeight": 1024,
  "apphash": "ABCD...",
  "chainId": "cosmoshub-1"
}
```

---
## Tx
## Update State
```shell
novad tx oracle update_state [from_key_or_address] [amount] [block_height] [app_hash] [chain_id]
```

Example:
```shell
novad tx oracle update_state nova1a2b... 10000uatom 1024 ABCD... cosmoshub-1
```