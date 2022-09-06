# Events

---

## EventDelegateFail
```protobuf
message EventDelegateFail {
  string msg_type_url = 1;
  string delegator_address = 2;
  string validator_address = 3;
  cosmos.base.v1beta1.Coin amount = 4 [
        (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
        (gogoproto.nullable) = false
    ];
}
```

`EventDelegateFail` is an event that occurs when a remote delegation fails.

## EventUndelegateFail
```protobuf
message EventUndelegateFail {
  string msg_type_url = 1;
  string delegator_address = 2;
  string validator_address = 3;
  cosmos.base.v1beta1.Coin amount = 4 [
        (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
        (gogoproto.nullable) = false
    ];
}
```

`EventUndelegateFail` is an event that occurs when a remote delegation fails.

## EventAutostakingFail
```protobuf
message EventAutostakingFail {
    string msg_type_url = 1;
    string delegator_address = 2;
    string validator_address = 3;
    cosmos.base.v1beta1.Coin amount = 4 [
        (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
        (gogoproto.nullable) = false
    ];
}
```

`EventAutostakingFail` is an event that occurs when auto-compounding fails.

## EventTransferFail
```protobuf
message EventTransferFail {
  option (gogoproto.equal)           = false;
  option (gogoproto.goproto_getters) = false;

  string msg_type_url = 1;
  string source_port = 2;
  string source_channel = 3;
  cosmos.base.v1beta1.Coin token = 4 [(gogoproto.nullable) = false];
  string sender = 5;
  string receiver = 6;
  string timeout_height = 7;
  uint64 timeout_timestamp = 8;
}

```

`EventTransferFail` is an event that occurs when an asset transfer between chains fails.