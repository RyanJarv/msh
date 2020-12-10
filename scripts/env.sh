#!/usr/bin/env bash
set -euo pipefail
export PATH="$PWD/out:$PATH"

exec "${@}"