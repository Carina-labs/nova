#!/bin/bash
rm -rf $HOME/.novad/

# make four osmosis directories
mkdir $HOME/.novad
mkdir $HOME/.novad/validator

# init all three validators
./build/novad init --chain-id=testing validator --home=$HOME/.novad/validator

# create keys for all three validators
./build/novad keys add validator --keyring-backend=test --home=$HOME/.novad/validator

# change staking denom to uatom
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uatom"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
./build/novad add-genesis-account $(./build/novad keys show validator -a --keyring-backend=test --home=$HOME/.novad/validator) 100000000000uatom,100000000000stake --home=$HOME/.novad/validator
./build/novad gentx validator 500000000uatom --keyring-backend=test --home=$HOME/.novad/validator --chain-id=testing
./build/novad collect-gentxs --home=$HOME/.novad/validator

# update staking genesis
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# update crisis variable to uatom
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uatom"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# udpate gov genesis
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uatom"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# validator
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/validator/config/config.toml

./build/novad start --home=$HOME/.novad/validator

