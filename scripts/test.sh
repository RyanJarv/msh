#!/usr/bin/env bash
set -euo pipefail
export PATH="$PWD/out:$PATH"

CMD=$1; REGEX=$2

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

function error() {
  echo "[ERROR] Line: $1, error code: $2"
}
trap 'error $LINENO $?' ERR

if [[ "$#" != "2" ]]; then
  echo "USAGE: $0 <cmd to execute> <regex of exected output>"
fi

echo "  * $CMD"

output="$(mktemp -d)/out"

DEBUG= script -q $output bash -c "$CMD" 1>/dev/null ||:

if cat $output | grep -qE "$REGEX"; then
  printf "      ${GREEN}[PASSED]${NC} %s\n" "$CMD"
else
  printf "${RED}[FAILED]${NC} %s
        Got:     '%s'
        Regex:   '%s'
" "${CMD}" "$(cat $output)" "$REGEX"
fi

rm $output