# State

---
## Shares

---
Indicates the stake of the user deposited through the GAL module.
This means the user's share of the total assets deposited through Nova to the host chain validator.

### DepositAccount
```protobuf
message DepositAccount {
  string denom = 1;
  repeated DepositInfo depositInfos = 2;
  int64 totalShare = 3;
  int64 lastBlockUpdate = 4;
}
```
`DepositAccount` specifies the accounts deposited, the destination chain, the entire stake, and the last updated block height.


### DepositInfo
```protobuf
message DepositInfo {
  string address = 1;
  int64 share = 2;
  int64 debt = 3;
}
```
`DepositInfo` specifies the user's deposit information.