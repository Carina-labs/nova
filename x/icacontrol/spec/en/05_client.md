# Client

---

## Query

### All Zone
```shell
novad query icacontrol all-zone
```

`All-zone` returns all chain information connected to nova chain.

Example:
```shell
novad query icacontrol all-zone
```

Example Output:
```json
{
    zones:
    - base_denom: uatom
      decimal: "6"
      deposit_max_entries: "100"
      ica_account:
        controller_address: nova1...
        host_address: cosmos1...
      ica_connection_info:
        connection_id: connection-0
        port_id: gaia.nova1...
      sn_denom: snuatom
      transfer_info:
        channel_id: channel-0
        port_id: transfer
      undelegate_max_entries: "100"
      validator_address: cosmosvaloper1...
      zone_id: gaia
    - base_denom: uosmo
      decimal: "6"
      deposit_max_entries: "100"
      ica_account:
        controller_address: nova1...
        host_address: osmo1...
      ica_connection_info:
        connection_id: connection-1
        port_id: osmosis.nova1...
      sn_denom: snuosmo
      transfer_info:
        channel_id: channel-1
        port_id: transfer
      undelegate_max_entries: "100"
      validator_address: osmovaloper1...
      zone_id: osmosis
}
```

### Zone
```shell
novad query icacontrol zone [zone-id]
```

`zone` returns information about the entered zone-id.

Example:
```shell
novad query icacontrol zone gaia
```

Example Output:
```json
{
 - base_denom: uatom
    decimal: "6"
    deposit_max_entries: "100"
    ica_account:
      controller_address: nova1...
      host_address: cosmos1...
    ica_connection_info:
      connection_id: connection-0
      port_id: gaia.nova1...
    sn_denom: snuatom
    transfer_info:
      channel_id: channel-0
      port_id: transfer
    undelegate_max_entries: "100"
    validator_address: cosmosvaloper1...
    zone_id: gaia
}
```

## Tx

### Register Zone
```shell
novad tx icacontrol register-zone [zone-id] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal] [deposit-max-entries] [undelegate-max-entries]
```

`register-zone` is a transaction that registers new Zone information. This transaction can only be submitted by a given signatory.

### Delete Zone
```shell
novad tx icacontrol delete-zone [zone-id]
```

`delete-zone` is a transaction that deletes the registered zone. This transaction can only be submitted by a given signatory.

### Change Zone
```shell
novad tx icacontrol change-zone [zone-id] [host-address] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal] [deposit-max-entries] [undelegate-max-entries]
```

`change-zone` is a transaction that modifies the registered zone. This transaction can only be submitted by a given signatory.

### ICA Delegate
```shell
novad tx icacontrol ica-delegate [zone-id] [amount]
```

`ica-delegate` is a transaction used for remote delegation using ICA. This transaction can only be submitted by a given signatory.

### ICA Undelegate
```shell
novad tx icacontrol ica-undelegate [zone-id] [amount]
```

`ica-undelegate` is a transaction used for remote de-delegation using ICA. This transaction can only be submitted by a given signatory.

### ICA Auto Staking
```shell
novad tx icacontrol ica-auto-staking [zone-id] [amount]
```

`ica-auto-staking` is a transaction used for auto-compounding using ICA. This transaction can only be submitted by a given signatory.

### ICA Transfer
```shell
novad tx icacontrol ica-transfer [zone-id] [receiver] [ica-transfer-port-id] [ica-transfer-channel-id] [amount]
```

`ica-transfer` is a transaction used to transfer assets between chains using ICA. This transaction can only be submitted by a given signatory.

### ICA Authz Grant
```shell
novad tx icacontrol ica-authz-grant [zone-id] [grantee-address] [authorization-type] --from [granter]
```

`ica-authz-grant` is a transaction used to transfer permissions between accounts using ICA. This transaction can only be submitted by a given signatory.

### ICA Authz Revoke
```shell
novad tx icacontrol ica-authz-revoke [zone-id] [grantee-address] [msg_type]
```

`ica-authz-revoke` is a transaction used to execute transferred permissions between accounts using ICA. This transaction can only be submitted by a given signatory.
