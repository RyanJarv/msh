#!/usr/bin/env bash

DIR=$1

mkdir -p "${DIR}"

cp ./out/msh "${DIR}/msh"

find ./cmd -type d -depth 1 | while read dir; do
  name=$(basename "${dir}")

  if [[ $name == "cron" ]]; then
     target="${DIR}/@${name}";
  else
     target="${DIR}/.${name}";
  fi
   echo ln "${DIR}/msh" "$target"
   ln "${DIR}/msh" "$target"
 done
