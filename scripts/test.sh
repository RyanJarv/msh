#!/usr/bin/env bash
export PATH="$PWD/out:$PATH"

CMD=$1
EXPECTED=$2
NOT_EXPECTED=$3

set -euo pipefail

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

function error() {
  echo "[ERROR] Line: $1, error code: $2"
  exit 1
}
trap 'error $LINENO $?' ERR

if [[ "$#" -lt "2" ]] || [[ "$#" -gt "3" ]]; then
  echo "USAGE: $0 <cmd> <expected regex> [not expected regex]"
fi

echo "  * $CMD"

output="$(mktemp -d)/out"

DEBUG= script -q $output bash -c "$CMD" 1>/dev/null ||:

if cat $output | grep -qE "$EXPECTED"; then
  printf "      ${GREEN}[PASSED]${NC} %s\n" "$CMD"
else
  printf "${RED}[FAILED]${NC} %s
        Got:     '%s'
        Expected:   '%s'
" "${CMD}" "$(cat $output)" "$EXPECTED"
fi

if [[ -n "$NOT_EXPECTED" ]]; then
  if cat $output | grep -vqE "$NOT_EXPECTED"; then
    printf "      ${GREEN}[PASSED]${NC} %s\n" "$CMD"
  else
    printf "${RED}[FAILED]${NC} %s
          Got:     '%s'
          Not Expected:   '%s'
  " "${CMD}" "$(cat $output)" "$NOT_EXPECTED"
  fi
fi

rm $output