#!/bin/bash

set -eo pipefail

echo "##################################"
echo "######## Cleaning API dir ########"
echo "##################################"
echo -e "\n"

find ./ -type f \( -iname \*.pb.go -o -iname \*.pb.gw.go \) -delete


echo "##################################"
echo "### Generating gogo proto code ###"
echo "##################################"

cd proto
proto_dirs=$(find ./nova -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      echo $file
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..

cp -r github.com/Carina-labs/nova/* ./
rm -rf github.com

go mod tidy -compat=1.18