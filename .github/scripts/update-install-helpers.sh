#!/usr/bin/env bash
# Update chart versions in install-helpers.sh
# Usage: update-install-helpers.sh <version> <file_path>

set -euo pipefail

VERSION="${1:-}"
FILE="${2:-./quick-start/install-helpers.sh}"

if [ -z "$VERSION" ]; then
  echo "Error: Version is required"
  exit 1
fi

if [ ! -f "$FILE" ]; then
  echo "Error: File not found: $FILE"
  exit 1
fi

# Create backup
cp "$FILE" "${FILE}.bak"

# Update AMP_CHART_VERSION
sed -i.bak "s/\(AMP_CHART_VERSION=\"\${AMP_CHART_VERSION:-\)[^}]*\(}\"\)/\1${VERSION}\2/" "$FILE"

# Update BUILD_CI_CHART_VERSION
sed -i.bak "s/\(BUILD_CI_CHART_VERSION=\"\${BUILD_CI_CHART_VERSION:-\)[^}]*\(}\"\)/\1${VERSION}\2/" "$FILE"

# Update OBSERVABILITY_CHART_VERSION
sed -i.bak "s/\(OBSERVABILITY_CHART_VERSION=\"\${OBSERVABILITY_CHART_VERSION:-\)[^}]*\(}\"\)/\1${VERSION}\2/" "$FILE"

# Remove backup files
rm -f "${FILE}.bak"

echo "Updated chart versions in $FILE to ${VERSION}"
echo "Verifying changes:"
grep -E "(AMP_CHART_VERSION|BUILD_CI_CHART_VERSION|OBSERVABILITY_CHART_VERSION)" "$FILE" | head -3

