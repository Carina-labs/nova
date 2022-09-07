## Core Concept

---
## Overview

---
The `poolincentive` module is responsible for incentives for decentralized exchanges to exchange Supernova's equity tokens and native coins.

This module manages the pool by dividing it into `candidate pool` and `incentive pool`. 
`Incentive pools` are rewarded with some of the new Nova coin issues, and users who provide liquidity to the pools can receive them. 

Also, which pool will be selected as the incentive pool will be decided by governance.

---

### Candidate Pool

Anyone can register `candidate pool`, but it is not an incentive.
Among the candidate pools, the pool designated to give incentives through governance becomes the incentive pool.

### Incentive Pool

The `incentive pool` is a pool that will be incentivized through governance.
How much incentive each pool receives is calculated based on the weight value.