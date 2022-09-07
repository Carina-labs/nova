# Messages

---
## MsgUpdateChainState
```protobuf
message MsgUpdateChainState {
    cosmos.base.v1beta1.Coin coin = 1 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.nullable) = false
    ];
    string operator = 2;
    int64 block_height = 3;
    bytes app_hash = 4;
    string chain_id = 5;
    }
```

`MsgUpdateChainState` is the message used to update the status of the zone stored in Oracle.

## MsgUpdateChainStateResponse
```protobuf
message MsgUpdateChainStateResponse {
  bool success = 1;
}
```

`MsgUpdateChainStateResponse` is a response message for `MsgUpdateChainState`.
If updating is success, the value of success is true.