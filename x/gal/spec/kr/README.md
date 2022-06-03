# GAL

---

## Abstract

---
GAL 모듈은 스테이킹 유동화를 위해 유저가 예치한 자산을 관리합니다.
예치된 자산은 전체 자산 대비 지분으로 저장됩니다.
GAL 모듈은 이 지분을 `stToken`으로 관리합니다.
그리고 주기적으로 호스트 체인에 쌓인 이자는 재위임되며, 유저가 받을 수 있는 총 자산을 다시 계산합니다.

## Content

---
1. [State](01_state.md)
2. [Keeper](02_keepers.md)
3. [Messages](03_messages.md)
4. [Events](04_events.md)
5. [Parameters](05_params.md)
6. [Client](06_client.md)