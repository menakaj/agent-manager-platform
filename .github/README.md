# GitHub Actions Workflows

This directory contains modular GitHub Actions workflows and reusable components for the AI Agent Management Platform release process.

## Structure

```
.github/
├── workflows/
│   └── release.yml          # Main release workflow
├── actions/
│   └── build-docker-image/   # Composite action for Docker builds
│       └── action.yml
└── scripts/
    ├── extract-version.sh           # Extract release tag and version
    ├── update-install-helpers.sh    # Update chart versions in install-helpers.sh
    ├── update-helm-chart.sh        # Update and push Helm charts
    └── create-release-artifacts.sh # Create release archives and checksums
```

## Components

### Composite Actions

#### `build-docker-image`
Reusable composite action for building and pushing Docker images to GitHub Container Registry.

**Inputs:**
- `service`: Service name (directory containing Dockerfile)
- `registry`: Container registry URL (default: `ghcr.io`)
- `registry-org`: Registry organization/namespace
- `tag`: Image tag (usually release tag)
- `github-token`: GitHub token for authentication

**Usage:**
```yaml
- uses: ./.github/actions/build-docker-image
  with:
    service: agent-management-platform
    registry: ghcr.io
    registry-org: wso2
    tag: v1.0.0
    github-token: ${{ secrets.GITHUB_TOKEN }}
```

### Scripts

#### `extract-version.sh`
Extracts release tag and version (removes 'v' prefix if present).

**Usage:**
```bash
bash .github/scripts/extract-version.sh "v1.2.3"
# Outputs: tag=v1.2.3, version=1.2.3
```

#### `update-install-helpers.sh`
Updates chart version variables in `install-helpers.sh`.

**Usage:**
```bash
bash .github/scripts/update-install-helpers.sh "1.2.3" "./quick-start/install-helpers.sh"
```

#### `update-helm-chart.sh`
Updates Helm chart version, image tags, packages and pushes to OCI registry.

**Usage:**
```bash
IMAGE_UPDATES='[
  {"path": ".agentManagerService.image", "repository": "ghcr.io/wso2/agent-management-platform", "tag": "v1.2.3"}
]'
bash .github/scripts/update-helm-chart.sh \
  "agent-management-platform" \
  "./deployment/helm-charts/agent-management-platform" \
  "1.2.3" \
  "v1.2.3" \
  "oci://ghcr.io/wso2" \
  "ghcr.io" \
  "wso2" \
  "$GITHUB_TOKEN" \
  "$IMAGE_UPDATES"
```

#### `create-release-artifacts.sh`
Creates archive and SHA256 checksum for release artifacts.

**Usage:**
```bash
bash .github/scripts/create-release-artifacts.sh "v1.2.3" "./quick-start" "quick-start"
```

## Workflow: Release

The `release.yml` workflow automates the complete release process:

1. **Build and Push Docker Images** - Builds and pushes images for all services
2. **Update and Push Helm Charts** - Updates chart versions and image tags, then pushes to OCI registry
3. **Update install-helpers.sh** - Updates chart version variables and commits changes
4. **Create Release Artifacts** - Creates archive and checksum, uploads to GitHub release

### Trigger

The workflow triggers automatically when a GitHub release is created.

### Requirements

- Dockerfiles in: `agent-management-platform/`, `console/`, `traces-observer-service/`
- Helm charts in: `deployment/helm-charts/{agent-management-platform,build-ci,observability-dataprepper}/`
- `GITHUB_TOKEN` (automatically provided by GitHub Actions)

## Best Practices

- **Modularity**: Logic is separated into reusable scripts and composite actions
- **Reusability**: Components can be used independently or in other workflows
- **Maintainability**: Scripts are simple, focused, and well-documented
- **Error Handling**: Scripts use `set -euo pipefail` for strict error handling
- **Logging**: Essential steps are logged for debugging

