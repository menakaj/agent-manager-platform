#!/usr/bin/env bash
# Generate changelog between two tags with PR links
# Usage: generate-changelog.sh <last-tag> <current-tag> <repo-url>
# Outputs changelog to stdout

set -euo pipefail

LAST_TAG="${1:-}"
CURRENT_TAG="${2:-}"
REPO_URL="${3:-${GITHUB_SERVER_URL:-https://github.com}/${GITHUB_REPOSITORY:-}}"

if [ -z "$LAST_TAG" ] || [ -z "$CURRENT_TAG" ]; then
  echo "Error: Missing required arguments"
  echo "Usage: generate-changelog.sh <last-tag> <current-tag> [repo-url]"
  exit 1
fi

if git rev-parse "$LAST_TAG" >/dev/null 2>&1; then
  if git rev-parse "$CURRENT_TAG" >/dev/null 2>&1; then
    # Get commits between tags
    COMMITS=$(git log --pretty=format:"%s|%h|%b" "$LAST_TAG".."$CURRENT_TAG" 2>/dev/null)
    
    if [ -z "$COMMITS" ]; then
      echo "- No changes"
    else
      # Process each commit
      echo "$COMMITS" | while IFS='|' read -r subject hash body; do
        # Extract PR number from commit message (looks for #123 or (#123) or fixes #123, etc.)
        PR_NUM=$(echo "$subject $body" | grep -oE '(#[0-9]+|\(#[0-9]+\)|fixes? #[0-9]+|closes? #[0-9]+)' | grep -oE '[0-9]+' | head -1)
        
        if [ -n "$PR_NUM" ]; then
          # Format with PR link only
          echo "- $subject ([#$PR_NUM]($REPO_URL/pull/$PR_NUM))"
        else
          # Format without PR link
          echo "- $subject"
        fi
      done
    fi
  else
    echo "- Initial release"
  fi
else
  echo "- Initial release"
fi
