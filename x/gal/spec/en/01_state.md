# State
---

## Overview

This document describes the states used by the `GAL` module.

---

### DepositRecord

```protobuf
message DepositRecord {
  string zone_id = 1;
  string depositor = 2;
  repeated DepositRecordContent records = 3;
}
```

`DepositRecord` manages historical data that you deposit in a particular zone.

### DepositRecordContent

```protobuf
message DepositRecordContent {
  string claimer = 1;
  cosmos.base.v1beta1.Coin amount = 2;
  int64 state = 3;
}
```

`DepositRecordContent` stores the user's Deposit details. sn-tokens are paid to the claimer.

### DelegateRecord

```protobuf
message DelegateRecord {
  string zone_id = 1;
  string claimer = 2;
  map <uint64, DelegateRecordContent> records = 3;
}
```

`DelegateRecord` manages historical data that you delegate in a particular zone.

### DelegateRecordContent

```protobuf
message DelegateRecordContent {
  cosmos.base.v1beta1.Coin amount = 1;
  int64 state = 2;
  uint64 oracle_version = 3;
}
```

`DelegateRecordContent` stores the user's delegate details.

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

Save the Oracle version at the time of the undelegation request and use it to calculate the amount of tokens that the
user receives back.

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
  cosmos.base.v1beta1.Coin unstaking_amount = 2;
  int64 state = 3;
  int64 oracle_version = 4;
  uint64 withdraw_version = 5;
  google.protobuf.Timestamp completion_time = 6[(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

`WithdrawRecordContent` manages detailed records of user withdrawal requests. The withdrawal request is made once every
few days, so the version of the last bot's action and the withdrawal completion time are saved together.

### AsstInfo

```protobuf
message AssetInfo {
  string zone_id = 1;
  string un_minted_w_asset = 2[(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int", (gogoproto.nullable) = false];
}
```

`AssetInfo` manages the information not minted snAsset.