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
  
}
```
`MsgDepositResponse` is a message used response for `MsgDepsit`

## MsgUnStaking

---
```protobuf
message MsgUnStaking {
  required string depositor = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
`MsgUnStaking` is a message used when user want to unstake their asset.


## MsgUnStakingResponse

---
```protobuf
message MsgUnStakingResponse {
  
}
```
`MsgUnStakingResponse` is a message used response for `MsgUnStaking`
As a result, the unstacked asset does not return to the user's wallet via IBC. If you just wait for the unstaking wait period,
you will be able to do IBC.
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