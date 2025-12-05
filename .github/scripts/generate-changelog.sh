#!/usr/bin/env bash
# Generate changelog between two tags
# Usage: generate-changelog.sh <last-tag> <current-tag>
# Outputs changelog to stdout

set -euo pipefail

LAST_TAG="${1:-}"
CURRENT_TAG="${2:-}"

if [ -z "$LAST_TAG" ] || [ -z "$CURRENT_TAG" ]; then
  echo "Error: Missing required arguments"
  echo "Usage: generate-changelog.sh <last-tag> <current-tag>"
  exit 1
fi

if git rev-parse "$LAST_TAG" >/dev/null 2>&1; then
  if git rev-parse "$CURRENT_TAG" >/dev/null 2>&1; then
    git log --pretty=format:"- %s (%h)" "$LAST_TAG".."$CURRENT_TAG" 2>/dev/null || echo "- No changes"
  else
    echo "- Initial release"
  fi
else
  echo "- Initial release"
fi

