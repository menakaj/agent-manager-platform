#!/usr/bin/env bash
# Update Helm chart version and image tags, then package and push
# Usage: update-helm-chart.sh <chart-name> <chart-dir> <version> <tag> <helm-registry> <registry> <registry-org> <github-token> [image-updates-json]

set -euo pipefail

CHART_NAME="${1:-}"
CHART_DIR="${2:-}"
VERSION="${3:-}"
TAG="${4:-}"
HELM_REGISTRY="${5:-}"
REGISTRY="${6:-ghcr.io}"
REGISTRY_ORG="${7:-}"
GITHUB_TOKEN="${8:-}"
IMAGE_UPDATES="${9:-[]}"

if [ -z "$CHART_NAME" ] || [ -z "$CHART_DIR" ] || [ -z "$VERSION" ] || [ -z "$TAG" ] || [ -z "$HELM_REGISTRY" ] || [ -z "$REGISTRY_ORG" ] || [ -z "$GITHUB_TOKEN" ]; then
  echo "Error: Missing required arguments"
  echo "Usage: update-helm-chart.sh <chart-name> <chart-dir> <version> <tag> <helm-registry> <registry> <registry-org> <github-token> [image-updates-json]"
  exit 1
fi

if [ ! -d "$CHART_DIR" ]; then
  echo "Error: Chart directory not found: $CHART_DIR"
  exit 1
fi

# Update Chart.yaml
if [ -f "$CHART_DIR/Chart.yaml" ]; then
  yq eval -i ".version = \"$VERSION\"" "$CHART_DIR/Chart.yaml"
  yq eval -i ".appVersion = \"$TAG\"" "$CHART_DIR/Chart.yaml"
  echo "Updated Chart.yaml: version=$VERSION, appVersion=$TAG"
else
  echo "Error: Chart.yaml not found in $CHART_DIR"
  exit 1
fi

# Update values.yaml images if provided
if [ -f "$CHART_DIR/values.yaml" ] && [ "$IMAGE_UPDATES" != "[]" ]; then
  echo "$IMAGE_UPDATES" | yq eval -P - | while IFS= read -r update; do
    if [ -n "$update" ]; then
      path=$(echo "$update" | yq eval '.path' -)
      repo=$(echo "$update" | yq eval '.repository' -)
      img_tag=$(echo "$update" | yq eval '.tag' -)
      
      if [ -n "$path" ] && [ -n "$repo" ] && [ -n "$img_tag" ]; then
        yq eval -i "$path.repository = \"$repo\"" "$CHART_DIR/values.yaml" || true
        yq eval -i "$path.tag = \"$img_tag\"" "$CHART_DIR/values.yaml" || true
        echo "Updated image: $path -> $repo:$img_tag"
      fi
    fi
  done
else
  echo "Skipping values.yaml update (file not found or no image updates specified)"
fi

# Log in to registry
# GITHUB_ACTOR is set by GitHub Actions, fallback to 'github-actions' if not set
ACTOR="${GITHUB_ACTOR:-github-actions}"
echo "$GITHUB_TOKEN" | helm registry login -u "$ACTOR" --password-stdin "$REGISTRY"

# Package and push
helm package "$CHART_DIR" --version "$VERSION"
helm push "${CHART_NAME}-${VERSION}.tgz" "${HELM_REGISTRY}/${CHART_NAME}"
echo "Pushed $CHART_NAME chart version $VERSION"

