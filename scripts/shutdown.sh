#!/bin/bash

set -o pipefail
set -e

got=$(ps -efl | grep 'novad' | awk '{print $2}')

echo "Shutdown novad validators"
echo ""
echo "[Notice]"
echo "You may see an error message that there is no process."
echo "But there is no problem with the behavior."
echo ""

for prc in $got
do
  kill -9 $prc || true
done