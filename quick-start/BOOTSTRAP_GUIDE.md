# Bootstrap Installation Guide

Complete one-command installation of the Agent Management Platform including Kind cluster and OpenChoreo.

## 🎯 What is Bootstrap?

The `bootstrap.sh` script provides a complete, automated installation that sets up:

1. **Kind Cluster** - Local Kubernetes cluster running in Docker
2. **OpenChoreo Platform** - Cloud-native application platform
3. **Agent Management Platform** - Your agent management system
4. **Observability Stack** - Complete tracing and monitoring

## ⚡ Quick Start

```bash
cd quick-start
./bootstrap.sh
```

That's it! Wait ~15-20 minutes and your platform will be ready.

## 📋 Prerequisites

Before running the bootstrap script, ensure you have:

- **Docker** - Docker Desktop or Colima running
- **kubectl** - Kubernetes command-line tool
- **Helm** - v3.8 or higher
- **kind** - Kubernetes in Docker

### Installing Prerequisites

**macOS:**
```bash
brew install kubectl helm kind
```

**Linux:**
```bash
# kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

**Docker:**
- macOS: Install [Docker Desktop](https://www.docker.com/products/docker-desktop) or use `brew install colima && colima start`
- Linux: Follow [Docker installation guide](https://docs.docker.com/engine/install/)

## 🚀 Installation Modes

### Full Installation (Default)

Installs all components including optional OpenChoreo features:

```bash
./bootstrap.sh
```

**Time:** ~15-20 minutes  
**Includes:** Core + Observability + Build Plane + Backstage + Identity Provider

### Minimal Installation (Faster)

Installs only core components:

```bash
./bootstrap.sh --minimal
```

**Time:** ~10-12 minutes  
**Includes:** Core + Observability (required for platform)

### Verbose Mode

See detailed progress and logs:

```bash
./bootstrap.sh --verbose
```

Useful for debugging or understanding what's happening.

## 🎛️ Advanced Options

### Skip Specific Steps

```bash
# Use existing Kind cluster
./bootstrap.sh --skip-kind

# Use existing OpenChoreo installation
./bootstrap.sh --skip-openchoreo

# Platform only (assumes Kind + OpenChoreo exist)
./bootstrap.sh --skip-kind --skip-openchoreo
```

### Custom Configuration

```bash
# Use custom platform configuration
./bootstrap.sh --config my-values.yaml
```

### Disable Auto Port-Forwarding

```bash
# Skip automatic port forwarding
./bootstrap.sh --no-port-forward

# Then manually start it later
./port-forward.sh
```

## 🌐 Accessing Your Platform

After installation completes, your platform is automatically accessible at:

| Service | URL | Description |
|---------|-----|-------------|
| **Console** | http://localhost:3000 | Web UI for managing agents |
| **API** | http://localhost:8080 | Agent Manager REST API |
| **Traces Observer** | http://localhost:9098 | Trace analysis service |
| **Data Prepper** | http://localhost:21893 | Trace ingestion pipeline |

### Open the Console

```bash
open http://localhost:3000
```

## 🔍 Verification

Check that everything is running:

```bash
# Check Kind cluster
kubectl get nodes

# Check OpenChoreo components
kubectl get pods -n openchoreo-control-plane
kubectl get pods -n openchoreo-data-plane
kubectl get pods -n openchoreo-observability-plane

# Check Agent Management Platform
kubectl get pods -n agent-management-platform
```

All pods should be in `Running` or `Completed` state.

## 🛑 Stopping and Starting

### Stop the Cluster (Save Resources)

```bash
# Stop Kind cluster containers
docker stop openchoreo-local-control-plane openchoreo-local-worker
```

### Start the Cluster Again

```bash
# Start Kind cluster containers
docker start openchoreo-local-control-plane openchoreo-local-worker

# Wait for pods to be ready
kubectl wait --for=condition=Ready nodes --all --timeout=120s
```

### Stop Port Forwarding

```bash
./stop-port-forward.sh
```

## 🧹 Uninstallation

### Remove Platform Only

```bash
./uninstall.sh
```

### Complete Cleanup (Everything)

```bash
# Uninstall platform
./uninstall.sh --force --delete-namespaces

# Delete Kind cluster
kind delete cluster --name openchoreo-local

# Clean up shared directory
rm -rf /tmp/kind-shared
```

## ❓ Troubleshooting

### Installation Fails

**Run with verbose output:**
```bash
./bootstrap.sh --verbose
```

**Check Docker:**
```bash
docker info
docker ps
```

**Check cluster:**
```bash
kubectl get nodes
kubectl get pods --all-namespaces
```

### Port Already in Use

If Kind cluster creation fails with "port already in use":

```bash
# Delete existing cluster
kind delete cluster --name openchoreo-local

# Try again
./bootstrap.sh
```

### Insufficient Resources

The installation requires:
- **RAM:** 4GB+ available
- **Disk:** 10GB+ free space
- **CPU:** 2+ cores

**Increase Docker resources:**
- Docker Desktop: Settings → Resources → Increase memory/CPU
- Colima: `colima stop && colima start --memory 8 --cpu 4`

### Pods Not Ready

If pods are stuck in `Pending` or `CrashLoopBackOff`:

```bash
# Check pod status
kubectl get pods -n agent-management-platform

# View pod logs
kubectl logs -n agent-management-platform <pod-name>

# Describe pod for events
kubectl describe pod -n agent-management-platform <pod-name>
```

### Services Not Accessible

If you can't access services after installation:

```bash
# Check port forwarding is running
ps aux | grep "kubectl port-forward"

# Restart port forwarding
./stop-port-forward.sh
./port-forward.sh
```

## 📚 Next Steps

1. **Deploy a Sample Agent**
   ```bash
   cd ../runtime/sample-agents/python-agent
   # Follow the README in that directory
   ```

2. **Explore the Console**
   - Navigate to http://localhost:3000
   - Create your first agent
   - View traces and logs

3. **Read the Documentation**
   - [Detailed Installation Guide](./README.md)
   - [Troubleshooting Guide](./TROUBLESHOOTING.md)
   - [Main Project README](../README.md)

## 💡 Tips

- **First time?** Use `--verbose` to see what's happening
- **Faster setup?** Use `--minimal` to skip optional components
- **Save resources?** Stop the cluster when not in use with `docker stop`
- **Multiple clusters?** Each bootstrap creates a cluster named `openchoreo-local`
- **Rerun safe:** The script is idempotent - running it again will upgrade/fix installations

## 🆘 Getting Help

If you encounter issues:

1. Check [TROUBLESHOOTING.md](./TROUBLESHOOTING.md)
2. Run with `--verbose` for detailed logs
3. Check pod status: `kubectl get pods --all-namespaces`
4. View logs: `kubectl logs -n <namespace> <pod-name>`
5. Open an issue on GitHub with logs and error messages

## 📊 Installation Timeline

**Full Installation (~15-20 minutes):**
- Prerequisites check: 10 seconds
- Kind cluster creation: 2-3 minutes
- OpenChoreo installation: 10-12 minutes
  - Cilium CNI: 1-2 minutes
  - Control Plane: 3-4 minutes
  - Data Plane: 2-3 minutes
  - Observability Plane: 4-5 minutes
- Platform installation: 3-5 minutes
- Port forwarding setup: 5 seconds

**Minimal Installation (~10-12 minutes):**
- Prerequisites check: 10 seconds
- Kind cluster creation: 2-3 minutes
- OpenChoreo core: 7-8 minutes
- Platform installation: 3-5 minutes
- Port forwarding setup: 5 seconds

---

**Ready to get started?**

```bash
./bootstrap.sh
```

Happy agent building! 🚀

