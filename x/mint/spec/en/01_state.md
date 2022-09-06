# State

---

## Minter
```protobuf
message Minter {
  // current annual inflation rate
  string inflation = 1
      [(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec", (gogoproto.nullable) = false];
  // current annual expected provisions
  string annual_provisions = 2 [
    (gogoproto.moretags)   = "yaml:\"annual_provisions\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
}
```

`Minter` represents the minting state.


# Params

---

```protobuf
message Params {
  option (gogoproto.goproto_stringer) = false;
  string mint_denom = 1;
  string inflation_rate_change = 2 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  string inflation_max = 3 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  string inflation_min = 4 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  string goal_bonded = 5 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  uint64 blocks_per_year = 6;
  DistributionProportions distribution_proportions = 7 [
    (gogoproto.nullable) = false 
  ];
}
```

`Params` manages variables related to Nova's new publication. These variables can be changed by governance.
