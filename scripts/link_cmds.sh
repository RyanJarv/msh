#!/usr/bin/env bash

find ./cmd -type d -depth 1 | while read dir; do
  name=$(basename "${dir}"); ln ./out/msh "./out/.${name}";
 done
