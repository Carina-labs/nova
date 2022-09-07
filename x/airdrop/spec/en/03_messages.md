# Messages

---

## MsgClaimAirdropRequest
```protobuf
message MsgClaimAirdropRequest {
  option (cosmos.msg.v1.signer) = "user_address";
  
  string user_address = 1 ;
  
  QuestType quest_type = 2;
}
```

`MsgClaimAirdropRequest` is an airdrop volume request message.

## MsgClaimAirdropResponse
```protobuf
message MsgClaimAirdropResponse {}
```

`MsgClaimAirdropResponse` is a response message for MsgClaimAirdropRequest.

## MsgMarkSocialQuestPerformedRequest
```protobuf
message MsgMarkSocialQuestPerformedRequest {
  option (cosmos.msg.v1.signer) = "controller_address";
  
  string controller_address = 1;
  
  repeated string user_addresses = 2;
}
```

`MsgMarkSocialQuestPerformedRequest` is a message that marks the completion of the SNS sharing mission by a specific user.

## MsgMarkSocialQuestPerformedResponse
```protobuf
message MsgMarkSocialQuestPerformedResponse {}
```

`MsgMarkSocialQuestPerformedResponse` is a response message for `MsgMarkSocialQuestPerformedRequest`.

## MsgMarkUserProvidedLiquidityRequest
```protobuf
message MsgMarkUserProvidedLiquidityRequest {
  option (cosmos.msg.v1.signer) = "controller_address";
  
  string controller_address = 1;
  
  repeated string user_addresses = 2;
}
```

`MsgMarkUserProvidedLiquidityRequest` is a message that marks a particular user as having provided liquidity to Supernova.

## MsgMarkUserProvidedLiquidityResponse
```protobuf
message MsgMarkUserProvidedLiquidityResponse {}
```

`MsgMarkUserProvidedLiquidityResponse` is a response message for `MsgMarkUserProvidedLiquidityRequest`.