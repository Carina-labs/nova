rm -rf ~/.novad

MONIKER=novatest
CHAINID=nova
KEYRING=test

#MONIKER : test, CHAINID : nova
novad init $MONIKER --chain-id $CHAINID

novad config keyring-backend test

novad keys add genkey
novad keys add relayer
novad keys add victor
novad keys add siu
novad keys add eric
novad keys add lucid
novad keys add harry
novad keys add johnhan
novad keys add sam
novad keys add sangyun
novad keys add sojeong
novad keys add ellie
novad keys add boosik

novad add-genesis-account $(novad keys show genkey -a) 10000000000unova
novad add-genesis-account $(novad keys show relayer -a) 10000000000unova
novad add-genesis-account $(novad keys show victor -a) 10000000000unova
novad add-genesis-account $(novad keys show siu -a) 10000000000unova
novad add-genesis-account $(novad keys show eric -a) 10000000000unova
novad add-genesis-account $(novad keys show lucid -a) 10000000000unova 
novad add-genesis-account $(novad keys show harry -a) 10000000000unova
novad add-genesis-account $(novad keys show johnhan -a) 10000000000unova 
novad add-genesis-account $(novad keys show sam -a) 10000000000unova 
novad add-genesis-account $(novad keys show sangyun -a) 10000000000unova
novad add-genesis-account $(novad keys show sojeong -a) 10000000000unova
novad add-genesis-account $(novad keys show ellie -a) 10000000000unova 
novad add-genesis-account $(novad keys show boosik -a) 10000000000unova

novad gentx genkey 1000000000unova --chain-id nova


novad collect-gentxs

# update staking genesis
cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
# cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="300s"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["staking"]["params"]["max_entries"]="10"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json

# update mint denom
cat $HOME/.novad/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json

# update crisis variable to unova
cat $HOME/.novad/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json

# udpate gov genesis
cat $HOME/.novad/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json
cat $HOME/.novad/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="unova"' > $HOME/.novad/config/tmp_genesis.json && mv $HOME/.novad/config/tmp_genesis.json $HOME/.novad/config/genesis.json