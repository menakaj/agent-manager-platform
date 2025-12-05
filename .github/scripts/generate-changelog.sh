#!/usr/bin/env bash
# Generate changelog between two tags with PR links
# Usage: generate-changelog.sh <last-tag> <current-tag> <repo-url>
# Outputs changelog to stdout

set -euo pipefail

LAST_TAG="${1:-}"
CURRENT_TAG="${2:-}"
REPO_URL="${3:-${GITHUB_SERVER_URL:-https://github.com}/${GITHUB_REPOSITORY:-}}"

if [ -z "$LAST_TAG" ] || [ -z "$CURRENT_TAG" ]; then
  echo "Error: Missing required arguments" >&2
  echo "Usage: generate-changelog.sh <last-tag> <current-tag> [repo-url]" >&2
  exit 1
fi

echo "Generating changelog from $LAST_TAG to $CURRENT_TAG" >&2
echo "Repository URL: $REPO_URL" >&2

# Check if last tag exists
if ! git rev-parse "$LAST_TAG" >/dev/null 2>&1; then
  echo "Last tag '$LAST_TAG' not found, treating as initial release" >&2
  echo "- Initial release"
  exit 0
fi

# Check if current tag exists
if ! git rev-parse "$CURRENT_TAG" >/dev/null 2>&1; then
  echo "Current tag '$CURRENT_TAG' not found, treating as initial release" >&2
  echo "- Initial release"
  exit 0
fi

echo "Both tags found, fetching commits..." >&2

# Get commits between tags
COMMITS=$(git log --pretty=format:"%s|%h|%b" "$LAST_TAG".."$CURRENT_TAG" 2>/dev/null || true)

if [ -z "$COMMITS" ]; then
  echo "No commits found between tags" >&2
  echo "- No changes"
  exit 0
fi

# Count commits
COMMIT_COUNT=$(echo "$COMMITS" | grep -c . || echo "0")
echo "Found $COMMIT_COUNT commit(s) to process" >&2

# Process each commit
PROCESSED=0
echo "$COMMITS" | while IFS='|' read -r subject hash body || [ -n "$subject" ]; do
  # Skip empty lines
  if [ -z "$subject" ]; then
    continue
  fi
  
  PROCESSED=$((PROCESSED + 1))
  
  # Extract PR number from commit message (looks for #123 or (#123) or fixes #123, etc.)
  PR_NUM=$(echo "$subject $body" | grep -oE '(#[0-9]+|\(#[0-9]+\)|fixes? #[0-9]+|closes? #[0-9]+)' 2>/dev/null | grep -oE '[0-9]+' 2>/dev/null | head -1 || true)
  
  if [ -n "$PR_NUM" ]; then
    # Format with PR link only
    echo "- $subject ([#$PR_NUM]($REPO_URL/pull/$PR_NUM))"
    echo "[INFO] Commit $PROCESSED/$COMMIT_COUNT: Found PR #$PR_NUM" >&2
  else
    # Format without PR link
    echo "- $subject"
    echo "[INFO] Commit $PROCESSED/$COMMIT_COUNT: No PR number found" >&2
  fi
done || true

echo "[INFO] Changelog generation completed successfully" >&2
