# Messages

---
## MsgDeposit

---
```protobuf
message MsgDeposit {
  required string depositor = 1;
  required string receiver = 2;
  repeated cosmos.base.v1beta1.Coin amount = 3;
}
```
`MsgDeposit` is a message used when depositing for asset liquidity.

## MsgDepositResponse

---
```protobuf
message MsgDepositResponse {
  string receiver = 1;
  string depositor = 2;
  cosmos.base.v1beta1.Coin deposited_amount = 3;
}
```
`MsgDepositResponse` is a message used response for `MsgDepsit`.

## MsgDelegate

---
```protobuf
message MsgDelegate {
  string zone_id = 1;
  string controller_address = 2;
}
```


## MsgUnDelegate

---
```protobuf
message MsgUnDelegate {
  string zone_id = 1;
  string controller_address = 2;
}
```
`MsgUndelegate` is the message used when requesting Undelegate. 
This request does not result in an immediate Undelegate request. 
Requests recorded in the Undelegate Record actually result in a undelegate request via `IcaUndelegate`.


## MsgUndelegateResponse

---
```protobuf
message MsgUnStakingResponse {
  string zone_id = 1;
  cosmos.base.v1beta1.Coin total_burn_asset = 2;
  cosmos.base.v1beta1.Coin total_undelegate_asset = 3;
}
```
`MsgUndelegateResponse` contains metadata for requests in response to MsgUndelegate.

## MsgWithdraw

---
```protobuf
message MsgWithdraw {
  required string withdrawer = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
`MsgWithdraw` is a message used when user want to withdraw their asset with IBC.

## MsgWithdrawResponse

---
```protobuf
message MsgWithdrawResponse {
  
}
```
`MsgWithdrawResponse` is a message used response for `MsgWithdraw`