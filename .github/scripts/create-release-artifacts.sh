#!/usr/bin/env bash
# Create release artifacts (archive and checksum)
# Usage: create-release-artifacts.sh <tag> <source_dir> <output_prefix>
# Outputs: archive and checksum filenames to GITHUB_OUTPUT

set -euo pipefail

TAG="${1:-}"
SOURCE_DIR="${2:-./quick-start}"
OUTPUT_PREFIX="${3:-quick-start}"

if [ -z "$TAG" ]; then
  echo "Error: Tag is required"
  exit 1
fi

if [ ! -d "$SOURCE_DIR" ]; then
  echo "⚠️ Source directory not found: $SOURCE_DIR, skipping artifacts"
  exit 0
fi

ARCHIVE_NAME="${OUTPUT_PREFIX}-${TAG}.tar.gz"
CHECKSUM_NAME="${ARCHIVE_NAME}.sha256"

# Create archive
tar -czf "$ARCHIVE_NAME" -C "$SOURCE_DIR" .

# Generate checksum
sha256sum "$ARCHIVE_NAME" > "$CHECKSUM_NAME"

echo "Created release artifacts:"
echo "  - $ARCHIVE_NAME"
echo "  - $CHECKSUM_NAME"

# Output artifact names for GitHub Actions
echo "archive=$ARCHIVE_NAME" >> "$GITHUB_OUTPUT"
echo "checksum=$CHECKSUM_NAME" >> "$GITHUB_OUTPUT"

