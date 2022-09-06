# State

---

## ChainInfo
```protobuf
message ChainInfo {
  cosmos.base.v1beta1.Coin coin = 1 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.moretags) = "yaml:\"coin\"",
    (gogoproto.nullable) = false
  ];
  string operator_address = 2 [ (gogoproto.moretags) = "yaml:\"operator_address\"" ];
  int64 last_block_height = 3 [ (gogoproto.moretags) = "yaml:\"last_block_height\"" ];
  bytes app_hash = 4 [ (gogoproto.moretags) = "yaml:\"app_hash\"" ];
  string chain_id = 5 [ (gogoproto.moretags) = "yaml:\"chain_id\"" ];
  uint64 oracle_version = 6;
}
```
`ChainInfo` refers to the state of the counterpart chain to be stored on the Oracle module.
Status includes the amount of coins delegated to Zone, AppHash, and block height.

Each time this message is submitted, the value of the oracle_version must increase sequentially.