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
novad keys add controller --recover
novad keys add front --recover

novad add-genesis-account $(novad keys show genkey -a) 10000000000unova
novad add-genesis-account $(novad keys show relayer -a) 10000000000unova
novad add-genesis-account $(novad keys show controller -a) 10000000000unova
novad add-genesis-account $(novad keys show front -a) 10000000000unova

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
# update icacontrol params
cat $HOME/.novad/config/genesis.json | jq '.app_state["icacontrol"]["params"]["controller_key_manager"][0]="nova1lds58drg8lvnaprcue2sqgfvjnz5ljlkq9lsyf"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["icacontrol"]["controller_address_info"][0]["zone_id"]="gaia"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["icacontrol"]["controller_address_info"][0]["controller_address"]="nova1lds58drg8lvnaprcue2sqgfvjnz5ljlkq9lsyf"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
#update oracle params
cat $HOME/.novad/config/genesis.json | jq '.app_state["oracle"]["params"]["oracle_key_manager"][0]="nova1lds58drg8lvnaprcue2sqgfvjnz5ljlkq9lsyf"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["oracle"]["oracle_address_info"][0]["zone_id"]="gaia"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["oracle"]["oracle_address_info"][0]["oracle_address"][0]="nova1lds58drg8lvnaprcue2sqgfvjnz5ljlkq9lsyf"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# validator
# enable rest api server & unsafe cors
sed -i -E 's|enable = false|enable = true|g' $HOME/.novad/config/app.toml
sed -i -E 's|swagger = false|swagger = true|g' $HOME/.novad/config/app.toml
sed -i -E 's|enabled-unsafe-cors = false|enabled-unsafe-cors = true|g' $HOME/.novad/config/app.toml
sed -i -E 's|minimum-gas-prices = \"\"|minimum-gas-prices = \"0unova\"|g' $HOME/.novad/config/app.toml
# allow duplicate ip
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"*\"\]|g' $HOME/.novad/config/config.toml
sed -i -E 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|g' $HOME/.novad/config/config.toml