# Client

---

## Query

TODO : Now, icacontrol module not serves query.

## Tx

### Register Zone
```shell
novad tx icacontrol register-zone [zone-id] [controller-address] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal]
```

`register-zone` is a transaction that registers new Zone information. This transaction can only be submitted by a given signatory.

### Delete Zone
```shell
novad tx icacontrol delete-zone [zone-id] [controller-address]
```

`delete-zone` is a transaction that deletes the registered zone. This transaction can only be submitted by a given signatory.

### Change Zone
```shell
novad tx icacontrol change-zone [zone-id] [host-address] [controller-address] [connection-id] [transfer-port-id] [transfer-channel-id] [validator_address] [base-denom] [decimal]
```

`change-zone` is a transaction that modifies the registered zone. This transaction can only be submitted by a given signatory.

### ICA Delegate
```shell
novad tx icacontrol ica-delegate [zone-id] [controller-address] [host-address] [amount]
```

`ica-delegate` is a transaction used for remote delegation using ICA. This transaction can only be submitted by a given signatory.

### ICA Undelegate
```shell
novad tx icacontrol ica-undelegate [zone-id] [controller-address] [host-address] [amount]
```

`ica-undelegate` is a transaction used for remote de-delegation using ICA. This transaction can only be submitted by a given signatory.

### ICA Auto Staking
```shell
novad tx icacontrol ica-auto-staking [zone-id] [controller-address] [amount]
```

`ica-auto-staking` is a transaction used for auto-compounding using ICA. This transaction can only be submitted by a given signatory.

### ICA Transfer
```shell
novad tx icacontrol ica-transfer [zone-id] [controller-address] [host-address] [receiver] [ica-transfer-port-id] [ica-transfer-channel-id] [amount]
```

`ica-transfer` is a transaction used to transfer assets between chains using ICA. This transaction can only be submitted by a given signatory.

### ICA Authz Grant
```shell
novad tx icacontrol ica-authz-grant [zone-id] [grantee-address] [authorization-type] --from [granter]
```

`ica-authz-grant` is a transaction used to transfer permissions between accounts using ICA. This transaction can only be submitted by a given signatory.

### ICA Authz Revoke
```shell
novad tx icacontrol ica-authz-revoke [zone-id] [grantee-address] [msg_type]  --from [controller-address]
```

`ica-authz-revoke` is a transaction used to execute transferred permissions between accounts using ICA. This transaction can only be submitted by a given signatory.