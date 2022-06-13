#!/bin/sh
MONIKER=test
KEY=mykey
KEYRING=test
CHAINID=novachain
novachaind init $MONIKER
novachaind config keyring-backend $KEYRING
novachaind keys add $KEY --keyring-backend=$KEYRING
novachaind add-genesis-account $KEY 1000000000stake,10000000000aphoton --keyring-backend=$KEYRING
novachaind gentx $KEY 1000000000stake --chain-id $CHAINID
novachaind collect-gentxs

