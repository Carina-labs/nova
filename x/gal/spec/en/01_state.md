# State
---

## Overview

---


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


### WithdrawInfo
```protobuf
 message WithdrawInfo {
  string address = 1;
  string denom = 2;
  int64 amount = 3;
  google.protobuf.Timestamp completion_time = 4[(gogoproto.nullable) = false, (gogoproto.stdtime)= true];
}
```
`WithdrawInfo` shows the information about the withdrawal you applied for and when it was completed.
Supernova collects withdrawal requests and processes them once every few days, so there is a completion time.


### RecordInfo
```protobuf
message RecordInfo {
  string zone_id = 1;
  uint64 delegate_version = 2;
  uint64 undelegate_version = 3;
  uint64 withdraw_version = 4;
}
```
All Supernova delegation, underegulation, and withdraw are processed in batch form and executed by an external bot.
Therefore, the GAL module manages the version of the last job executed by the bot.

### DepositRecord
```protobuf
message DepositRecord {
  string zone_id = 1;
  string claimer = 2;
  repeated DepositRecordContent records = 3;
}
```
`DepositRecord` manages historical data that you deposit in a particular zone.
sn-tokens are paid to the claimer.


### DepositRecordContent
```protobuf
message DepositRecordContent {
  string depositor = 1;
  cosmos.base.v1beta1.Coin amount = 2;
  int64 state = 3;
  uint64 oracle_version = 4;
  uint64 delegate_version = 5;
}
```
`DepositRecordContent` stores the user's Deposit details.

### UndelegateRecord
```protobuf
message UndelegateRecord {
  string zone_id = 1;
  string delegator = 2;
  repeated UndelegateRecordContent records = 3;
}
```
`UndelegateRecord` manages the history of the user's undelegation requests.

### UndelegateRecordContent
```protobuf
message UndelegateRecordContent {
  string withdrawer = 1;
  cosmos.base.v1beta1.Coin sn_asset_amount = 2;
  string withdraw_amount = 3[(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int", (gogoproto.nullable) = false];
  int64 state = 4;
  uint64 oracle_version = 5;
  uint64 undelegate_version = 6;
}
```
UndelegateRecordContent manages information about each of the user's undelegated requests.

Save the Oracle version at the time of the undelegation request and use it to calculate the amount of tokens that the user receives back.

### WithdrawRecord
```protobuf
message WithdrawRecord {
  string zone_id = 1;
  string withdrawer = 2;
  map <uint64, WithdrawRecordContent> records = 3;
}
```
`WithdrawRecord` manages records of user withdrawal requests.

### WithdrawRecordContent
```protobuf
message WithdrawRecordContent {
  string amount = 1[(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int", (gogoproto.nullable) = false];
  int64 state = 2;
  int64 oracle_version = 3;
  uint64 withdraw_version = 4;
  google.protobuf.Timestamp completion_time = 5[(gogoproto.nullable) = false, (gogoproto.stdtime)= true];
}
```
`WithdrawRecordContent` manages detailed records of user withdrawal requests. The withdrawal request is made once every few days, so the version of the last bot's action and the withdrawal completion time are saved together.
