#!/bin/bash

buf generate

cp -r ./gen/proto/go/novachain/gal/v1/* ./x/gal/types
rm -rf gen