#!/usr/bin/env bash
# Cleanup script for atomic release - removes pushed images and charts on failure
# Usage: cleanup-release.sh <registry> <registry-org> <tag> <helm-registry> <github-token> [services] [charts]

set -euo pipefail

REGISTRY="${1:-ghcr.io}"
REGISTRY_ORG="${2:-}"
TAG="${3:-}"
HELM_REGISTRY="${4:-}"
GITHUB_TOKEN="${5:-}"
SERVICES="${6:-}"
CHARTS="${7:-}"

if [ -z "$REGISTRY" ] || [ -z "$REGISTRY_ORG" ] || [ -z "$TAG" ] || [ -z "$GITHUB_TOKEN" ]; then
  echo "Error: Missing required arguments for cleanup"
  exit 1
fi

echo "=== Starting atomic cleanup for failed release ==="
echo "Registry: $REGISTRY"
echo "Organization: $REGISTRY_ORG"
echo "Tag: $TAG"
echo ""

# Extract repository owner from GITHUB_REPOSITORY (format: owner/repo)
REPO_OWNER="${GITHUB_REPOSITORY%%/*}"

# Cleanup Docker images using GitHub API
if [ -n "$SERVICES" ]; then
  echo "Cleaning up Docker images..."
  OLD_IFS="$IFS"
  IFS='|' read -ra SERVICE_ARRAY <<< "$SERVICES"
  
  for service in "${SERVICE_ARRAY[@]}"; do
    if [ -n "$service" ]; then
      IMAGE="$REGISTRY/$REGISTRY_ORG/$service:$TAG"
      echo "  Attempting to delete package: $service (tag: $TAG)"
      
      # Get package versions using GitHub API
      API_RESPONSE=$(curl -s \
        -H "Authorization: Bearer $GITHUB_TOKEN" \
        -H "Accept: application/vnd.github+json" \
        "https://api.github.com/orgs/$REPO_OWNER/packages/container/$service/versions" 2>/dev/null || echo "[]")
      
      if [ "$API_RESPONSE" != "[]" ] && [ -n "$API_RESPONSE" ]; then
        # Extract version IDs
        PACKAGE_VERSIONS=$(echo "$API_RESPONSE" | \
          grep -o "\"id\":[0-9]*" | \
          cut -d: -f2 || echo "")
        
        if [ -n "$PACKAGE_VERSIONS" ]; then
          for version_id in $PACKAGE_VERSIONS; do
            # Check if this version has the tag we're looking for
            VERSION_INFO=$(curl -s \
              -H "Authorization: Bearer $GITHUB_TOKEN" \
              -H "Accept: application/vnd.github+json" \
              "https://api.github.com/orgs/$REPO_OWNER/packages/container/$service/versions/$version_id" 2>/dev/null || echo "{}")
            
            if echo "$VERSION_INFO" | grep -q "\"$TAG\""; then
              echo "    Deleting package version ID: $version_id (tag: $TAG)"
              DELETE_RESPONSE=$(curl -s -w "%{http_code}" -o /dev/null \
                -X DELETE \
                -H "Authorization: Bearer $GITHUB_TOKEN" \
                -H "Accept: application/vnd.github+json" \
                "https://api.github.com/orgs/$REPO_OWNER/packages/container/$service/versions/$version_id")
              
              if [ "$DELETE_RESPONSE" = "204" ]; then
                echo "    ✅ Successfully deleted version $version_id"
              else
                echo "    ⚠️  Failed to delete version $version_id (HTTP $DELETE_RESPONSE)"
              fi
            fi
          done
        else
          echo "    ⚠️  No package versions found for service $service"
        fi
      else
        echo "    ⚠️  Could not fetch package versions (API error or no versions)"
      fi
      
      echo "    Image cleanup attempted: $IMAGE"
    fi
  done
  IFS="$OLD_IFS"
else
  echo "No services specified for cleanup"
fi

# Cleanup Helm charts (OCI registry deletion is limited)
if [ -n "$CHARTS" ] && [ -n "$HELM_REGISTRY" ]; then
  echo ""
  echo "Cleaning up Helm charts..."
  echo "⚠️  Note: OCI registry chart deletion via API is limited"
  echo "⚠️  Charts may need manual deletion from registry"
  
  OLD_IFS="$IFS"
  IFS='|' read -ra CHART_ARRAY <<< "$CHARTS"
  
  for chart in "${CHART_ARRAY[@]}"; do
    if [ -n "$chart" ]; then
      CHART_REF="$HELM_REGISTRY/$chart:$TAG"
      echo "  Chart marked for cleanup: $CHART_REF"
      echo "    Note: Manual deletion may be required from GitHub Packages"
    fi
  done
  IFS="$OLD_IFS"
else
  echo "No charts specified for cleanup"
fi

echo ""
echo "=== Cleanup Summary ==="
if [ -n "$SERVICES" ]; then
  echo ""
  echo "Docker Images cleanup attempted:"
  IFS='|' read -ra SERVICE_ARRAY <<< "$SERVICES"
  for service in "${SERVICE_ARRAY[@]}"; do
    [ -n "$service" ] && echo "  - $REGISTRY/$REGISTRY_ORG/$service:$TAG"
  done
fi
if [ -n "$CHARTS" ]; then
  echo ""
  echo "Helm Charts (may require manual cleanup):"
  IFS='|' read -ra CHART_ARRAY <<< "$CHARTS"
  for chart in "${CHART_ARRAY[@]}"; do
    [ -n "$chart" ] && echo "  - $HELM_REGISTRY/$chart:$TAG"
  done
fi
echo ""
echo "⚠️  Please verify cleanup was successful in GitHub Packages"
echo "⚠️  If cleanup failed, manually delete the packages listed above"
