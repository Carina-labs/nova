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
`MsgDeposit`는 자산 유동화를 위해 예치할 때 사용하는 메시지입니다.

## MsgDepositResponse

---
```protobuf
message MsgDepositResponse {
  
}
```
`MsgDepositResponse`는 `MsgDeposit`의 결과 메시지입니다.

## MsgUnStaking

---
```protobuf
message MsgUnStaking {
  required string depositor = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
`MsgUnStaking`는 유저가 예치된 자산을 해제할 때 사용하는 메시지입니다.

## MsgUnStakingResponse

---
```protobuf
message MsgUnStakingResponse {
  
}
```
`MsgUnStakingResponse`는 `MsgUnStaking`의 결과로 사용되는 메시지입니다.
이 결과를 받았다고해서 호스트 체인의 자산이 IBC 전송을 통해 노바 체인의 계정으로 넘어오지 않습니다.
만약 스테이킹 해제 대기 기간을 전부 기다렸다면, IBC 전송을 할 수 있습니다.


## MsgWithdraw

---
```protobuf
message MsgWithdraw {
  required string withdrawer = 1;
  repeated cosmos.base.v1beta1.Coin amount = 2;
}
```
`MsgWithdraw`는 유저가 스테이킹 해제된 자산을 IBC를 통해 노바 체인으로 가져올 때 사용하는 메시지입니다.

## MsgWithdrawResponse

---
```protobuf
message MsgWithdrawResponse {
  
}
```
`MsgWithdrawResponse`는 `MsgWithdraw`의 결과로 사용되는 메시지입니다.