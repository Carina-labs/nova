#!/bin/bash
rm -rf $HOME/.novad/

# make four osmosis directories
mkdir $HOME/.novad
mkdir $HOME/.novad/validator1
mkdir $HOME/.novad/validator2
mkdir $HOME/.novad/validator3

# init all three validators
./build/novad init --chain-id=testing validator1 --home=$HOME/.novad/validator1
./build/novad init --chain-id=testing validator2 --home=$HOME/.novad/validator2
./build/novad init --chain-id=testing validator3 --home=$HOME/.novad/validator3
# create keys for all three validators
./build/novad keys add validator1 --keyring-backend=test --home=$HOME/.novad/validator1
./build/novad keys add validator2 --keyring-backend=test --home=$HOME/.novad/validator2
./build/novad keys add validator3 --keyring-backend=test --home=$HOME/.novad/validator3

# change staking denom to unova
cat $HOME/.novad/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="unova"' > $HOME/.novad/validator1/config/tmp_genesis.json && mv $HOME/.novad/validator1/config/tmp_genesis.json $HOME/.novad/validator1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
./build/novad add-genesis-account $(./build/novad keys show validator1 -a --keyring-backend=test --home=$HOME/.novad/validator1) 100000000000unova,100000000000stake --home=$HOME/.novad/validator1
./build/novad gentx validator1 500000000unova --keyring-backend=test --home=$HOME/.novad/validator1 --chain-id=testing
./build/novad collect-gentxs --home=$HOME/.novad/validator1

# update staking genesis
cat $HOME/.novad/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.novad/validator1/config/tmp_genesis.json && mv $HOME/.novad/validator1/config/tmp_genesis.json $HOME/.novad/validator1/config/genesis.json

# update crisis variable to unova
cat $HOME/.novad/validator1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="unova"' > $HOME/.novad/validator1/config/tmp_genesis.json && mv $HOME/.novad/validator1/config/tmp_genesis.json $HOME/.novad/validator1/config/genesis.json

# udpate gov genesis
cat $HOME/.novad/validator1/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.novad/validator1/config/tmp_genesis.json && mv $HOME/.novad/validator1/config/tmp_genesis.json $HOME/.novad/validator1/config/genesis.json
cat $HOME/.novad/validator1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="unova"' > $HOME/.novad/validator1/config/tmp_genesis.json && mv $HOME/.novad/validator1/config/tmp_genesis.json $HOME/.novad/validator1/config/genesis.json

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $HOME/.novad/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.novad/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.novad/validator2/config/app.toml

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $HOME/.novad/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.novad/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.novad/validator3/config/app.toml

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/validator1/config/config.toml

# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $HOME/.novad/validator2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $HOME/.novad/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $HOME/.novad/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.novad/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/validator2/config/config.toml

# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $HOME/.novad/validator3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $HOME/.novad/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.novad/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.novad/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.novad/validator3/config/config.toml

# copy validator1 genesis file to validator2-3
cp $HOME/.novad/validator1/config/genesis.json $HOME/.novad/validator2/config/genesis.json
cp $HOME/.novad/validator1/config/genesis.json $HOME/.novad/validator3/config/genesis.json

# copy tendermint node id of validator1 to persistent peers of validator2-3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(./build/novad tendermint show-node-id --home=$HOME/.novad/validator1)@localhost:26656\"|g" $HOME/.novad/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(./build/novad tendermint show-node-id --home=$HOME/.novad/validator1)@localhost:26656\"|g" $HOME/.novad/validator3/config/config.toml

nohup ./build/novad start --home=$HOME/.novad/validator1 > nova1.log &
nohup ./build/novad start --home=$HOME/.novad/validator2 > nova2.log &
nohup ./build/novad start --home=$HOME/.novad/validator3 > nova3.log &

