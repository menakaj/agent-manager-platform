#!/usr/bin/env bash
# Update Helm chart versions and image tags
# Usage: update-helm-charts.sh <target-version> <release-tag>

set -euo pipefail

TARGET_VERSION="${1:-}"
RELEASE_TAG="${2:-}"

if [ -z "$TARGET_VERSION" ] || [ -z "$RELEASE_TAG" ]; then
  echo "Error: Missing required arguments"
  echo "Usage: update-helm-charts.sh <target-version> <release-tag>"
  exit 1
fi

# Install yq if not available
if ! command -v yq &> /dev/null; then
  echo "Installing yq..."
  wget -qO /usr/local/bin/yq https://github.com/mikefarah/yq/releases/download/v4.40.5/yq_linux_amd64
  chmod +x /usr/local/bin/yq
fi

# Update wso2-ai-agent-management-platform chart
if [ -d "./deployments/helm-charts/wso2-ai-agent-management-platform" ]; then
  yq eval -i ".version = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-ai-agent-management-platform/Chart.yaml"
  yq eval -i ".appVersion = \"$RELEASE_TAG\"" "./deployments/helm-charts/wso2-ai-agent-management-platform/Chart.yaml"
  if [ -f "./deployments/helm-charts/wso2-ai-agent-management-platform/values.yaml" ]; then
    yq eval -i ".agentManagerService.image.tag = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-ai-agent-management-platform/values.yaml" || true
    yq eval -i ".console.image.tag = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-ai-agent-management-platform/values.yaml" || true
  fi
  echo "Updated wso2-ai-agent-management-platform chart"
fi

# Update build-ci chart
if [ -d "./deployments/helm-charts/wso2-amp-build-extension" ]; then
  yq eval -i ".version = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-amp-build-extension/Chart.yaml"
  yq eval -i ".appVersion = \"$RELEASE_TAG\"" "./deployments/helm-charts/wso2-amp-build-extension/Chart.yaml"
  echo "Updated build-ci chart"
fi

# Update observability-dataprepper chart
if [ -d "./deployments/helm-charts/wso2-amp-observability-extension" ]; then
  yq eval -i ".version = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-amp-observability-extension/Chart.yaml"
  yq eval -i ".appVersion = \"$RELEASE_TAG\"" "./deployments/helm-charts/wso2-amp-observability-extension/Chart.yaml"
  if [ -f "./deployments/helm-charts/wso2-amp-observability-extension/values.yaml" ]; then
    yq eval -i ".tracesObserverService.image.tag = \"$TARGET_VERSION\"" "./deployments/helm-charts/wso2-amp-observability-extension/values.yaml" || true
  fi
  echo "Updated wso2-amp-observability-extension chart"
fi

echo "✅ Updated all Helm chart versions"

