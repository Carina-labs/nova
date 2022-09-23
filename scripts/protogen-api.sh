#!/bin/bash

set -eo pipefail

echo "##################################"
echo "######## Cleaning API dir ########"
echo "##################################"
echo -e "\n"


if [ -d "api" ]; then
  cd api
else
  mkdir api
  cd api
fi

find ./ -type f \( -iname \*.pb.go -o -iname \*.cosmos_orm.go -o -iname \*.pb.gw.go \) -delete
find . -empty -type d -delete


cd ..
echo "##################################"
echo "#### Generating proto API set ####"
echo "##################################"
cd proto
buf mod update
buf generate --template buf.gen.api.yaml buf.build/cosmos/cosmos-sdk
buf generate --template buf.gen.api.yaml

go mod tidy -compat=1.18