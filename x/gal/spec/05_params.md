# Parameters

---
```protobuf
message Params {
  repeated string sn_token_denoms = 1;
  required string mint_denom = 2;
  required uint64 blocks_per_year = 3;
}
```

The parameter manages the rate of block creation and the whitelist chain list.
