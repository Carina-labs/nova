# State

---
## Shares

---

GAL 모듈을 통해 예치된 유저의 지분을 나타냅니다.

### DepositAccount
```protobuf
message DepositAccount {
  string denom = 1;
  repeated DepositInfo depositInfos = 2;
  int64 totalShare = 3;
  int64 lastBlockUpdate = 4;
}
```
`DepositAccount`는 타겟 체인의 이름과 계정 별 지분 정보, 전체 지분과 마지막 업데이트 된 블록 높이를 나타냅니다.

### DepositInfo
```protobuf
message DepositInfo {
  string address = 1;
  int64 share = 2;
  int64 debt = 3;
}
```
`DepositInfo`는 유저의 예치 상태를 나타냅니다.