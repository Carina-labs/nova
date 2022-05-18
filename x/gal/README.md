# GAL Module

---
The GAL(General Asset Liquidity) module is reponsible for exchanging the native token with the Supernova token.

## Deposit

---
The user may deposit the token using the `GAL` module.
The deposited token is transmitted to each native chain and delegate to the validator.
Staking interest is periodically synchronized and the bond token is paid as compensation.

## Withdraw

---
The user can withdraw the assets deposited using the `GAL` module. 
The bond token is incinerated, and in this process, it must wait for an unstaking period.
