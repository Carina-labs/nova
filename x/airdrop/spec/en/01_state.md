# State

---

## AirdropInfo
```protobuf
message AirdropInfo {
  google.protobuf.Timestamp snapshot_timestamp = 1;
  google.protobuf.Timestamp airdrop_start_timestamp = 2;
  google.protobuf.Timestamp airdrop_end_timestamp = 3;
  string airdrop_denom = 4;
  int32 quests_count = 5;
  string controller_address = 6;
  string maximum_token_alloc_per_user = 7;
}
```

`AirdropInfo` stores metadata related to airdrop. This includes airdrop start and end times, maximum acceptable quantity, etc.

## UserState
```protobuf
message UserState {
  string recipient = 1;
  string total_amount = 2;
  map<int32, QuestState> quest_states = 3;
}
```

`UserState` indicates the status of the user performing the quest.

## QuestState
```protobuf
message QuestState {
  QuestStateType state = 1;
  google.protobuf.Timestamp achieved_at = 2;
  google.protobuf.Timestamp claimed_at = 3;
}
```

`QuestState` indicates the status of the quest that users perform.
The status of each quest is defined by the `QuestStateType`.

### QuestStateType
```protobuf
enum QuestStateType {
  QUEST_STATE_NOT_STARTED = 0;
  QUEST_STATE_CLAIMABLE = 1;
  QUEST_STATE_CLAIMED = 2;
}
```

QuestStateType is a variable that represents the user's quest performance status.

`QUEST_STATE_NOT_STARTED` : The quest has not been started.

`QUEST_STATE_CLAIMABLE` : You can perform a quest and receive an airdrop.

`QUEST_STATE_CLAIMED` : We have completed the quest and received the airdrop.

