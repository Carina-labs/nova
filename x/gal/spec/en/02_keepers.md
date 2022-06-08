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
Deposit the coin and send it to the other chain via IBC transfer.
The remote deposit proceeds with an ACK response.

## Unstaking

---
```go
func (Keeper) UnStaking(ctx sdk.Context) error {}
```
Un-stack the coins you deposited.
At this point, you must wait for the unstaking period.


## WithdrawCoin

---
```go
func (Keeper) WithdrawCoin(ctx sdk.Context,
	withdrawer string,
	amt sdk.Coins) error {
	
}
```
Out of assets not staked in the host chain account, withdraw the withdrawable assets.
