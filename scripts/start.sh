#!/bin/bash
rm -rf $HOME/.novad
./build/novad unsafe-reset-all
MONIKER=test
KEY=mykey
KEYRING=test
CHAINID=novachain
novad init $MONIKER
novad config keyring-backend $KEYRING
novad keys add $KEY --keyring-backend=$keyring
novad add-genesis-account $KEY 1000000000stake,10000000000aphoton
novad gentx $KEY 1000000000stake --chain-id $CHAINID
novad collect-gentxs
novad start