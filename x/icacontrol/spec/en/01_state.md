# State

---

## Overview
This document describes the states used by the `ibcstaking` module.

---

### RegisteredZone

`RegisteredZone` stores metadata for the relative chain associated with the Supernova chain. 

The metadata includes the relative chain's denom, validator address, connection information, and so on.

```protobuf
message RegisteredZone {
  string zone_id = 1;
  IcaConnectionInfo ica_connection_info = 2;
  IcaAccount ica_account = 3;
  TransferConnectionInfo transfer_info = 4;
  string validator_address = 5;
  string base_denom = 6;
  string sn_denom = 7;
  int64 decimal = 8;
}
```

### IcaAccount

`IcaAccount` stores information about the account you use to connect to the relative chain and ICA.

```protobuf
message IcaAccount {
  string controller_address = 1;
  string host_address = 2;
  cosmos.base.v1beta1.Coin balance = 3 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.nullable) = false
  ];
}

```

### IcaConnectionInfo

`IcaConnectionInfo` stores connection information for ICA connections. 

Connection information includes connection-id and port-id.

```protobuf
message IcaConnectionInfo {
  string connection_id = 1;
  string port_id = 2;
}
```

### TransferConnectionInfo

`TransferConnectionInfo` stores connection information for using IBC Transfer.

Connection information includes channel_id and port_id.

```protobuf
message TransferConnectionInfo {
  string channel_id = 1;
  string port_id = 2;
}
```