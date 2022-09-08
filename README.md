# Supernova Protocol

---

## Overview

Supernova is a liquid staking platform for the cosmos ecosystem. 
Using `IBC` and `ICA`, tokens from multiple app chains in the Cosmos ecosystem can be staked and equity tokens can be minted.
In addition, Supernova can securely trade liquidated assets through a decentralized exchange that allows you to trade tokens that match equity tokens.

Supernova the token shares onto the platform to deploy a variety of utility.
Please refer to our [Roadmap](https://medium.com/supernovazone/product-roadmap-2ac43cef5343) for details.

---

## Modules

Supernova is an **App-Chain based on Cosmos-SDK**.
Therefore, we have written the following modules for smooth liquid-staking.

### GAL

The GAL module manages deposit records, undelegation and withdrawal records of users who want to use liquid-staking.
For more information, see [GAL SPEC](x/gal/README.md).

### IcaControl

The IcaControl module manages the Interchain Account (ICA) required to ensure accurate operation of the liquid stacking.
For more information, see [IcaControl SPEC](x/ibcstaking/README.md).

### Oracle

The Oracle module manages the status (total delegation) of the validator of the counterpart zone to be delegated by Supernova.
The reason why this information is needed is to calculate the equity when issuing equity tokens(snAsset).
For more information, see [Oracle SPEC](x/oracle/README.md).

### Mint

The Mint modules are responsible for minting and distributing the Supernova's governance coin, Nova.
For more information, see [Mint SPEC](x/mint/README.md).

### Pool-Incentive

The Pool-Incentive module manages information to provide incentives to Supernova's liquidity providers.
For more information, see [Mint SPEC](x/mint/README.md).

### Airdrop

The Airdrop module is a module that manages information to incentivize early participants in Supernova.

---

## Scripts

### Build
If you want to build nova, use `make` scripts.

### Build Proto
If you want to build proto, use `make protogen-all` scripts.

### For testing
If you want to test nova, use `scripts/genesis_setup.sh` scripts.
This will create 3 validators with test genesis file.

---

## Community

Join our community!

|         |                                                        |
|---------|--------------------------------------------------------|
| Discord | [Go to Discord](https://discord.com/invite/2gj8fScWqD) |
| Twitter | [Go to Twitter](https://twitter.com/Supernovazone)     |
