# Messages

---

## MsgRegisterZone
```protobuf
message MsgRegisterZone {
  option (gogoproto.equal) = false;

  string zone_id = 1;
  IcaConnectionInfo ica_info = 2;
  IcaAccount ica_account = 3;
  TransferConnectionInfo transfer_info = 4;
  string validator_address = 5;
  string base_denom = 6;
  int64 decimal = 7;
}
```

`MsgRegisterZone` is the message you use to register a new zone.

## MsgRegisterZoneResponse
```protobuf
message MsgRegisterZoneResponse {}
```

`MsgRegisterZoneResponse` is a response message for `MsgRegisterZone`.

## MsgChangeRegisteredZoneInfo
```protobuf
message MsgChangeRegisteredZoneInfo{
  option (gogoproto.equal) = false;

  string zone_id = 1;
  IcaConnectionInfo ica_info = 2;
  IcaAccount ica_account = 3;
  TransferConnectionInfo transfer_info = 4;
  string validator_address = 5;
  string base_denom = 6;
  int64 decimal = 7;
}
```

`MsgChangeRegisteredZoneInfo` modifies the information in the registeredZone.

## MsgChangeRegisteredZoneInfoResponse
```protobuf
message MsgChangeRegisteredZoneInfoResponse{}
```

`MsgChangeRegisteredZoneInfoResponse` is a response message for `MsgChangeRegisteredZone`.

## MsgDeleteRegisteredZone
```protobuf
message MsgDeleteRegisteredZone{
  string zone_id = 1;
  string controller_address = 2;
}
```

`MsgDeleteRegisteredZone` deletes registered Zone information.

## MsgDeleteRegisteredZoneResponse
```protobuf
message MsgDeleteRegisteredZoneResponse{}
```

`MsgDeleteRegisteredZoneResponse` is a response message for `MsgDeleteRegisteredZone`.

## MsgIcaDelegate
```protobuf
message MsgIcaDelegate{
  string zone_id = 1;
  string host_address = 2;
  string controller_address = 3;
  cosmos.base.v1beta1.Coin amount = 4[(gogoproto.nullable) = false];
}
```

`MsgIcaDelegate` is a message used for remote delegation using ICA.

## MsgIcaDelegateResponse
```protobuf
message MsgIcaDelegateResponse{}
```

`MsgIcaDelegateResponse` is a response message for `MsgIcaDelegate`.

## MsgIcaUndelegate
```protobuf
message MsgIcaUndelegate{
  string zone_id = 1;
  string host_address = 2;
  string controller_address = 3;
  cosmos.base.v1beta1.Coin amount = 4[(gogoproto.nullable) = false];
}
```

`MsgIcaUndelegate` is a message used to de-delegate remote using ICA.

## MsgIcaUndelegateResponse
```protobuf
message MsgIcaUndelegateResponse{}
```

`MsgIcaUndelegateResponse` is a response message for `MsgIcaUndelegate`.

## MsgIcaAutoStaking
```protobuf
message MsgIcaAutoStaking{
  string zone_id = 1;
  string host_address = 2;
  string daomodifier_address = 3;
  cosmos.base.v1beta1.Coin amount = 4[(gogoproto.nullable) = false];
}
```

`MsgIcaAutoStaking` is a message for remote auto-compound using ICA.

## MsgIcaAutoStakingResponse
```protobuf
message MsgIcaAutoStakingResponse{}
```

`MsgIcaAutoStakingResponse` is a response message for `MsgIcaAutoStaking`.

## MsgIcaTransfer
```protobuf
message MsgIcaTransfer{
  string zone_id = 1;
  string host_address = 2;
  string daomodifier_address = 3;
  string receiver_address = 4;
  string ica_transfer_port_id = 5;
  string ica_transfer_channel_id = 6;
  cosmos.base.v1beta1.Coin amount = 7[(gogoproto.nullable) = false];
}
```

`MsgIcaTransfer` is a message for IBC transfer from the counterpart to the Supernova chain using ICA.

## MsgIcaTransferResponse
```protobuf
message MsgIcaTransferResponse{}
```

`MsgIcaTransferResponse` is a response message for `MsgIcaTransfer`.

## MsgIcaAuthzGrant
```protobuf
message MsgIcaAuthzGrant{
  string zone_id = 1;
  string grantee = 2;
  string controller_address = 3;
  cosmos.authz.v1beta1.Grant grant = 4 [(gogoproto.nullable) = false];
}
```

`MsgIcaAuthzGrant` is a message used to transfer authority through authz between two accounts in the other chain using ICA.

## MsgIcaAuthzGrantResponse
```protobuf
message MsgIcaAuthzGrantResponse{}
```

`MsgIcaAuthzGrantResponse` is a response messasge for `MsgIcaAuthzGrant`.

## MsgIcaAuthzRevoke
```protobuf
message MsgIcaAuthzRevoke{
  string zone_id = 1;
  string grantee = 2;
  string controller_address = 3;
  string msg_type_url = 4;
}
```

`MsgIcaAuthzRevoke` is a message used for proxy execution between accounts linked to authz via ICA.

## MsgIcaAuthzRevokeReponse
```protobuf
message MsgIcaAuthzRevokeResponse{}
```

`MsgIcaAuthzRevokeResponse` is a response message for `MsgIcaAuthzRevoke`.