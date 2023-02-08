# Messages

---

## MsgDeposit

---

```protobuf
message MsgDeposit {
  string zone_id = 1;
  string depositor = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string claimer = 3;
  cosmos.base.v1beta1.Coin amount = 4 [(gogoproto.nullable) = false];
  uint64 timeout_timestamp = 5;
}
```

`MsgDeposit` is a message used when depositing for asset liquidity.

## MsgDepositResponse

---

```protobuf
message MsgDepositResponse {
  string receiver = 1;
  string depositor = 2;
  cosmos.base.v1beta1.Coin deposited_amount = 3 [(gogoproto.nullable) = false];
}
```

`MsgDepositResponse` is a message used response for `MsgDepsit`.

## MsgDelegate

---

```protobuf
message MsgDelegate {
  string zone_id = 1;
  string controller_address = 2;
  uint64 version = 3;
  uint64 timeout_timestamp = 4;
}
```

`MsgDelegate` is the message the bot requests for a delegate.

## MsgPendingUndelegate

---
```protobuf
message MsgPendingUndelegate {
  string zone_id = 1;
  string delegator = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string withdrawer = 3 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  cosmos.base.v1beta1.Coin amount = 4 [(gogoproto.nullable) = false];
}
```
`MsgPendingUndelegate` is the message the user requests to undelegate.
This request does not result in an immediate Undelegate request.
Requests recorded in the Undelegate Record actually result in a undelegate request via `Undelegate`.

## MsgPendingUndelegateResponse

---
```protobuf
message MsgPendingUndelegateResponse {
  string zone_id = 1;
  string delegator = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string withdrawer = 3 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  cosmos.base.v1beta1.Coin burn_asset = 4 [(gogoproto.nullable) = false];
  cosmos.base.v1beta1.Coin undelegate_asset = 5 [(gogoproto.nullable) = false];
}
```
`MsgPendingUndelegateResponse` contains metadata for requests in response to MsgPendingUndelegate.

## MsgUnDelegate

---
```protobuf
message MsgUndelegate {
  string zone_id = 1;
  string controller_address = 2;
  uint64 version = 3;
  uint64 timeout_timestamp = 4;
}
```
`MsgUndelegate` is the message the bot requests to undelegate.

## MsgUndelegateResponse

---
```protobuf
message MsgUndelegateResponse{
  string zone_id = 1;
  cosmos.base.v1beta1.Coin total_burn_asset = 2 [(gogoproto.nullable) = false];
  cosmos.base.v1beta1.Coin total_undelegate_asset = 3 [(gogoproto.nullable) = false];
}
```
`MsgUndelegateResponse` contains metadata for requests in response to MsgUndelegate.

## MsgWithdraw

---
```protobuf
message MsgWithdraw {
  string zone_id = 1;
  string withdrawer = 2[(cosmos_proto.scalar) = "cosmos.AddressString"];
  string from_address = 3 [(cosmos_proto.scalar) = "cosmos.AddressString"];
}
```
`MsgWithdraw` is a message used when user want to withdraw their asset with IBC.

## MsgWithdrawResponse

---
```protobuf
message MsgWithdrawResponse {
  string withdrawer = 1;
  string withdraw_amount = 2[(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int", (gogoproto.nullable) = false];
}
```
`MsgWithdrawResponse` is a message used response for `MsgWithdraw`

## MsgClaimSnAsset

---
```protobuf
message MsgClaimSnAsset {
  string zone_id = 1;
  string claimer = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string from_address = 3 [(cosmos_proto.scalar) = "cosmos.AddressString"];
}
```
`MsgClaimSnAsset` receives the snAsset, which is the equity token for the assets deposited.

## MsgClaimSnAssetResponse

---
```protobuf
message MsgClaimSnAssetResponse {
  string claimer = 1;
  cosmos.base.v1beta1.Coin minted = 2
  [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"];
}
```

## MsgIcaWithdraw

---
```protobuf
message MsgIcaWithdraw{
  string zone_id = 1;
  string controller_address = 2;
  string ica_transfer_port_id = 3;
  string ica_transfer_channel_id = 4;
  google.protobuf.Timestamp chain_time = 5[(gogoproto.stdtime) = true, (gogoproto.nullable) = false];
  uint64 version = 6;
  uint64 timeout_timestamp = 7;
}
```
`MsgIcaWithdraw` remotely transfers the undelegated assets in the other zone to the Supernova chain.

## MsgIcaWithdrawResponse

---
```protobuf
message MsgIcaWithdrawResponse {
}
```

## MsgClaimAllSnAsset

---
```protobuf
message MsgClaimAllSnAsset {
  string zone_id = 1;
  string from_address = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
}

message MsgClaimAllSnAssetResponse {
  string claimer = 1;
}
```
`MsgClaimAllSnAsset` claims snAsset for all users.