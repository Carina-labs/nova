# Keeper

---
## DepositCoin

---
```go
func (Keeper) DepositCoin(ctx sdk.Context,
	depository string,
	receiver string,
	sourcePort string,
	sourceChannel string,
	amt sdk.Coins) error {}
```
IBC 전송을 통해 `Coin`을 상대 체인으로 보내고 예치합니다.
원격 예치는 IBC ACK 응답을 받으면 진행합니다.

## Unstaking

---
```go
func (Keeper) UnStaking(ctx sdk.Context) error {}
```
예치한 코인을 해제합니다.
여기서 사용자는 해제 대기 기간을 기다려야합니다.

## WithdrawCoin

---
```go
func (Keeper) WithdrawCoin(ctx sdk.Context,
	withdrawer string,
	amt sdk.Coins) error {
	
}
```
호스트체인에서 스테이킹 해제 된 자산을 IBC를 통해 유저의 계정으로 가져옵니다.