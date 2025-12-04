#!/usr/bin/env bash
# Create release artifacts (archive and checksum)
# Usage: create-release-artifacts.sh <tag> <source_dir> <output_prefix>

set -euo pipefail

TAG="${1:-}"
SOURCE_DIR="${2:-./quick-start}"
OUTPUT_PREFIX="${3:-quick-start}"

if [ -z "$TAG" ]; then
  echo "Error: Tag is required"
  exit 1
fi

if [ ! -d "$SOURCE_DIR" ]; then
  echo "Error: Source directory not found: $SOURCE_DIR"
  exit 1
fi

ARCHIVE_NAME="${OUTPUT_PREFIX}-${TAG}.tar.gz"
CHECKSUM_NAME="${ARCHIVE_NAME}.sha256"

# Create archive
cd "$SOURCE_DIR"
tar -czf "../${ARCHIVE_NAME}" .
cd ..

# Generate checksum
sha256sum "$ARCHIVE_NAME" > "$CHECKSUM_NAME"

echo "Created release artifacts:"
echo "  - $ARCHIVE_NAME"
echo "  - $CHECKSUM_NAME"
ls -lh "$ARCHIVE_NAME" "$CHECKSUM_NAME"

# Output artifact names for GitHub Actions
echo "archive=${ARCHIVE_NAME}" >> "$GITHUB_OUTPUT"
echo "checksum=${CHECKSUM_NAME}" >> "$GITHUB_OUTPUT"

