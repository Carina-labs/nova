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

## EventRegisterZone
```protobuf
message EventRegisterZone{
  string zone_id = 1;
  IcaConnectionInfo ica_connection_info = 2;
  IcaAccount ica_account = 3;
  TransferConnectionInfo transfer_info = 4;
  string validator_address = 5;
  string base_denom = 6;
  string sn_denom = 7;
  int64 decimal = 8;
  int64 undelegate_max_entries = 9;
  int64 deposit_max_entries = 10;
}
```

`EventRegisterZone` is an event that occurs when the RegisterZone message is executed.

## EventChangeRegisteredZone
```protobuf
message EventChangeRegisteredZone{
  string zone_id = 1;
  IcaConnectionInfo ica_info = 2;
  IcaAccount ica_account = 3;
  TransferConnectionInfo transfer_info = 4;
  string validator_address = 5;
  string base_denom = 6;
  string sn_denom = 7;
  int64 decimal = 8;
  int64 undelegate_max_entries = 9;
  int64 deposit_max_entries = 10;
}
```

`EventChangeRegisteredZone` is an event that occurs when the ChangeRegisteredZone message is executed.

## EventDeleteZone
```protobuf
message EventDeleteZone{
  string zone_id = 1;
  string controller_address = 2;
}
```

`EventDeleteZone` is an event that occurs when the DeleteZone message is executed.

## EventIcaDelegate
```protobuf
message EventIcaDelegate{
  string zone_id = 1;
  string controller_address = 2;
  cosmos.base.v1beta1.Coin amount = 3[(gogoproto.nullable) = false];
}
```

`EventIcaDelegate` is an event that occurs when the IcaDelegate message is executed.

## EventIcaUndelegate
```protobuf
message EventIcaUndelegate{
  string zone_id = 1;
  string controller_address = 2;
  cosmos.base.v1beta1.Coin amount = 3[(gogoproto.nullable) = false];
}
```

`EventIcaUndelegate` is an event that occurs when the IcaUndelegate message is executed.

## EventIcaAutoStaking
```protobuf
message EventIcaAutoStaking{
  string zone_id = 1;
  string controller_address = 2;
  cosmos.base.v1beta1.Coin amount = 3[(gogoproto.nullable) = false];
}
```

`EventIcaAutoStaking` is an event that occurs when the IcaAutoStaking message is executed.

## EventIcaTransfer
```protobuf
message EventIcaTransfer{
  string zone_id = 1;
  string controller_address = 2;
  string receiver_address = 3;
  string ica_transfer_port_id = 4;
  string ica_transfer_channel_id = 5;
  cosmos.base.v1beta1.Coin amount = 6[(gogoproto.nullable) = false];
}
```

`EventIcaTransfer` is an event that occurs when the IcaTransfer message is executed.

## EventIcaAuthzGrant
```protobuf
message EventIcaAuthzGrant{
  string zone_id = 1;
  string grantee = 2;
  string controller_address = 3;
  cosmos.authz.v1beta1.Grant grant = 4 [(gogoproto.nullable) = false];
}
```

`EventIcaAuthzGrant` is an event that occurs when the IcaAuthzGrant message is executed.

## EventIcaAuthzRevoke
```protobuf
message EventIcaAuthzRevoke{
  string zone_id = 1;
  string grantee = 2;
  string controller_address = 3;
  string msg_type_url = 4;
}
```

`EventIcaAuthzRevoke` is an event that occurs when the IcaAuthzRevoke message is executed.

## EventRegisterControllerAddress
```protobuf
message EventRegisterControllerAddress{
  string zone_id = 1;
  string controller_address = 2;
  string from_address = 3;
}
```

`EventRegisterControllerAddress` is an event that occurs when the RegisterControllerAddress message is executed.
