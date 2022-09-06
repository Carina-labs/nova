## Core Concept

---

## Overview

The `oracle` module manages the status of the zones associated with the Supernova protocol.
The status includes the amount of coins delegated to the Zone's Validator, block height, proof, and so on.

This information is injected by the bot at a short interval (about 15 minutes).
For the integrity of the information, the `AppHash` and `block height` are provided together.
Therefore, the functionality of the oracle module can only be invoked by the specified `controller address`.
The controller address is stored in Parameters on the module.

---

### Validator

The Supernova protocol operates a single `validator` in each zone for remote liquid-staking.

`Validator` operates the blockchain by creating or verifying blocks and participating in governance.


### Controller Address

`Controller Address` is an external bot account operated by the Supernova team.

This account injects an external chain state into the Oracle module using a secure key.

Since the ability to query via IBC has not yet been developed, Supernova uses external bots.