#!/usr/bin/env bash
export PATH="$PWD/out:$PATH"

CMD=$1
EXPECTED=$2
NOT_EXPECTED=$3

set -euo pipefail

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color


if [[ "$#" -lt "1" ]] || [[ "$#" -gt "3" ]]; then
  echo "USAGE: $0 <cmd> <expected regex> [not expected regex]"
  exit 1
fi

echo "  * $CMD"

output="$(mktemp -d)/out"

set +euo pipefail
script -q $output bash -c "$CMD" 1>/dev/null
status_code=$?
set -euo pipefail

function error() {
  echo "[ERROR] Line: $1, error code: $2"
  exit 1
}
trap 'error $LINENO $?' ERR

if [[ $status_code -ne 0 ]]; then
  printf "${RED}[FAILED]${NC} %s
        StatusCode:     '%s'
        Got:     '%s'
  " "${CMD}" "$status_code" "$(cat $output)"
fi

if [[ "$#" -eq "1" ]]; then
  exit 0
fi

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