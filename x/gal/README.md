# GAL

---

## Abstract

---
The GAL module manages the assets deposited by the user for liquid staking.
Assets deposited by the user are calculated as the percentage of total assets deposited.
The GAL module manages this asset as stToken.
And periodically, using the principal and interest accumulated in the host chain,
we calculate the total amount of tokens that the user can receive.

## Content

---
0. [Core Concept](spec/en/00_core_concept.md)
1. [State](spec/en/01_state.md)
2. [Keeper](spec/en/02_keeper.md)
3. [Messages](spec/en/03_messages.md)
4. [Events](spec/en/04_events.md)
5. [Client](spec/en/05_client.md)