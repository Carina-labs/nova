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