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
## MsgDepositResponse

---
```protobuf
message MsgDepositResponse {
  
}
```

## MsgUnStaking

---
```protobuf
message MsgUnStaking {
  required string depositor = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
## MsgUnStakingResponse

---
```protobuf
message MsgUnStakingResponse {
  
}
```
## MsgWithdraw

---
```protobuf
message MsgWithdraw {
  required string withdrawer = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
## MsgWithdrawResponse

---
```protobuf
message MsgWithdrawResponse {
  
}
```