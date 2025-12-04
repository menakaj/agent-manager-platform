# Agent Management Platform - Installation Guide

This directory contains installation scripts for the Agent Management Platform.

> **🐳 NEW!** For the simplest experience, use the [Dev Container](#option-1-dev-container-single-docker-command) - everything in one Docker container!

## Installation Options

### Option 1: Dev Container (Single Docker Command) 🐳

**Simplest option - everything runs in a single container:**

```bash
cd quick-start
./docker-run.sh --build
```

Complete Docker-in-Docker environment with Kind, OpenChoreo, and Platform.

**Prerequisites:** Docker only (4GB+ RAM)  
**Time:** ~20-25 minutes (includes build)

[📚 Full Dev Container Guide →](./DEVCONTAINER_GUIDE.md)

---

### Option 2: Complete Bootstrap (Recommended for Development)

**Single command installs everything:**
- Kind cluster (local Kubernetes)
- OpenChoreo platform
- Agent Management Platform
- Observability stack

```bash
./bootstrap.sh
```

**Prerequisites:**
- Docker (Docker Desktop or Colima)
- kubectl
- Helm v3.8+
- kind

**Time:** ~15-20 minutes

---

### Option 3: Platform Only (For Existing OpenChoreo Clusters)

If you already have an OpenChoreo cluster with Observability Plane:

```bash
./install.sh
```

**Prerequisites:**
- **OpenChoreo cluster (v0.3.2+)** with Observability Plane installed
- **kubectl** configured with access to the cluster
- **Helm** v3.8+ installed
- Sufficient permissions to create namespaces and deploy resources

**Time:** ~5-8 minutes

## What Gets Installed

The installation includes:

1. ✅ **Agent Management Platform** - Core platform with PostgreSQL, Agent Manager Service, and Console
2. ✅ **Observability Stack** - DataPrepper and Traces Observer (always included)
3. ⚪ **Build CI** - Workflow templates for building container images (optional)

**Note**: Observability is a core component and is always installed, not optional.

---

## Bootstrap Installation (Complete Setup)

### Basic Bootstrap

```bash
./bootstrap.sh
```

This creates a complete local environment with Kind, OpenChoreo, and the Agent Management Platform.

### Bootstrap Options

```bash
# Minimal installation (core components only, faster)
./bootstrap.sh --minimal

# Verbose output for debugging
./bootstrap.sh --verbose

# Use existing Kind cluster
./bootstrap.sh --skip-kind

# Use existing OpenChoreo installation
./bootstrap.sh --skip-openchoreo

# Custom platform configuration
./bootstrap.sh --config my-values.yaml
```

### Bootstrap Flags

| Option | Description |
|--------|-------------|
| `--minimal, --core-only` | Install only core OpenChoreo components (faster) |
| `--verbose, -v` | Show detailed installation output |
| `--skip-kind` | Skip Kind cluster creation (use existing cluster) |
| `--skip-openchoreo` | Skip OpenChoreo installation (platform only) |
| `--no-port-forward` | Skip automatic port forwarding |
| `--config FILE` | Use custom configuration file |
| `--help, -h` | Show help message |

---

## Platform-Only Installation

### Simple Installation (Recommended)

```bash
./install.sh
```

This installs the platform with observability in the `agent-management-platform` namespace on an existing OpenChoreo cluster.

**What it does:**
- ✅ Validates prerequisites (including OpenChoreo Observability Plane)
- ✅ Installs Agent Management Platform
- ✅ Installs Observability components (DataPrepper + Traces Observer)
- ✅ Automatically configures port forwarding for all 4 services

**After installation, access at:**
- Console: http://localhost:3000
- API: http://localhost:8080
- Traces Observer: http://localhost:9098
- Data Prepper: http://localhost:21893

### Installation with Custom Configuration

```bash
./install.sh --config custom-values.yaml
```

### Verbose Installation (for debugging)

```bash
./install.sh --verbose
```

### Installation without Auto Port-Forward

```bash
./install.sh --no-port-forward
```

Then manually start port forwarding:
```bash
./port-forward.sh
```

## Installation Options

| Option | Description |
|--------|-------------|
| `--verbose, -v` | Show detailed installation output |
| `--no-port-forward` | Skip automatic port forwarding |
| `--config FILE` | Use custom configuration file |
| `--help, -h` | Show help message |

---

## Port Forwarding

### Automatic (Default)

Port forwarding starts automatically after installation for all 4 services:
- Console: 3000
- Agent Manager API: 8080
- Traces Observer: 9098
- Data Prepper: 21893

### Manual Control

```bash
# Start port forwarding
./port-forward.sh

# Stop port forwarding
./stop-port-forward.sh
```

---

## Validation

Installation includes built-in validation. To manually check the deployment:

```bash
# Check pod status
kubectl get pods -n agent-management-platform
kubectl get pods -n openchoreo-observability-plane

# Check services
kubectl get svc -n agent-management-platform
kubectl get svc -n openchoreo-observability-plane

# Check Helm releases
helm list -n agent-management-platform
helm list -n openchoreo-observability-plane
```

---

## Uninstallation

### Interactive Uninstall

```bash
./uninstall.sh
```

### Force Uninstall (no confirmation)

```bash
./uninstall.sh --force
```

### Complete Cleanup (including namespaces)

```bash
./uninstall.sh --force --delete-namespaces
```

**Note**: The observability namespace (`openchoreo-observability-plane`) is shared with OpenChoreo and will not be deleted.

## Uninstallation Options

| Option | Description |
|--------|-------------|
| `--force, -f` | Skip confirmation prompts |
| `--delete-namespaces` | Delete Agent Management Platform namespace after uninstalling |
| `--help, -h` | Show help message |

---

## Advanced Configuration

### Custom Values File

Create a custom values file (e.g., `my-values.yaml`):

```yaml
agentManagerService:
  replicaCount: 2
  resources:
    requests:
      memory: 512Mi
      cpu: 500m

console:
  replicaCount: 2
  
postgresql:
  auth:
    password: "my-secure-password"
```

Then install:
```bash
./install.sh --config my-values.yaml
```

### Environment Variables

You can override default namespaces:

```bash
export AMP_NS=my-custom-namespace
export OBSERVABILITY_NS=my-observability-namespace
./install.sh
```

---

## Troubleshooting

For common issues and solutions, see [TROUBLESHOOTING.md](TROUBLESHOOTING.md)

### Quick Diagnostics

```bash
# Check logs
kubectl logs -n agent-management-platform deployment/agent-manager-service
kubectl logs -n agent-management-platform deployment/console
kubectl logs -n openchoreo-observability-plane deployment/data-prepper

# Check events
kubectl get events -n agent-management-platform --sort-by='.lastTimestamp'

# Check Helm release status
helm status agent-management-platform -n agent-management-platform
helm status amp-observability-traces -n openchoreo-observability-plane
```

### Verbose Installation

If installation fails, run with verbose mode to see detailed output:

```bash
./install.sh --verbose
```

---

## Default Configuration

### Namespaces
- Agent Management Platform: `agent-management-platform`
- Observability: `openchoreo-observability-plane` (shared with OpenChoreo)
- Build CI: `agent-build-ci` (optional)

### Ports
- Console: 3000
- Agent Manager API: 8080
- Traces Observer: 9098
- Data Prepper: 21893

### Helm Charts
Charts are pulled from GitHub Container Registry (GHCR):
- `ghcr.io/agent-mgt-platform/agent-management-platform:0.1.0`
- `ghcr.io/agent-mgt-platform/amp-observability-traces:0.1.1`
- `ghcr.io/agent-mgt-platform/agent-manager-build-ci-workflows:0.1.0`

---

## Files in This Directory

| File | Purpose |
|------|---------|
| `docker-run.sh` | **NEW!** Dev container runner (Docker-in-Docker) |
| `Dockerfile.devcontainer` | **NEW!** Dev container image definition |
| `entrypoint.sh` | **NEW!** Dev container entrypoint script |
| `bootstrap.sh` | Complete installation (Kind + OpenChoreo + Platform) |
| `install.sh` | Platform installation (requires existing OpenChoreo) |
| `uninstall.sh` | Uninstallation script |
| `install-helpers.sh` | Helper functions for installation |
| `port-forward.sh` | Port forwarding for all services |
| `stop-port-forward.sh` | Stop port forwarding |
| `DEVCONTAINER_GUIDE.md` | **NEW!** Complete dev container documentation |
| `BOOTSTRAP_GUIDE.md` | Complete bootstrap documentation |
| `QUICK_START.md` | Ultra-simple guide |
| `README.md` | This detailed guide |
| `TROUBLESHOOTING.md` | Common issues and solutions |
| `example-values.yaml` | Example custom configuration |

---

## Notes

- The scripts are idempotent - running them multiple times will upgrade existing installations
- PostgreSQL is deployed as part of the Agent Management Platform chart
- Observability is always installed as a core component
- Default credentials are set in the values files - change them for production
- All scripts include proper error handling and logging
- Port forwarding runs in the background and can be stopped with `./stop-port-forward.sh`

---

## See Also

- [Quick Start Guide](QUICK_START.md) - Simple 2-step installation
- [Troubleshooting Guide](TROUBLESHOOTING.md) - Common issues and solutions
- [Main README](../README.md) - Project overview and architecture
