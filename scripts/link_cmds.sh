#!/usr/bin/env bash

find ./cmd -type d -depth 1 | while read dir; do
  name=$(basename "${dir}")

  if [[ $name == "cron" ]]; then
     target="./out/@${name}";
  else
     target="./out/.${name}";
  fi
   echo ln ./out/msh "$target"
   ln ./out/msh "$target"
 done
