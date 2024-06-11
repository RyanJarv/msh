#!/usr/bin/env bash

arn=$(
  ./lambda.py | msh build
)

aws lambda invoke \
  --function-name "$arn" \
  --payload '{"key1":"value1", "key2":"value2", "key3":"value3"}'
