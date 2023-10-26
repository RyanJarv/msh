#!/usr/bin/env bash

function error() {
  echo "[ERROR] Line: $1, error code: $2"
  exit 1
}
trap 'error $LINENO $?' ERR

set -euo pipefail

names=$(aws schemas list-schemas --registry-name aws.events --query 'Schemas[*].SchemaName' --out text)

mkdir -p schema/json

for name in $names; do
#  aws schemas get-code-binding-source \
#    --language Go1 \
#    --registry-name aws.events \
#    --schema-name "$name" \
#    /dev/stdout \
#  | tar -xvf - schema
  aws schemas describe-schema \
    --registry-name aws.events \
    --schema-name "$name" \
    --query Content \
    --out text \
    > "schema/json/${name}.json"
 done
