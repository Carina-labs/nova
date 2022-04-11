#!/bin/bash
rm -rf $HOME/.novachaind/


# make four osmosis directories
mkdir $HOME/.novachaind
mkdir $HOME/.novachaind/validator1
mkdir $HOME/.novachaind/validator2
mkdir $HOME/.novachaind/validator3

# init all three validators
./build/bin/novachaind init --chain-id=testing validator1 --home=$HOME/.novachaind/validator1
./build/bin/novachaind init --chain-id=testing validator2 --home=$HOME/.novachaind/validator2
./build/bin/novachaind init --chain-id=testing validator3 --home=$HOME/.novachaind/validator3
# create keys for all three validators
./build/bin/novachaind keys add validator1 --keyring-backend=test --home=$HOME/.novachaind/validator1
./build/bin/novachaind keys add validator2 --keyring-backend=test --home=$HOME/.novachaind/validator2
./build/bin/novachaind keys add validator3 --keyring-backend=test --home=$HOME/.novachaind/validator3

# change staking denom to uatom
cat $HOME/.novachaind/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uatom"' > $HOME/.novachaind/validator1/config/tmp_genesis.json && mv $HOME/.novachaind/validator1/config/tmp_genesis.json $HOME/.novachaind/validator1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
./build/bin/novachaind add-genesis-account $(./build/bin/novachaind keys show validator1 -a --keyring-backend=test --home=$HOME/.novachaind/validator1) 100000000000uatom,100000000000stake --home=$HOME/.novachaind/validator1
./build/bin/novachaind gentx validator1 500000000uatom --keyring-backend=test --home=$HOME/.novachaind/validator1 --chain-id=testing
./build/bin/novachaind collect-gentxs --home=$HOME/.novachaind/validator1


# update staking genesis
cat $HOME/.novachaind/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.novachaind/validator1/config/tmp_genesis.json && mv $HOME/.novachaind/validator1/config/tmp_genesis.json $HOME/.novachaind/validator1/config/genesis.json

# update crisis variable to uatom
cat $HOME/.novachaind/validator1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uatom"' > $HOME/.novachaind/validator1/config/tmp_genesis.json && mv $HOME/.novachaind/validator1/config/tmp_genesis.json $HOME/.novachaind/validator1/config/genesis.json

# udpate gov genesis
cat $HOME/.novachaind/validator1/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.novachaind/validator1/config/tmp_genesis.json && mv $HOME/.novachaind/validator1/config/tmp_genesis.json $HOME/.novachaind/validator1/config/genesis.json
cat $HOME/.novachaind/validator1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uatom"' > $HOME/.novachaind/validator1/config/tmp_genesis.json && mv $HOME/.novachaind/validator1/config/tmp_genesis.json $HOME/.novachaind/validator1/config/genesis.json

# copy validator1 genesis file to validator2-3
cp $HOME/.novachaind/validator1/config/genesis.json $HOME/.novachaind/validator2/config/genesis.json
cp $HOME/.novachaind/validator1/config/genesis.json $HOME/.novachaind/validator3/config/genesis.json