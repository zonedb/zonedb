#!/usr/bin/env bash
#
# Update Go and TinyGo versions in go.mod and go.yaml.
# Runs go-versions-check.sh first to detect required versions.
#
# Requires: curl, jq, gh (authenticated), sed, go
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$(git -C "$SCRIPT_DIR" rev-parse --show-toplevel)"

# Run check script and extract KEY=VALUE pairs
CHECK_OUTPUT=$("$SCRIPT_DIR/go-versions-check.sh" 2>&1) || true
echo "$CHECK_OUTPUT" | grep -v '^[A-Z_]*='
eval "$(echo "$CHECK_OUTPUT" | grep '^[A-Z_]*=')"

if [ "$NEEDS_UPDATE" = "false" ]; then
  echo "Everything is up to date."
  exit 0
fi

echo ""
echo "=== Applying updates ==="

# Update go.mod minimum version
sed -i.bak "s/^go .*/go ${GO_PREV}.0/" go.mod && rm -f go.mod.bak
echo "Updated go.mod: go ${GO_PREV}.0"

# Update go.yaml TinyGo matrix: Go version and TinyGo versions
sed -i.bak "s/big: \"${CUR_BIG}\"/big: \"${TINYGO_GO}\"/g" .github/workflows/go.yaml && rm -f .github/workflows/go.yaml.bak
sed -i.bak "s/tiny: \"${CUR_TINY_PREV}\"/tiny: \"${TINYGO_PREV}\"/g" .github/workflows/go.yaml && rm -f .github/workflows/go.yaml.bak
sed -i.bak "s/tiny: \"${CUR_TINY_LATEST}\"/tiny: \"${TINYGO_LATEST}\"/g" .github/workflows/go.yaml && rm -f .github/workflows/go.yaml.bak
echo "Updated go.yaml: Go ${TINYGO_GO}, TinyGo ${TINYGO_PREV}/${TINYGO_LATEST}"

go mod tidy
echo "Ran go mod tidy"

echo ""
echo "Done. Review changes with: git diff"
