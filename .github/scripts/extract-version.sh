#!/usr/bin/env bash
# Extract release tag and version from GitHub release event
# Usage: extract-version.sh <tag_name>
# Outputs: tag and version to GitHub Actions outputs

set -euo pipefail

TAG="${1:-}"

if [ -z "$TAG" ]; then
  echo "Error: Tag name is required"
  exit 1
fi

# Remove 'v' prefix if present, otherwise use tag as-is
if [[ "$TAG" =~ ^v ]]; then
  VERSION="${TAG#v}"
else
  VERSION="$TAG"
fi

echo "tag=$TAG" >> "$GITHUB_OUTPUT"
echo "version=$VERSION" >> "$GITHUB_OUTPUT"

echo "Release tag: $TAG"
echo "Chart version: $VERSION"

