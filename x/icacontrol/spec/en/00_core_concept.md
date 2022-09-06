# Core Concept

---

## Overview

The `icacontrol` module stores metadata about the relative chain associated with the Supernova protocol.

And this module also implements an interface that propagates IBC or ICA messages.


### Zone

`Zone` means each chain to the IBC. Supernova can also be thought of as a zone


### IBC (Inter-Blockchain Communication)
`IBC` is a protocol that sends and receives messages between zones. For more information on IBC, see https://ibc.cosmos.network/.


### ICA (Inter-Chain Account)
`ICA` is one of the use-case of `IBC`, a technology that allows messages to be delivered remotely by connecting the accounts of the controller chain (Supernova) and the host chain.

Supernova uses `ICA` to delegate the amount in the host chain using the controller chain. At this time, ICA allows you to send a delegation message to the controller chain to manipulate the same behavior in the host chain.
