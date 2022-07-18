#!/bin/bash
set -eo pipefail

WORKDIR=$(pwd)

if [ -d "./tmp-swagger-gen" ] 
then
    echo "temp dir were not deleted, let this try to remove it"
    rm -rf ./tmp-swagger-gen
    echo "done"
fi

mkdir -p ./tmp-swagger-gen

# cosmos chain generate docs
cd tmp-swagger-gen
git clone https://github.com/cosmos/cosmos-sdk.git --branch v0.45.4

# nova chain generate docs
cd $WORKDIR/proto
proto_dirs=$(find ./nova ../tmp-swagger-gen/cosmos-sdk/proto/cosmos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd $WORKDIR
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# move swagger.yaml to openapi.yml
mv ./client/docs/swagger-ui/swagger.yaml ./client/docs/swagger-ui/openapi.yml

# clean swagger files
rm -rf ./tmp-swagger-gen