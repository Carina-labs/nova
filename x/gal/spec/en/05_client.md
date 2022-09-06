# Client

---
## Cli

---
A user can query and interact with the GAL module using the CLI.

### Query
The query commands allow users to query GAL state.
```shell
novad query gal --help
```

#### Claimable Asset
`Claimable Asset` returns the amount of equity tokens(sn-token) that can be received for the assets of the `address` deposited 
in the zone corresponding to the `zone-id`.
```shell
novad query gal claimable_asset [zone-id] [address]
```

Example:
```shell
novad query gal claimable_asset cosmoshub-1 nova1a2b...
```

Example Output:
```json
{
    "denom": "snAtom",
    "amount": 1000
}
```

#### Pending Withdrawals
`pending_withdrawals` returns the amount of pending withdrawal asset of user corresponding to zone-id.
```shell
novad query gal pending_withdrawals [zone-id] [address]
```

Example:
```shell
novad query gal pending_withdrawals cosmoshub-1 nova1a2b...
```

Example Output:
```json
{
    "denom": "uatom",
    "amount": 1000
}
```

#### Active Withdrawals
`active_withdrawals` returns the amount of active withdrawal asset of user corresponding to zone-id.
```shell
novad query gal active_withdrawals [zone-id] [address]
```

Example:
```shell
novad query gal active_withdrawals [zone-id] [address]
```

Example Output:
```json
{
  "denom": "uatom",
  "amount": 1000
}
```

### Deposit Records
`deposit_records` returns the records of deposit for user corresponding to zone-id.
```shell
novad query gal deposit_records [zone-id] [address]
```

Example:
```shell
novad query gal deposit_records cosmoshub-1 nova1a2b...
```

Example Output:
```json
{
  "zoneId": "cosmoshub-1",
  "claimer": "nova1a2b...",
  "records": [
    {
      "depositor": "nova1a2b...",
      "amount": {
        "denom": "uatom",
        "amount": 1000
      },
      "state": 1,
      "oracleVersion": 100,
      "delegateVersion": 124
    }
  ]
}
```

#### Undelegate Records
`undelegate_records` returns the records of undelegation for user corresponding to zond-id.

```shell
novad query gal undelegate_records [zone-id] [address]
```

Example:
```shell
novad query gal undelegate_records cosmoshub-1 nova1a2b...
```

Example Output:

```json
{
  "zoneId": "cosmoshub-1",
  "delegator": "nova1a2b...",
  "records": [
    {
      "withdrawer": "nova1a2b...",
      "snAssetAmount": {
        "denom": "snAtom",
        "amount": 1000
      },
      "withdrawAmount": 500,
      "state": 1,
      "oracleVersion": 100,
      "undelegateVersion": 125
    }
  ]
}
```

#### Withdraw Records
`withdraw_records` returns the records of withdraw for user corresponding to zond-id.

```shell
novad query withdraw_records [zone-id] [address]
```

Example:
```shell
novad query withdraw_records cosmoshub-1 nova1a2b...
```

Example Output:

```json
{
  "zoneId": "cosmoshub-1",
  "withdrawer": "nova1a2b...",
  "records": {
    "1": {
      "Amount": 1000,
      "state": 1,
      "oracleVersion": 1000,
      "withdrawVersion": 1024,
      "CompletionTime": {
        "well": 1000,
        "ext": 1000,
        "loc": {
          "name": "UTC",
          "zone": null,
          "tx": [],
          "extend": null
        }
      }
    }
  }
}
```

### Transaction

#### Deposit
```shell
novad tx gal deposit [zone-id] [depositor] [clamier] [amount]
```

Example:
```shell
novad tx gal deposit cosmoshub-1 nova1a2b... nova1a2b... 10000uatom
```

#### Pending Undelegate
```shell
novad tx gal pendingundelegate [zone-id] [delegator] [withdrawer] [amount]
```

Example:
```shell
novad tx gal pendingundelegate cosmoshub-1 nova1a2b... nova9a8b... 10000uatom
```

#### Undelegate
```shell
novad tx gal undelegate [zone-id] [controller-address]
```

Example:
```shell
novad tx gal undelegate cosmoshub-1 nova7ba2...
```

#### Withdraw
```shell
novad tx gal withdraw [zone-id] [withdrawer]
```

Example:
```shell
novad tx gal withdraw cosmoshub-1 nova1a2b...
```

#### Claim sn-Token
```shell
novad tx gal claimsntoken [zone-id] [claimer-address]
```

Example:
```shell
novad tx gal claimsntoken cosmoshub-1 nova1a2b...
```

#### Pending Withdraw
```shell
novad tx gal pendingwithdraw [zone-id] [controller-address] [ica-transfer-port-id]
```

Example:
```shell
novad tx gal pendingwithdraw cosmoshub-1 nova1a2b... port-1
```

#### Delegate
```shell
novad tx gal delegate [zone-id] [controller-address]
```

Example:
```shell
novad tx gal delegate cosmoshub-1 nova1a2b...
```

## gRPC

---
A user can query the `gal` module using gRPC endpoints.

### Estimate sn-Asset
The `estimate_sn_asset` endpoint allows user to calculate the amount of asset if they deposit their asset.
```
nova.gal.v1.Query/estimate_sn_asset
```

### Claimable Amount
The `claimable_amount` endpoint allows user to query the amount of sn-asset they can claim.
```
nova.gal.v1.Query/claimable_amount
```

### Pending Withdrawal
The `pending_withdrawal` endpoint allows user to query the amount of pending withdraw.
```
nova.gal.v1.Query/pending_withdrawals
```

### Active Withdrawal
The `active_withdrawals` endpoint allows user to query the amount of active withdraw.
```
nova.gal.v1.Query/active_withdrawals
```

### Deposit Records
The `deposit_records` endpoint allows user to query the deposit records of user.
```
nova.gal.v1.Query/deposit_records
```

### Undelegate Records
The `undelegate_records` endpoint allows user to query the undelegate records of user.
```
nova.gal.v1.Query/undelegate_records
```

### Withdraw Records
The `withdraw_records` endpoint allows user to query the withdraw records of user.
```
nova.gal.v1.Query/withdraw_records
```