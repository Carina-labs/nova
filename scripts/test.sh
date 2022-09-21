#!/bin/bash
rm -rf $HOME/.novad/
MONIKER=novatest
CHAINID=nova
KEYRING=test
#MONIKER : test, CHAINID : nova
novad init $MONIKER --chain-id $CHAINID
novad config keyring-backend $KEYRING
novad keys add genkey --recover
novad keys add relayer --recover
novad keys add daomodifier --recover
novad add-genesis-account $(novad keys show genkey -a) 10000000000unova
novad add-genesis-account $(novad keys show relayer -a) 10000000000unova
novad add-genesis-account $(novad keys show daomodifier -a) 10000000000unova
novad gentx genkey 1000000000unova --chain-id nova
novad collect-gentxs
# update staking genesis
cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="600s"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["max_entries"]="1000000"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# update mint denom
cat $HOME/.novad/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="nova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["mint"]["params"]["distribution_proportions"]["staking"]="0.5"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["mint"]["params"]["distribution_proportions"]["lp_incentives"]="0.4"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["mint"]["params"]["distribution_proportions"]["community_pool"]="0.1"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# update crisis variable to unova
cat $HOME/.novad/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# udpate gov genesis
cat $HOME/.novad/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# update inter-tx params
cat $HOME/.novad/config/genesis.json | jq '.app_state["icacontrol"]["params"]["controller_address"][0]="nova1ygwjtmvtxqpjy7jklt5qqa9ykp2ya25dxshuvt"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
#update oracle params
cat $HOME/.novad/config/genesis.json | jq '.app_state["oracle"]["params"]["oracle_operators"][0]="nova1ygwjtmvtxqpjy7jklt5qqa9ykp2ya25dxshuvt"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# validator
# enable rest api server & unsafe cors
sed -i -E '|enable = false|enable = true|g' $HOME/.novad/config/app.toml
sed -i -E '|enabled-unsafe-cors = false|enabled-unsafe-cors = true|g' $HOME/.novad/config/app.toml
# allow duplicate ip
sed -i -E '|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"*\"\]|g' $HOME/.novad/config/config.toml
novad start