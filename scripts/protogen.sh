#!/bin/bash

#buf generate
#
#cp -r ./gen/proto/go/novachain/gal/v1/* ./x/gal/types
#rm -rf gen

set -eo pipefail

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./novachain -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..

cp -r github.com/Carina-labs/novachain/* ./
rm -rf github.com

go mod tidy -compat=1.18