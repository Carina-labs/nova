#!/bin/sh
rm -rf $HOME/.novad/

# make four osmosis directories
mkdir $HOME/.novad
mkdir $HOME/.novad/validator

# init all three validators
/bin/novad init --chain-id=novatest validator --home=$HOME/.novad/validator

# create keys for all three validators
/bin/novad keys add validator --keyring-backend=test --home=$HOME/.novad/validator

# change staking denom to unova
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="unova"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
/bin/novad add-genesis-account $(/bin/novad keys show validator -a --keyring-backend=test --home=$HOME/.novad/validator) 100000000000unova,100000000000stake --home=$HOME/.novad/validator
/bin/novad gentx validator 500000000unova --keyring-backend=test --home=$HOME/.novad/validator --chain-id=novatest
/bin/novad collect-gentxs --home=$HOME/.novad/validator

# update staking genesis
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# update crisis variable to unova
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="unova"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# udpate gov genesis
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="unova"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# update oracle params
cat $HOME/.novad/validator/config/genesis.json | jq '.app_state["oracle"]["params"]["oracle_operators"][0]="nova1lds58drg8lvnaprcue2sqgfvjnz5ljlkq9lsyf"' > $HOME/.novad/validator/config/tmp_genesis.json && mv $HOME/.novad/validator/config/tmp_genesis.json $HOME/.novad/validator/config/genesis.json

# validator
# enable rest api server & unsafe cors
sed -i -E 's|enable = false|enable = true|g' $HOME/.novad/validator/config/app.toml
sed -i -E 's|enabled-unsafe-cors = false|enabled-unsafe-cors = true|g' $HOME/.novad/validator/config/app.toml

# allow duplicate ip
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/validator/config/config.toml

/bin/novad start --home=$HOME/.novad/validator

