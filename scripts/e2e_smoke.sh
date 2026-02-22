#!/usr/bin/env bash
set -euo pipefail

BIN="${BIN:-./bin/forge}"
TMP="$(mktemp -d)"

"$BIN" version
"$BIN" doctor --no-interactive || true
"$BIN" new --framework django --name e2e-dj --path "$TMP" --python-version 3.11 --env-manager venv --extras drf,docker,ci --no-interactive --dry-run
"$BIN" new --framework next --name e2e-next --path "$TMP" --extras tailwind --no-interactive --dry-run

echo "E2E smoke complete"
