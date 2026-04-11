#!/usr/bin/env bash
#
# Detect current and required Go/TinyGo versions.
# Prints a summary and exits 0 if up to date, 1 if updates are needed.
#
# Requires: curl, jq, gh (authenticated), sed
#
set -euo pipefail

cd "$(git -C "$(dirname "$0")" rev-parse --show-toplevel)"

# Two most recent stable Go major.minor versions
# e.g. "1.25\n1.26"
GO_VERSIONS=$(curl -sS 'https://go.dev/dl/?mode=json' | \
  jq -r '[.[] | select(.stable) | .version | ltrimstr("go") | split(".")[0:2] | join(".")] | unique | sort | .[-2:][]')
GO_PREV=$(echo "$GO_VERSIONS" | head -1)   # e.g. "1.25"
GO_CURRENT=$(echo "$GO_VERSIONS" | tail -1) # e.g. "1.26"

# Two most recent TinyGo releases with distinct minor versions
# e.g. "0.39.0\n0.40.1" (latest patch of each minor)
TINYGO_VERSIONS=$(gh api repos/tinygo-org/tinygo/releases --jq \
  '[.[] | select(.prerelease | not) | .tag_name | ltrimstr("v")] | group_by(split(".")[0:2] | join(".")) | map(sort | last) | sort | .[-2:][]')
TINYGO_PREV=$(echo "$TINYGO_VERSIONS" | head -1)   # e.g. "0.39.0"
TINYGO_LATEST=$(echo "$TINYGO_VERSIONS" | tail -1) # e.g. "0.40.1"

# Highest Go minor version TinyGo supports, e.g. "1.25"
# Parsed from TinyGo source: `const minorMax = 25` -> "1.25"
TINYGO_MAX_GO="1.$(curl -sS 'https://raw.githubusercontent.com/tinygo-org/tinygo/release/builder/config.go' | \
  sed -n 's/.*minorMax[[:space:]]*=[[:space:]]*\([0-9]*\).*/\1/p')"

# TinyGo tests should use the oldest supported Go that TinyGo can handle
if [ "$(printf '%s\n%s' "$GO_PREV" "$TINYGO_MAX_GO" | sort -V | head -1)" = "$GO_PREV" ]; then
  TINYGO_GO="$GO_PREV"
else
  echo "WARNING: TinyGo $TINYGO_LATEST only supports up to Go $TINYGO_MAX_GO (older than previous stable $GO_PREV)"
  TINYGO_GO="$TINYGO_MAX_GO"
fi

# What's currently in the repo
CUR_BIG=$(sed -n 's/.*big: "\([0-9.]*\)".*/\1/p' .github/workflows/go.yaml | head -1)
CUR_TINY_PREV=$(sed -n 's/.*tiny: "\([0-9.]*\)".*/\1/p' .github/workflows/go.yaml | sort -V | uniq | head -1)
CUR_TINY_LATEST=$(sed -n 's/.*tiny: "\([0-9.]*\)".*/\1/p' .github/workflows/go.yaml | sort -V | uniq | tail -1)
CUR_GOMOD=$(sed -n 's/^go \([0-9.]*\).*/\1/p' go.mod)

echo "=== Go & TinyGo Version Check ==="
echo "Go stable releases:       $GO_PREV, $GO_CURRENT"
echo "TinyGo releases:          $TINYGO_PREV, $TINYGO_LATEST (supports up to Go $TINYGO_MAX_GO)"
echo ""
echo "go.mod minimum:           $CUR_GOMOD -> ${GO_PREV}.0"
echo "go.yaml CI Go (TinyGo):   $CUR_BIG -> $TINYGO_GO"
echo "go.yaml CI TinyGo:        $CUR_TINY_PREV, $CUR_TINY_LATEST -> $TINYGO_PREV, $TINYGO_LATEST"

# Check if updates are needed
NEEDS_UPDATE=false
if [ "$CUR_BIG" != "$TINYGO_GO" ]; then NEEDS_UPDATE=true; fi
if [ "$CUR_TINY_LATEST" != "$TINYGO_LATEST" ]; then NEEDS_UPDATE=true; fi
if [ "$CUR_GOMOD" != "${GO_PREV}.0" ]; then NEEDS_UPDATE=true; fi
echo ""
echo "Needs update: $NEEDS_UPDATE"

# Machine-readable values (used by go-versions-update.sh)
echo ""
echo "GO_PREV=$GO_PREV"
echo "TINYGO_GO=$TINYGO_GO"
echo "TINYGO_PREV=$TINYGO_PREV"
echo "TINYGO_LATEST=$TINYGO_LATEST"
echo "CUR_BIG=$CUR_BIG"
echo "CUR_TINY_PREV=$CUR_TINY_PREV"
echo "CUR_TINY_LATEST=$CUR_TINY_LATEST"
echo "NEEDS_UPDATE=$NEEDS_UPDATE"

if [ "$NEEDS_UPDATE" = "true" ]; then exit 1; else exit 0; fi
