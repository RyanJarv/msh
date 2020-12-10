#!/usr/bin/env bash
set -euo pipefail
export PATH="$PWD/out:$PATH"

CMD=$1; REGEX=$2

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

function error() {
  result="$( $CMD )"
  printf "
${RED}[FAILED]${NC} ${CMD}
  Got:    '$result'
  Regex:  '$REGEX'

"
}
trap 'error $1' ERR

if [[ "$#" != "2" ]]; then
  echo "USAGE: $0 <cmd to execute> <regex of exected output>"
fi

$CMD | grep -qE "$REGEX"
printf "\n${GREEN}[PASSED]${NC} $CMD\n\n"