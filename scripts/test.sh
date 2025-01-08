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
tmp="$(mktemp -d)"
stdout="${tmp}/stdout"
stderr="${tmp}/stderr"

set +euo pipefail
bash -c "$CMD" 2> $stderr 1> $stdout

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
  " "${CMD}" "$status_code" "$(cat $stdout)"
fi

if [[ "$#" -eq "1" ]]; then
  exit 0
fi

if cat $stdout | grep -qE "$EXPECTED"; then
  printf "      ${GREEN}[PASSED]${NC} %s\n" "$CMD"
else
  if [[ -n $(cat "$stderr") ]]; then
    echo "** stderr **"
    cat "$stderr"
  fi

  printf "${RED}[FAILED]${NC} %s
        Got:     '%s'
        Expected:   '%s'
" "${CMD}" "$(cat $stdout)" "$EXPECTED"
fi

if [[ -n "$NOT_EXPECTED" ]]; then
  if cat $stdout | grep -vqE "$NOT_EXPECTED"; then
    printf "      ${GREEN}[PASSED]${NC} %s\n" "$CMD"
  else
    printf "${RED}[FAILED]${NC} %s
          Got:     '%s'
          Not Expected:   '%s'
  " "${CMD}" "$(cat $stdout)" "$NOT_EXPECTED"
  fi
fi

rm $stdout