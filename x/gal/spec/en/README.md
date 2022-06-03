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
1. [State](01_state.md)
2. [Keeper](02_keepers.md)
3. [Messages](03_messages.md)
4. [Events](04_events.md)
5. [Parameters](05_params.md)
6. [Client](06_client.md)